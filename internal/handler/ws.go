package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // dev 放行;生产收紧为白名单
}

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = 30 * time.Second
)

// WS 长连接:应用层心跳(服务端 ping + 客户端 pong)+ echo。
// 这套「保活 + 探活」逻辑与设备通信里的心跳同源:
//   - pingPeriod 定期 ping 保活(中间设备不清理连接)
//   - pongWait 超时未收到 pong 判死(读 deadline 到期自动断开)
func WS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("ws upgrade: %v", err)
		return
	}
	defer conn.Close()

	// 读超时 = pongWait:到期未收到 pong 即判死,触发客户端重连
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	// 服务端心跳:每 pingPeriod 发一次 ping,失败则结束连接
	go func() {
		ticker := time.NewTicker(pingPeriod)
		defer ticker.Stop()
		for range ticker.C {
			if err := conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait)); err != nil {
				return
			}
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
}

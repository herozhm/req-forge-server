package main

import (
	"log"
	"net/http"
	"req-forge-server/internal/handler"
	"req-forge-server/internal/middleware"
	"req-forge-server/internal/store"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := store.InitDB("req-forge.db")
	if err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}

	r := gin.Default()

	// CORS:前端 dev 跨域(生产可收紧为白名单)
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	h := handler.NewDecompositionHandler(db)

	// 公开路由:登录签发 JWT
	r.POST("/api/login", handler.Login)

	// 受保护路由:需要 Authorization: Bearer <token>
	api := r.Group("/api")
	api.Use(middleware.Auth())
	{
		api.GET("/decompositions", h.List)
		api.POST("/decompositions", h.Create)
		api.GET("/decompositions/:id", h.Get)
		api.DELETE("/decompositions/:id", h.Delete)
	}

	// WebSocket 长连接(心跳 + echo)
	r.GET("/ws", handler.WS)

	log.Println("req-forge-server 监听 :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

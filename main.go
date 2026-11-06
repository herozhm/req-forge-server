package main

import (
	"log"
	"net/http"
	"req-forge-server/internal/handler"
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

	api := r.Group("/api")
	{
		api.GET("/decompositions", h.List)          // 列表(最近 50 条)
		api.POST("/decompositions", h.Create)       // 保存一次拆解
		api.GET("/decompositions/:id", h.Get)       // 详情
		api.DELETE("/decompositions/:id", h.Delete) // 删除
	}

	log.Println("req-forge-server 监听 :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

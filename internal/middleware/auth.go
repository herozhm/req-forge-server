package middleware

import (
	"net/http"
	"strings"

	"req-forge-server/internal/auth"

	"github.com/gin-gonic/gin"
)

// Auth JWT 鉴权中间件:校验 Authorization: Bearer <token>
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if !strings.HasPrefix(h, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "缺少 token"})
			return
		}
		tokenStr := strings.TrimPrefix(h, "Bearer ")
		if _, err := auth.Verify(tokenStr); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token 无效或过期"})
			return
		}
		c.Next()
	}
}

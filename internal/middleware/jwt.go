package middleware

import (
	"github.com/gin-gonic/gin"
	"job-platform-go/pkg/e"
	"job-platform-go/pkg/response"
	"job-platform-go/pkg/utils"
	"strings"
)

// JWTAuth 鉴权中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//1.获取 Authorization Header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.ErrorWithStatus(c, 401, e.ERROR_UNAUTHORIZED, "未携带 Token")
			c.Abort()
			return

		}
		//2.验证格式 "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.ErrorWithStatus(c, 401, e.ERROR_UNAUTHORIZED, "Token 格式错误")
			c.Abort()
			return
		}
		//3.解析 Token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			response.ErrorWithStatus(c, 401, e.ERROR_UNAUTHORIZED, "Token 无效或已过期")
			c.Abort()
			return
		}

		//4. 将用户信息存入上下文 (Context)
		//后续 Handler 可以通过 c.GetInt("userID") 获取
		c.Set("userID", claims.ID)
		c.Set("userRole", claims.Role)
		c.Set("userEmail", claims.Subject)
		c.Next()

	}
}

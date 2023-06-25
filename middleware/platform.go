package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/service"
	"net/http"
	"time"
)

var secretKeyService = service.ServiceGroupApp.OpenfishServiceGroup.SecretKeyService

func PlatformOpenApiSkAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/v1/chat/completions" && (c.Request.Method == "OPTIONS" || c.Request.Method == "options") {
			c.Next()
		}
		sk := c.Request.Header.Get("Authorization")
		if sk == "" || sk[:6] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": map[string]string{
					"message": "未获得sk或非法访问",
					"type":    "invalid_request_error",
					"param":   "",
					"code":    "",
				},
			})
			c.Abort()
			return
		}
		sk = sk[7:]
		key, err := secretKeyService.GetSecretKeyBySk(sk)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": map[string]string{
					"message": "您的sk有误，请检查是否存在或已禁用",
					"type":    "invalid_request_error",
					"param":   "",
					"code":    "invalid_api_key",
				},
			})
			c.Abort()
			return
		}
		// 判断是否过期
		unixMilli := time.Now().UnixMilli()
		if *key.Expire != 0 && *key.Expire < unixMilli {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": map[string]string{
					"message": "您的sk已过期",
					"type":    "invalid_request_error",
					"param":   "",
					"code":    "invalid_api_key",
				},
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

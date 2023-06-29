package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/service"
	"net/http/httputil"
	"net/url"
	"sync"
)

type OpenApi struct{}

var userRequestStatus sync.Map
var secretKeyService = service.ServiceGroupApp.OpenfishServiceGroup.SecretKeyService

func (o *OpenApi) ForwardChatCompletionsApi(c *gin.Context) {
	// 获取用户ID
	var userID uint
	sk := c.Request.Header.Get("Authorization")
	sk = sk[7:]
	key, err := secretKeyService.GetSecretKeyBySk(sk)
	if err != nil {
		userID = 0
	}
	userID = key.CreatedBy
	// 检查用户的请求状态
	_, loaded := userRequestStatus.LoadOrStore(userID, true)
	if loaded {
		c.JSON(429, gin.H{"msg": "太多请求了"})
		return
	}
	defer userRequestStatus.Delete(userID) // 在处理完毕后删除用户的请求状态
	// 创建目标URL
	targetURL := "http://127.0.0.1:8080" // 更改为您实际的目标服务URL

	// 创建反向代理器
	target, _ := url.Parse(targetURL)
	proxy := httputil.NewSingleHostReverseProxy(target)

	// 更改请求头中的Host字段
	c.Request.Host = target.Host

	// 进行反向代理
	proxy.ServeHTTP(c.Writer, c.Request)
}

func (o *OpenApi) ForwardOptionsChatCompletionsApi(c *gin.Context) {
	// Set headers for CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Headers", "*")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/service"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type OpenApi struct{}

var userRequestStatus sync.Map
var secretKeyService = service.ServiceGroupApp.OpenfishServiceGroup.SecretKeyService
var mailAccountService = service.ServiceGroupApp.OpenfishServiceGroup.MailAccountService

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
	defer userRequestStatus.Delete(userID) // 在处理完毕后删除用户的请求状态
	if loaded {
		c.JSON(429, gin.H{"error": map[string]interface{}{
			"message": "您的请求过多，系统限制并发请求为1",
			"type":    "requests",
			"param":   nil,
			"code":    nil,
		}})
		return
	}
	// 创建目标URL
	server, _ := mailAccountService.GetServerNodeByUpdatedAtAsc() // 更改为您实际的目标服务URL
	mailAccount, _ := mailAccountService.GetAccessTokenByUpdatedAtAsc()

	// 创建反向代理器
	target, _ := url.Parse(server)
	proxy := httputil.NewSingleHostReverseProxy(target)

	// 更改请求头中的Host字段
	c.Request.Host = target.Host
	c.Request.Header.Set("Authorization", "Bearer "+mailAccount.OpenaiAccessToken)
	// 创建一个自定义的 http.Client 并设置超时时间
	client := &http.Client{
		Timeout: time.Second * 3600,
	}
	// 使用自定义的 http.Client 进行反向代理
	proxy.Transport = client.Transport
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

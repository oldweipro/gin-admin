package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/chatgpt2api"
)

type OpenApi struct{}

func (o *OpenApi) ForwardChatCompletionsApi(c *gin.Context) {
	chatgpt2api.Nightmare(c, "http://127.0.0.1:7890")
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

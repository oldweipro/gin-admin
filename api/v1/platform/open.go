package platform

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type OpenApi struct{}

func (o *OpenApi) ForwardChatCompletionsApi(c *gin.Context) {
	// 获取原始请求的数据
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无法读取请求体"})
		return
	}

	// 构建转发请求
	forwardURL := "http://127.0.0.1:8080/v1/chat/completions" // 第三方接口的URL
	req, err := http.NewRequest(c.Request.Method, forwardURL, bytes.NewBuffer(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建转发请求"})
		return
	}

	// 设置原始请求的Header
	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// 发送转发请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法发送转发请求"})
		return
	}
	defer resp.Body.Close()

	// 读取转发响应的数据
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法读取转发响应"})
		return
	}

	// 将转发响应返回给客户端
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), respBody)
}

func (o *OpenApi) ForwardOptionsChatCompletionsApi(c *gin.Context) {
	// 获取原始请求的数据
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无法读取请求体"})
		return
	}

	// 构建转发请求
	forwardURL := "http://127.0.0.1:8080/v1/chat/completions" // 第三方接口的URL
	req, err := http.NewRequest(c.Request.Method, forwardURL, bytes.NewBuffer(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建转发请求"})
		return
	}

	// 设置原始请求的Header
	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// 发送转发请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法发送转发请求"})
		return
	}
	defer resp.Body.Close()

	// 读取转发响应的数据
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法读取转发响应"})
		return
	}

	// 将转发响应返回给客户端
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), respBody)
}

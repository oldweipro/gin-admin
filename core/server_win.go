//go:build windows
// +build windows

package core

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    70 * time.Second,
		WriteTimeout:   70 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

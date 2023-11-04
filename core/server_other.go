//go:build !windows
// +build !windows

// Package core
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-11 16:35
package core

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func initServer(router *gin.Engine) server {
	//s := endless.NewServer(address, router)
	//s.ReadHeaderTimeout = 30 * time.Minute
	//s.WriteTimeout = 30 * time.Minute
	//s.MaxHeaderBytes = 1 << 23
	//return s

	return &http.Server{
		Handler:        router,
		ReadTimeout:    30 * time.Minute,
		WriteTimeout:   30 * time.Minute,
		MaxHeaderBytes: 1 << 23,
	}
}

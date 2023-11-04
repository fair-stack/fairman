// Package middleware
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2023-04-30 21:53
package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func EventMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("H1")

		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Next()
	}
}

func EventAddClientMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//traceId := c.Query("traceId")
		//fmt.Println("traceId2: ", traceId)
		//
		//global.BaseSEEChan.Message[traceId] = make(chan string)
		//clientChan := make(chan string)
		//// Initialize client channel
		//
		//// Send new connection to event server
		//global.BaseSEEChan.NewClients <- clientChan
		//
		//defer func() {
		//	// Send closed connection to event server
		//	global.BaseSEEChan.ClosedClients <- clientChan
		//}()

		//c.Set("clientChan", global.BaseSEEChan.Message[traceId])

		c.Next()
	}
}

// Package middleware
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-16 17:51
package middleware

import (
	"bytes"
	"cnic/fairman-gin/global"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LogLayout journallayout
type LogLayout struct {
	Time      time.Time
	Metadata  map[string]interface{} // Store custom raw data
	Path      string                 // Access Path
	Query     string                 // carryquery
	Body      string                 // carrybodycarry
	IP        string                 // ipaddress
	UserAgent string                 // agent
	Error     string                 // error
	Cost      time.Duration          // Spend time
	Source    string                 // source
}

type Logger struct {
	// Filter User defined filtering
	Filter func(c *gin.Context) bool
	// FilterKeyword Keyword filtering(key)
	FilterKeyword func(layout *LogLayout) bool
	// AuthProcess Authentication processing
	AuthProcess func(c *gin.Context, layout *LogLayout)
	// Log processing
	Print func(LogLayout)
	// Source Service Unique Identification
	Source string
}

func (l Logger) SetLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		var body []byte
		if l.Filter != nil && !l.Filter(c) {
			body, _ = c.GetRawData()
			// Replace the originalbodyReplace the original
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
		c.Next()
		cost := time.Since(start)
		layout := LogLayout{
			Time:      time.Now(),
			Path:      path,
			Query:     query,
			IP:        c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
			Error:     strings.TrimRight(c.Errors.ByType(gin.ErrorTypePrivate).String(), "\n"),
			Cost:      cost,
			Source:    l.Source,
		}
		if l.Filter != nil && !l.Filter(c) {
			layout.Body = string(body)
		}
		// Information required for processing authentication
		if l.AuthProcess != nil {
			l.AuthProcess(c, &layout)
		}
		if l.FilterKeyword != nil {
			// Self judgmentkey/value Self judgment
			l.FilterKeyword(&layout)
		}
		// Self processing logs
		l.Print(layout)
	}
}

func DefaultLogger() gin.HandlerFunc {
	return Logger{
		Print: func(layout LogLayout) {
			// standard output ,k8sstandard output
			v, _ := json.Marshal(layout)
			fmt.Println(string(v))
		},
		Source: "FairManServer",
	}.SetLoggerMiddleware()
}

// GinLogger receiveginreceive
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		global.FairLog.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

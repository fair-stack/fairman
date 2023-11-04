// Package middleware
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-16 17:46
package middleware

import (
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/model/v2/common/response"
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GinRecovery recoverPossible occurrences of dropped itemspanic
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					global.FairLog.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					//c.Error(err.(error)) // nolint: errcheck
					_ = c.Error(errors.New("Internal system error")) // nolint: errcheck
					response.ResultAll(http.StatusInternalServerError, gin.H{}, "Internal system error-1", c)
					c.Abort()
					return
				}

				if stack {
					global.FairLog.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
					response.ResultAll(http.StatusInternalServerError, gin.H{}, "Internal system error-2", c)
				} else {
					global.FairLog.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					response.ResultAll(http.StatusInternalServerError, gin.H{}, "Internal system error-3", c)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

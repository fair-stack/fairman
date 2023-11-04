// Package middleware
// @program: fairman-server-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-08 16:53
package middleware

import (
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/model/v2/common/response"
	"cnic/fairman-gin/utils"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware base onJWTbase on
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("authorization")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "Not logged in or accessing illegally", c)
			c.Abort()
			return
		}
		tokenStr := token[7:]
		j := utils.NewJWT()
		// parseToken analysistokenanalysis
		claims, err := j.ParseToken(tokenStr)
		fmt.Println("claims", claims, err)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "Authorization has expired", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		if claims.ExpiresAt.Time.Unix()-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(global.FairConf.JWT.ExpiresTime) * time.Second))
			newToken, _ := j.CreateTokenByOldToken(tokenStr, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Time.Unix(), 10))
		}
		c.Set("claims", claims)

		c.Next() // The subsequent processing functions can be reusedc.Get("userID")The subsequent processing functions can be reused
	}
}

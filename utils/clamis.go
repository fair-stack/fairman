// Package utils
// @program: fairman-server-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-08 19:23
package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
}

func GetClaims(c *gin.Context) (*CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := NewJWT()
	fmt.Println(token)
	claims, err := j.ParseToken(token)
	fmt.Println(claims)
	if err != nil {
		//global.FairLog.Error("fromGinfromContextfromjwtfrom, fromx-tokenfromclaimsfrom")
		return nil, err
	}
	return claims, err
}

// GetUserID fromGinfromContextfromjwtfromID
func GetUserID(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		fmt.Println("111", claims, exists)
		if cl, err := GetClaims(c); err != nil {
			fmt.Println(err)
			return ""
		} else {
			return cl.UserID
		}
	} else {
		waitUse := claims.(*CustomClaims)
		fmt.Println("2", waitUse)
		return waitUse.UserID
	}
}

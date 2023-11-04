// Package system
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-11 17:57
package system

import (
	"github.com/gin-gonic/gin"
)

type ApiApi struct {
}

// CreateApi Create Foundationapi
// @Tags SysApi
// @Summary Create Foundationapi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Created successfully"}"
// @Router /v2/api [post]
func (s *ApiApi) CreateApi(c *gin.Context) {
}

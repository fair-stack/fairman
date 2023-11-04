// Package software
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-28 10:56
package software

import (
	"github.com/gin-gonic/gin"
)

type TemplateRouter struct {
}

func (s *TemplateRouter) InitTemplateRouter(Router *gin.RouterGroup) {
	//templateRouterWithoutRecord := Router.Group("template")
	//var templateRouterApi = v2.ApiV2GroupApp.Software.TemplateApi
	//{
	//templateRouterWithoutRecord.GET("config/:id", templateRouterApi.GetTemplateConfigById)
	//templateRouterWithoutRecord.GET("/:id", templateRouterApi.GetTemplateById)
	//}
}

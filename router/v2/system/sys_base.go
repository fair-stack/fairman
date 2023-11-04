// Package system
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-16 17:14
package system

import (
	v2 "cnic/fairman-gin/api/v2"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct {
}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	baseRouterWithoutRecord := Router.Group("base")
	var baseRouterApi = v2.ApiV2GroupApp.System.BaseApi
	{
		//baseRouterWithoutRecord.POST("login", baseRouterApi.Login)
		//baseRouterWithoutRecord.POST("restart", baseRouterApi.Restart)
		//baseRouterWithoutRecord.GET("test", baseRouterApi.Test)
		baseRouterWithoutRecord.GET("logs", baseRouterApi.BaseLogs)
		baseRouterWithoutRecord.GET("unique", baseRouterApi.GetUnique)
	}
	{
		baseRouterWithoutRecord.POST("update", baseRouterApi.UpdateApp)
		baseRouterWithoutRecord.POST("file", baseRouterApi.UploadFile)

	}
}

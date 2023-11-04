// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-04 16:10
package docker

import (
	v2 "cnic/fairman-gin/api/v2"

	"github.com/gin-gonic/gin"
)

type EndpointRoute struct {
}

func (s *EndpointRoute) InitEndpointRouter(Router *gin.RouterGroup) {
	endpointRouterWithoutRecord := Router.Group("endpoint")
	var endpointRouterApi = v2.ApiV2GroupApp.Docker.EndpointApi
	{
		//endpointRouterWithoutRecord.POST("", endpointRouterApi.CreateEndpoint)
		endpointRouterWithoutRecord.POST("test", endpointRouterApi.TestEndpoint)
		endpointRouterWithoutRecord.POST("test/one", endpointRouterApi.TestOneEndpoint)
		//endpointRouterWithoutRecord.POST("config", endpointRouterApi.SaveEndpointConfig)
	}
	{
		//endpointRouterWithoutRecord.GET("/:id", endpointRouterApi.GetEndpoint)
		//endpointRouterWithoutRecord.GET("", endpointRouterApi.GetEndpointAll)
	}
}

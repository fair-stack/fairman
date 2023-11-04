// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-10 16:47
package docker

import (
	v2 "cnic/fairman-gin/api/v2"

	"github.com/gin-gonic/gin"
)

type SoftwareRouter struct {
}

func (s *SoftwareRouter) InitSoftwareRouter(Router *gin.RouterGroup) {
	softwareRouterWithoutRecord := Router.Group("software")
	var softwareRouterApi = v2.ApiV2GroupApp.Software.SoftwareApi
	{
		softwareRouterWithoutRecord.POST("", softwareRouterApi.CreateSoftware)
		softwareRouterWithoutRecord.POST("status", softwareRouterApi.GetSoftwareStatus)
		softwareRouterWithoutRecord.POST("stop", softwareRouterApi.StopSoftware)
		softwareRouterWithoutRecord.POST("start", softwareRouterApi.StartSoftware)
		softwareRouterWithoutRecord.PUT("", softwareRouterApi.UpdateSoftware)
		softwareRouterWithoutRecord.DELETE("", softwareRouterApi.RemoveSoftware)
		softwareRouterWithoutRecord.POST("endpoint", softwareRouterApi.SaveEndpoint)
		softwareRouterWithoutRecord.PUT("inspect", softwareRouterApi.UpdateSoftwareInspect)
		softwareRouterWithoutRecord.PUT("inspect1", softwareRouterApi.UpdateSoftwareInspect1)
		//softwareRouterWithoutRecord.DELETE("service", softwareRouterApi.RemoveSoftwareService)
		//softwareRouterWithoutRecord.PUT("replicas", softwareRouterApi.UpdateSoftwareReplicas)
		//softwareRouterWithoutRecord.GET("", softwareRouterApi.GetSoftware)
		//softwareRouterWithoutRecord.GET("test222", softwareRouterApi.Test)
		//softwareRouterWithoutRecord.POST("test", softwareRouterApi.Test1)
		//softwareRouterWithoutRecord.POST("stop", softwareRouterApi.StopSoftware)
		//softwareRouterWithoutRecord.POST("start", softwareRouterApi.StartSoftware)
		//softwareRouterWithoutRecord.GET("details", softwareRouterApi.GetSoftwareDetails)
		//softwareRouterWithoutRecord.GET("inspect", softwareRouterApi.GetSoftwareInspect)
		//softwareRouterWithoutRecord.PUT("/:id/service", softwareRouterApi.UpdateSoftwareService)
		//softwareRouterWithoutRecord.GET("/service/logs", softwareRouterApi.GetSoftwareLogs)
	}
}

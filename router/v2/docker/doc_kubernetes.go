// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-12-09 07:32
package docker

import (
	v2 "cnic/fairman-gin/api/v2"

	"github.com/gin-gonic/gin"
)

type KubernetesRoute struct {
}

func (s *KubernetesRoute) InitKubernetesRouter(Router *gin.RouterGroup) {
	kubernetesRouterWithoutRecord := Router.Group("kubernetes")
	var kubernetesRouterApi = v2.ApiV2GroupApp.Docker.KubernetesApi
	{
		//kubernetesRouter.POST("container/exec/create", kubernetesRouterApi.ContainerExecCreate)
	}
	{
		//kubernetesRouter.GET("containers", kubernetesRouterApi.ContainerList)
		kubernetesRouterWithoutRecord.GET("pods", kubernetesRouterApi.GetPodList)
		kubernetesRouterWithoutRecord.GET("pod/list/details", kubernetesRouterApi.GetPodListDetails)
		kubernetesRouterWithoutRecord.GET("pod/terminal", kubernetesRouterApi.PodTerminal) // websocket
		kubernetesRouterWithoutRecord.GET("pod/logs", kubernetesRouterApi.PodLog)          // websocket
	}
}

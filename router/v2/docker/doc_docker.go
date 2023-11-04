// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-29 15:21
package docker

import (
	v2 "cnic/fairman-gin/api/v2"

	"github.com/gin-gonic/gin"
)

type DockerRoute struct {
}

func (s *DockerRoute) InitDockerRouter(Router *gin.RouterGroup) {
	dockerRouter := Router.Group("docker")
	var dockerRouterApi = v2.ApiV2GroupApp.Docker.DockerApi
	{
		dockerRouter.POST("container/exec/create", dockerRouterApi.ContainerExecCreate)
		dockerRouter.POST("container/exec/resize", dockerRouterApi.ContainerExecResize)
		dockerRouter.PUT("service", dockerRouterApi.UpdateService)
		dockerRouter.PUT("container/inspect", dockerRouterApi.ContainerInspectUpdate)
	}
	{
		//softwareRouterWithoutRecord.POST("containers", softwareRouterApi.GetSoftwareContainers)
		dockerRouter.GET("containers", dockerRouterApi.ContainerList)
		dockerRouter.GET("container/terminal", dockerRouterApi.ContainerTerminal) // websocket
		dockerRouter.GET("container/logs", dockerRouterApi.ContainerLogs)
		dockerRouter.GET("container/inspect", dockerRouterApi.ContainerInspect)
		dockerRouter.GET("services", dockerRouterApi.ServiceList)
		dockerRouter.GET("service/logs", dockerRouterApi.ServiceLogs)
		dockerRouter.GET("service/inspect", dockerRouterApi.ServiceInspect)
		dockerRouter.GET("service/runtime", dockerRouterApi.ServiceRunningTime)
		dockerRouter.GET("tasks", dockerRouterApi.TaskList)

		//dockerRouter.GET("/:id", s.GetDocker)
	}
}

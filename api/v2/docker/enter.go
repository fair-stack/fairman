// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-04 15:20
package docker

import "cnic/fairman-gin/service"

type ApiGroup struct {
	EndpointApi
	DockerApi
	KubernetesApi
}

var endpointService = service.ServiceGroupApp.Docker.EndpointService
var softwareService = service.ServiceGroupApp.Docker.SoftwareService
var dockerService = service.ServiceGroupApp.Docker.DockerService
var KubernetesService = service.ServiceGroupApp.Docker.KubernetesService

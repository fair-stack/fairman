// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-04 15:24
package docker

type ServiceGroup struct {
	EndpointService
	SoftwareService
	DockerService
	KubernetesService
}

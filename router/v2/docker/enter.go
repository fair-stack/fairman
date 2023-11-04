// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-04 16:10
package docker

type RouterGroup struct {
	EndpointRoute
	SoftwareRouter
	DockerRoute
	KubernetesRoute
}

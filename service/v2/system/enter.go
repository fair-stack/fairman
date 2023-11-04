// Package system
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-16 17:17
package system

type ServiceGroup struct {
	BaseService
	EndpointService
}

var (
	AddEndpoint = "%s/v2/client/endpoint"
)

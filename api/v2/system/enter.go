// Package system
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-11 17:56
package system

import "cnic/fairman-gin/service"

type ApiGroup struct {
	ApiApi
	BaseApi
	EndpointApi
}

var baseService = service.ServiceGroupApp.System.BaseService
var endpointService = service.ServiceGroupApp.System.EndpointService

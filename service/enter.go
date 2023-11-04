// Package service
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-17 15:22
package service

import (
	"cnic/fairman-gin/service/v2/docker"
	"cnic/fairman-gin/service/v2/software"
	"cnic/fairman-gin/service/v2/system"
)

type ServiceGroup struct {
	Software software.ServiceGroup
	System   system.ServiceGroup
	Docker   docker.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

// Package software
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-10 16:28
package software

import (
	"cnic/fairman-gin/api/v2/docker"
)

type ApiGroup struct {
	docker.SoftwareApi
	TemplateApi
}

//var templateService = service.ServiceGroupApp.Software.TemplateService

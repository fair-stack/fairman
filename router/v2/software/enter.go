// Package software
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-10 16:47
package software

import "cnic/fairman-gin/router/v2/docker"

type RouterGroup struct {
	docker.SoftwareRouter
	TemplateRouter
}

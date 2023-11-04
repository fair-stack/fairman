// Package v2
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-11 17:59
package v2

import (
	"cnic/fairman-gin/router/v2/docker"
	"cnic/fairman-gin/router/v2/software"
	"cnic/fairman-gin/router/v2/system"
)

type GroupV2 struct {
	System   system.RouterGroup
	Software software.RouterGroup
	Docker   docker.RouterGroup
}

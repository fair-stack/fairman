// Package v2
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-11 17:56
package v2

import (
	"cnic/fairman-gin/api/v2/docker"
	"cnic/fairman-gin/api/v2/software"
	"cnic/fairman-gin/api/v2/system"
)

type ApiV2Group struct {
	System   system.ApiGroup
	Software software.ApiGroup
	Docker   docker.ApiGroup
}

var ApiV2GroupApp = new(ApiV2Group)

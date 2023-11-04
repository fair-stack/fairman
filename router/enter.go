// Package router
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-30 21:33
package router

import (
	v2 "cnic/fairman-gin/router/v2"
)

type Group struct {
	V2 v2.GroupV2
}

var GroupApp = new(Group)

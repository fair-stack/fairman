// Package global
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-10 16:23
package global

import "sync"

var (
	DockerCmd        = "cmd/docker"
	DockerComposeCmd = "cmd/docker-compose"
	KubectlCmd       = "cmd/kubectl"
	KomposeCmd       = "cmd/kompose"
	HelmCmd          = "cmd/helm"
)

var (
	BaseChan    chan ChanLogModal
	BaseSEEChan sync.Map

	Test sync.Map
)

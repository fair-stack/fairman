// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-10 18:38
package docker

import (
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/model/v2/docker"
	"context"
	"sync"
)

type StackDeployer interface {
	DeploySwarmStack(stack docker.DocSoftware, prune bool) error
	DeployComposeStack(stack docker.DocSoftware) error
	//DeployKubernetesStack(stack *portainer.Stack, endpoint *portainer.Endpoint, user *portainer.User) error
}

type stackDeployer struct {
	lock                *sync.Mutex
	swarmStackManager   global.SwarmStackManager
	composeStackManager global.ComposeStackManager
	//kubernetesDeployer  portainer.KubernetesDeployer
}

// NewStackDeployer inits a stackDeployer struct with a SwarmStackManager, a ComposeStackManager and a KubernetesDeployer
func NewStackDeployer(swarmStackManager global.SwarmStackManager, composeStackManager global.ComposeStackManager) *stackDeployer {
	return &stackDeployer{
		lock:                &sync.Mutex{},
		swarmStackManager:   swarmStackManager,
		composeStackManager: composeStackManager,
	}
}

func (d *stackDeployer) DeploySwarmStack(stack docker.DocSoftware, prune bool) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.swarmStackManager.Login(stack.Template.Config.SwarmRegistry, stack.Endpoint) // To be completed
	defer d.swarmStackManager.Logout(stack.Endpoint)                               // To be completed

	return d.swarmStackManager.Deploy(stack, prune)
}

func (d *stackDeployer) DeployComposeStack(stack docker.DocSoftware, isLog bool) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.composeStackManager.Login(stack.Template.Config.SwarmRegistry, stack.Endpoint) // TODO: To be completed
	defer d.composeStackManager.Logout(stack.Endpoint)                               // TODO: To be completed

	return d.composeStackManager.Up(context.TODO(), stack, isLog)
}

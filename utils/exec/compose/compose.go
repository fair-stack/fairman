// Package compose
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-15 15:32
package compose

import (
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/utils/exec/compose/errors"
	"cnic/fairman-gin/utils/exec/compose/initernal/composebinary"
	"cnic/fairman-gin/utils/exec/compose/initernal/composeplugin"
)

// NewComposeDeployer will try to create a wrapper for docker-compose binary
// if it's not availbale it will try to create a wrapper for docker-compose plugin
func NewComposeDeployer(binaryPath, configPath string) (global.ComposeDeployer, error) {
	deployer, err := composebinary.NewComposeWrapper(binaryPath, configPath)
	if err == nil {
		return deployer, nil
	}

	if err == errors.ErrBinaryNotFound {
		return composeplugin.NewPluginWrapper(binaryPath, configPath)
	}

	return nil, err
}

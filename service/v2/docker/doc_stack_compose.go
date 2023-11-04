// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-15 15:19
package docker

import (
	"bytes"
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/model/v2/docker"
	dockerV2Req "cnic/fairman-gin/model/v2/docker/request"
	"cnic/fairman-gin/utils"
	"cnic/fairman-gin/utils/exec"
	"cnic/fairman-gin/utils/filesystem"
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"

	"github.com/pkg/errors"
)

const (
	// ProjectLabel allow to track resource related to a compose project
	ProjectLabel = "com.docker.compose.project"
	// ContainerLabel allow to track resource related to a compose container
	ContainerLabel = "com.docker.container.name"
	// ServiceLabel allow to track resource related to a compose service
	ServiceLabel = "com.docker.compose.service"
	// ConfigHashLabel stores configuration hash for a compose service
	ConfigHashLabel = "com.docker.compose.config-hash"
	// ContainerNumberLabel stores the container index of a replicated service
	ContainerNumberLabel = "com.docker.compose.container-number"
	// VolumeLabel allow to track resource related to a compose volume
	VolumeLabel = "com.docker.compose.volume"
	// NetworkLabel allow to track resource related to a compose network
	NetworkLabel = "com.docker.compose.network"
	// WorkingDirLabel stores absolute path to compose project working directory
	WorkingDirLabel = "com.docker.compose.project.working_dir"
	// ConfigFilesLabel stores absolute path to compose project configuration files
	ConfigFilesLabel = "com.docker.compose.project.config_files"
	// EnvironmentFileLabel stores absolute path to compose project env file set by `--env-file`
	EnvironmentFileLabel = "com.docker.compose.project.environment_file"
	// OneoffLabel stores value 'True' for one-off containers created by `compose run`
	OneoffLabel = "com.docker.compose.oneoff"
	// SlugLabel stores unique slug used for one-off container identity
	SlugLabel = "com.docker.compose.slug"
	// ImageDigestLabel stores digest of the container image used to run service
	ImageDigestLabel = "com.docker.compose.image"
	// DependenciesLabel stores service dependencies
	DependenciesLabel = "com.docker.compose.depends_on"
	// VersionLabel stores the compose tool version used to build/run application
	VersionLabel = "com.docker.compose.version"
	// ImageBuilderLabel stores the builder (classic or BuildKit) used to produce the image.
	ImageBuilderLabel = "com.docker.compose.image.builder"
)

// createComposeStackFormFileContent create a compose stack from a compose file content
func (s *SoftwareService) createComposeStackFromFileContent(software docker.DocSoftware) error {
	// 1. initialization docker compose
	fmt.Println(global.FairConf.Docker)
	manager, err := exec.NewComposeStackManager(global.FairConf.Docker.BinaryPath, "")
	if err != nil {
		return utils.Errorf(err)
	}

	// Processing software name Processing software name
	software.Name = manager.NormalizeStackName(software.Name)

	// Determine if the software name exists
	isUnique, err := s.checkUniqueComposeStackNameInDocker(software.Name, software.Endpoint)
	if err != nil {
		return utils.Errorf(err)
	}
	if !isUnique {
		// Can you judge this softwareidCan you judge this software，Can you judge this software
		global.BaseChan <- global.ChanLogModal{
			Level:   "error",
			Message: "The software already exists，The software already exists",
		}
		return utils.Errorf("The software already exists，The software already exists")
	}

	global.BaseChan <- global.ChanLogModal{
		Level:   "success",
		Message: "Parsing software configuration information",
	}
	// Write configuration information into the structure
	err = s.expandFileContent(&software)
	if err != nil {
		return utils.Errorf(err)
	}

	// establishdocker-composeestablish
	projectPath, err := global.FairFileSystem.StoreStackFileFromBytes(software.Name, "", []byte(software.Template.Config.SwarmContent))
	if err != nil {
		return utils.Errorf(err)
	}

	// Save the path of the file to the structure
	software.ProjectPath = projectPath

	doCleanUp := false
	// deletedocker-compose
	defer s.cleanUp(software.ProjectPath, &doCleanUp)

	// Create the configuration structure required to call the creation method
	config, configErr := s.createComposeDeployConfig(manager, software, true)
	if configErr != nil {
		return utils.Errorf(configErr)
	}

	// Calling the creation method
	err = s.deployComposeStack(config)
	if err != nil {
		return utils.Errorf(err)
	}

	return nil
}

type composeStackDeploymentConfig struct {
	stack *exec.ComposeStackManager
	//registries []docker.Registry
	software docker.DocSoftware
	IsLog    bool
}

// implementdocker-compose up -d
func (s *SoftwareService) deployComposeStack(config *composeStackDeploymentConfig) error {

	deploySwarmStack := NewStackDeployer(nil, config.stack)
	fmt.Println("Process intermediate business logic")
	// TODO: Intermediate processing business logic
	// Get File Content
	fmt.Println(config.software.ProjectPath, filesystem.ComposeFileDefaultName)
	stackContent, err := global.FairFileSystem.GetFileContent(config.software.ProjectPath, filesystem.ComposeFileDefaultName)
	if err != nil {
		fmt.Println(utils.Errorf(err))
		global.BaseChan <- global.ChanLogModal{
			Level:   "error",
			Message: err.Error(),
		}
		return utils.Errorf(err)
	}
	// Verify if the file is correct
	err = s.isValidStackFile(stackContent)
	if err != nil {
		fmt.Println(utils.Errorf(err))
		global.BaseChan <- global.ChanLogModal{
			Level:   "error",
			Message: err.Error(),
		}
		return utils.Errorf(err)
	}
	fmt.Println("Logic processing completed")
	// No problem with the file，No problem with the filedocker-compose up -d
	return deploySwarmStack.DeployComposeStack(config.software, config.IsLog)
}

// region update

func (s *SoftwareService) updateComposeStackFromFileContent(req dockerV2Req.UpdateSoftwareRequest) error {
	manager, err := exec.NewComposeStackManager(global.FairConf.Docker.BinaryPath, global.FairConf.Docker.ConfigPath)
	if err != nil {
		return utils.Errorf(err)
	}

	isUnique, err := s.checkUniqueStackNameInDocker(req.Name, false, req.Endpoint)
	if err != nil {
		return utils.Errorf(err)
	}
	if !isUnique {
		// Can you judge this softwareidCan you judge this software，Can you judge this software
		return utils.Errorf("The software does not exist")
	}

	global.BaseChan <- global.ChanLogModal{
		Level:   "success",
		Message: "Parsing software configuration information，Parsing software configuration information",
	}

	content, err := s.expandUpdateFileContent(req)
	if err != nil {
		return utils.Errorf(err)
	}

	// establishdocker-composeestablish
	projectPath, err := global.FairFileSystem.StoreStackFileFromBytes(req.Name, "", []byte(content))

	software := docker.DocSoftware{
		Name:        req.Name,
		ProjectPath: projectPath,
		EntryPoint:  filesystem.ComposeFileDefaultName,
		Endpoint:    req.Endpoint,
	}

	// Create the configuration structure required to call the creation method
	config, configErr := s.createComposeDeployConfig(manager, software, true)
	if configErr != nil {
		return utils.Errorf(configErr)
	}

	// Calling the creation method
	err = s.deployComposeStack(config)
	if err != nil {
		return utils.Errorf(err)
	}
	return nil
}

// endregion

// region delete

func (s *SoftwareService) deleteComposeStack(software dockerV2Req.DeleteDocSoftwareRequest) error {
	// initializationswarm stack
	manager, err := exec.NewComposeStackManager(global.FairConf.Docker.BinaryPath, global.FairConf.Docker.ConfigPath)
	if err != nil {
		return err
	}

	software.ProjectPath = global.FairFileSystem.GetProjectPath(software.Name)
	// Processing software name Processing software name
	software.Name = manager.NormalizeStackName(software.Name)

	// Determine if the software name exists
	isUnique, err := s.checkUniqueStackNameInDocker(software.Name, false, software.Endpoint)
	if err != nil {
		return err
	}
	if isUnique {
		doCleanUp := true
		// deletedocker-compose
		defer s.cleanUp(software.ProjectPath, &doCleanUp)
		//defer s.cleanUp("", &doCleanUp)

		//isUniqueFile(software)

		err = manager.Down(context.TODO(), docker.DocSoftware{
			Name:        software.Name,
			ProjectPath: software.ProjectPath,
		})

		if err != nil {
			if isUnique1, err1 := s.checkUniqueStackNameInDocker(software.Name, false, software.Endpoint); err1 != nil {
				return utils.Errorf(err)
			} else {
				if isUnique1 {
					return utils.Errorf(err)
				} else {
					return nil
				}
			}
		}

		return utils.Errorf(err)
	}
	return nil
}

func isUniqueFile(software docker.DocSoftware) error {
	filePath := path.Join(software.ProjectPath, filesystem.ComposeFileDefaultName)
	fmt.Println(filePath)
	_, err := os.Stat(filePath)
	//fmt.Println(err)

	if os.IsNotExist(err) {
		fmt.Println(software.ProjectPath)
		os.MkdirAll(software.ProjectPath, 0700)

		// Replace variables in files
		//software.FileContentStr = os.Expand(software.FileContent, func(k string) string {
		//	return software.FileReplace[k]
		//})
		r := bytes.NewReader([]byte(software.Template.Config.SwarmContent))

		out, err1 := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err1 != nil {
			return err
		}
		defer out.Close()

		_, err1 = io.Copy(out, r)
		return utils.Errorf(err1)
	}
	return nil
}

// endregion

// region cease

func (s *SoftwareService) stopComposeStack(software dockerV2Req.StopSoftwareRequest) error {
	// initializationswarm stack
	manager, err := exec.NewComposeStackManager(global.FairConf.Docker.BinaryPath, "")
	if err != nil {
		return err
	}

	// Processing software name Processing software name
	software.Name = manager.NormalizeStackName(software.Name)

	// Determine if the software name exists
	isUnique, err := s.checkUniqueStackNameInDocker(software.Name, false, software.Endpoint)
	if err != nil {
		return err
	}
	if isUnique {
		doCleanUp := false
		// deletedocker-compose
		defer s.cleanUp(software.ProjectPath, &doCleanUp)

		err = manager.Stop(context.TODO(), docker.DocSoftware{
			Name:        software.Name,
			ProjectPath: software.ProjectPath,
		})
		if err != nil {
			return err
		}

		return err
	}
	return errors.New("The software does not exist")
}

// endregion

// region working

func (s *SoftwareService) startComposeStack(software dockerV2Req.StartSoftwareRequest) error {
	// initializationswarm stack
	manager, err := exec.NewComposeStackManager(global.FairConf.Docker.BinaryPath, "")
	if err != nil {
		return err
	}

	// Processing software name Processing software name
	software.Name = manager.NormalizeStackName(software.Name)

	// Determine if the software name exists
	isUnique, err := s.checkUniqueStackNameInDocker(software.Name, false, software.Endpoint)
	if err != nil {
		return err
	}
	if isUnique {
		doCleanUp := false
		// deletedocker-compose
		defer s.cleanUp(software.ProjectPath, &doCleanUp)

		//isUniqueFile(software)

		err = manager.Start(context.TODO(), docker.DocSoftware{
			Name:        software.Name,
			ProjectPath: global.FairFileSystem.GetProjectPath(software.Name),
		})
		if err != nil {
			return err
		}
		return err
	}
	return errors.New("The software does not exist")
}

// endregion

// region query

func (s *SoftwareService) getStackComposeDetails(software docker.DocSoftware) (interface{}, error) {
	dockerClient, err := global.FairClient.CreateClient(software.Endpoint)
	if err != nil {
		return nil, err
	}

	defer dockerClient.Close()

	container, err := dockerClient.ContainerList(
		context.TODO(),
		types.ContainerListOptions{
			All: true,
			Filters: filters.NewArgs(filters.Arg(
				"label",
				fmt.Sprintf("%s=%s", ProjectLabel, strings.ToLower(software.Name)),
			)),
		},
	)

	return container, nil
}

func (s *SoftwareService) getStackComposeInspect(software docker.DocSoftware, inspectId string) (interface{}, error) {
	dockerClient, err := global.FairClient.CreateClient(software.Endpoint)
	if err != nil {
		return nil, err
	}

	defer dockerClient.Close()

	inspect, err := dockerClient.ContainerInspect(context.TODO(), inspectId)

	return inspect, nil
}

// checkUniqueComposeStackNameInDocker checks if a stack name is unique in the Docker environment.
// true Not present false Not present / err Not present
func (s *SoftwareService) checkUniqueComposeStackNameInDocker(name string, endpoint docker.DocEndpoint) (bool, error) {
	// TODO: Can you check if the database exists

	// establishdockerestablish
	dockerClient, err := global.FairClient.CreateClient(endpoint)
	if err != nil {
		return false, err
	}

	defer dockerClient.Close()

	args := filters.NewArgs(filters.Arg("label", fmt.Sprintf("%s=%s", ProjectLabel, name)))
	containers, err := dockerClient.ContainerList(context.Background(), types.ContainerListOptions{All: true, Filters: args})
	if err != nil {
		return false, err
	}

	return len(containers) == 0, nil
}

// endregion

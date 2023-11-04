// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-10 18:36
package docker

import (
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/model/v2/common/response"
	"cnic/fairman-gin/model/v2/docker"
	docRequst "cnic/fairman-gin/model/v2/docker/request"
	"cnic/fairman-gin/utils"
	"cnic/fairman-gin/utils/exec"
	"cnic/fairman-gin/utils/filesystem"
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"

	"github.com/docker/cli/cli/compose/loader"
	cliTypes "github.com/docker/cli/cli/compose/types"
	"github.com/pkg/errors"
)

// region establish

func (s *SoftwareService) createSwarmStackFromFileContent(software docker.DocSoftware) error {
	// initializationswarm stack
	manager, err := exec.NewSwarmStackManager("cmd", "data")
	if err != nil {
		return utils.Errorf(err)
	}
	// Processing software name Processing software name
	software.Name = manager.NormalizeStackName(software.Name)

	// Determine if the software name exists
	isUnique, err := s.checkUniqueStackNameInDocker(software.Name, true, software.Endpoint)
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

	//// Replace variables in files
	//software.FileContentStr = os.Expand(software.FileContent, func(k string) string {
	//	return software.FileReplace[k]
	//})

	global.BaseChan <- global.ChanLogModal{
		Level:   "success",
		Message: "Scanning port",
	}
	err = s.expandFileContent(&software)
	if err != nil {
		// There is an issue with the configuration parameters，There is an issue with the configuration parameters
		global.BaseChan <- global.ChanLogModal{
			Level:   "error-back",
			Message: "acc_hostoracc_portor",
		}
		return utils.Errorf(err)
	}

	//fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	//fmt.Println(software.Template.Config.SwarmContent)
	//fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

	// establishdocker-compose
	projectPath, err := global.FairFileSystem.StoreStackFileFromBytes(software.Name, "", []byte(software.Template.Config.SwarmContent))
	if err != nil {
		return utils.Errorf(err)
	}
	software.ProjectPath = projectPath

	doCleanUp := false
	// deletedocker-compose
	defer s.cleanUp(software.ProjectPath, &doCleanUp)
	fmt.Println(projectPath)
	config, configErr := s.createSwarmDeployConfig(manager, software)
	if configErr != nil {
		return utils.Errorf(configErr)
	}
	err = s.deploySwarmStack(config)
	if err != nil {
		return utils.Errorf(err)
	}
	return nil
}

func (s *SoftwareService) deploySwarmStack(config *swarmStackDeploymentConfig) error {

	deploySwarmStack := NewStackDeployer(config.stack, nil)

	// TODO: Intermediate processing business logic
	// Verify if the file is correct
	stackContent, err := global.FairFileSystem.GetFileContent(config.software.ProjectPath, filesystem.ComposeFileDefaultName)
	if err != nil {
		return utils.Errorf(err)
	}
	err = s.isValidStackFile(stackContent)
	if err != nil {
		return utils.Errorf(err)
	}

	return deploySwarmStack.DeploySwarmStack(config.software, true)
}

type swarmStackDeploymentConfig struct {
	stack *exec.SwarmStackManager
	//registries []docker.Registry
	software docker.DocSoftware
}

func (s *SoftwareService) createSwarmDeployConfig(stack *exec.SwarmStackManager, software docker.DocSoftware) (*swarmStackDeploymentConfig, error) {
	config := &swarmStackDeploymentConfig{
		stack:    stack,
		software: software,
	}
	return config, nil
}

func (s *SoftwareService) isValidStackFile(stackFileContent []byte) error {
	composeConfigYAML, err := loader.ParseYAML(stackFileContent)
	if err != nil {
		return utils.Errorf(err)
	}

	composeConfigFile := cliTypes.ConfigFile{
		Config: composeConfigYAML,
	}

	composeConfigDetails := cliTypes.ConfigDetails{
		ConfigFiles: []cliTypes.ConfigFile{composeConfigFile},
		Environment: map[string]string{},
	}

	_, err = loader.Load(composeConfigDetails, func(options *loader.Options) {
		options.SkipValidation = true
		options.SkipInterpolation = true
	})
	if err != nil {
		return utils.Errorf(err)
	}

	return nil
}

// endregion

// region modify

func (s *SoftwareService) updateSwarmStackFromFileContent(software docker.DocSoftware) error {
	// initializationswarm stack
	manager, err := exec.NewSwarmStackManager("cmd", "data")
	if err != nil {
		return utils.Errorf(err)
	}

	// Processing software name Processing software name
	software.Name = manager.NormalizeStackName(software.Name)

	fmt.Println(software, software.Name)

	// Determine if the software name exists
	isUnique, err := s.checkUniqueStackNameInDocker(software.Name, true, software.Endpoint)
	if err != nil {
		return utils.Errorf(err)
	}
	if isUnique {
		// Can you judge this softwareidCan you judge this software，Can you judge this software
		return utils.Errorf("The software does not exist")
	}

	//software.EntryPoint = "docker-compose.yaml"

	// Replace variables in files
	//software.FileContent = os.Expand(software.FileContent, func(k string) string {
	//	return software.FileReplace[k]
	//})

	err = s.expandFileContent(&software)
	if err != nil {
		return utils.Errorf(err)
	}

	// establishdocker-compose
	projectPath, err := global.FairFileSystem.StoreStackFileFromBytes(software.Name, filesystem.ComposeFileDefaultName, []byte(software.Template.Config.SwarmContent))
	if err != nil {
		return utils.Errorf(err)
	}

	software.ProjectPath = projectPath

	doCleanUp := false
	// deletedocker-compose
	defer s.cleanUp(software.ProjectPath, &doCleanUp)

	config, configErr := s.createSwarmDeployConfig(manager, software)
	if configErr != nil {
		return utils.Errorf(configErr)
	}

	err = s.deploySwarmStack(config)
	if err != nil {
		return utils.Errorf(err)
	}
	//
	//var resp response.Response
	//err = utils.Post(fmt.Sprintf(AddSoftware, global.FairConf.Service.Url), &software, "", &resp)
	//if err != nil {
	//	return utils.Errorf(err)
	//}

	return nil
}

// endregion

// region delete

func (s *SoftwareService) deleteSwarmStack(software docRequst.DeleteDocSoftwareRequest) error {
	// initializationswarm stack
	manager, err := exec.NewSwarmStackManager("cmd", "")
	if err != nil {
		return err
	}

	// Processing software name Processing software name Processing software name
	software.Name = manager.NormalizeStackName(software.Name)

	// Determine if the software name exists
	isUnique, err := s.checkUniqueStackNameInDocker(software.Name, true, software.Endpoint)
	if err != nil {
		return err
	}
	if !isUnique {
		err = manager.Remove(software.Name, software.Endpoint)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("The software does not exist")
}

// endregion

// region query

func (s *SoftwareService) getStackSwarmDetails(software docker.DocSoftware) (interface{}, error) {
	var details []response.SoftwareDetailsResp
	dockerClient, err := global.FairClient.CreateClient(software.Endpoint)
	if err != nil {
		return nil, err
	}

	defer dockerClient.Close()

	services, err := dockerClient.ServiceList(
		context.TODO(),
		types.ServiceListOptions{Filters: filters.NewArgs(filters.Arg(
			"label",
			fmt.Sprintf("com.docker.stack.namespace=%s", software.Name),
		))},
	)

	for _, service := range services {
		taskList, err := dockerClient.TaskList(
			context.TODO(),
			types.TaskListOptions{Filters: filters.NewArgs(
				filters.Arg(
					"label",
					fmt.Sprintf("com.docker.stack.namespace=%s", software.Name),
				),
				filters.Arg(
					"service",
					fmt.Sprintf(service.Spec.Name),
				),
			)},
		)
		if err != nil {
			utils.Errorf(err)
			break
		}

		details = append(details, response.SoftwareDetailsResp{Service: service, Children: taskList})
	}

	return details, nil
}

func (s *SoftwareService) getStackSwarmInspect(software docker.DocSoftware, inspectId string) (interface{}, error) {
	dockerClient, err := global.FairClient.CreateClient(software.Endpoint)
	if err != nil {
		return nil, err
	}

	defer dockerClient.Close()

	inspect, _, err := dockerClient.ServiceInspectWithRaw(context.TODO(), inspectId, types.ServiceInspectOptions{InsertDefaults: true})

	return inspect, err
}

// endregion

// region start-up

func (s *SoftwareService) startSwarmStack(software docker.DocSoftware) error {
	// initializationswarm stack
	manager, err := exec.NewSwarmStackManager(global.FairConf.Docker.BinaryPath, "")
	if err != nil {
		return err
	}

	// Processing software name Processing software name
	software.Name = manager.NormalizeStackName(software.Name)

	// Determine if the software name exists
	isUnique, err := s.checkUniqueStackNameInDocker(software.Name, true, software.Endpoint)
	if err != nil {
		return err
	}
	if isUnique {
		err = manager.Deploy(software, true)
		if err != nil {
			return err
		}

		return nil
	}
	return errors.New("Software started")
}

// endregion

// region cease

func (s *SoftwareService) stopSwarmStack(software docker.DocSoftware) error {
	// initializationswarm stack
	manager, err := exec.NewSwarmStackManager(global.FairConf.Docker.BinaryPath, "")
	if err != nil {
		return err
	}

	// Processing software name Processing software name
	software.Name = manager.NormalizeStackName(software.Name)

	// Determine if the software name exists
	isUnique, err := s.checkUniqueStackNameInDocker(software.Name, true, software.Endpoint)
	if err != nil {
		return err
	}
	if !isUnique {
		err = manager.Remove(software.Name, software.Endpoint)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("The software does not exist")
}

// endregion

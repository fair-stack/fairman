// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-10 16:35
package docker

import (
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/model/v2/docker"
	dockerV2Req "cnic/fairman-gin/model/v2/docker/request"
	"cnic/fairman-gin/utils"
	"cnic/fairman-gin/utils/exec"
	"cnic/fairman-gin/utils/filesystem"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/docker/docker/api/types/filters"

	dockerYml "github.com/lliuhuan/compose-yml"

	"github.com/docker/docker/api/types"

	"github.com/gin-gonic/gin"
)

type SoftwareService struct {
}

// CreateSoftware Install software
func (s *SoftwareService) CreateSoftware(software docker.DocSoftware, c *gin.Context) error {
	// Get and Setid
	software.UserId = utils.GetUserID(c)

	// Set File Name
	if software.EntryPoint == "" {
		software.EntryPoint = filesystem.ComposeFileDefaultName
	}

	switch global.StackType(software.Endpoint.Type) {
	case global.DockerStack:
		return s.createDockerStack(software)
	case global.DockerSwarmStack:
		return s.createComposeStackFromFileContent(software)
	case global.DockerComposeStack:
		return s.createComposeStackFromFileContent(software)
	case global.KubernetesStack:
		return s.createKubernetesStackFromFileContent(software)
	}
	return utils.Errorf("Please select a server")
}

func (s *SoftwareService) UpdateSoftware(req dockerV2Req.UpdateSoftwareRequest) error {
	if val, ok := docker.EndpointMap[req.EndpointId]; !ok {
		return utils.Errorf("endpointId: %s not found", req.EndpointId)
	} else {
		req.Endpoint = val
	}

	switch global.StackType(req.Endpoint.Type) {
	case global.DockerStack:
		fmt.Println("Docker")
		return s.updateDockerStack(req)
	case global.DockerSwarmStack:
		//req.Template.Config.SwarmConfig = req.Software.Template.Config.SwarmConfig
		//
		//req.Software.Template = req.Template
		return s.updateComposeStackFromFileContent(req)
	case global.DockerComposeStack:
		fmt.Println("Compose")
		return s.updateComposeStackFromFileContent(req)
	case global.KubernetesStack:
		fmt.Println("Kubernetes")
		return s.updateKubernetesStackFromFileContent(req)
	}
	return nil
}

// RemoveSoftware Remove software
func (s *SoftwareService) RemoveSoftware(software dockerV2Req.DeleteDocSoftwareRequest) error {
	if val, ok := docker.EndpointMap[software.EndpointId]; !ok {
		return utils.Errorf("endpointId: %s not found", software.EndpointId)
	} else {
		software.Endpoint = val
	}
	switch global.StackType(software.Endpoint.Type) {
	case global.DockerStack:
		fmt.Println("Docker")
		return s.deleteDockerStack(software)
		//return s.deleteComposeStack(software)
		//return s.deleteSwarmStack(software)
	case global.DockerSwarmStack:
		return s.deleteComposeStack(software)
		//return s.deleteSwarmStack(software)
	case global.DockerComposeStack:
		fmt.Println("compose")
		return s.deleteComposeStack(software)
	//return s.deleteSwarmStack(software)
	case global.KubernetesStack:
		return s.deleteKubernetesStack(software)
	}
	return nil
}

// StopSoftware Stop software
func (s *SoftwareService) StopSoftware(software dockerV2Req.StopSoftwareRequest) error {
	if val, ok := docker.EndpointMap[software.EndpointId]; !ok {
		return utils.Errorf("endpointId: %s not found", software.EndpointId)
	} else {
		software.Endpoint = val
	}
	switch global.StackType(software.Endpoint.Type) {
	case global.DockerStack:
		fmt.Println("Docker")
		return s.stopDockerStack(software)
		//return s.stopComposeStack(software)
	case global.DockerSwarmStack:
		return s.stopComposeStack(software)
	case global.DockerComposeStack:
		fmt.Println("compose")
		return s.stopComposeStack(software)
	case global.KubernetesStack:
		return s.stopKubernetesStack(software)
	}
	return nil
}

// StartSoftware execute
func (s *SoftwareService) StartSoftware(software dockerV2Req.StartSoftwareRequest) error {
	if val, ok := docker.EndpointMap[software.EndpointId]; !ok {
		return utils.Errorf("endpointId: %s not found", software.EndpointId)
	} else {
		software.Endpoint = val
	}
	switch global.StackType(software.Endpoint.Type) {
	case global.DockerStack:
		fmt.Println("Docker")
		return s.startDockerStack(software)
		//return s.startComposeStack(software)
	case global.DockerSwarmStack:
		return s.startComposeStack(software)
	case global.DockerComposeStack:
		fmt.Println("compose")
		return s.startComposeStack(software)
	case global.KubernetesStack:
		return s.startKubernetesStack(software)
	}
	return nil
}

func (s *SoftwareService) checkUniqueStackNameInDocker(name string, swarmMode bool, endpoint docker.DocEndpoint) (bool, error) {
	// TODO: Can you check if the database exists

	// TODO: Link to other servers，Link to other servers，Link to other servers
	dockerClient, err := global.FairClient.CreateClient(endpoint)
	if err != nil {
		return false, err
	}

	defer dockerClient.Close()

	if swarmMode {
		services, err := dockerClient.ServiceList(context.Background(), types.ServiceListOptions{})
		if err != nil {
			return false, err
		}
		for _, service := range services {
			//fmt.Printf("%s - %v \n", service.Spec.Labels["com.docker.stack.namespace"], service)
			serviceNS, ok := service.Spec.Labels["com.docker.stack.namespace"]
			if ok && serviceNS == name {
				return false, nil
			}
		}
	}

	args := filters.NewArgs(filters.Arg("label", fmt.Sprintf("%s=%s", ProjectLabel, name)))

	containers, err := dockerClient.ContainerList(context.Background(), types.ContainerListOptions{All: true, Filters: args})
	if err != nil {
		return false, err
	}

	if len(containers) > 0 {
		return true, nil
	}

	return false, nil
}

func (s *SoftwareService) checkStackNameStatus(name string, swarmMode bool, endpoint docker.DocEndpoint) (string, error) {
	var status string
	// TODO: Can you check if the database exists

	// TODO: Link to other servers，Link to other servers，Link to other servers
	dockerClient, err := global.FairClient.CreateClient(endpoint)
	if err != nil {
		return "error", err
	}

	defer dockerClient.Close()

	// TODO: Is it necessary to change the method used to verify whether it is running based on the environment in the later stage

	containers, err := dockerClient.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return "error", err
	}

	for _, container := range containers {
		containerComposeNS, ComposeOK := container.Labels[ProjectLabel]
		containerSwarmNS, SwarmOK := container.Labels["com.docker.stack.namespace"]

		if (ComposeOK && containerComposeNS == name) || (SwarmOK && containerSwarmNS == name) {
			status = "running"
		}
	}

	return status, nil
}

func (s *SoftwareService) cleanUp(projectPath string, doCleanUp *bool) error {
	if !*doCleanUp {
		return nil
	}

	err := global.FairFileSystem.RemoveDirectory(projectPath)
	if err != nil {
		//log.Printf("http error: Unable to cleanup stack creation (err=%s)\n", err)
		global.FairLog.Error("http error: Unable to cleanup stack creation (err=" + err.Error() + ")\n")
	}
	return nil
}

func (s *SoftwareService) createComposeDeployConfig(stack *exec.ComposeStackManager, software docker.DocSoftware, isLog bool) (*composeStackDeploymentConfig, error) {
	config := &composeStackDeploymentConfig{
		stack:    stack,
		software: software,
		IsLog:    isLog,
		//registries: resp.Data,
	}
	return config, nil
}

//func (s *SoftwareService) getSoftwareById(softwareId string) (software docker.DocSoftware) {
//	var resp response.Response
//	url := fmt.Sprintf(GetSoftwareById, global.FairConf.Service.Url, softwareId)
//	err := utils.Get(url, &resp)
//	if err != nil {
//		return
//	}
//
//	if resp.Code == 0 {
//		err1 := utils.DataJson(resp.Data, &software)
//		if err1 != nil {
//			return
//		}
//
//		// Set software installation location
//		var resp1 response.Response
//		url1 := fmt.Sprintf(GetEndpointById, global.FairConf.Service.Url, software.EndpointId)
//		err1 = utils.Get(url1, &resp1)
//		if resp1.Code == 0 {
//			err1 = utils.DataJson(resp1.Data, &software.Endpoint)
//		}
//		if err1 != nil || resp1.Code != 0 || resp1.Data == nil {
//			// TODO: Need to make some modifications Need to make some modificationsWindows
//			software.Endpoint.Socket = true
//			software.Endpoint.URL = "unix:///var/run/docker.sock"
//		}
//
//		return software
//	}
//
//	return
//}

func (s *SoftwareService) expandFileContent(software *docker.DocSoftware) error {
	accCount := 0
	accMap := make(map[string]string, 2)
	var yml dockerYml.Yml
	err := dockerYml.UnmarshalYaml(software.Template.Config.SwarmContent, &yml)
	if err != nil {
		return utils.Errorf("Unable to parse YAML file: %s", err)
	}

	for key, service := range yml.Services {
		service.Ports = software.Template.Config.SwarmConfig[key].Ports
		service.Volumes = software.Template.Config.SwarmConfig[key].Mounts
		service.Environment = software.Template.Config.SwarmConfig[key].Environment
		service.HostName = software.Template.Config.SwarmConfig[key].Hostname
		for _, environment := range service.Environment {
			if strings.ToLower(environment.Key) == "acc_host" {
				accCount++
				accMap["acc_host"] = environment.Val
			}
			if strings.ToLower(environment.Key) == "acc_port" {
				accCount++
				accMap["acc_port"] = environment.Val
			}
		}
	}

	software.Template.Config.SwarmContent, err = dockerYml.MarshalYaml(yml)

	if accCount == 2 {
		software.HomeHost = accMap["acc_host"]
		software.HomePort = accMap["acc_port"]
	} else {
		return utils.Errorf("acc_hostoracc_portor")
	}

	software.DockerYml = yml

	global.BaseChan <- global.ChanLogModal{
		Level:    "success",
		Message:  "Software template file parsing successful",
		Progress: 0,
		Steps:    len(yml.Services),
	}

	return nil
}

func (s *SoftwareService) expandOneFileContent(filePath string) (yml dockerYml.Yml, err error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return yml, utils.Errorf("Unable to read file: %s", err)
	}
	err = dockerYml.UnmarshalYaml(string(fileContent), &yml)
	if err != nil {
		return yml, utils.Errorf("Unable to parse YAML file: %s", err)
	}

	return yml, nil
}

// expandUpdateFileContent Generate the final configuration file based on the template file and configuration file
func (s *SoftwareService) expandUpdateFileContent(req dockerV2Req.UpdateSoftwareRequest) (string, error) {
	var yml dockerYml.Yml
	err := dockerYml.UnmarshalYaml(req.NewContent, &yml)
	if err != nil {
		return "", utils.Errorf("Unable to parse YAML file: %s", err)
	}

	for key, service := range yml.Services {
		yml.Services[key].Ports = req.OldConfig[key].Ports
		if len(req.OldConfig[key].Mounts) != len(req.NewConfig[key].Mounts) {
			if len(req.OldConfig[key].Mounts) > 0 {
				oldPath := strings.Split(req.OldConfig[key].Mounts[0].Source, "fairman")[0]

				for index, config := range req.NewConfig[key].Mounts {
					newPath := strings.Split(config.Source, "fairman")
					newPath[0] = oldPath
					req.NewConfig[key].Mounts[index].Source = strings.Join(newPath, "fairman")
				}
			}
			service.Volumes = req.NewConfig[key].Mounts
		} else {
			service.Volumes = req.OldConfig[key].Mounts
		}
		yml.Services[key].Environment = req.OldConfig[key].Environment
		yml.Services[key].HostName = req.NewConfig[key].Hostname
	}

	return dockerYml.MarshalYaml(yml)
}

// GetSoftwareStatus Obtain software status
func (s *SoftwareService) GetSoftwareStatus(softwareList []dockerV2Req.SoftwareStatusRequest) ([]dockerV2Req.SoftwareStatusRequest, error) {
	for index, software := range softwareList {
		if software.Endpoint.Type == 4 {
			kubeconfig := path.Join("data", "config", software.Endpoint.Name, "config")

			config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
			if err != nil {
				return nil, utils.Errorf(err)
			}

			clientset, err := kubernetes.NewForConfig(config)
			if err != nil {
				return nil, utils.Errorf(err)
			}

			pods, err := clientset.CoreV1().Pods(software.Name).List(context.Background(), metav1.ListOptions{
				//LabelSelector: fmt.Sprintf("app=%s", software.Name),
				LabelSelector: fmt.Sprintf("app=%s", software.Name+"baas"),
			})
			if len(pods.Items) != 0 {
				softwareList[index].Status = "running"
			} else {
				softwareList[index].Status = "stopped"
			}

			//k8sConfig, err := clientcmd.BuildConfigFromFlags("", path.Join("data", "config", software.Endpoint.Name, "config"))
			//if err != nil {
			//	continue
			//}
			//
			//k8sClient, err := kubernetes.NewForConfig(k8sConfig)
			//if err != nil {
			//	continue
			//}
			//
			//fmt.Println(k8sClient.AppsV1().Deployments(software.Name).Get(context.Background(), software.Name, metav1.GetOptions{}))
			//deployment, err := k8sClient.AppsV1().Deployments()
		} else {
			dockerClient, err := global.FairClient.CreateClient(software.Endpoint)
			if err != nil {
				return nil, utils.Errorf(err)
			}

			ctx := context.Background()
			filter := filters.NewArgs(filters.Arg("label", fmt.Sprintf("%s=%s", ProjectLabel, software.Name)))
			services, err := dockerClient.ContainerList(ctx, types.ContainerListOptions{
				Filters: filter,
			})
			if err != nil {
				return nil, utils.Errorf(err)
			}

			if len(services) != 0 {
				softwareList[index].Status = "running"
			} else {
				softwareList[index].Status = "stopped"
			}

			dockerClient.Close()
		}
	}

	//endpoints, err := global.FairStore.Endpoint().Endpoints()
	//fmt.Println(endpoints, err)
	//s.getT1()
	return softwareList, nil
}

// SaveEndpoint preserveendpoint
func (s *SoftwareService) SaveEndpoint(req dockerV2Req.SaveEndpointRequest) error {
	if _, ok := docker.EndpointMap[req.Endpoint.Id]; !ok {
		docker.EndpointMap[req.Endpoint.Id] = req.Endpoint
	}
	return nil
}

// UpdateSoftwareInspect Update software information
func (s *SoftwareService) UpdateSoftwareInspect(req dockerV2Req.UpdateSoftwareInspectRequest) error {
	if val, ok := docker.EndpointMap[req.EndpointId]; !ok {
		return utils.Errorf("endpointId: %s not found", req.EndpointId)
	} else {
		req.Endpoint = val
	}

	dockerClient, err := global.FairClient.CreateClient(req.Endpoint)
	if err != nil {
		return err
	}
	defer dockerClient.Close()

	ctx := context.Background()

	// 0. Parse the old software configuration file，Parse the old software configuration file
	// 0.1 Splicing address
	filePath := filesystem.JoinPaths("data", "compose", req.SoftwareName, "docker-compose.yml")
	yml, err := s.expandOneFileContent(filePath)
	if err != nil {
		return utils.Errorf(err)
	}

	// 0.2 replace content
	yml.Services[req.ContainerName].Ports = req.ContainerConfig.Ports
	yml.Services[req.ContainerName].Environment = req.ContainerConfig.Environment
	yml.Services[req.ContainerName].Volumes = req.ContainerConfig.Mounts
	yml.Services[req.ContainerName].HostName = req.ContainerConfig.Hostname

	// 1. Query software information
	args := filters.NewArgs(
		filters.Arg("label", fmt.Sprintf("%s=%s", ProjectLabel, strings.ToLower(req.SoftwareName))),
	)
	list, err := dockerClient.ContainerList(ctx, types.ContainerListOptions{All: true, Filters: args})

	if err != nil {
		return err
	}

	if len(list) == 0 {
		return utils.Errorf("container not found")
	}

	for _, containerInfo := range list {
		if containerInfo.Labels[ContainerLabel] != strings.ToLower(req.ContainerName) ||
			containerInfo.Labels[ServiceLabel] != strings.ToLower(req.ContainerName) {
			fmt.Println("#############################################")
			fmt.Println(containerInfo)
			fmt.Println("#############################################")
			continue
		}

		// 2. Stopping the software
		dockerClient.ContainerStop(ctx, containerInfo.ID, nil)
		// 3. Restart a new configured software
		var config container.Config
		var hostConfig container.HostConfig
		if err := loadConfigInfo(*yml.Services[req.ContainerName], &config, &hostConfig); err != nil {
			// TODO: Rollback software
			//removeAll(dockerClient, software.Name)
			return utils.Errorf(fmt.Sprintf("%s Failed to generate configuration information", req.ContainerName))
		}
		config.Labels = map[string]string{
			"com.docker.container.namespace": req.SoftwareName,
			ProjectLabel:                     strings.ToLower(req.SoftwareName),
			"com.docker.stack.namespace":     strings.ToLower(req.SoftwareName),
			ContainerLabel:                   req.ContainerName,
		}

		// TODO: Have you compared adding all containers to thefairmanHave you compared adding all containers to the
		// docker run
		networkName := strings.ToLower(fmt.Sprintf("fair_%s", req.SoftwareName))
		containerCreate, err := dockerClient.ContainerCreate(
			context.Background(),
			&config,
			&hostConfig,
			&network.NetworkingConfig{
				EndpointsConfig: map[string]*network.EndpointSettings{
					networkName: {
						Aliases: []string{req.ContainerName},
						//NetworkID: create.ID,
					},
				},
			},
			nil,
			fmt.Sprintf("%s_%s_tmp", req.SoftwareName, req.ContainerName),
		)

		if err != nil {
			// TODO: Rollback software
			dockerClient.ContainerRemove(ctx, containerInfo.ID, types.ContainerRemoveOptions{Force: true})
			return utils.Errorf(err)
		}

		// Created successfully，Created successfully
		if err := dockerClient.ContainerStart(context.TODO(), containerCreate.ID, types.ContainerStartOptions{}); err != nil {
			dockerClient.ContainerRemove(ctx, containerInfo.ID, types.ContainerRemoveOptions{Force: true})
			dockerClient.ContainerStart(context.Background(), containerInfo.ID, types.ContainerStartOptions{})
			return utils.Errorf(err, "create docker container success, but start container error")
		}

		// 4. If there are no issues after startup, remove the old software
		dockerClient.ContainerRemove(ctx, containerInfo.ID, types.ContainerRemoveOptions{Force: true})

		// 5. Rename the new software to the old software name
		dockerClient.ContainerRename(ctx, containerCreate.ID, fmt.Sprintf("%s_%s", req.SoftwareName, req.ContainerName))

		// 6. Restart all components once
		args = filters.NewArgs(
			//filters.Arg("label", fmt.Sprintf("com.docker.container.namespace=%s", req.SoftwareName)),
			filters.Arg("label", fmt.Sprintf("%s=%s", ProjectLabel, strings.ToLower(req.SoftwareName))),
		)
		list, err = dockerClient.ContainerList(ctx, types.ContainerListOptions{All: true, Filters: args})

		// Handle the startup sequence
		instList := make([]string, 0)
		for _, service := range yml.Services {
			// take outcontainerId
			containerId := ""
			for _, t := range list {
				if strings.Contains(t.Image, service.Image.Name) {
					containerId = t.ID
					break
				}
			}
			// Check the currentserviceNameCheck the current
			if arrayIndexOf(containerId, instList) == -1 {
				instList = append(instList, containerId)
			}
			for _, depend := range service.DependsOn {
				if arrayIndexOf(depend, instList) == -1 {
					sub := arrayIndexOf(containerId, instList)
					temp := append([]string{}, instList[sub:]...)
					instList = append(instList[:sub], depend)
					instList = append(instList, temp...)
				}
			}
		}

		for _, id := range instList {
			dockerClient.ContainerStop(ctx, id, nil)
		}

		for _, id := range instList {
			dockerClient.ContainerStart(ctx, id, types.ContainerStartOptions{})
		}
	}

	return nil
}

// UpdateSoftwareInspect1 Update software information
func (s *SoftwareService) UpdateSoftwareInspect1(req dockerV2Req.UpdateSoftwareInspectRequest1) error {
	if val, ok := docker.EndpointMap[req.EndpointId]; !ok {
		return utils.Errorf("endpointId: %s not found", req.EndpointId)
	} else {
		req.Endpoint = val
	}

	dockerClient, err := global.FairClient.CreateClient(req.Endpoint)
	if err != nil {
		return err
	}
	defer dockerClient.Close()

	ctx := context.Background()

	// 0. Parse the old software configuration file，Parse the old software configuration file
	// 0.1 Splicing address
	filePath := filesystem.JoinPaths("data", "compose", req.SoftwareName, "docker-compose.yml")
	yml, err := s.expandOneFileContent(filePath)
	if err != nil {
		return utils.Errorf(err)
	}

	// Do you want to modify it
	updList := make([]string, 0)

	for serviceKey, service := range yml.Services {
		oldPort, _ := json.Marshal(service.Ports)
		oldEnv, _ := json.Marshal(service.Environment)
		oldMount, _ := json.Marshal(service.Volumes)
		oldHostName, _ := json.Marshal(service.HostName)

		newPort, _ := json.Marshal(req.ContainerConfig[serviceKey].Ports)
		newEnv, _ := json.Marshal(req.ContainerConfig[serviceKey].Environment)
		newMount, _ := json.Marshal(req.ContainerConfig[serviceKey].Mounts)
		newHostName, _ := json.Marshal(req.ContainerConfig[serviceKey].Hostname)

		fmt.Println(string(oldPort), string(newPort))
		fmt.Println(string(oldEnv), string(newEnv))
		fmt.Println(string(oldMount), string(newMount))
		fmt.Println(string(oldHostName), string(newHostName))

		if string(oldPort) != string(newPort) || string(oldEnv) != string(newEnv) || string(oldMount) != string(newMount) || string(oldHostName) != string(newHostName) {
			// Normal execution of modifications
			yml.Services[serviceKey].Ports = req.ContainerConfig[serviceKey].Ports
			yml.Services[serviceKey].Environment = req.ContainerConfig[serviceKey].Environment
			yml.Services[serviceKey].Volumes = req.ContainerConfig[serviceKey].Mounts
			yml.Services[serviceKey].HostName = req.ContainerConfig[serviceKey].Hostname

			updList = append(updList, serviceKey)
		}
	}

	if len(updList) == 0 {
		return errors.New("The configuration information has not changed")
	}

	// 1. Query software information
	args := filters.NewArgs(
		filters.Arg("label", fmt.Sprintf("%s=%s", ProjectLabel, strings.ToLower(req.SoftwareName))),
	)
	list, err := dockerClient.ContainerList(ctx, types.ContainerListOptions{All: true, Filters: args})

	if err != nil {
		return err
	}

	if len(list) == 0 {
		return utils.Errorf("container not found")
	}

	for _, updName := range updList {
		containerName := fmt.Sprintf("%s_%s", strings.ToLower(req.SoftwareName), strings.ToLower(updName))
		var inspect types.ContainerJSON
		if containerInspect, err := dockerClient.ContainerInspect(ctx, containerName); err == nil {
			inspect = containerInspect
		}

		if containerInspect, err := dockerClient.ContainerInspect(ctx, fmt.Sprintf("%s_1", containerName)); err == nil {
			inspect = containerInspect
		}

		if inspect.ID == "" {
			return utils.Errorf("container not found")
		}

		// 2. Stopping the software
		dockerClient.ContainerStop(ctx, inspect.ID, nil)
		// 3. Restart a new configured software
		var config container.Config
		var hostConfig container.HostConfig
		if err := loadConfigInfo(*yml.Services[updName], &config, &hostConfig); err != nil {
			// TODO: Rollback software
			//removeAll(dockerClient, software.Name)
			return utils.Errorf(fmt.Sprintf("%s Failed to generate configuration information", inspect.Name))
		}
		config.Labels = map[string]string{
			"com.docker.container.namespace": req.SoftwareName,
			ProjectLabel:                     strings.ToLower(req.SoftwareName),
			"com.docker.stack.namespace":     strings.ToLower(req.SoftwareName),
			ContainerLabel:                   inspect.Name,
		}

		// TODO: Have you compared adding all containers to thefairmanHave you compared adding all containers to the
		// docker run
		networkName := strings.ToLower(fmt.Sprintf("fair_%s", req.SoftwareName))
		containerCreate, err := dockerClient.ContainerCreate(
			context.Background(),
			&config,
			&hostConfig,
			&network.NetworkingConfig{
				EndpointsConfig: map[string]*network.EndpointSettings{
					networkName: {
						Aliases: []string{inspect.Name},
						//NetworkID: create.ID,
					},
				},
			},
			nil,
			fmt.Sprintf("%s_tmp", inspect.Name),
		)

		if err != nil {
			// TODO: Rollback software
			dockerClient.ContainerRemove(ctx, inspect.ID, types.ContainerRemoveOptions{Force: true})
			return utils.Errorf(err)
		}

		// Created successfully，Created successfully
		if err := dockerClient.ContainerStart(context.TODO(), containerCreate.ID, types.ContainerStartOptions{}); err != nil {
			dockerClient.ContainerRemove(ctx, inspect.ID, types.ContainerRemoveOptions{Force: true})
			dockerClient.ContainerStart(context.Background(), inspect.ID, types.ContainerStartOptions{})
			return utils.Errorf(err, "create docker container success, but start container error")
		}

		// 4. If there are no issues after startup, remove the old software
		dockerClient.ContainerRemove(ctx, inspect.ID, types.ContainerRemoveOptions{Force: true})

		// 5. Rename the new software to the old software name
		dockerClient.ContainerRename(ctx, containerCreate.ID, inspect.Name)

		// 6. Restart all components once
		args = filters.NewArgs(
			//filters.Arg("label", fmt.Sprintf("com.docker.container.namespace=%s", req.SoftwareName)),
			filters.Arg("label", fmt.Sprintf("%s=%s", ProjectLabel, strings.ToLower(req.SoftwareName))),
		)
		list, err = dockerClient.ContainerList(ctx, types.ContainerListOptions{All: true, Filters: args})
	}

	// Handle the startup sequence
	instList := make([]string, 0)
	for _, service := range yml.Services {
		// take outcontainerId
		containerId := ""
		for _, t := range list {
			if strings.Contains(t.Image, service.Image.Name) {
				containerId = t.ID
				break
			}
		}
		// Check the currentserviceNameCheck the current
		if arrayIndexOf(containerId, instList) == -1 {
			instList = append(instList, containerId)
		}
		for _, depend := range service.DependsOn {
			if arrayIndexOf(depend, instList) == -1 {
				sub := arrayIndexOf(containerId, instList)
				temp := append([]string{}, instList[sub:]...)
				instList = append(instList[:sub], depend)
				instList = append(instList, temp...)
			}
		}
	}

	for _, id := range instList {
		dockerClient.ContainerStop(ctx, id, nil)
	}

	for _, id := range instList {
		dockerClient.ContainerStart(ctx, id, types.ContainerStartOptions{})
	}

	// establishdocker-composeestablish
	yaml, _ := dockerYml.MarshalYaml(yml)
	global.FairFileSystem.StoreStackFileFromBytes(req.SoftwareName, filesystem.ComposeFileDefaultName, []byte(yaml))

	return nil
}

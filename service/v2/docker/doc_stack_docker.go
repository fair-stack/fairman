// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-04-06 17:31
package docker

import (
	"bufio"
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/model/v2/docker"
	dockerV2Req "cnic/fairman-gin/model/v2/docker/request"
	"cnic/fairman-gin/utils"
	"cnic/fairman-gin/utils/common"
	"cnic/fairman-gin/utils/filesystem"
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/docker/docker/client"

	"github.com/docker/docker/api/types/network"
	dockerYml "github.com/lliuhuan/compose-yml"

	"github.com/docker/go-connections/nat"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
)

// createDockerStack establishdocker stack
func (s *SoftwareService) createDockerStack(software docker.DocSoftware) error {
	software.Name = strings.ToLower(software.Name)

	// 1. Check if the software exists
	isUnique, err := s.checkUniqueStackNameInDocker(software.Name, false, software.Endpoint)
	if err != nil {
		return utils.Errorf(err)
	}
	if isUnique {
		global.BaseChan <- global.ChanLogModal{
			Level:    "error",
			Message:  "The software already exists，The software already exists",
			Progress: 1,
		}
		return utils.Errorf("The software already exists，The software already exists")
	}

	global.BaseChan <- global.ChanLogModal{
		Level:    "success",
		Message:  "Parsing software configuration information",
		Progress: 0,
	}
	// 2. Parse the content，Parse the content
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
	// 3. Create software
	return s.createContainer(software)

	//// 1. Change the software name to lowercase
	//software.Name = normalizeStackName(software.Name)
	//// 2. Determine if the software has been created
	//if s.isStackExist(software.Name, software.Endpoint) {
	//	return utils.Errorf("The software already exists")
	//}

}

// deleteDockerStack deletedocker stack
func (s *SoftwareService) deleteDockerStack(req dockerV2Req.DeleteDocSoftwareRequest) error {
	// Check if it is running，Check if it is running
	dockerClient, err := global.FairClient.CreateClient(req.Endpoint)
	if err != nil {
		return utils.Errorf(err, "create docker client error")
	}

	defer dockerClient.Close()

	// 		config.Labels = map[string]string{
	//			"com.docker.container.namespace": software.Name,
	//			"com.docker.compose.project":     strings.ToLower(software.Name),
	//			"com.docker.stack.namespace":     strings.ToLower(software.Name),
	//		}

	args := filters.NewArgs(
		//filters.Arg("label", fmt.Sprintf("com.docker.container.namespace=%s", req.Name)),
		filters.Arg("label", fmt.Sprintf("%s=%s", ProjectLabel, strings.ToLower(req.Name))),
		//filters.Arg("label", fmt.Sprintf("com.docker.stack.namespace=%s", strings.ToLower(req.Name))),
	)
	list, err := dockerClient.ContainerList(
		context.Background(),
		types.ContainerListOptions{
			All:     true,
			Filters: args,
		},
	)

	if err != nil {
		return utils.Errorf(err)
	}

	for _, container := range list {
		dockerClient.ContainerStop(context.Background(), container.ID, nil)
		err := dockerClient.ContainerRemove(context.Background(), container.ID, types.ContainerRemoveOptions{Force: true})
		if err != nil {
			return utils.Errorf(err)
		}
	}

	return nil
}

// stopDockerStack ceasedocker stack
func (s *SoftwareService) stopDockerStack(req dockerV2Req.StopSoftwareRequest) error {
	// Check if it is running
	dockerClient, err := global.FairClient.CreateClient(req.Endpoint)
	if err != nil {
		return utils.Errorf(err, "create docker client error")
	}

	defer dockerClient.Close()

	args := filters.NewArgs(
		//filters.Arg("label", fmt.Sprintf("com.docker.container.namespace=%s", req.Name)),
		filters.Arg("label", fmt.Sprintf("%s=%s", ProjectLabel, strings.ToLower(req.Name))),
		//filters.Arg("label", fmt.Sprintf("com.docker.stack.namespace=%s", strings.ToLower(req.Name))),
	)
	list, err := dockerClient.ContainerList(
		context.Background(),
		types.ContainerListOptions{
			All:     true,
			Filters: args,
		},
	)

	if err != nil {
		return utils.Errorf(err)
	}

	for _, item := range list {
		err := dockerClient.ContainerStop(context.Background(), item.ID, nil)
		if err != nil {
			return utils.Errorf(err)
		}
	}

	return nil
}

// startDockerStack start-updocker stack
func (s *SoftwareService) startDockerStack(req dockerV2Req.StartSoftwareRequest) error {
	// Check if it is running
	dockerClient, err := global.FairClient.CreateClient(req.Endpoint)
	if err != nil {
		return utils.Errorf(err, "create docker client error")
	}

	defer dockerClient.Close()

	args := filters.NewArgs(
		//filters.Arg("label", fmt.Sprintf("com.docker.container.namespace=%s", req.Name)),
		filters.Arg("label", fmt.Sprintf("%s=%s", ProjectLabel, strings.ToLower(req.Name))),
		//filters.Arg("label", fmt.Sprintf("com.docker.stack.namespace=%s", strings.ToLower(req.Name))),
	)
	list, err := dockerClient.ContainerList(
		context.Background(),
		types.ContainerListOptions{
			All:     true,
			Filters: args,
		},
	)

	if err != nil {
		return utils.Errorf(err)
	}

	for _, item := range list {
		err := dockerClient.ContainerStart(context.Background(), item.ID, types.ContainerStartOptions{})
		if err != nil {
			return utils.Errorf(err)
		}
	}

	return nil
}

// updateDockerStack updatedocker stack
func (s *SoftwareService) updateDockerStack(req dockerV2Req.UpdateSoftwareRequest) error {
	// 1. Check if the software exists
	isUnique, err := s.checkUniqueStackNameInDocker(req.Name, false, req.Endpoint)
	if err != nil {
		return utils.Errorf(err)
	}
	if !isUnique {
		// TODO: The software does not exist，The software does not exist

		return nil
	}

	global.BaseChan <- global.ChanLogModal{
		Level:   "success",
		Message: "Parsing software configuration information，Parsing software configuration information",
	}

	software := docker.DocSoftware{
		Name:       req.Name,
		EntryPoint: filesystem.ComposeFileDefaultName,
		Endpoint:   req.Endpoint,
	}
	// Merge the original configuration file with the new configuration file
	content, err := s.expandUpdateFileContent(req)
	if err != nil {
		global.BaseChan <- global.ChanLogModal{
			Level:   "error",
			Message: "Merge configuration file failed",
		}
		return utils.Errorf(err)
	}

	// Write a new configuration file into the structure
	software.Template.Config.SwarmContent = content
	software.Template.Config.SwarmConfig = req.OldConfig

	// establishdocker-composeestablish
	projectPath, err := global.FairFileSystem.StoreStackFileFromBytes(req.Name, "", []byte(content))

	// 2. Parse the content，Parse the content
	err = s.expandFileContent(&software)
	if err != nil {
		return utils.Errorf(err)
	}

	// Save the path of the file to the structure
	software.ProjectPath = projectPath

	return s.updateContainer(software)
}

// Convert software names to lowercase
var stackNameNormalizeRegex = regexp.MustCompile("[^-_a-z0-9]+")

// normalizeStackName Convert software names to lowercase
func normalizeStackName(name string) string {
	return stackNameNormalizeRegex.ReplaceAllString(strings.ToLower(name), "")
}

// isStackExist Determine if this container already exists
func (s *SoftwareService) isStackExist(name string, endpoint docker.DocEndpoint) bool {
	dockerClient, err := global.FairClient.CreateClient(endpoint)
	if err != nil {
		fmt.Println(utils.Errorf(err, "create docker client error"))
		return false
	}

	defer dockerClient.Close()

	arg := filters.NewArgs(filters.Arg("name", name))
	opt := types.ContainerListOptions{
		All:     true,
		Filters: arg,
	}
	containerList, err := dockerClient.ContainerList(context.TODO(), opt)
	if err != nil {
		fmt.Println(utils.Errorf(err, "create docker client error"))
		return false
	}

	if len(containerList) > 0 {
		for _, containerInfo := range containerList {
			for _, containerName := range containerInfo.Names {
				if containerName == "/"+name {
					return true
				}
			}
		}
	}

	return false
}

func (s *SoftwareService) createContainer(software docker.DocSoftware) error {
	dockerClient, err := global.FairClient.CreateClient(software.Endpoint)
	if err != nil {
		return utils.Errorf(err, "create docker client error")
	}

	defer dockerClient.Close()

	global.BaseChan <- global.ChanLogModal{
		Level:   "success",
		Message: "Connected toDocker APi",
	}

	// Installation sequence list
	//InstallationSequenceList := make(map[string]*dockerYml.Service, 0)
	// Installation sequence
	instList := make([]string, 0)
	// base onDependsOnbase on
	for serviceName, service := range software.DockerYml.Services {
		// Check the currentserviceNameCheck the current
		if arrayIndexOf(serviceName, instList) == -1 {
			instList = append(instList, serviceName)
		}
		for _, depend := range service.DependsOn {
			if arrayIndexOf(depend, instList) == -1 {
				sub := arrayIndexOf(serviceName, instList)
				temp := append([]string{}, instList[sub:]...)
				instList = append(instList[:sub], depend)
				instList = append(instList, temp...)
			}
		}
	}

	// Determine if the network exists，Determine if the network exists
	networkName := strings.ToLower(fmt.Sprintf("fair_%s", software.Name))

	global.BaseChan <- global.ChanLogModal{
		Level:   "success",
		Message: fmt.Sprintf("Creating network [%s]", networkName),
	}

	_, err = createNetwork(networkName, dockerClient)
	if err != nil {
		global.BaseChan <- global.ChanLogModal{
			Level:   "error",
			Message: fmt.Sprintf("network[%s]network [%s]", networkName, err.Error()),
		}
		return utils.Errorf(err)
	}

	// TODO: This is for processing networks，This is for processing networksfairmanThis is for processing networks，This is for processing networks（This is for processing networks），This is for processing networks
	//netFilter := filters.NewArgs(filters.Arg("name", "fairman"))
	//networkList, _ := dockerClient.NetworkList(context.Background(), types.NetworkListOptions{Filters: netFilter})

	// Start processing
	for _, serverName := range instList {
		service := software.DockerYml.Services[serverName]

		global.BaseChan <- global.ChanLogModal{
			Level:   "success",
			Message: fmt.Sprintf("Processing [%s_%s]", software.Name, serverName),
		}

		// Pulling a tertiary mirror image
		for i := 0; i < 3; i++ {
			reader, err := dockerClient.ImagePull(context.Background(), fmt.Sprintf("%s:%s", service.Image.Name, service.Image.Tag), types.ImagePullOptions{})
			if err != nil {
				global.BaseChan <- global.ChanLogModal{
					Level:   "error",
					Message: fmt.Sprintf("%s:%sMirror pull failed；[%s:%s], error: %s", software.Name, serverName, service.Image.Name, service.Image.Tag, err.Error()),
				}
				break
			}

			scanner := bufio.NewScanner(reader)
			const maxCapacity = 512 * 1024
			buf := make([]byte, maxCapacity)
			scanner.Buffer(buf, maxCapacity)

			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}

			if err != nil {
				global.BaseChan <- global.ChanLogModal{
					Level:   "error",
					Message: fmt.Sprintf("Section[%d]Section Section[%s:%s]Section，Section，Section：(%s)", i, service.Image.Name, service.Image.Tag, err.Error()),
				}
				continue
			} else {
				global.BaseChan <- global.ChanLogModal{
					Level:   "success",
					Message: fmt.Sprintf("[%s:%s]Download completed，Download completed...", service.Image.Name, service.Image.Tag),
				}
				break
			}
		}

		global.BaseChan <- global.ChanLogModal{
			Level:   "success",
			Message: fmt.Sprintf("%s Generate Configuration Information", serverName),
		}

		var config container.Config
		var hostConfig container.HostConfig
		if err := loadConfigInfo(*service, &config, &hostConfig); err != nil {
			global.BaseChan <- global.ChanLogModal{
				Level:   "error",
				Message: fmt.Sprintf("%s Failed to generate configuration information [%s]", serverName, err.Error()),
			}

			global.BaseChan <- global.ChanLogModal{
				Level:   "error",
				Message: "Rolling back",
			}

			// TODO: Rollback software
			removeAll(dockerClient, software.Name)
			return utils.Errorf(fmt.Sprintf("%s Failed to generate configuration information", serverName))
		}
		config.Labels = map[string]string{
			"com.docker.container.namespace": software.Name,
			ProjectLabel:                     software.Name,
			"com.docker.stack.namespace":     software.Name,
			"com.docker.container.name":      serverName,
		}

		global.BaseChan <- global.ChanLogModal{
			Level:   "success",
			Message: fmt.Sprintf("Joining to create software [%s]", fmt.Sprintf("%s_%s", software.Name, serverName)),
		}

		// docker run
		containerCreate, err := dockerClient.ContainerCreate(
			context.Background(),
			&config,
			&hostConfig,
			&network.NetworkingConfig{
				EndpointsConfig: map[string]*network.EndpointSettings{
					networkName: {
						Aliases: []string{serverName},
						//NetworkID: create.ID,
					},
				},
			},
			nil,
			fmt.Sprintf("%s_%s", software.Name, serverName),
		)

		if err != nil {
			// TODO: Rollback software
			global.BaseChan <- global.ChanLogModal{
				Level:   "error",
				Message: fmt.Sprintf("%s Creation failed [%s]，Creation failed", serverName, err.Error()),
			}
			removeAll(dockerClient, software.Name)
			return utils.Errorf(err)
		}

		global.BaseChan <- global.ChanLogModal{
			Level:   "success",
			Message: fmt.Sprintf("%s Created successfully", serverName),
		}

		// Created successfully，Created successfully
		if err := dockerClient.ContainerStart(context.TODO(), containerCreate.ID, types.ContainerStartOptions{}); err != nil {
			// TODO: Delete container if startup fails
			//_ = operation.RemoveContainer(containerCreate.ID, types.ContainerRemoveOptions{Force: true})
			global.BaseChan <- global.ChanLogModal{
				Level:   "error",
				Message: fmt.Sprintf("%s Start failed [%s]，Start failed", serverName, err.Error()),
			}
			removeAll(dockerClient, software.Name)
			return utils.Errorf(err, "create docker container success, but start container error")
		}

		global.BaseChan <- global.ChanLogModal{
			Level:   "success",
			Message: fmt.Sprintf("%s Successfully started", serverName),
		}

		//for _, net := range networkList {
		//	global.BaseChan <- global.ChanLogModal{
		//		Level:   "success",
		//		Message: fmt.Sprintf("%s Join the unified scheduling network [%s]", serverName, net.Name),
		//	}
		//
		//	if err := dockerClient.NetworkConnect(context.Background(), net.ID, containerCreate.ID, nil); err != nil {
		//		fmt.Println(err)
		//		global.BaseChan <- global.ChanLogModal{
		//			Level:   "error",
		//			Message: fmt.Sprintf("%s Join the unified scheduling network [%s] Join the unified scheduling network", serverName, net.Name),
		//		}
		//		continue
		//	}
		//}
	}

	global.BaseChan <- global.ChanLogModal{
		Level:   "success",
		Message: fmt.Sprintf("software [%s] software", software.Name),
	}

	//imageName := software.Template.Config.DockerImage
	//
	//operation := exec.ContainerOperation{
	//	Endpoint: software.Endpoint,
	//}
	//// TODO: Link to other servers，Link to other servers，Link to other servers
	//
	//// 1. Processing environment variables
	//var envList []string
	//for _, env := range software.Template.Config.DockerConfig["default"].Environment {
	//	if env.Key == "" {
	//		continue
	//	}
	//	envList = append(envList, env.Key+"="+env.Val)
	//}
	//// 2. Process Port Mapping
	//portBindings := make(nat.PortMap)
	//exposedPorts := map[nat.Port]struct{}{}
	//for _, val := range software.Template.Config.DockerConfig["default"].Ports {
	//	if val.Target == "" {
	//		continue
	//	}
	//	portBinding := nat.PortBinding{HostIP: "", HostPort: val.Published}
	//	portBindings[nat.Port(val.Target+"/"+val.Protocol)] = []nat.PortBinding{portBinding}
	//	exposedPorts[nat.Port(val.Target+"/"+val.Protocol)] = struct{}{}
	//}
	//
	//containerConfig := container.Config{
	//	Image:        imageName,
	//	ExposedPorts: exposedPorts,
	//	Labels: map[string]string{
	//		"com.docker.container.namespace": software.Name,
	//	},
	//}
	//if software.Template.Config.DockerConfig["default"].Hostname != "" {
	//	containerConfig.Hostname = software.Template.Config.DockerConfig["default"].Hostname
	//}
	//if len(envList) > 0 {
	//	containerConfig.Env = envList
	//}
	//
	//// 3. Configuring Storage Volumes
	//var volumeList []string
	//var mountList []mount.Mount
	//mounts := software.Template.Config.DockerConfig["default"].Mounts
	//for key := range mounts {
	//	if mounts[key].Target == "" {
	//		continue
	//	}
	//	volumeList = append(volumeList, mounts[key].Source+":"+mounts[key].Target)
	//	mountList = append(mountList, mount.Mount{
	//		Type:     mount.Type(mounts[key].Type),
	//		Source:   mounts[key].Source,
	//		Target:   mounts[key].Target,
	//		ReadOnly: mounts[key].ReadOnly,
	//	})
	//}
	//
	//restartPolicy := container.RestartPolicy{Name: "always"}
	//
	//containerHostConfig := container.HostConfig{
	//	RestartPolicy: restartPolicy,
	//	PortBindings:  portBindings,
	//	Binds:         volumeList,
	//	Mounts:        mountList,
	//	Privileged:    true,
	//}
	//
	//containerNetworkConfig := network.NetworkingConfig{}
	//
	//// docker run
	//containerCreate, err := dockerClient.ContainerCreate(
	//	context.TODO(),
	//	&containerConfig,
	//	&containerHostConfig,
	//	&containerNetworkConfig,
	//	nil,
	//	software.Name,
	//)

	//// Delete container if startup fails
	//if err != nil {
	//	_ = operation.RemoveContainer(containerCreate.ID, types.ContainerRemoveOptions{Force: true})
	//	return utils.Errorf(err, "create docker container error")
	//}
	//
	//// Created successfully，Created successfully
	//startConfig := types.ContainerStartOptions{}
	//if err1 := dockerClient.ContainerStart(context.TODO(), containerCreate.ID, startConfig); err1 != nil {
	//	// Delete container if startup fails
	//	_ = operation.RemoveContainer(containerCreate.ID, types.ContainerRemoveOptions{Force: true})
	//	return utils.Errorf(err, "create docker container success, but start container error")
	//}

	return nil
}

func (s *SoftwareService) updateContainer(software docker.DocSoftware) error {
	dockerClient, err := global.FairClient.CreateClient(software.Endpoint)
	if err != nil {
		return utils.Errorf(err, "create docker client error")
	}

	defer dockerClient.Close()

	fmt.Println("Connected toDocker APi")

	global.BaseChan <- global.ChanLogModal{
		Level:   "success",
		Message: "Connected toDocker APi",
	}

	// Installation sequence list
	//InstallationSequenceList := make(map[string]*dockerYml.Service, 0)
	// Installation sequence
	instList := make([]string, 0)
	// base onDependsOnbase on
	for serviceName, service := range software.DockerYml.Services {
		// Check the currentserviceNameCheck the current
		if arrayIndexOf(serviceName, instList) == -1 {
			instList = append(instList, serviceName)
		}
		for _, depend := range service.DependsOn {
			if arrayIndexOf(depend, instList) == -1 {
				sub := arrayIndexOf(serviceName, instList)
				temp := append([]string{}, instList[sub:]...)
				instList = append(instList[:sub], depend)
				instList = append(instList, temp...)
			}
		}
	}

	type RemoveContainer struct {
		ContainerName string
		ContainerId   string
	}
	// Containers that need to be deleted after startup is completed
	removeContainer := make([]RemoveContainer, 0)

	// Determine if the network exists，Determine if the network exists
	networkName := strings.ToLower(fmt.Sprintf("fair_%s", software.Name))

	_, err = createNetwork(networkName, dockerClient)
	if err != nil {
		global.BaseChan <- global.ChanLogModal{
			Level:   "error",
			Message: fmt.Sprintf("network[%s]network [%s]", networkName, err.Error()),
		}
		return utils.Errorf(err)
	}

	// Start processing
	for _, serverName := range instList {
		isBreak := false
		service := software.DockerYml.Services[serverName]
		containerName := fmt.Sprintf("%s_%s", software.Name, serverName)
		fmt.Println(containerName)

		global.BaseChan <- global.ChanLogModal{
			Level:   "success",
			Message: fmt.Sprintf("Processing%s", containerName),
		}

		// Determine if the current image exists，Determine if the current image exists，Determine if the current image exists，Determine if the current image exists
		imgList, _ := dockerClient.ImageList(context.Background(), types.ImageListOptions{All: true})

		for _, imageSummary := range imgList {
			for _, tag := range imageSummary.RepoTags {
				if tag == service.Image.Name+":"+service.Image.Tag {
					// Delete the image，Delete the image，Delete the image
					// If deleted, proceed with execution，If deleted, proceed with execution
					_, err := dockerClient.ImageRemove(context.Background(), imageSummary.ID, types.ImageRemoveOptions{Force: true})
					if err != nil {
						fmt.Println(fmt.Sprintf("%s No changes at present", containerName))

						// Mainly for handlingv2.0.1Mainly for handling
						// If there are no changes，If there are no changesnetworkIf there are no changes，If there are no changesnetworkNameIf there are no changes，If there are no changesnetworkNameIf there are no changes
						list, err := dockerClient.ContainerList(context.Background(), types.ContainerListOptions{
							All:     true,
							Filters: filters.NewArgs(filters.KeyValuePair{Key: "ancestor", Value: tag}),
						})
						if err != nil {
							// Unable to find the container for this image，Unable to find the container for this image，Unable to find the container for this image
							break
						}

						networkInspect, err := dockerClient.NetworkInspect(context.Background(), networkName, types.NetworkInspectOptions{})
						if err != nil {
							// If you can't find it online, it means there's a problem
							fmt.Println("###################################################################")
							fmt.Println("Unable to find network，Unable to find network。。。。。")
							fmt.Println("###################################################################")
							break
						}

						// Add the detected container to thenetworkAdd the detected container to the
						for _, containerInfo := range list {
							err := dockerClient.NetworkConnect(context.Background(), networkInspect.ID, containerInfo.ID, &network.EndpointSettings{
								Aliases: []string{serverName},
							})
							if err != nil {
								fmt.Println(utils.Errorf(err, "Container failed to join the network"))
							}
						}

						global.BaseChan <- global.ChanLogModal{
							Level:   "success",
							Message: fmt.Sprintf("[%s] No changes at present", containerName),
						}
						isBreak = true
						break
					}
				}
			}
		}

		if isBreak {
			continue
		}

		// Pulling a tertiary mirror image
		for i := 0; i < 3; i++ {
			reader, err := dockerClient.ImagePull(context.Background(), fmt.Sprintf("%s:%s", service.Image.Name, service.Image.Tag), types.ImagePullOptions{})
			if err != nil {
				global.BaseChan <- global.ChanLogModal{
					Level:   "error",
					Message: fmt.Sprintf("%s:%sMirror pull failed；[%s:%s]", software.Name, serverName, service.Image.Name, service.Image.Tag),
				}
				break
			}

			scanner := bufio.NewScanner(reader)
			const maxCapacity = 512 * 1024
			buf := make([]byte, maxCapacity)
			scanner.Buffer(buf, maxCapacity)

			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}

			if err != nil {
				global.BaseChan <- global.ChanLogModal{
					Level:   "error",
					Message: fmt.Sprintf("Section[%d]Section Section[%s:%s]Section，Section，Section：(%s)", i, service.Image.Name, service.Image.Tag, err.Error()),
				}
				continue
			} else {
				global.BaseChan <- global.ChanLogModal{
					Level:   "success",
					Message: fmt.Sprintf("[%s:%s]Download completed，Download completed...", service.Image.Name, service.Image.Tag),
				}
				break
			}
		}

		global.BaseChan <- global.ChanLogModal{
			Level:   "success",
			Message: fmt.Sprintf("%s Generate Configuration Information", serverName),
		}

		var config container.Config
		var hostConfig container.HostConfig
		if err := loadConfigInfo(*service, &config, &hostConfig); err != nil {
			global.BaseChan <- global.ChanLogModal{
				Level:   "error",
				Message: fmt.Sprintf("%s Failed to generate configuration information [%s]", serverName, err.Error()),
			}

			global.BaseChan <- global.ChanLogModal{
				Level:   "error",
				Message: "Rolling back",
			}

			// TODO: Rollback software
			removeAll(dockerClient, software.Name)
			return utils.Errorf(fmt.Sprintf("%s Failed to generate configuration information", serverName))
		}

		fmt.Println(hostConfig)

		config.Labels = map[string]string{
			"com.docker.container.namespace": software.Name,
			ProjectLabel:                     software.Name,
			"com.docker.stack.namespace":     software.Name,
			"com.docker.container.name":      serverName,
		}

		global.BaseChan <- global.ChanLogModal{
			Level:   "success",
			Message: fmt.Sprintf("Creating software [%s]", fmt.Sprintf("%s_%s", software.Name, serverName)),
		}

		// docker run
		containerCreate, err := dockerClient.ContainerCreate(
			context.Background(),
			&config,
			&hostConfig,
			&network.NetworkingConfig{
				EndpointsConfig: map[string]*network.EndpointSettings{
					networkName: {
						Aliases: []string{serverName},
						//NetworkID: create.ID,
					},
				},
			},
			nil,
			fmt.Sprintf("%s_new", containerName),
		)

		if err != nil {
			// TODO: Rollback software
			global.BaseChan <- global.ChanLogModal{
				Level:   "error",
				Message: fmt.Sprintf("%s Creation failed [%s]，Creation failed", serverName, err.Error()),
			}
			removeAll(dockerClient, software.Name)
			return utils.Errorf(err)
		}

		removeContainer = append(removeContainer, RemoveContainer{
			ContainerName: containerName,
			ContainerId:   containerCreate.ID,
		})

		global.BaseChan <- global.ChanLogModal{
			Level:   "success",
			Message: fmt.Sprintf("%s Created successfully", serverName),
		}
	}
	fmt.Println("Containers that need to be deleted：", removeContainer)
	// Delete old containers，Delete old containers
	for _, con := range removeContainer {
		list, err := dockerClient.ContainerList(context.Background(), types.ContainerListOptions{All: true, Filters: filters.NewArgs(filters.Arg("name", con.ContainerName))})
		if err != nil {
			global.BaseChan <- global.ChanLogModal{
				Level:   "error",
				Message: fmt.Sprintf("%s Update failed [%s]", con.ContainerName, err.Error()),
			}
			return utils.Errorf(err.Error())
		}

		for _, con1 := range list {
			if strings.Index(con1.Names[0], "new") >= 0 {
				continue
			}
			err = dockerClient.ContainerRemove(context.Background(), con1.ID, types.ContainerRemoveOptions{Force: true})
			if err != nil {
				fmt.Println(utils.Errorf(err))
			}
			err = dockerClient.ContainerRename(context.Background(), con.ContainerId, con.ContainerName)
			if err != nil {
				fmt.Println(utils.Errorf(err))
			}
		}
	}

	// Restart all containers
	for _, serverName := range instList {
		list, err := dockerClient.ContainerList(context.Background(), types.ContainerListOptions{
			All:     true,
			Filters: filters.NewArgs(filters.Arg("name", fmt.Sprintf("%s_%s", software.Name, serverName))),
		})
		if err != nil {
			fmt.Println(utils.Errorf(err))
		}
		if len(list) > 0 {
			err = dockerClient.ContainerRestart(context.Background(), list[0].ID, nil)
			if err != nil {
				global.BaseChan <- global.ChanLogModal{
					Level:   "error",
					Message: fmt.Sprintf("%s Update failed [%s].", serverName, err.Error()),
				}
				return utils.Errorf(err)
			}
		}
	}

	global.BaseChan <- global.ChanLogModal{
		Level:   "success",
		Message: fmt.Sprintf("software [%s] software", software.Name),
	}
	return nil
}

func loadConfigInfo(req dockerYml.Service, config *container.Config, hostConf *container.HostConfig) error {
	if req.HostName != "" {
		config.Hostname = req.HostName
	}
	portMap, err := checkPortStats(req.Ports)
	if err != nil {
		return err
	}
	exposeds := make(nat.PortSet)
	for port := range portMap {
		exposeds[port] = struct{}{}
	}
	config.Image = fmt.Sprintf("%s:%s", req.Image.Name, req.Image.Tag)
	//config.Cmd = req.Cmd
	// 1. Processing environment variables
	config.Env = []string{}
	for _, env := range req.Environment {
		config.Env = append(config.Env, env.Key+"="+env.Val)
	}
	config.ExposedPorts = exposeds

	//hostConf.AutoRemove = req.AutoRemove // Automatic deletion
	//hostConf.CPUShares = req.CPUShares // limitCPU
	//hostConf.PublishAllPorts = req.PublishAllPorts // Automatically map one for all open ports，Automatically map one for all open ports
	hostConf.RestartPolicy = container.RestartPolicy{Name: "always"}
	//hostConf.RestartPolicy = container.RestartPolicy{Name: "on-failure"}
	//if req.RestartPolicy == "on-failure" {
	//hostConf.RestartPolicy.MaximumRetryCount = 5
	//}
	if len(req.Ports) != 0 {
		hostConf.PortBindings = portMap
	}
	hostConf.Binds = []string{}
	//hostConf.Mounts = []mount.Mount{}
	if len(req.Volumes) != 0 {
		config.Volumes = make(map[string]struct{})
		for _, volume := range req.Volumes {
			config.Volumes[volume.Target] = struct{}{}
			hostConf.Binds = append(hostConf.Binds, fmt.Sprintf("%s:%s:rw", volume.Source, volume.Target))
			//hostConf.Mounts = append(hostConf.Mounts, mount.Mount{
			//	Type:     mount.Type(volume.Type),
			//	Source:   volume.Source,
			//	Target:   volume.Target,
			//	ReadOnly: true,
			//})
		}
	}
	hostConf.Privileged = true
	return nil
}

func checkPortStats(ports []dockerYml.Port) (nat.PortMap, error) {
	portMap := make(nat.PortMap)
	if len(ports) == 0 {
		return portMap, nil
	}
	for _, port := range ports {
		if strings.Contains(port.Target, "-") {
			if !strings.Contains(port.Published, "-") {
				return portMap, utils.Errorf("Port inconsistency，Port inconsistency")
			}
			hostStart, _ := strconv.Atoi(strings.Split(port.Published, "-")[0])
			hostEnd, _ := strconv.Atoi(strings.Split(port.Published, "-")[1])
			containerStart, _ := strconv.Atoi(strings.Split(port.Target, "-")[0])
			containerEnd, _ := strconv.Atoi(strings.Split(port.Target, "-")[1])
			if (hostEnd-hostStart) <= 0 || (containerEnd-containerStart) <= 0 {
				return portMap, utils.Errorf("Inconsistent range ports，Inconsistent range ports")
			}
			if (containerEnd - containerStart) != (hostEnd - hostStart) {
				return portMap, utils.Errorf("Inconsistent range ports，Inconsistent range ports")
			}
			for i := 0; i <= hostEnd-hostStart; i++ {
				bindItem := nat.PortBinding{HostPort: strconv.Itoa(hostStart + i), HostIP: "0.0.0.0"}
				portMap[nat.Port(fmt.Sprintf("%d/%s", containerStart+i, port.Protocol))] = []nat.PortBinding{bindItem}
			}
			for i := hostStart; i <= hostEnd; i++ {
				if common.ScanPort(i, port.Target) {
					return portMap, utils.Errorf(fmt.Sprintf("%dPort is occupied", i))
				}
			}
		} else {
			portItem := 0
			if strings.Contains(port.Published, "-") {
				portItem, _ = strconv.Atoi(strings.Split(port.Published, "-")[0])
			} else {
				portItem, _ = strconv.Atoi(port.Published)
			}
			if common.ScanPort(portItem, port.Target) {
				return portMap, utils.Errorf(fmt.Sprintf("%dPort is occupied", portItem))
			}
			bindItem := nat.PortBinding{HostPort: strconv.Itoa(portItem), HostIP: "0.0.0.0"}
			portMap[nat.Port(fmt.Sprintf("%s/%s", port.Target, port.Protocol))] = []nat.PortBinding{bindItem}
		}
	}
	return portMap, nil
}

// Take temporary measures，Take temporary measures
func arrayIndexOf(key string, list []string) int {
	for i, s := range list {
		if key == s {
			return i
		}
	}

	return -1
}

// Rollback software
func removeAll(dockerClient *client.Client, name string) error {
	args := filters.NewArgs(
		//filters.Arg("label", fmt.Sprintf("com.docker.container.namespace=%s", name)),
		filters.Arg("label", fmt.Sprintf("%s=%s", ProjectLabel, strings.ToLower(name))),
		//filters.Arg("label", fmt.Sprintf("com.docker.stack.namespace=%s", strings.ToLower(name))),
	)
	list, err := dockerClient.ContainerList(context.Background(), types.ContainerListOptions{All: true, Filters: args})
	if err != nil {
		return utils.Errorf(err)
	}

	for _, con := range list {
		dockerClient.ContainerStop(context.Background(), con.ID, nil)
		dockerClient.ContainerRemove(context.Background(), con.ID, types.ContainerRemoveOptions{})
	}
	return nil
}

func createNetwork(networkName string, dockerClient *client.Client) (string, error) {
	// Check if the network exists，Check if the network exists
	arg := filters.NewArgs(
		filters.Arg("label", fmt.Sprintf("com.docker.network.namespace=%s", networkName)),
		filters.Arg("label", fmt.Sprintf("com.docker.network.project=%s", networkName)),
	)
	if list, err := dockerClient.NetworkList(context.Background(), types.NetworkListOptions{Filters: arg}); err != nil {
		return "", utils.Errorf(err)
	} else if len(list) == 0 {
		create, err := dockerClient.NetworkCreate(context.Background(), networkName, types.NetworkCreate{Labels: map[string]string{
			"com.docker.network.project":   networkName,
			"com.docker.network.namespace": networkName,
		}})
		if err != nil {
			return create.ID, utils.Errorf(err)
		}
		return create.ID, nil
	}

	return "", nil
}

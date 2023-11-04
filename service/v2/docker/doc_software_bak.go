// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-10 16:35
package docker

//
//import (
//	"cnic/fairman-gin/global"
//	"cnic/fairman-gin/model/v2/docker"
//	docRequst "cnic/fairman-gin/model/v2/docker/request"
//	"cnic/fairman-gin/utils"
//	"cnic/fairman-gin/utils/exec"
//	"cnic/fairman-gin/utils/filesystem"
//	"context"
//	"fmt"
//	"strings"
//
//	"github.com/docker/docker/api/types/filters"
//
//	dockerYml "github.com/lliuhuan/compose-yml"
//
//	"github.com/docker/docker/api/types"
//
//	"github.com/gin-gonic/gin"
//)
//
//type SoftwareService struct {
//}
//
////func (s *SoftwareService) Test(id string) (swarm.Service, string, error) {
////	dockerClient, err := global.FairClient.CreateClient(docker.DocEndpoint{IsSocket: true})
////	if err != nil {
////		return swarm.Service{}, "", err
////	}
////	defer dockerClient.Close()
////
////	service, data, err := dockerClient.ServiceInspectWithRaw(context.TODO(), id, types.ServiceInspectOptions{InsertDefaults: true})
////	return service, string(data), err
////}
////
////func (s *SoftwareService) Test1(service swarm.Service) (types.ServiceUpdateResponse, error) {
////	dockerClient, err := global.FairClient.CreateClient(docker.DocEndpoint{IsSocket: true})
////	if err != nil {
////		return types.ServiceUpdateResponse{}, err
////	}
////	defer dockerClient.Close()
////
////	serviceUpdate, err := dockerClient.ServiceUpdate(context.TODO(), service.ID, service.Version, service.Spec, types.ServiceUpdateOptions{Rollback: "previous"})
////
////	return serviceUpdate, err
////}
////
////// GetSoftware Obtaining software
////func (s *SoftwareService) GetSoftware(c *gin.Context) (software []response.SoftwareResp, err error) {
////	//id := utils.GetUserID(c)
////	id := "0db56233-76a2-445b-a60d-5f2590cb8672"
////	// Get and Setip
////	ip, _ := utils.ExternalIP()
////	fmt.Println("UserId: ", id)
////	url := fmt.Sprintf(GetSoftwareByUserId, global.FairConf.Service.Url, id)
////	var resp response.Response
////	var softwareResp response.SoftwareByUserIdResp
////	err = utils.Get(url, &resp)
////	if err != nil {
////		return software, utils.Errorf(err)
////	}
////
////	if resp.Code == 0 {
////		err := utils.DataJson(resp.Data, &softwareResp)
////		if err != nil {
////			return software, utils.Errorf(err)
////		}
////
////		for i := 0; i < len(softwareResp.Template); i++ {
////			for j := 0; j < len(softwareResp.Software); j++ {
////				fmt.Printf(
////					"templateId: [%s] == softwareId: [%s]   %v \n",
////					softwareResp.Template[i].Id,
////					softwareResp.Software[j].TemplateId,
////					softwareResp.Template[i].Id == softwareResp.Software[j].TemplateId,
////				)
////				fmt.Println(softwareResp.Software[j].Status)
////				if softwareResp.Template[i].Id == softwareResp.Software[j].TemplateId {
////					if softwareResp.Software[j].UserIp == ip.String() {
////						softwareResp.Template[i].IsInstall = true
////						// Set software installation location
////						var resp1 response.Response
////						url1 := fmt.Sprintf(GetEndpointById, global.FairConf.Service.Url, softwareResp.Software[j].EndpointId)
////						err1 := utils.Get(url1, &resp1)
////						if resp1.Code == 0 {
////							err1 = utils.DataJson(resp1.Data, &softwareResp.Software[j].Endpoint)
////						}
////						if err1 != nil || resp1.Code != 0 || resp1.Data == nil {
////							softwareResp.Software[j].Endpoint.Name = "Local"
////						}
////						software = append(software, response.SoftwareResp{
////							SoftwareId:         softwareResp.Software[j].Id,
////							CreateAt:           softwareResp.Software[j].CreateAt,
////							UpdateAt:           softwareResp.Software[j].UpdateAt,
////							SoftwareName:       softwareResp.Software[j].Name,
////							TemplateLogo:       softwareResp.Template[i].Logo,
////							TemplateScreenshot: softwareResp.Template[i].Screenshot,
////							TemplateTypes:      softwareResp.Template[i].Types,
////							IsInstall:          true,
////							TemplateDownload:   softwareResp.Template[i].Download,
////							TemplateVersion:    softwareResp.Template[i].Version,
////							TemplateScore:      softwareResp.Template[i].Score,
////							TemplateId:         softwareResp.Template[i].Id,
////							EndpointName:       softwareResp.Software[j].Endpoint.Name,
////							TemplateType:       softwareResp.Template[i].TemplateType,
////							HomeHost:           softwareResp.Template[i].HomeHost,
////							HomePort:           softwareResp.Template[i].HomePort,
////							Status:             softwareResp.Software[j].Status,
////						})
////
////						break
////					}
////				}
////			}
////		}
////
////		for _, template := range softwareResp.Template {
////			if !template.IsInstall {
////				software = append(software, response.SoftwareResp{
////					CreateAt:           template.CreateAt,
////					UpdateAt:           template.UpdateAt,
////					SoftwareName:       template.Name,
////					TemplateLogo:       template.Logo,
////					TemplateScreenshot: template.Screenshot,
////					TemplateTypes:      template.Types,
////					IsInstall:          false,
////					TemplateDownload:   template.Download,
////					TemplateVersion:    template.Version,
////					TemplateScore:      template.Score,
////					TemplateType:       template.TemplateType,
////					TemplateId:         template.Id,
////					HomeHost:           template.HomeHost,
////					HomePort:           template.HomePort,
////				})
////			}
////		}
////
////		//for _, templateInfo := range softwareResp.Template {
////		//	for _, softwareInfo := range softwareResp.Software {
////		//		if templateInfo.Id == softwareInfo.TemplateId {
////		//			// Set software installation location
////		//			var resp1 response.Response
////		//			url1 := fmt.Sprintf(GetEndpointById, global.FairConf.Service.Url, softwareInfo.EndpointId)
////		//			err1 := utils.Get(url1, &resp1)
////		//			if resp1.Code == 0 {
////		//				err1 = utils.DataJson(resp1.Data, &softwareInfo.Endpoint)
////		//			}
////		//			if err1 != nil || resp1.Code != 0 || resp1.Data == nil {
////		//				softwareInfo.Endpoint.Socket = true
////		//				softwareInfo.Endpoint.URL = "unix:///var/run/docker.sock"
////		//			}
////		//
////		//			dockerClient, err1 := global.FairClient.CreateClient(softwareInfo.Endpoint)
////		//			if err1 != nil {
////		//				continue
////		//			}
////		//
////		//			software.SoftwareId = softwareInfo.Id
////		//			software.CreateAt = softwareInfo.CreateAt
////		//		}
////		//	}
////		//}
////	}
////	return software, nil
////}
////
////// GetSoftwareDetails querycontainerquerytaskquery
////func (s *SoftwareService) GetSoftwareDetails(softwareId string) (interface{}, error) {
////	var resp response.Response
////	var software docker.DocSoftware
////	url := fmt.Sprintf(GetSoftwareById, global.FairConf.Service.Url, softwareId)
////	err := utils.Get(url, &resp)
////	if err != nil {
////		return nil, utils.Errorf(err)
////	}
////
////	if resp.Code == 0 {
////		err1 := utils.DataJson(resp.Data, &software)
////		if err1 != nil {
////			return nil, utils.Errorf(err1)
////		}
////
////		fmt.Println(software.Name)
////
////		// Set software installation location
////		var resp1 response.Response
////		url1 := fmt.Sprintf(GetEndpointById, global.FairConf.Service.Url, software.EndpointId)
////		err1 = utils.Get(url1, &resp1)
////		if resp1.Code == 0 {
////			err1 = utils.DataJson(resp1.Data, &software.Endpoint)
////		}
////		if err1 != nil || resp1.Code != 0 || resp1.Data == nil {
////			// TODO: Need to make some modifications Need to make some modificationsWindows
////			software.Endpoint.Socket = true
////			software.Endpoint.URL = "unix:///var/run/docker.sock"
////		}
////
////		switch global.StackType(software.TemplateType) {
////		case global.DockerStack:
////			fmt.Println("Docker")
////		case global.DockerSwarmStack:
////			return s.getStackSwarmDetails(software)
////		case global.DockerComposeStack:
////			return s.getStackComposeDetails(software)
////		}
////	}
////	return nil, nil
////}
////
////func (s *SoftwareService) GetSoftwareInspect(inspect request.SoftwareInspectReq) (interface{}, error) {
////
////	software := s.getSoftwareById(inspect.SoftwareId)
////
////	switch global.StackType(software.TemplateType) {
////	case global.DockerStack:
////		fmt.Println("Docker")
////	case global.DockerSwarmStack:
////		return s.getStackSwarmInspect(software, inspect.Id)
////	case global.DockerComposeStack:
////		return s.getStackComposeInspect(software, inspect.Id)
////	}
////
////	return nil, nil
////}
//
//// CreateSoftware Install software
//func (s *SoftwareService) CreateSoftware(software docker.DocSoftware, c *gin.Context) error {
//	// Get and Setid
//	software.UserId = utils.GetUserID(c)
//
//	// Set File Name
//	if software.EntryPoint == "" {
//		software.EntryPoint = filesystem.ComposeFileDefaultName
//	}
//
//	switch global.StackType(software.Endpoint.Type) {
//	case global.DockerStack:
//		//return s.createDockerStack(software)
//		return s.createSwarmStackFromFileContent(software)
//	case global.DockerSwarmStack:
//		return s.createComposeStackFromFileContent(software)
//		//return s.createSwarmStackFromFileContent(software)
//	case global.DockerComposeStack:
//		//return s.createComposeStackFromFileContent(software)
//		return s.createSwarmStackFromFileContent(software)
//	case global.KubernetesStack:
//		return s.createKubernetesStackFromFileContent(software)
//	}
//	return nil
//}
//
////func (s *SoftwareService) GetSoftwareLogs(l request.GetServiceLog) ([]string, error) {
////	software := s.getSoftwareById(l.SoftwareId)
////	var logList []string
////
////	dockerClient, err := global.FairClient.CreateClient(software.Endpoint)
////	if err != nil {
////		return nil, utils.Errorf(err)
////	}
////	defer dockerClient.Close()
////
////	logs, err := dockerClient.ServiceLogs(
////		context.TODO(),
////		l.ServiceId,
////		types.ContainerLogsOptions{
////			Tail:       l.Tail,
////			Since:      l.Since,
////			ShowStderr: true,
////			ShowStdout: true,
////			Timestamps: l.Timestamps,
////		},
////	)
////	if err != nil {
////		fmt.Println("err, ", err.Error())
////	}
////
////	buf := bufio.NewReader(logs)
////	for {
////		line, err := buf.ReadString('\n')
////		line = strings.TrimSpace(line)
////		if err != nil {
////			if err == io.EOF {
////				return logList, nil
////			}
////			return nil, err
////		}
////		logList = append(logList, line)
////	}
////}
//
//func (s *SoftwareService) UpdateSoftware(req docker.DocSoftware) error {
//	switch global.StackType(req.Endpoint.Type) {
//	case global.DockerStack:
//		fmt.Println("Docker")
//		return s.updateSwarmStackFromFileContent(req)
//	case global.DockerSwarmStack:
//		//req.Template.Config.SwarmConfig = req.Software.Template.Config.SwarmConfig
//		//
//		//req.Software.Template = req.Template
//		return s.updateSwarmStackFromFileContent(req)
//	case global.DockerComposeStack:
//		fmt.Println("Compose")
//		return s.updateSwarmStackFromFileContent(req)
//	}
//	return nil
//}
//
//// RemoveSoftware Remove software
//func (s *SoftwareService) RemoveSoftware(software docRequst.DeleteDocSoftwareRequest) error {
//	// Query software information
//	endpoint, err := global.FairStore.Endpoint().Endpoint(software.EndpointId)
//	if err != nil {
//		return utils.Errorf(err)
//	}
//
//	software.Endpoint = endpoint
//
//	switch global.StackType(endpoint.Type) {
//	case global.DockerStack:
//		fmt.Println("Docker")
//		return s.deleteSwarmStack(software)
//	case global.DockerSwarmStack:
//		//return s.deleteComposeStack(software)
//		return s.deleteSwarmStack(software)
//	case global.DockerComposeStack:
//		fmt.Println("compose")
//		//return s.deleteComposeStack(software)
//		return s.deleteSwarmStack(software)
//	}
//	return nil
//}
//
////func (s *SoftwareService) RemoveSoftwareService(l request.DelSoftwareService) error {
////	software := s.getSoftwareById(l.SoftwareId)
////
////	dockerClient, err := global.FairClient.CreateClient(software.Endpoint)
////	if err != nil {
////		return err
////	}
////	defer dockerClient.Close()
////
////	err = dockerClient.ServiceRemove(context.TODO(), l.ServiceId)
////	if err != nil {
////		return err
////	}
////
////	// Determine if the software name exists
////	isUnique, err := s.checkUniqueStackNameInDocker(software.Name, true, software.Endpoint)
////	if err != nil {
////		return err
////	}
////	if !isUnique {
////		return nil
////	}
////
////	// All software services have been removed，All software services have been removed
////	var res response.Response
////	err = utils.Delete(fmt.Sprintf(global.DelSoftware, global.FairConf.Service.Url, software.Id), nil, "", &res)
////	if err != nil {
////		return err
////	}
////
////	if res.Code == 0 {
////		return nil
////	}
////
////	return utils.Errorf("Code != 0")
////}
////
////// StopSoftware Stop software
////func (s *SoftwareService) StopSoftware(softwareId string) error {
////	// Query software information
////	software := s.getSoftwareById(softwareId)
////	fmt.Println(software.TemplateType)
////	switch global.StackType(software.TemplateType) {
////	case global.DockerStack:
////		fmt.Println("Docker")
////	case global.DockerSwarmStack:
////		fmt.Println("Swarm")
////		return s.stopSwarmStack(software)
////	case global.DockerComposeStack:
////		fmt.Println("compose")
////		return s.stopComposeStack(software)
////	}
////	return nil
////}
////
////func (s *SoftwareService) StartSoftware(softwareId string) error {
////	// Query software information
////	software := s.getSoftwareById(softwareId)
////
////	switch global.StackType(software.TemplateType) {
////	case global.DockerStack:
////		fmt.Println("Docker")
////	case global.DockerSwarmStack:
////		return s.startSwarmStack(software)
////	case global.DockerComposeStack:
////		fmt.Println("compose")
////		return s.startComposeStack(software)
////	}
////	return nil
////}
////
////func (s *SoftwareService) UpdateSoftwareReplicas(req request.UpdateSoftwareReplicasReq) (types.ServiceUpdateResponse, error) {
////	dockerClient, err := global.FairClient.CreateClient(docker.DocEndpoint{Socket: true})
////	if err != nil {
////		return types.ServiceUpdateResponse{}, err
////	}
////	defer dockerClient.Close()
////
////	serviceList, _ := dockerClient.ServiceList(context.Background(), types.ServiceListOptions{Filters: filters.NewArgs(filters.Arg("id", req.Id))})
////
////	serviceList[0].Spec.Mode.Replicated.Replicas = &req.Replicas
////
////	update, err := dockerClient.ServiceUpdate(context.Background(), req.Id, serviceList[0].Version, serviceList[0].Spec, types.ServiceUpdateOptions{})
////
////	if err != nil {
////		err = utils.Errorf(err)
////		return update, err
////	}
////	return update, nil
////}
//
//func (s *SoftwareService) checkUniqueStackNameInDocker(name string, swarmMode bool, endpoint docker.DocEndpoint) (bool, error) {
//	// TODO: Can you check if the database exists
//
//	// TODO: Link to other servers，Link to other servers，Link to other servers
//	dockerClient, err := global.FairClient.CreateClient(endpoint)
//	if err != nil {
//		return false, err
//	}
//
//	defer dockerClient.Close()
//
//	if swarmMode {
//		services, err := dockerClient.ServiceList(context.Background(), types.ServiceListOptions{})
//		if err != nil {
//			return false, err
//		}
//		for _, service := range services {
//			//fmt.Printf("%s - %v \n", service.Spec.Labels["com.docker.stack.namespace"], service)
//			serviceNS, ok := service.Spec.Labels["com.docker.stack.namespace"]
//			if ok && serviceNS == name {
//				return false, nil
//			}
//		}
//	}
//	containers, err := dockerClient.ContainerList(context.Background(), types.ContainerListOptions{All: true})
//	if err != nil {
//		return false, err
//	}
//
//	for _, container := range containers {
//		containerNS, ok := container.Labels["com.docker.compose.project"]
//
//		if ok && containerNS == name {
//			return false, nil
//		}
//	}
//
//	return true, nil
//}
//
//func (s *SoftwareService) checkStackNameStatus(name string, swarmMode bool, endpoint docker.DocEndpoint) (string, error) {
//	var status string
//	// TODO: Can you check if the database exists
//
//	// TODO: Link to other servers，Link to other servers，Link to other servers
//	dockerClient, err := global.FairClient.CreateClient(endpoint)
//	if err != nil {
//		return "error", err
//	}
//
//	defer dockerClient.Close()
//
//	// TODO: Is it necessary to change the method used to verify whether it is running based on the environment in the later stage
//
//	containers, err := dockerClient.ContainerList(context.Background(), types.ContainerListOptions{All: true})
//	if err != nil {
//		return "error", err
//	}
//
//	for _, container := range containers {
//		containerComposeNS, ComposeOK := container.Labels["com.docker.compose.project"]
//		containerSwarmNS, SwarmOK := container.Labels["com.docker.stack.namespace"]
//
//		if (ComposeOK && containerComposeNS == name) || (SwarmOK && containerSwarmNS == name) {
//			status = "running"
//			//status = container.State
//			//fmt.Println(container.Names, container.State)
//			//if status != "running" {
//			//	return status, nil
//			//}
//		}
//	}
//
//	return status, nil
//}
//
//func (s *SoftwareService) cleanUp(projectPath string, doCleanUp *bool) error {
//	if !*doCleanUp {
//		return nil
//	}
//
//	err := global.FairFileSystem.RemoveDirectory(projectPath)
//	if err != nil {
//		//log.Printf("http error: Unable to cleanup stack creation (err=%s)\n", err)
//		global.FairLog.Error("http error: Unable to cleanup stack creation (err=" + err.Error() + ")\n")
//	}
//	return nil
//}
//
//func (s *SoftwareService) createComposeDeployConfig(stack *exec.ComposeStackManager, software docker.DocSoftware) (*composeStackDeploymentConfig, error) {
//	config := &composeStackDeploymentConfig{
//		stack:    stack,
//		software: software,
//		//registries: resp.Data,
//	}
//	return config, nil
//}
//
////func (s *SoftwareService) getSoftwareById(softwareId string) (software docker.DocSoftware) {
////	var resp response.Response
////	url := fmt.Sprintf(GetSoftwareById, global.FairConf.Service.Url, softwareId)
////	err := utils.Get(url, &resp)
////	if err != nil {
////		return
////	}
////
////	if resp.Code == 0 {
////		err1 := utils.DataJson(resp.Data, &software)
////		if err1 != nil {
////			return
////		}
////
////		// Set software installation location
////		var resp1 response.Response
////		url1 := fmt.Sprintf(GetEndpointById, global.FairConf.Service.Url, software.EndpointId)
////		err1 = utils.Get(url1, &resp1)
////		if resp1.Code == 0 {
////			err1 = utils.DataJson(resp1.Data, &software.Endpoint)
////		}
////		if err1 != nil || resp1.Code != 0 || resp1.Data == nil {
////			// TODO: Need to make some modifications Need to make some modificationsWindows
////			software.Endpoint.Socket = true
////			software.Endpoint.URL = "unix:///var/run/docker.sock"
////		}
////
////		return software
////	}
////
////	return
////}
//
//func (s *SoftwareService) expandFileContent(software *docker.DocSoftware) (docker.DocSoftware, error) {
//	accCount := 0
//	accMap := make(map[string]string, 2)
//	var yml dockerYml.Yml
//	err := dockerYml.UnmarshalYaml(software.Template.Config.SwarmContent, &yml)
//	if err != nil {
//		return docker.DocSoftware{}, utils.Errorf("Unable to parse YAML file: %s", err)
//	}
//
//	for key, service := range yml.Services {
//		service.Ports = software.Template.Config.SwarmConfig[key].Ports
//		service.Volumes = software.Template.Config.SwarmConfig[key].Mounts
//		service.Environment = software.Template.Config.SwarmConfig[key].Environment
//		service.HostName = software.Template.Config.SwarmConfig[key].Hostname
//		for _, environment := range service.Environment {
//			if strings.ToLower(environment.Key) == "acc_host" {
//				accCount++
//				accMap["acc_host"] = environment.Val
//			}
//			if strings.ToLower(environment.Key) == "acc_port" {
//				accCount++
//				accMap["acc_port"] = environment.Val
//			}
//		}
//	}
//
//	software.Template.Config.SwarmContent, err = dockerYml.MarshalYaml(yml)
//
//	if accCount == 2 {
//		software.HomeHost = accMap["acc_host"]
//		software.HomePort = accMap["acc_port"]
//	} else {
//		return docker.DocSoftware{}, utils.Errorf("acc_hostoracc_portor")
//	}
//
//	fmt.Println(software.Template.Config.SwarmConfig)
//	fmt.Println(software.Template.Config.SwarmContent)
//	return *software, nil
//}
//
//// GetSoftwareStatus Obtain software status
//func (s *SoftwareService) GetSoftwareStatus(softwareList []docRequst.SoftwareStatusRequest) ([]docRequst.SoftwareStatusRequest, error) {
//	for index, software := range softwareList {
//		//err := global.FairStore.Endpoint().Create(&software.Endpoint)
//		//if err != nil {
//		//	return nil, err
//		//}
//
//		if software.Endpoint.Type == 4 {
//			//k8sConfig, err := clientcmd.BuildConfigFromFlags("", path.Join("data", "config", software.Endpoint.Name, "config"))
//			//if err != nil {
//			//	continue
//			//}
//			//
//			//k8sClient, err := kubernetes.NewForConfig(k8sConfig)
//			//if err != nil {
//			//	continue
//			//}
//			//
//			//fmt.Println(k8sClient.AppsV1().Deployments(software.Name).Get(context.Background(), software.Name, metav1.GetOptions{}))
//			//deployment, err := k8sClient.AppsV1().Deployments()
//		} else {
//			dockerClient, err := global.FairClient.CreateClient(software.Endpoint)
//			if err != nil {
//				return nil, utils.Errorf(err)
//			}
//
//			ctx := context.Background()
//			filter := filters.NewArgs(filters.Arg("label", fmt.Sprintf("%s=%s", "com.docker.stack.namespace", software.Name)))
//			services, err := dockerClient.ServiceList(ctx, types.ServiceListOptions{
//				Filters: filter,
//				Status:  true,
//			})
//			if err != nil {
//				return nil, utils.Errorf(err)
//			}
//
//			if len(services) != 0 {
//				softwareList[index].Status = "running"
//			} else {
//				softwareList[index].Status = "stopped"
//			}
//
//			dockerClient.Close()
//		}
//	}
//
//	endpoints, err := global.FairStore.Endpoint().Endpoints()
//	fmt.Println(endpoints, err)
//
//	return softwareList, nil
//}

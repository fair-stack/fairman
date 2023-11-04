// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-09-20 19:50
package docker

import (
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/model/v2/docker"
	docRequst "cnic/fairman-gin/model/v2/docker/request"
	"cnic/fairman-gin/utils"
	"cnic/fairman-gin/utils/exec"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	gyaml "github.com/ghodss/yaml"
	appsV1 "k8s.io/api/apps/v1"
	k8sYml "k8s.io/apimachinery/pkg/util/yaml"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func (s *SoftwareService) createKubernetesStackFromFileContent(software docker.DocSoftware) error {
	// initializationswarm stack
	manager := exec.NewKubernetesDeployer("cmd", "data", software.Endpoint.Name)

	//buffer := yaml_k8s.IsJSONBuffer([]byte(software.Template.Config.K8sContent))
	//if buffer {
	//	deployJson, err := yaml_k8s.ToJSON([]byte(software.Template.Config.K8sContent))
	//	if err != nil {
	//		return utils.Errorf(err)
	//	}
	//	var Dep core_v1.Service
	//	err = json.Unmarshal(deployJson, &Dep)
	//	if err != nil {
	//		return utils.Errorf(err)
	//	}
	//
	//	fmt.Println(Dep.Spec.Ports)
	//}

	// establishdocker-compose
	//projectPath, err := global.FairFileSystem.StoreStackFileFromBytes(software.Name, "", []byte(software.Template.Config.K8sContent))
	files, err := global.FairFileSystem.CreateKubernetesFile(software.Name, software.Template.Config.K8sConfig)
	if err != nil {
		return utils.Errorf(err)
	}
	//software.ProjectPath = projectPath

	//doCleanUp := false
	// deletedocker-compose
	//defer s.cleanUp(software.ProjectPath, &doCleanUp)

	software.Name = manager.NormalizeStackName(software.Name)
	// establishnamespace
	err = manager.CreateNamespace(software.Name)
	// Processing software name Processing software name
	return manager.Deploy(software.Name, files, true)
}

// updateKubernetesStackFromFileContent updatekubernetes stack
func (s *SoftwareService) updateKubernetesStackFromFileContent(req docRequst.UpdateSoftwareRequest) error {
	//fmt.Println(req)
	manager := exec.NewKubernetesDeployer("cmd", "data", req.Endpoint.Name)

	deploymentPath, err := s.updateDeploymentFile(req)
	if err != nil {
		return utils.Errorf(err)
	}

	return manager.Deploy(req.Name, []string{deploymentPath}, true)
}

//func (s *SoftwareService) deployKubernetesStack(software docker.DocSoftware) error {
//
//}

func (s *SoftwareService) deleteKubernetesStack(software docRequst.DeleteDocSoftwareRequest) error {
	// initializationswarm stack
	manager := exec.NewKubernetesDeployer("cmd", "data", software.Endpoint.Name)

	// Processing software name Processing software name
	//software.Name = manager.NormalizeStackName(software.Name)
	files, err := s.getFiles(path.Join("data", "compose", software.Name))
	if err != nil {
		return utils.Errorf(err)
	}
	return manager.Remove(files, software.Name)
}

func (s *SoftwareService) stopKubernetesStack(software docRequst.StopSoftwareRequest) error {
	//software.Name = "datalabbaas"
	kubeconfig := path.Join("data", "config", software.Endpoint.Name, "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return utils.Errorf(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return utils.Errorf(err)
	}

	namespace := software.Name
	deployment, err := clientset.AppsV1().Deployments(namespace).Get(context.Background(), software.Name+"deployment", metav1.GetOptions{})
	if err != nil {
		return utils.Errorf(err)
	}

	var replicas int32 = 0
	deployment.Spec.Replicas = &replicas
	_, err = clientset.AppsV1().Deployments(namespace).Update(context.Background(), deployment, metav1.UpdateOptions{})
	if err != nil {
		return utils.Errorf(err)
	}

	return nil
}

func (s *SoftwareService) startKubernetesStack(software docRequst.StartSoftwareRequest) error {
	//software.Name = "datalabbaas"

	kubeconfig := path.Join("data", "config", software.Endpoint.Name, "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return utils.Errorf(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return utils.Errorf(err)
	}

	namespace := software.Name
	deployment, err := clientset.AppsV1().Deployments(namespace).Get(context.Background(), software.Name+"deployment", metav1.GetOptions{})
	if err != nil {
		return utils.Errorf(err)
	}

	var replicas int32 = 1
	deployment.Spec.Replicas = &replicas
	_, err = clientset.AppsV1().Deployments(namespace).Update(context.Background(), deployment, metav1.UpdateOptions{})
	if err != nil {
		return utils.Errorf(err)
	}

	return nil
}

func (s *SoftwareService) getImage(imageStr string) (img, ver string, err error) {
	imageSplit := strings.Split(imageStr, ":")
	if len(imageSplit) == 3 {
		img = imageSplit[0] + ":" + imageSplit[1]
		ver = imageSplit[2]
	} else if len(imageSplit) == 2 {
		img = imageSplit[0]
		ver = imageSplit[1]
	} else {
		return "", "", utils.Errorf("image format error")
	}
	return img, ver, nil
}

func (s *SoftwareService) getFiles(folder string) ([]string, error) {
	paths := make([]string, 0)
	files, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			newPaths, err := s.getFiles(folder + "/" + file.Name())
			if err != nil {
				return nil, err
			}

			paths = append(paths, newPaths...)
		} else {
			paths = append(paths, folder+"/"+file.Name())
		}
	}

	return paths, nil
}

func (s *SoftwareService) updateDeploymentFile(req docRequst.UpdateSoftwareRequest) (string, error) {
	rootPath := path.Join("data", "compose", req.Name)
	fileName := "deployment.yaml"
	content, err := global.FairFileSystem.GetFileContent(rootPath, fileName)
	if err != nil {
		return "", utils.Errorf(err)
	}

	deployJson, err := k8sYml.ToJSON(content)
	if err != nil {
		return "", utils.Errorf(err)
	}

	var deployment appsV1.Deployment
	err = json.Unmarshal(deployJson, &deployment)
	if err != nil {
		return "", utils.Errorf(err)
	}

	for i, c := range deployment.Spec.Template.Spec.Containers {
		img, ver, err := s.getImage(c.Image)
		if err != nil {
			return "", utils.Errorf(err)
		}
		for _, container := range req.KubernetesConfig.Deployment.Spec.Template.Spec.Containers {
			img1, ver1, err1 := s.getImage(container.Image)
			if err1 != nil {
				return "", utils.Errorf(err)
			}

			if img == img1 && ver != ver1 {
				deployment.Spec.Template.Spec.Containers[i].Image = fmt.Sprintf("%s:%s", img1, ver1)
			}
		}
	}

	fmt.Println(deployment.Spec.Template.Spec.Containers[0].Image)
	deployJson, err = json.Marshal(deployment)
	if err != nil {
		return "", utils.Errorf(err)
	}

	yamlByte, err := gyaml.JSONToYAML(deployJson)
	if err != nil {
		return "", utils.Errorf(err)
	}

	// preserveyamlBytepreserve
	deploymentFile, err := os.OpenFile(path.Join(rootPath, fileName), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return "", err
	}
	deploymentFile.WriteString(string(yamlByte))
	deploymentFile.Close()
	return path.Join(rootPath, fileName), nil
}

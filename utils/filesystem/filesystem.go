// Package filesystem
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-17 19:00
package filesystem

import (
	"bytes"
	"cnic/fairman-gin/model/v2/docker"
	"cnic/fairman-gin/utils"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	gyaml "github.com/ghodss/yaml"
)

const (
	// ComposeStorePath represents the subfolder where compose files are stored in the file store folder.
	ComposeStorePath = "compose"
	// ComposeFileDefaultName represents the default name of a compose file.
	ComposeFileDefaultName = "docker-compose.yml"
)

// Service represents a service for managing files and directories.
type Service struct {
	dataStorePath string
	fileStorePath string
}

// JoinPaths takes a trusted root path and a list of untrusted paths and joins
// them together using directory separators while enforcing that the untrusted
// paths cannot go higher up than the trusted root
func JoinPaths(trustedRoot string, untrustedPaths ...string) string {
	if trustedRoot == "" {
		trustedRoot = "."
	}

	p := filepath.Join(trustedRoot, filepath.Join(append([]string{"/"}, untrustedPaths...)...))

	// avoid setting a volume name from the untrusted paths
	vnp := filepath.VolumeName(p)
	if filepath.VolumeName(trustedRoot) == "" && vnp != "" {
		return strings.TrimPrefix(strings.TrimPrefix(p, vnp), `\`)
	}

	return p
}

func NewService(dataStorePath, fileStorePath string) (*Service, error) {
	service := &Service{
		dataStorePath: dataStorePath,
		fileStorePath: JoinPaths(dataStorePath, fileStorePath),
	}

	err := os.MkdirAll(dataStorePath, 0755)
	if err != nil {
		return nil, err
	}

	err = service.createDirectoryInStore(ComposeStorePath)
	if err != nil {
		return nil, err
	}

	return service, nil
}

// StoreStackFileFromBytes creates a subfolder in the ComposeStorePath and stores a new file from bytes.
// It returns the path to the folder where the file is stored.
func (service *Service) StoreStackFileFromBytes(stackIdentifier, fileName string, data []byte) (string, error) {
	if fileName == "" {
		fileName = ComposeFileDefaultName
	}
	stackStorePath := JoinPaths(ComposeStorePath, stackIdentifier)
	err := service.createDirectoryInStore(stackStorePath)
	if err != nil {
		return "", err
	}

	composeFilePath := JoinPaths(stackStorePath, fileName)
	r := bytes.NewReader(data)

	fmt.Println(composeFilePath)
	err = service.createFileInStore(composeFilePath, r)
	if err != nil {
		return "", err
	}

	return service.wrapFileStore(stackStorePath), nil
}

func (service *Service) CreateKubernetesFile(stackIdentifier string, conf docker.KubernetesConfig) (files []string, err error) {
	stackStorePath := JoinPaths(ComposeStorePath, stackIdentifier)
	err = service.createDirectoryInStore(stackStorePath)
	if err != nil {
		return nil, err
	}

	if conf.Deployment.APIVersion != "" {
		if conf.Deployment.Spec.Strategy.Type == "" {
			conf.Deployment.Spec.Strategy.Type = "RollingUpdate"
		}
		marshal, err := json.Marshal(conf.Deployment)
		if err != nil {
			return nil, utils.Errorf(err)
		}

		fmt.Println(string(marshal))

		toYAML, err := gyaml.JSONToYAML(marshal)
		if err != nil {
			return nil, utils.Errorf(err)
		}

		r := bytes.NewReader(toYAML)
		filePath := JoinPaths(stackStorePath, "deployment.yaml")
		err = service.createFileInStore(filePath, r)
		if err != nil {
			return nil, utils.Errorf(err)
		}
		files = append(files, service.wrapFileStore(filePath))
	}

	if conf.Service.APIVersion != "" {
		marshal, err := json.Marshal(conf.Service)
		if err != nil {
			return nil, utils.Errorf(err)
		}

		fmt.Println(string(marshal))

		toYAML, err := gyaml.JSONToYAML(marshal)
		if err != nil {
			return nil, utils.Errorf(err)
		}

		r := bytes.NewReader(toYAML)
		filePath := JoinPaths(stackStorePath, "service.yaml")
		err = service.createFileInStore(filePath, r)
		if err != nil {
			return nil, utils.Errorf(err)
		}
		files = append(files, service.wrapFileStore(filePath))
	}

	//if conf.PersistentVolumeClaim.APIVersion != "" {
	//	//if conf.PersistentVolumeClaim.Spec.Resources. == "" {
	//	//	conf.PersistentVolumeClaim.Spec.Strategy.Type = "RollingUpdate"
	//	//}
	//	marshal, err := json.Marshal(conf.PersistentVolumeClaim)
	//	if err != nil {
	//		return nil, utils.Errorf(err)
	//	}
	//
	//	fmt.Println(string(marshal))
	//
	//	toYAML, err := gyaml.JSONToYAML(marshal)
	//	if err != nil {
	//		return nil, utils.Errorf(err)
	//	}
	//
	//	r := bytes.NewReader(toYAML)
	//	filePath := JoinPaths(stackStorePath, "pvc.yaml")
	//	err = service.createFileInStore(filePath, r)
	//	if err != nil {
	//		return nil, utils.Errorf(err)
	//	}
	//	files = append(files, service.wrapFileStore(filePath))
	//}

	return
}

func (service *Service) GetProjectPath(fileName string) string {
	return JoinPaths(service.dataStorePath, ComposeStorePath, fileName, ComposeFileDefaultName)
}

func (service *Service) RemoveDirectory(directoryPath string) error {
	return os.RemoveAll(directoryPath)
}

// createFileInStore createFile creates a new file in the file store with the content from r.
func (service *Service) createFileInStore(filePath string, r io.Reader) error {
	path := service.wrapFileStore(filePath)

	out, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	//defer out.Close()

	_, err = io.Copy(out, r)
	return err
}

// createDirectoryInStore creates a new directory in the file store
func (service *Service) createDirectoryInStore(name string) error {
	path := service.wrapFileStore(name)
	return os.MkdirAll(path, 0700)
}

func (service *Service) wrapFileStore(filepath string) string {
	return JoinPaths(service.fileStorePath, filepath)
}

// GetFileContent returns the content of a file as bytes.
func (service *Service) GetFileContent(trustedRoot, filePath string) ([]byte, error) {
	content, err := os.ReadFile(JoinPaths(trustedRoot, filePath))
	if err != nil {
		if filePath == "" {
			filePath = trustedRoot
		}
		return nil, fmt.Errorf("could not get the contents of the file '%s'", filePath)
	}

	return content, nil
}

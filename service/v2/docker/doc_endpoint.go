// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-04 15:24
package docker

import (
	"cnic/fairman-gin/model/v2/docker"
	"cnic/fairman-gin/utils"
	"cnic/fairman-gin/utils/exec"
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/docker/docker/client"
)

type EndpointService struct{}

//// CreateEndpoint Creating an environment for native storage
//func (e *EndpointService) CreateEndpoint(endpoint *docker.DocEndpoint) error {
//	return global.FairStore.Endpoint().Create(endpoint)
//}
//
//// GetEndpoint Obtain the environment for local storage
//func (e *EndpointService) GetEndpoint(ID string) (*docker.DocEndpoint, error) {
//	return global.FairStore.Endpoint().Endpoint(ID)
//}
//
//// GetEndpointAll Obtain the environment for all local storage
//func (e *EndpointService) GetEndpointAll() ([]docker.DocEndpoint, error) {
//	return global.FairStore.Endpoint().Endpoints()
//}

// SaveEndpointConfig Save the environment for local storage Save the environment for local storagek8sSave the environment for local storageconfigSave the environment for local storage
func (e *EndpointService) SaveEndpointConfig(req docker.DocEndpoint) error {
	err := os.MkdirAll(path.Join("data", "config", req.Name), os.ModePerm)

	if err != nil {
		return utils.Errorf(err)
	}

	err = os.WriteFile(path.Join("data", "config", req.Name, "config"), []byte(req.K8sConfig), os.ModePerm)

	if err != nil {
		return utils.Errorf(err)
	}

	return nil
}

// TestEndpoint Testing the environment for local storage
func (e *EndpointService) TestEndpoint(endpoints []docker.DocEndpoint) ([]docker.DocEndpoint, error) {
	for key, endpoint := range endpoints {
		if endpoint.Type == 2 {
			var opts []client.Opt
			opts = append(opts, client.FromEnv)
			opts = append(opts, client.WithAPIVersionNegotiation())
			if endpoint.IsSocket {
				opts = append(opts, client.WithHost("unix:///var/run/docker.sock"))
			} else {
				opts = append(opts, client.WithHost(endpoint.EnvURL))
				if endpoint.IsTLS {
					opts = append(opts, client.WithTLSClientConfig(
						utils.TextAddUri(endpoint.TLSConfig.TLSCACertPath),
						utils.TextAddUri(endpoint.TLSConfig.TLSCertPath),
						utils.TextAddUri(endpoint.TLSConfig.TLSKeyPath),
					))
				}
			}
			clientWithOpts, err := client.NewClientWithOpts(
				opts...,
			)
			if err != nil {
				fmt.Println(utils.Errorf(err))
				endpoints[key].Connected = false
			}
			ping, err := clientWithOpts.Ping(context.Background())
			if err != nil {
				fmt.Println(utils.Errorf(err))
				endpoints[key].Connected = false
			}
			if ping.APIVersion != "" {
				endpoints[key].Connected = true
			} else {
				endpoints[key].Connected = false
			}
		} else if endpoint.Type == 4 {
			// TODO: testk8stest
			// takek8stake
			err := e.SaveEndpointConfig(endpoint)
			if err != nil {
				return nil, utils.Errorf(err)
			}

			manager := exec.NewKubernetesDeployer("cmd", "data", endpoint.Name)
			err = manager.Version()
			if err != nil {
				endpoints[key].Connected = false
			} else {
				endpoints[key].Connected = true
			}
		} else if endpoint.Type == 5 {
			sshClient, err := exec.NewSSHDeployer(endpoint.EnvURL, endpoint.UserName, endpoint.Password)
			if err != nil {
				endpoints[key].Connected = false
				continue
			}

			session, err := sshClient.NewSession()
			if err != nil {
				fmt.Println(err)
				endpoints[key].Connected = false
			} else {
				endpoints[key].Connected = true
			}
			defer session.Close()
		}
	}
	return endpoints, nil
}

func (e *EndpointService) TestOneEndpoint(req docker.DocEndpoint) (interface{}, error) {
	// != 4 Represented bydockerRepresented by
	if req.Type == 2 {
		var opts []client.Opt
		opts = append(opts, client.FromEnv)
		opts = append(opts, client.WithAPIVersionNegotiation())
		if req.IsSocket {
			opts = append(opts, client.WithHost("unix:///var/run/docker.sock"))
		} else {
			// Determine if it ishttps
			if strings.Index(req.EnvURL, "http") < 0 && strings.Index(req.EnvURL, "https") < 0 {
				req.EnvURL = "https://" + req.EnvURL
			}

			opts = append(opts, client.WithHost(req.EnvURL))
			if req.IsTLS {
				opts = append(opts, client.WithTLSClientConfig(
					utils.TextAddUri(req.TLSConfig.TLSCACertPath),
					utils.TextAddUri(req.TLSConfig.TLSCertPath),
					utils.TextAddUri(req.TLSConfig.TLSKeyPath),
				))
			}
		}
		clientWithOpts, err := client.NewClientWithOpts(
			opts...,
		)
		if err != nil {
			return req, utils.Errorf(err)
		}

		ping, err := clientWithOpts.Ping(context.Background())
		if err != nil {
			return req, utils.Errorf(err)
		}
		if ping.APIVersion == "" {
			return req, utils.Errorf("connection failed, connection failed")
		}
	} else if req.Type == 4 {
		// takek8stake
		err := e.SaveEndpointConfig(req)
		if err != nil {
			return req, utils.Errorf(err)
		}

		manager := exec.NewKubernetesDeployer("cmd", "data", req.Name)
		err = manager.Version()
		if err != nil {
			return req, utils.Errorf(err)
		}
	} else if req.Type == 5 {

	}
	return req, nil
}

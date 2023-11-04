// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-24 16:43
package docker

import (
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/model/v2/docker"
	"cnic/fairman-gin/utils"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/docker/docker/client"
)

const (
	defaultDockerRequestTimeout = 60
	dockerClientVersion         = "1.41"
)

type ClientFactory struct {
}

type registryAuthenticationHeader struct {
	Username      string `json:"username,omitempty"`
	Password      string `json:"password,omitempty"`
	Serveraddress string `json:"serveraddress,omitempty"`
}

func NewClientFactory() *ClientFactory {
	if strings.ToLower(global.FairConf.Docker.Type) == "local" {
		global.FairConf.Docker.Url = "unix:///var/run/docker.sock"
		if runtime.GOOS == "windows" {
			global.FairConf.Docker.Url = "npipe:////./pipe/docker_engine"
		}
	}

	return &ClientFactory{}
}

func (c *ClientFactory) CreateClient(endpoint docker.DocEndpoint) (*client.Client, error) {
	// TODO: Later on, it will be changed to support multiple methodsDocker，Later on, it will be changed to support multiple methods
	// Agent/Edge/Local
	if endpoint.IsSocket {
		return createLocalClient()
	}

	if strings.HasPrefix(endpoint.EnvURL, "unix://") || strings.HasPrefix(endpoint.EnvURL, "npipe://") {
		return createLocalClient()
	}

	fmt.Println(endpoint.EnvURL)

	return createTCPClient(endpoint)
}

func createLocalClient() (*client.Client, error) {
	if global.FairConf.Docker.Url == "" {
		global.FairConf.Docker.Url = "unix:///var/run/docker.sock"
	}
	return client.NewClientWithOpts(
		client.FromEnv,
		client.WithHost(global.FairConf.Docker.Url),
		client.WithAPIVersionNegotiation(),
	)
}

func createTCPClient(endpoint docker.DocEndpoint) (*client.Client, error) {
	httpCli, err := httpClient()
	if err != nil {
		return nil, err
	}

	if !strings.HasPrefix(endpoint.EnvURL, "tcp://") && !strings.HasPrefix(endpoint.EnvURL, "http://") && !strings.HasPrefix(endpoint.EnvURL, "https://") {
		endpoint.EnvURL = "tcp://" + endpoint.EnvURL
	}

	var opts = []client.Opt{
		client.FromEnv,
		client.WithHost(endpoint.EnvURL),
		client.WithHTTPClient(httpCli),
		client.WithVersion(dockerClientVersion),
		client.WithAPIVersionNegotiation(),
	}
	// Skip or notTLSSkip or not，Skip or not Skip or notTLSSkip or not
	if endpoint.IsTLS {
		opts = append(opts, client.WithTLSClientConfig(
			utils.TextAddUri(endpoint.TLSConfig.TLSCACertPath),
			utils.TextAddUri(endpoint.TLSConfig.TLSCertPath),
			utils.TextAddUri(endpoint.TLSConfig.TLSKeyPath)))
	}

	return client.NewClientWithOpts(
		opts...,
	)
	//WithHost(host), WithVersion(version), WithHTTPClient(client), WithHTTPHeaders(httpHeaders)
	//return client.NewClient(endpoint.EnvURL, dockerClientVersion, httpCli, nil)
}

func httpClient() (*http.Client, error) {
	transport := &http.Transport{}

	//if endpoint.TLSConfig.TLS {
	//	tlsConfig, err := crypto.CreateTLSConfigurationFromDisk(endpoint.TLSConfig.TLSCACertPath, endpoint.TLSConfig.TLSCertPath, endpoint.TLSConfig.TLSKeyPath, endpoint.TLSConfig.TLSSkipVerify)
	//	if err != nil {
	//		return nil, err
	//	}
	//	transport.TLSClientConfig = tlsConfig
	//}

	return &http.Client{
		Transport: transport,
		Timeout:   defaultDockerRequestTimeout * time.Second,
	}, nil
}

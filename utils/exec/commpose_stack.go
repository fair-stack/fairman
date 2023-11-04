// Package exec
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-15 15:22
package exec

import (
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/model/v2/docker"
	"cnic/fairman-gin/utils/exec/compose"
	"cnic/fairman-gin/utils/internal/stackutils"
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"

	"github.com/docker/cli/cli/compose/loader"

	"github.com/pkg/errors"
)

// ComposeStackManager is a wrapper for docker-compose binary
type ComposeStackManager struct {
	deployer   global.ComposeDeployer
	binaryPath string
	configPath string
}

// NewComposeStackManager returns a docker-compose wrapper if corresponding binary present, otherwise nil
func NewComposeStackManager(binaryPath string, configPath string) (*ComposeStackManager, error) {
	deployer, err := compose.NewComposeDeployer(binaryPath, configPath)
	if err != nil {
		return nil, err
	}

	return &ComposeStackManager{
		deployer:   deployer,
		binaryPath: binaryPath,
		configPath: configPath,
	}, nil
}

func (c *ComposeStackManager) Up(ctx context.Context, software docker.DocSoftware, isLog bool) error {
	// TODO: Pending processing
	url := software.Endpoint.EnvURL
	if software.Endpoint.IsSocket {
		url = ""
	}

	envFilePath, err := createEnvFile(software)
	if err != nil {
		return errors.Wrap(err, "failed to create env file")
	}

	filePaths := stackutils.GetStackFilePaths(software)
	err = c.deployer.Deploy(ctx, software.ProjectPath, url, software.Name, filePaths, envFilePath, isLog)

	if err != nil {
		err1 := c.Down(ctx, software)
		//err := wrapper.Remove(context.Background(), workingDir, host, projectName, []string{
		//	global.FairFileSystem.GetProjectPath(projectName),
		//})
		global.BaseChan <- global.ChanLogModal{
			Level:   "error",
			Message: "Abnormal installation，Abnormal installation",
		}

		if err1 != nil {
			fmt.Println("Rollback failed", err1)
			global.BaseChan <- global.ChanLogModal{
				Level:   "error",
				Message: "Rollback failed，Rollback failed",
			}
		}
		return err
	}
	return err
}

func (c *ComposeStackManager) Down(ctx context.Context, software docker.DocSoftware) error {
	//url, _, err := c.fetchEndpointProxy()
	url := software.Endpoint.EnvURL
	if software.Endpoint.IsSocket {
		url = ""
	}

	if err := updateNetworkEnvFile(software); err != nil {
		return err
	}

	filePaths := stackutils.GetStackFilePaths(software)

	err := c.deployer.Remove(ctx, software.ProjectPath, url, software.Name, filePaths)
	return errors.Wrap(err, "failed to remove a stack")
}

func (c *ComposeStackManager) Stop(ctx context.Context, software docker.DocSoftware) error {
	//url, _, err := c.fetchEndpointProxy()
	url := software.Endpoint.EnvURL
	if software.Endpoint.IsSocket {
		url = ""
	}

	//if err := updateNetworkEnvFile(software); err != nil {
	//	return err
	//}

	//filePaths := stackutils.GetStackFilePaths(software)

	err := c.deployer.Stop(ctx, software.ProjectPath, url, software.Name, []string{
		global.FairFileSystem.GetProjectPath(software.Name),
	})
	return errors.Wrap(err, "failed to remove a stack")
}

func (c *ComposeStackManager) Start(ctx context.Context, software docker.DocSoftware) error {
	//url, _, err := c.fetchEndpointProxy()
	url := software.Endpoint.EnvURL
	if software.Endpoint.IsSocket {
		url = ""
	}

	if err := updateNetworkEnvFile(software); err != nil {
		return err
	}

	filePaths := stackutils.GetStackFilePaths(software)

	err := c.deployer.Start(ctx, software.ProjectPath, url, software.Name, filePaths)
	return errors.Wrap(err, "failed to remove a stack")
}

// Login executes the docker login command against a list of registries (including DockerHub).
func (c *ComposeStackManager) Login(registries []docker.DocRegistry, endpoint docker.DocEndpoint) error {
	command, args, err := c.prepareDockerCommandAndArgs(c.binaryPath, c.configPath, endpoint)
	if err != nil {
		return err
	}
	for _, registry := range registries {
		if registry.Authentication {
			registryArgs := append(args, "login", "--username", registry.Username, "--password", registry.Password, registry.URL)
			runCommandAndCaptureStdErr(command, registryArgs, nil, "")
		}
	}
	return nil
}

// Logout executes the docker logout command.
func (c *ComposeStackManager) Logout(endpoint docker.DocEndpoint) error {
	command, args, err := c.prepareDockerCommandAndArgs(c.binaryPath, c.configPath, endpoint)
	if err != nil {
		return err
	}
	args = append(args, "logout")
	return runCommandAndCaptureStdErr(command, args, nil, "")
}

func (c *ComposeStackManager) prepareDockerCommandAndArgs(binaryPath, configPath string, endpoint docker.DocEndpoint) (string, []string, error) {
	// Assume Linux as a default
	command := path.Join(binaryPath, "docker")

	if runtime.GOOS == "windows" {
		command = path.Join(binaryPath, "docker.exe")
	}

	args := make([]string, 0)
	// --config string      Location of client configuration files (default  "/Users/liuhuan/.docker")
	//args = append(args, "--config", configPath)

	if !strings.HasPrefix(endpoint.EnvURL, "tcp://") {
		endpoint.EnvURL = "tcp://" + endpoint.EnvURL
	}
	if endpoint.IsSocket {
		endpoint.EnvURL = "unix:///var/run/docker.sock"
		if runtime.GOOS == "windows" {
			endpoint.EnvURL = "npipe:////./pipe/docker_engine"
		}
	}

	endpointURL := endpoint.EnvURL
	//if endpoint.Type == portainer.EdgeAgentOnDockerEnvironment {
	//	tunnel, err := manager.reverseTunnelService.GetActiveTunnel(endpoint)
	//	if err != nil {
	//		return "", nil, err
	//	}
	//	endpointURL = fmt.Sprintf("tcp://127.0.0.1:%d", tunnel.Port)
	//}

	args = append(args, "-H", endpointURL)

	//if endpoint.IsTLS {
	//	args = append(args, "--tls")
	//
	//	if !endpoint.TLSConfig.TLSSkipVerify {
	//		args = append(args, "--tlsverify", "--tlscacert", endpoint.TLSConfig.TLSCACertPath)
	//	} else {
	//		args = append(args, "--tlscacert", "''")
	//	}
	//
	//	if endpoint.TLSConfig.TLSCertPath != "" && endpoint.TLSConfig.TLSKeyPath != "" {
	//		args = append(args, "--tlscert", endpoint.TLSConfig.TLSCertPath, "--tlskey", endpoint.TLSConfig.TLSKeyPath)
	//	}
	//}

	return command, args, nil
}

func fileNotExist(filePath string) bool {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return true
	}

	return false
}

func updateNetworkEnvFile(software docker.DocSoftware) error {
	envFilePath := path.Join(software.ProjectPath, ".env")
	stackFilePath := path.Join(software.ProjectPath, "stack.env")
	if fileNotExist(envFilePath) {
		if fileNotExist(stackFilePath) {
			return nil
		}

		flags := os.O_WRONLY | os.O_SYNC | os.O_CREATE
		envFile, err := os.OpenFile(envFilePath, flags, 0666)
		if err != nil {
			return err
		}

		defer envFile.Close()

		stackFile, err := os.Open(stackFilePath)
		if err != nil {
			return err
		}

		defer stackFile.Close()

		_, err = io.Copy(envFile, stackFile)
		return err
	}

	return nil
}

// NormalizeStackName returns a new stack name with unsupported characters replaced
func (c *ComposeStackManager) NormalizeStackName(name string) string {
	return stackNameNormalizeRegex.ReplaceAllString(strings.ToLower(name), "")
}

// ComposeSyntaxMaxVersion returns the maximum supported version of the docker compose syntax
func (c *ComposeStackManager) ComposeSyntaxMaxVersion() string {
	return "3.9"
}

// TODO: The second parameter is given toagentThe second parameter is given to，The second parameter is given to
func (c *ComposeStackManager) fetchEndpointProxy() (string, interface{}, error) {
	if strings.HasPrefix(global.FairConf.Docker.Url, "unix://") || strings.HasPrefix(global.FairConf.Docker.Url, "npipe://") {
		return "", nil, nil
	}

	//proxy, err := manager.proxyManager.CreateAgentProxyServer(endpoint)
	//if err != nil {
	//	return "", nil, err
	//}
	//
	//return fmt.Sprintf("tcp://127.0.0.1:%d", proxy.Port), proxy, nil
	return "", nil, nil
}

func createEnvFile(req docker.DocSoftware) (string, error) {
	// workaround for EE-1862. It will have to be removed when
	// docker/compose upgraded to v2.x.
	if err := createNetworkEnvFile(req); err != nil {
		return "", errors.Wrap(err, "failed to create network env file")
	}

	if req.Env == nil || len(req.Env) == 0 {
		return "", nil
	}

	envFilePath := path.Join(req.ProjectPath, "stack.env")

	envfile, err := os.OpenFile(envFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return "", err
	}

	for _, v := range req.Env {
		envfile.WriteString(fmt.Sprintf("%s=%s\n", v.Name, v.Value))
	}
	envfile.Close()

	return "stack.env", nil
}

func createNetworkEnvFile(software docker.DocSoftware) error {
	networkNameSet := NewStringSet()

	for _, filePath := range stackutils.GetStackFilePaths(software) {
		networkNames, err := extractNetworkNames(filePath)
		if err != nil {
			return errors.Wrap(err, "failed to extract network name")
		}

		if networkNames == nil || networkNames.Len() == 0 {
			continue
		}

		networkNameSet.Union(networkNames)
	}

	for _, s := range networkNameSet.List() {
		if _, ok := os.LookupEnv(s); ok {
			networkNameSet.Remove(s)
		}
	}

	if networkNameSet.Len() == 0 && software.Env == nil {
		return nil
	}

	envfile, err := os.OpenFile(path.Join(software.ProjectPath, ".env"),
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return errors.Wrap(err, "failed to open env file")
	}

	defer envfile.Close()

	var scanEnvSettingFunc = func(name string) (string, bool) {
		if software.Env != nil {
			for _, v := range software.Env {
				if name == v.Name {
					return v.Value, true
				}
			}
		}

		return "", false
	}

	for _, s := range networkNameSet.List() {
		if _, ok := scanEnvSettingFunc(s); !ok {
			software.Env = append(software.Env, docker.Pair{
				Name:  s,
				Value: "None",
			})
		}
	}

	if software.Env != nil {
		for _, v := range software.Env {
			envfile.WriteString(
				fmt.Sprintf("%s=%s\n", v.Name, v.Value))
		}
	}

	return nil
}

func extractNetworkNames(filePath string) (StringSet, error) {
	if info, err := os.Stat(filePath); errors.Is(err,
		os.ErrNotExist) || info.IsDir() {
		return nil, nil
	}

	stackFileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open yaml file")
	}

	config, err := loader.ParseYAML(stackFileContent)
	if err != nil {
		// invalid stack file
		return nil, errors.Wrap(err, "invalid stack file")
	}

	var version string
	if _, ok := config["version"]; ok {
		version, _ = config["version"].(string)
	}

	var networks map[string]interface{}
	if value, ok := config["networks"]; ok {
		if value == nil {
			return nil, nil
		}

		if networks, ok = value.(map[string]interface{}); !ok {
			return nil, nil
		}
	} else {
		return nil, nil
	}

	networkContent, err := loader.LoadNetworks(networks, version)
	if err != nil {
		return nil, nil // skip the error
	}

	re := regexp.MustCompile(`^\$\{?([^\}]+)\}?$`)
	networkNames := NewStringSet()

	for _, v := range networkContent {
		matched := re.FindAllStringSubmatch(v.Name, -1)
		if matched != nil && matched[0] != nil {
			if strings.Contains(matched[0][1], ":-") {
				continue
			}

			if strings.Contains(matched[0][1], "?") {
				continue
			}

			if strings.Contains(matched[0][1], "-") {
				continue
			}

			networkNames.Add(matched[0][1])
		}
	}

	if networkNames.Len() == 0 {
		return nil, nil
	}

	return networkNames, nil
}

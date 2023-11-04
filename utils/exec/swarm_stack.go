// Package exec
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-17 17:30
package exec

import (
	"bufio"
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/model/v2/docker"
	"cnic/fairman-gin/utils"
	"cnic/fairman-gin/utils/internal/stackutils"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

// SwarmStackManager Represents a service used to manage the stack。
type SwarmStackManager struct {
	binaryPath string
	configPath string
}

// NewSwarmStackManager Initialize newSwarmStackManagerInitialize new。
// It also updates the configuration of the Docker CLI binary.
func NewSwarmStackManager(
	binaryPath, configPath string,
) (*SwarmStackManager, error) {
	manager := &SwarmStackManager{
		binaryPath: binaryPath, // docker Run Path
		configPath: configPath, // TODO: docker configuration file configuration file
	}

	// TODO: changedocker config

	return manager, nil
}

// Login executes the docker login command against a list of registries (including DockerHub).
func (manager *SwarmStackManager) Login(registries []docker.DocRegistry, endpoint docker.DocEndpoint) error {
	command, args, err := manager.prepareDockerCommandAndArgs(manager.binaryPath, manager.configPath, endpoint)
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
func (manager *SwarmStackManager) Logout(endpoint docker.DocEndpoint) error {
	command, args, err := manager.prepareDockerCommandAndArgs(manager.binaryPath, manager.configPath, endpoint)
	if err != nil {
		return err
	}
	args = append(args, "logout")
	return runCommandAndCaptureStdErr(command, args, nil, "")
}

func (manager *SwarmStackManager) Ls(stack docker.DocSoftware) error {
	//if stack.Endpoint.IsSocket {
	//	dockerClient, err := global.FairClient.CreateClient(stack.Endpoint)
	//	if err != nil {
	//		return err
	//	}
	//
	//	defer dockerClient.Close()
	//
	//	inspect, err := dockerClient.SwarmInspect(context.TODO())
	//	if err != nil {
	//		return err
	//	}
	//
	//	fmt.Println(inspect)
	//	return nil
	//} else {
	command, args, err := manager.prepareDockerCommandAndArgs(manager.binaryPath, manager.configPath, stack.Endpoint)
	if err != nil {
		return err
	}

	args = append(args, "stack", "ls")

	return runCommand(command, args)
	//}
}

func (manager *SwarmStackManager) Init(stack docker.DocSoftware) error {
	fmt.Println("initialization")
	if err := manager.Ls(stack); err != nil {
		fmt.Println(err)
		fmt.Println("uninitialized")
		command, args, err := manager.prepareDockerCommandAndArgs(manager.binaryPath, manager.configPath, stack.Endpoint)
		if err != nil {
			return err
		}

		args = append(args, "swarm", "init")

		fmt.Println(command, args)
		return runCommand(command, args)
		//}
	}
	fmt.Println("Initialized")
	return nil
}

// Deploy implementdocker stack deployimplement。
// docker stack deploy [OPTIONS] STACK
// OPTIONS
// --compose-file , -c      =   (v1.25+)Compose Path to file，Path to file“-”
// --namespace              =   (k8s, Abandonment)Abandonment Kubernetes Abandonment
// --prune                  =   (v1.27+,swarm)Delete services that are no longer referenced
// --resolve-image	always  =   (v1.30+,swarm)Query the registry to parse image abstracts and supported platforms ("always"|"changed"|"never")
// --with-registry-auth     =   (swarm)towardsSwarmtowards
// --kubeconfig             =   (k8s， Abandonment)Kubernetes Abandonment
// --orchestrator           =   (Abandonment)Abandonment (swarm|kubernetes|all)
func (manager *SwarmStackManager) Deploy(stack docker.DocSoftware, prune bool) error {
	filePaths := stackutils.GetStackFilePaths(stack)
	command, args, err := manager.prepareDockerCommandAndArgs(manager.binaryPath, manager.configPath, stack.Endpoint)
	if err != nil {
		return utils.Errorf(err)
	}
	if prune {
		args = append(args, "stack", "deploy", "--prune", "--with-registry-auth")
	} else {
		args = append(args, "stack", "deploy", "--with-registry-auth")
	}

	args = configureFilePaths(args, filePaths)
	args = append(args, stack.Name)
	// ENV Temporarily useless
	env := make([]string, 0)
	for _, envvar := range stack.Env {
		env = append(env, envvar.Name+"="+envvar.Value)
	}
	fmt.Println(command)
	fmt.Println(args)
	fmt.Println(env)
	fmt.Println(stack.ProjectPath)
	return runCommandAndCaptureStdErr(command, args, env, stack.ProjectPath)
}

// Remove executes the docker stack rm command.
func (manager *SwarmStackManager) Remove(name string, endpoint docker.DocEndpoint) error {
	command, args, err := manager.prepareDockerCommandAndArgs(manager.binaryPath, manager.configPath, endpoint)
	if err != nil {
		return err
	}
	args = append(args, "stack", "rm", name)
	fmt.Println(command)
	fmt.Println(args)
	return runCommandAndCaptureStdErr(command, args, nil, "")
}

func runCommand(command string, args []string) error {
	cmd := exec.Command(command, args...)
	err := cmd.Run()
	if err != nil {
		//return utils.Errorf(stderr.String())
		return err
	}
	return nil
}

func runCommandAndCaptureStdErr(command string, args []string, env []string, workingDir string) error {
	//var stderr bytes.Buffer
	//var stdout bytes.Buffer
	cmd := exec.Command(command, args...)
	//cmd.Stdout = &stdout
	//cmd.Stderr = &stderr
	//cmd.Dir = workingDir

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(utils.Errorf(err))
	}

	go SendOutWsLog("success", stdout)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println(utils.Errorf(err))
	}

	go SendOutWsLog("error", stderr)

	if env != nil {
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, env...)
	}

	err = cmd.Run()
	if err != nil {
		return utils.Errorf(err)
		//return err
	}

	return nil
}

func SendOutWsLog(statusStr string, closer io.ReadCloser) {
	scanner := bufio.NewScanner(closer)
	const maxCapacity = 512 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	fmt.Println(statusStr, "Start reading download logs")
	for scanner.Scan() {
		text := scanner.Text()

		if text != "" {
			global.BaseChan <- global.ChanLogModal{
				Level:   statusStr,
				Message: text,
			}
		}
		fmt.Println(statusStr, scanner.Text())
	}
}

func (manager *SwarmStackManager) prepareDockerCommandAndArgs(binaryPath, configPath string, endpoint docker.DocEndpoint) (string, []string, error) {
	// Assume Linux as a default
	command := path.Join(binaryPath, "docker")

	if runtime.GOOS == "windows" {
		command = path.Join(binaryPath, "docker.exe")
	}

	args := make([]string, 0)
	// TODO: Use when installing on other servers，Use when installing on other servers
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

	// TODO: TLS The files will need to be downloaded at that time
	//if endpoint.TLSConfig.TLSSkipVerify {
	//	//args = append(args, "--tls")
	//
	//	args = append(args, "--tlsverify", "--tlscacert", endpoint.TLSConfig.TLSCACertPath)
	//
	//	if endpoint.TLSConfig.TLSCertPath != "" && endpoint.TLSConfig.TLSKeyPath != "" {
	//		args = append(args, "--tlscert", endpoint.TLSConfig.TLSCertPath, "--tlskey", endpoint.TLSConfig.TLSKeyPath)
	//	}
	//}

	return command, args, nil
}

// NormalizeStackName Standardized Service Name
func (manager *SwarmStackManager) NormalizeStackName(name string) string {
	return stackNameNormalizeRegex.ReplaceAllString(strings.ToLower(name), "")
}

func configureFilePaths(args []string, filePaths []string) []string {
	for _, path := range filePaths {
		args = append(args, "--compose-file", path)
	}
	return args
}

// Package composeplugin
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-15 15:43
package composeplugin

import (
	"bytes"
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/utils"
	composeerrors "cnic/fairman-gin/utils/exec/compose/errors"
	"context"

	"github.com/pkg/errors"

	"os"
	"os/exec"
	"path"
	"strings"
)

var (
	MissingDockerComposePluginErr = errors.New("docker-compose plugin is missing from config path")
)

// PluginWrapper provide a type for managing docker compose commands
type PluginWrapper struct {
	binaryPath string
	configPath string
}

// NewPluginWrapper initializes a new ComposeWrapper service with local docker-compose binary.
func NewPluginWrapper(binaryPath, configPath string) (global.ComposeDeployer, error) {
	if !utils.IsBinaryPresent(utils.ProgramPath(binaryPath, "docker")) {
		return nil, composeerrors.ErrBinaryNotFound
	}

	if configPath == "" {
		homePath, err := os.UserHomeDir()
		if err != nil {
			return nil, errors.WithMessage(err, "failed fetching user home directory")
		}
		configPath = path.Join(homePath, ".docker")
	}

	dockerPluginsPath := path.Join(configPath, "cli-plugins")
	if !utils.IsBinaryPresent(utils.ProgramPath(dockerPluginsPath, "docker-compose")) {
		pluginPath := utils.ProgramPath(binaryPath, "docker-compose.plugin")
		if !utils.IsBinaryPresent(pluginPath) {
			return nil, MissingDockerComposePluginErr
		}

		err := os.MkdirAll(dockerPluginsPath, 0755)
		if err != nil {
			return nil, errors.WithMessage(err, "failed creating plugins path")
		}

		err = utils.Copy(pluginPath, utils.ProgramPath(dockerPluginsPath, "docker-compose"))
		if err != nil {
			return nil, err
		}

	}

	return &PluginWrapper{binaryPath: binaryPath, configPath: configPath}, nil
}

// Deploy Up create and start containers
func (wrapper *PluginWrapper) Deploy(ctx context.Context, workingDir, host, projectName string, filePaths []string, envFilePath string, isLog bool) error {
	output, err := wrapper.command(newUpCommand(filePaths), workingDir, host, projectName, envFilePath)
	if len(output) != 0 {
		return utils.Errorf("[libstack,composebinary] [message: finish deploying] [output: %s] [err: %s]", output, err)
	}

	return err
}

// Remove Down stop and remove containers
func (wrapper *PluginWrapper) Remove(ctx context.Context, workingDir, host, projectName string, filePaths []string) error {
	output, err := wrapper.command(newDownCommand(filePaths), workingDir, host, projectName, "")
	if len(output) != 0 {
		return utils.Errorf("[libstack,composebinary] [message: finish deploying] [output: %s] [err: %s]", output, err)
	}

	return err
}

// Stop containers
func (wrapper *PluginWrapper) Stop(ctx context.Context, workingDir, host, projectName string, filePaths []string) error {
	output, err := wrapper.command(newStopCommand(filePaths), workingDir, host, projectName, "")
	if len(output) != 0 {
		return utils.Errorf("[libstack,composebinary] [message: finish deploying] [output: %s] [err: %s]", output, err)
	}

	return err
}

// Start containers
func (wrapper *PluginWrapper) Start(ctx context.Context, workingDir, host, projectName string, filePaths []string) error {
	output, err := wrapper.command(newStartCommand(filePaths), workingDir, host, projectName, "")
	if len(output) != 0 {
		return utils.Errorf("[libstack,composebinary] [message: finish deploying] [output: %s] [err: %s]", output, err)
	}

	return err
}

// command Command exectue a docker-compose commanåd
func (wrapper *PluginWrapper) command(command composeCommand, workingDir, url, projectName, envFilePath string) ([]byte, error) {
	program := utils.ProgramPath(wrapper.binaryPath, "docker")

	if projectName != "" {
		command.WithProjectName(projectName)
	}

	if envFilePath != "" {
		command.WithEnvFilePath(envFilePath)
	}

	var stderr bytes.Buffer

	args := []string{}

	if wrapper.configPath != "" {
		args = append(args, "--config", wrapper.configPath)
	}

	if url != "" {
		args = append(args, "-H", url)
	}

	args = append(args, "compose")
	args = append(args, command.ToArgs()...)

	cmd := exec.Command(program, args...)
	cmd.Dir = workingDir

	cmd.Stderr = &stderr

	output, err := cmd.Output()
	if err != nil {
		return nil, errors.New(stderr.String())
	}

	return output, nil
}

type composeCommand struct {
	command []string
	args    []string
}

func newCommand(command []string, filePaths []string) composeCommand {
	var args []string
	for _, path := range filePaths {
		args = append(args, "-f")
		args = append(args, strings.TrimSpace(path))
	}
	return composeCommand{
		args:    args,
		command: command,
	}
}

func newUpCommand(filePaths []string) composeCommand {
	return newCommand([]string{"up", "-d"}, filePaths)
}

func newDownCommand(filePaths []string) composeCommand {
	return newCommand([]string{"down", "--remove-orphans"}, filePaths)
}

func newStopCommand(filePaths []string) composeCommand {
	return newCommand([]string{"stop"}, filePaths)
}

func newStartCommand(filePaths []string) composeCommand {
	return newCommand([]string{"start"}, filePaths)
}

func (command *composeCommand) WithConfigPath(configPath string) {
	command.args = append(command.args, "--config", configPath)
}

func (command *composeCommand) WithProjectName(projectName string) {
	command.args = append(command.args, "-p", projectName)
}

func (command *composeCommand) WithEnvFilePath(envFilePath string) {
	command.args = append(command.args, "--env-file", envFilePath)
}

func (command *composeCommand) WithURL(url string) {
	command.args = append(command.args, "-H", url)
}

func (command *composeCommand) ToArgs() []string {
	return append(command.args, command.command...)
}

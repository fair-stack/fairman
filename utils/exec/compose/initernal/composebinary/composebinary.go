// Package composebinary
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-15 15:33
package composebinary

import (
	"bufio"
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/utils"
	composeerrors "cnic/fairman-gin/utils/exec/compose/errors"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// ComposeWrapper provide a type for managing docker compose commands
type ComposeWrapper struct {
	binaryPath string
	configPath string
}

// NewComposeWrapper initializes a new ComposeWrapper service with local docker-compose binary.
func NewComposeWrapper(binaryPath, configPath string) (global.ComposeDeployer, error) {
	if !utils.IsBinaryPresent(utils.ProgramPath(binaryPath, "docker-compose")) {
		return nil, composeerrors.ErrBinaryNotFound
	}

	return &ComposeWrapper{binaryPath: binaryPath, configPath: configPath}, nil
}

// Deploy creates and starts containers
func (wrapper *ComposeWrapper) Deploy(ctx context.Context, workingDir, host, projectName string, filePaths []string, envFilePath string, isLog bool) error {
	output, err := wrapper.Command(newUpCommand(filePaths), workingDir, host, projectName, envFilePath, isLog)
	fmt.Println(output, err)
	if len(output) != 0 {
		return utils.Errorf("[libstack,composebinary] [message: finish deploying] [output: %s] [err: %s]", output, err)
	}
	fmt.Println(err)

	return err
}

// Remove stops and removes containers
func (wrapper *ComposeWrapper) Remove(ctx context.Context, workingDir, host, projectName string, filePaths []string) error {
	output, err := wrapper.Command(newDownCommand(filePaths), workingDir, host, projectName, "", false)
	if len(output) != 0 {
		return utils.Errorf("[libstack,composebinary] [message: finish deploying] [output: %s] [err: %s]", output, err.Error())
	}

	return err
}

func (wrapper *ComposeWrapper) Stop(ctx context.Context, workingDir, host, projectName string, filePaths []string) error {
	output, err := wrapper.Command(newStopCommand(filePaths), workingDir, host, projectName, "", false)
	if len(output) != 0 {
		return utils.Errorf("[libstack,composebinary] [message: finish deploying] [output: %s] [err: %s]", output, err.Error())
	}

	return err
}

func (wrapper *ComposeWrapper) Start(ctx context.Context, workingDir, host, projectName string, filePaths []string) error {
	output, err := wrapper.Command(newStartCommand(filePaths), workingDir, host, projectName, "", false)
	if len(output) != 0 {
		return utils.Errorf("[libstack,composebinary] [message: finish deploying] [output: %s] [err: %s]", output, err.Error())
	}

	return err
}

// Command executes a docker-compose command
func (wrapper *ComposeWrapper) Command(command composeCommand, workingDir, host, projectName, envFilePath string, isLog bool) ([]byte, error) {
	program := utils.ProgramPath(wrapper.binaryPath, "docker-compose")

	if projectName != "" {
		command.WithProjectName(projectName)
	}

	if envFilePath != "" {
		command.WithEnvFilePath(envFilePath)
	}

	if host != "" {
		command.WithHost(host)
	}

	//var stderr bytes.Buffer
	cmd := exec.Command(program, command.ToArgs()...)
	//cmd.Dir = workingDir

	if wrapper.configPath != "" {
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, fmt.Sprintf("DOCKER_CONFIG=%s", wrapper.configPath))
	}

	//cmd.Stderr = &stderr

	//output, err := cmd.Output()
	//if err != nil {
	//	return nil, errors.Wrap(err, stderr.String())
	//}

	if isLog {
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return nil, err
		}

		stderr, err := cmd.StderrPipe()
		if err != nil {
			return nil, err
		}

		if err := cmd.Start(); err != nil {
			return nil, err
		}

		go wrapper.SendOutWsLog("success", stdout)
		go wrapper.SendOutWsLog("error", stderr)

		if err := cmd.Wait(); err != nil {
			fmt.Println(workingDir, host, projectName)
			return nil, err
		}
	} else {
		output, err := cmd.Output()
		if err != nil {
			return nil, err
		}
		return output, nil
	}

	return []byte(""), nil
}

func (wrapper *ComposeWrapper) SendOutWsLog(statusStr string, closer io.ReadCloser) {
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
		fmt.Println(statusStr, text)
	}
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
	return newCommand([]string{"up", "-d", "--force-recreate"}, filePaths)
}

func newDownCommand(filePaths []string) composeCommand {
	return newCommand([]string{"down", "--remove-orphans"}, filePaths)
}

func newStopCommand(filePats []string) composeCommand {
	return newCommand([]string{"stop"}, filePats)
}

func newStartCommand(filePats []string) composeCommand {
	return newCommand([]string{"start"}, filePats)
}

func (command *composeCommand) WithProjectName(projectName string) {
	command.args = append(command.args, "-p", projectName)
}

func (command *composeCommand) WithEnvFilePath(envFilePath string) {
	command.args = append(command.args, "--env-file", envFilePath)
}

func (command *composeCommand) WithHost(host string) {
	command.args = append(command.args, "-H", host)
}

func (command *composeCommand) ToArgs() []string {
	return append(command.args, command.command...)
}

// Package exec
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-09-20 15:00
package exec

import (
	"bufio"
	"bytes"
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/utils"
	"context"
	"fmt"
	"io"
	"os/exec"
	"path"
	"runtime"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/pkg/errors"
)

// KubernetesDeployer Represented inKubernetesRepresented in（Represented in）Represented in。
type KubernetesDeployer struct {
	binaryPath   string
	configPath   string
	endpointName string
}

// NewKubernetesDeployer initializes a new KubernetesDeployer service.
func NewKubernetesDeployer(binaryPath string, configPath string, endpointName string) *KubernetesDeployer {
	return &KubernetesDeployer{
		binaryPath:   binaryPath,
		configPath:   configPath,
		endpointName: endpointName,
	}
}

// NormalizeStackName returns a new stack name with unsupported characters replaced
func (deployer *KubernetesDeployer) NormalizeStackName(name string) string {
	return stackNameNormalizeRegex.ReplaceAllString(strings.ToLower(name), "")
}

// CreateNamespace create namespace
func (deployer *KubernetesDeployer) CreateNamespace(namespace string) error {
	// kubectl create namespace fairman
	return deployer.commandCreateNamespace("create", namespace)
}

// Deploy upserts Kubernetes resources defined in manifest(s)
func (deployer *KubernetesDeployer) Deploy(namespace string, manifestFiles []string, isLog bool) error {
	return deployer.command("apply", namespace, manifestFiles, isLog)
}

// Remove deletes Kubernetes resources defined in manifest(s)
func (deployer *KubernetesDeployer) Remove(manifestFiles []string, namespace string) error {
	return deployer.command("delete", namespace, manifestFiles, false)
}

// Version returns the version of the Kubernetes cluster
func (deployer *KubernetesDeployer) Version() error {
	return deployer.command("version", "", nil, false)
}

func (deployer *KubernetesDeployer) command(operation string, namespace string, manifestFiles []string, isLog bool) error {
	command := path.Join(deployer.binaryPath, "kubectl")
	if runtime.GOOS == "windows" {
		command = path.Join(deployer.binaryPath, "kubectl.exe")
	}

	args := make([]string, 0)
	// args := []string{"--token", token}

	args = append(args, "--kubeconfig", path.Join("data", "config", deployer.endpointName, "config"))

	if namespace != "" {
		args = append(args, "--namespace", namespace)
	}

	if operation == "delete" {
		args = append(args, "--ignore-not-found=true")
	}

	args = append(args, operation)
	for _, file := range manifestFiles {
		args = append(args, "-f", strings.TrimSpace(file))
	}

	if operation == "apply" {
		//args = append(args, "--force")
		args = append(args, "--record")
	}
	fmt.Println(command, args)
	//var stderr bytes.Buffer
	cmd := exec.Command(command, args...)

	if isLog {
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return utils.Errorf(err)
		}

		stderr, err := cmd.StderrPipe()
		if err != nil {
			return utils.Errorf(err)
		}

		if err := cmd.Start(); err != nil {
			return utils.Errorf(err)
		}

		go deployer.SendOutWsLog("success", stdout)
		go deployer.SendOutWsLog("error", stderr)

		if err := cmd.Wait(); err != nil {
			//fmt.Println(workingDir, host, projectName)
			return utils.Errorf(err)
		}
	} else {
		_, err := cmd.Output()
		if err != nil {
			return utils.Errorf(err)
		}
		return utils.Errorf(err)
	}

	//cmd.Stderr = &stderr
	//
	//output, err := cmd.Output()
	//if err != nil {
	//	global.BaseChan <- global.ChanLogModal{
	//		Level:   "error",
	//		Message: stderr.String(),
	//	}
	//	return utils.Errorf(errors.Wrapf(err, "failed to execute kubectl command: %q", stderr.String()))
	//}
	//
	//fmt.Println(string(output))

	return nil
}

func (deployer *KubernetesDeployer) SendOutWsLog(statusStr string, closer io.ReadCloser) {
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

func (deployer *KubernetesDeployer) commandCreateNamespace(operation string, namespace string) error {
	command := path.Join(deployer.binaryPath, "kubectl")
	if runtime.GOOS == "windows" {
		command = path.Join(deployer.binaryPath, "kubectl.exe")
	}

	args := make([]string, 0)
	// args := []string{"--token", token}
	kubeconfig := path.Join("data", "config", deployer.endpointName, "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return utils.Errorf(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return utils.Errorf(err)
	}

	_, err = clientset.CoreV1().Namespaces().Get(context.Background(), namespace, metav1.GetOptions{})
	if err != nil {
		args = append(args, "--kubeconfig", kubeconfig)

		args = append(args, operation, "ns", namespace)
		fmt.Println(command, args)
		var stderr bytes.Buffer
		cmd := exec.Command(command, args...)
		cmd.Stderr = &stderr

		output, err := cmd.Output()
		if err != nil {
			return utils.Errorf(errors.Wrapf(err, "failed to execute kubectl command: %q", stderr.String()))
		}

		fmt.Println(string(output))
		return err
	}

	return nil
}

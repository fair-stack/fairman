// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-12-09 07:35
package docker

import (
	"bytes"
	requestV2 "cnic/fairman-gin/model/v2/docker/request"
	"cnic/fairman-gin/utils"
	"fmt"
	"io"
	"path"
	"strings"

	"k8s.io/client-go/tools/remotecommand"

	responseDoc "cnic/fairman-gin/model/v2/docker/response"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/gin-gonic/gin"
)

type KubernetesService struct {
}

// GetPodList obtaink8sPodobtain
func (k *KubernetesService) GetPodList(req requestV2.PodListRequest, ctx *gin.Context) ([]responseDoc.PodListResponse, error) {
	var podList = make([]responseDoc.PodListResponse, 0)
	kubeconfig := path.Join("data", "config", req.Endpoint.Name, "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, utils.Errorf(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, utils.Errorf(err)
	}

	namespace := req.Name
	// TODO: datalabbaas
	pods, err := clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{
		LabelSelector: fmt.Sprintf("app=%s", req.Name+"baas"),
	})

	if err != nil {
		return nil, utils.Errorf(err)
	}

	for _, item := range pods.Items {
		for _, container := range item.Spec.Containers {
			fmt.Println(item.Name, item.Namespace)
			pod := responseDoc.PodListResponse{
				Namespace:     item.Namespace,
				PodName:       item.Name,
				ContainerName: container.Name,
				Status:        item.Status.Phase,
				CreatedAt:     item.CreationTimestamp.Format("2006-01-02 15:04:05"),
			}

			podList = append(podList, pod)
		}
	}
	return podList, nil
}

// GetPodListDetails obtaink8sPodobtain
func (k *KubernetesService) GetPodListDetails(req requestV2.PodListRequest, ctx *gin.Context) (*corev1.PodList, error) {
	kubeconfig := path.Join("data", "config", req.Endpoint.Name, "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, utils.Errorf(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, utils.Errorf(err)
	}

	namespace := req.Name
	// TODO: datalabbaas
	pods, err := clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{
		LabelSelector: fmt.Sprintf("app=%s", req.Name+"baas"),
	})

	if err != nil {
		return nil, utils.Errorf(err)
	}

	return pods, nil
}

// PodTerminal k8sPodterminal
func (k *KubernetesService) PodTerminal(req requestV2.PodTerminalRequest, ctx *gin.Context) error {
	pty, err := NewTerminalSession(ctx)
	if err != nil {
		return utils.Errorf(err)
	}

	kubeconfig := path.Join("data", "config", req.Endpoint.Name, "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return utils.Errorf(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return utils.Errorf(err)
	}

	namespace := req.Namespace
	podName := req.PodName
	containerName := req.Container

	reqExec := clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec")
	reqExec.VersionedParams(&corev1.PodExecOptions{
		Container: containerName,
		Command:   []string{"/bin/sh"},
		Stdin:     !(pty.Stdin() == nil),
		Stdout:    !(pty.Stdout() == nil),
		Stderr:    !(pty.Stderr() == nil),
		TTY:       pty.Tty(),
	}, scheme.ParameterCodec)

	executor, err := remotecommand.NewSPDYExecutor(config, "POST", reqExec.URL())
	if err != nil {
		return utils.Errorf(err)
	}
	err = executor.Stream(remotecommand.StreamOptions{
		Stdin:             pty.Stdin(),
		Stdout:            pty.Stdout(),
		Stderr:            pty.Stderr(),
		TerminalSizeQueue: pty,
		Tty:               pty.Tty(),
	})
	fmt.Println(utils.Errorf(err))
	return err

	// adoptapiadoptpodadoptterminal
	//
	// 1. obtainpodobtainexec
	// 2. obtainpodobtainattach
	// 3. obtainpodobtainportforward
}

// PodLog k8sPodjournal
func (k *KubernetesService) PodLog(req requestV2.PodLogRequest, ctx *gin.Context) (logs []string, err error) {
	fmt.Println("#######################################")
	kubeconfig := path.Join("data", "config", req.Endpoint.Name, "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, utils.Errorf(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, utils.Errorf(err)
	}

	namespace := req.Namespace
	podName := req.PodName
	containerName := req.ContainerName

	podLogOpts := corev1.PodLogOptions{
		Container: containerName,
		Follow:    false,
	}

	reqLog := clientset.CoreV1().Pods(namespace).GetLogs(podName, &podLogOpts)
	podLogs, err := reqLog.Stream(ctx)
	if err != nil {
		return nil, utils.Errorf(err)
	}
	defer podLogs.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return nil, utils.Errorf(err)
	}
	str := buf.String()
	logs = strings.Split(str, "\n")
	return logs, nil
}

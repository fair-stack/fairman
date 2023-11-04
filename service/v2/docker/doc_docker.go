// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-29 15:11
package docker

import (
	"bytes"
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/model/v2/docker"
	"cnic/fairman-gin/model/v2/docker/request"
	"cnic/fairman-gin/utils"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/swarm"

	"github.com/docker/docker/api/types"
)

type DockerService struct {
}

// ContainerList Obtain all software containers
func (d *DockerService) ContainerList(req request.ContainerListRequest) (interface{}, error) {
	if val, ok := docker.EndpointMap[req.EndpointId]; !ok {
		return nil, utils.Errorf("endpointId: %s not found", req.EndpointId)
	} else {
		req.Endpoint = val
	}

	dockerClient, err := global.FairClient.CreateClient(req.Endpoint)
	if err != nil {
		return nil, err
	}

	defer dockerClient.Close()

	var filter = filters.NewArgs()
	// TODO: Cannot use multiplelabelCannot use multiple
	//filter.Add("label", fmt.Sprintf("com.docker.container.namespace=%s", req.Name))
	filter.Add("label", fmt.Sprintf("%s=%s", ProjectLabel, strings.ToLower(req.Name)))
	//filter.Add("label", fmt.Sprintf("com.docker.stack.namespace=%s", strings.ToLower(req.Name)))
	containers, err := dockerClient.ContainerList(context.Background(), types.ContainerListOptions{Filters: filter, All: true})
	if err != nil {
		return nil, err
	}

	// Change the creation time to the difference between the creation time and the current time
	for i := range containers {
		containers[i].Created = time.Now().Unix() - containers[i].Created
	}

	return containers, nil
}

// ContainerTerminal obtaindockerobtain
func (d *DockerService) ContainerTerminal(req request.ContainerTerminalRequest) (hr types.HijackedResponse, err error) {
	if val, ok := docker.EndpointMap[req.EndpointId]; !ok {
		return hr, utils.Errorf("endpointId: %s not found", req.EndpointId)
	} else {
		req.Endpoint = val
	}

	dockerClient, err := global.FairClient.CreateClient(req.Endpoint)
	if err != nil {
		return
	}
	defer dockerClient.Close()

	ctx := context.Background()

	hr, err = dockerClient.ContainerExecAttach(ctx, req.ExecId, types.ExecStartCheck{
		Detach: false, Tty: true,
	})

	if err != nil {
		return hr, utils.Errorf(err)
	}

	if req.Width != 0 && req.Height != 0 {
		dockerClient.ContainerExecResize(ctx, req.ExecId, types.ResizeOptions{
			Height: req.Height,
			Width:  req.Width,
		})
	}
	return
}

// ContainerExecCreate Create adockerCreate a
func (d *DockerService) ContainerExecCreate(req request.ContainerExecCreateRequest) (execId string, err error) {
	if val, ok := docker.EndpointMap[req.EndpointId]; !ok {
		return execId, utils.Errorf("endpointId: %s not found", req.EndpointId)
	} else {
		req.Endpoint = val
	}

	dockerClient, err := global.FairClient.CreateClient(req.Endpoint)
	if err != nil {
		return "", utils.Errorf(err)
	}
	defer dockerClient.Close()
	ctx := context.Background()
	idRep, err := dockerClient.ContainerExecCreate(ctx, req.ContainerId, types.ExecConfig{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          []string{req.Cmd},
		Tty:          true,
	})
	if err != nil {
		return "", utils.Errorf(err)
	}
	return idRep.ID, nil
}

// ContainerExecResize changedockerchange
func (d *DockerService) ContainerExecResize(req request.ContainerExecResizeRequest) (err error) {
	if val, ok := docker.EndpointMap[req.EndpointId]; !ok {
		return utils.Errorf("endpointId: %s not found", req.EndpointId)
	} else {
		req.Endpoint = val
	}

	dockerClient, err := global.FairClient.CreateClient(req.Endpoint)
	if err != nil {
		return
	}
	defer dockerClient.Close()

	ctx := context.Background()
	err = dockerClient.ContainerExecResize(ctx, req.ExecId, types.ResizeOptions{
		Height: req.Height,
		Width:  req.Width,
	})
	if err != nil {
		return err
	}
	return
}

// ContainerLogs obtaindockerobtain
func (d *DockerService) ContainerLogs(req request.ContainerLogsRequest) (io.ReadCloser, error) {
	if val, ok := docker.EndpointMap[req.EndpointId]; !ok {
		return nil, utils.Errorf("endpointId: %s not found", req.EndpointId)
	} else {
		req.Endpoint = val
	}
	dockerClient, err := global.FairClient.CreateClient(req.Endpoint)
	if err != nil {
		return nil, err
	}
	defer dockerClient.Close()

	ctx := context.Background()

	return dockerClient.ContainerLogs(ctx, req.ContainerId, types.ContainerLogsOptions{
		//Details:    true,
		ShowStderr: true,
		ShowStdout: true,
		//Since:      "0",
		//Timestamps: req.Timestamps,
		Timestamps: true,
		Follow:     true,
		Tail:       req.Tail,
	})
}

// ContainerInspect Obtain detailed information about the container
func (d *DockerService) ContainerInspect(req request.ContainerLogsInspect, ctx *gin.Context) (inspect types.ContainerJSON, err error) {
	dockerClient, err := global.FairClient.CreateClient(req.Endpoint)
	if err != nil {
		return inspect, err
	}
	defer dockerClient.Close()

	inspect, err = dockerClient.ContainerInspect(ctx, req.ContainerId)
	if err != nil {
		return inspect, utils.Errorf(err)
	}

	return inspect, err
}

// ServiceList obtaindockerobtain
func (d *DockerService) ServiceList(req request.ServiceListRequest) (services []swarm.Service, err error) {
	dockerClient, err := global.FairClient.CreateClient(req.Endpoint)
	if err != nil {
		return nil, utils.Errorf(err)
	}
	defer dockerClient.Close()

	ctx := context.Background()
	filter := filters.NewArgs(filters.Arg("label", fmt.Sprintf("%s=%s", "com.docker.stack.namespace", req.Name)))
	services, err = dockerClient.ServiceList(ctx, types.ServiceListOptions{
		Filters: filter,
		Status:  true,
	})
	if err != nil {
		return nil, utils.Errorf(err)
	}

	for _, service := range services {
		tasks, err := dockerClient.TaskList(ctx, types.TaskListOptions{
			Filters: filters.NewArgs(filters.Arg("service", service.Spec.Name)),
		})

		if err != nil {
			return nil, utils.Errorf(err)
		}

		for _, task := range tasks {
			logs, err := dockerClient.TaskLogs(ctx, task.ID, types.ContainerLogsOptions{Details: true, ShowStderr: true,
				ShowStdout: true})
			if err != nil {
				return nil, utils.Errorf(err)
			}

			buf := new(bytes.Buffer)
			_, err = buf.ReadFrom(logs)
		}
	}
	return
}

// ServiceLogs obtaindockerobtain
func (d *DockerService) ServiceLogs(req request.ServiceLogsRequest) ([]string, error) {

	if val, ok := docker.EndpointMap[req.EndpointId]; !ok {
		return nil, utils.Errorf("endpointId: %s not found", req.EndpointId)
	} else {
		req.Endpoint = val
	}

	dockerClient, err := global.FairClient.CreateClient(req.Endpoint)
	if err != nil {
		return nil, err
	}
	defer dockerClient.Close()

	ctx := context.Background()

	logs, err := dockerClient.ServiceLogs(ctx, req.ServiceId, types.ContainerLogsOptions{
		//Details:    true,
		ShowStderr: true,
		ShowStdout: true,
		//Since:      "0",
		Timestamps: req.Timestamps,
		Tail:       req.Tail,
	})

	if err != nil {
		return nil, err
	}
	var logList []string

	hdr := make([]byte, 8)
	for {
		_, err := logs.Read(hdr)
		if err != nil || err == io.EOF {
			break
		}
		//var w io.Writer
		//fmt.Println(string(hdr))
		//switch hdr[0] {
		//case 1:
		//	w = os.Stdout
		//default:
		//	w = os.Stderr
		//}
		count := binary.BigEndian.Uint32(hdr[4:])
		dat := make([]byte, count)
		_, err = logs.Read(dat)
		if err != nil {
			break
		}

		logList = append(logList, string(dat))
	}

	return logList, nil
}

// ServiceInspect obtaindockerobtain
func (d *DockerService) ServiceInspect(req request.ServiceInspectRequest) (service swarm.Service, err error) {
	dockerClient, err := global.FairClient.CreateClient(req.Endpoint)
	if err != nil {
		return
	}
	defer dockerClient.Close()

	ctx := context.Background()
	service, _, err = dockerClient.ServiceInspectWithRaw(ctx, req.ServiceId, types.ServiceInspectOptions{})
	if err != nil {
		return
	}
	return
}

// UpdateService modifyServicemodify
func (d *DockerService) UpdateService(req request.UpdateServiceRequest) error {
	dockerClient, err := global.FairClient.CreateClient(req.Endpoint)
	if err != nil {
		return utils.Errorf(err)
	}
	defer dockerClient.Close()

	_, err = dockerClient.ServiceUpdate(
		context.TODO(),
		req.Service.ID,
		req.Service.Version,
		req.Service.Spec,
		types.ServiceUpdateOptions{},
	)
	if err != nil {
		return utils.Errorf(err)
	}
	return nil
}

// ServiceRunningTime Run time
func (d *DockerService) ServiceRunningTime(req request.ServiceListRequest) (timeStr float64, err error) {
	newTime := time.Now()
	services, err := d.ServiceList(req)
	if err != nil {
		return 0, err
	}

	for _, service := range services {
		if time.Now().After(service.CreatedAt) {
			newTime = service.CreatedAt
		}
	}
	return time.Now().Sub(newTime).Seconds(), err
}

// TaskList obtaindockerobtain
func (d *DockerService) TaskList(req request.TaskListRequest) (tasks []swarm.Task, err error) {
	dockerClient, err := global.FairClient.CreateClient(req.Endpoint)
	if err != nil {
		return nil, utils.Errorf(err)
	}
	defer dockerClient.Close()

	ctx := context.Background()
	filter := filters.NewArgs()
	if req.Name != "" {
		filter.Add("service", req.Name)
	}
	tasks, err = dockerClient.TaskList(ctx, types.TaskListOptions{
		Filters: filter,
	})
	if err != nil {
		return nil, utils.Errorf(err)
	}
	return
}

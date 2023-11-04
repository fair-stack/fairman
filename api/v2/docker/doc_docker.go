// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-29 15:00
package docker

import (
	"cnic/fairman-gin/model/v2/common/response"
	"cnic/fairman-gin/model/v2/docker/request"
	"cnic/fairman-gin/utils"
	"cnic/fairman-gin/utils/socket"
	"fmt"

	"github.com/gin-gonic/gin"
)

type DockerApi struct {
}

// ContainerList Obtain all software containers
func (d *DockerApi) ContainerList(c *gin.Context) {
	var service request.ContainerListRequest
	if errStr, err := utils.BaseValidatorQuery(&service, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	containers, err := dockerService.ContainerList(service)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(containers, c)
}

// ContainerTerminal obtaindockerobtain
func (d *DockerApi) ContainerTerminal(c *gin.Context) {
	var req request.ContainerTerminalRequest
	if errStr, err := utils.BaseValidatorQuery(&req, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	ws, err := socket.Upgrade(c)
	if err != nil {
		response.FailWithMessage("websocketinitialization failed！", c)
		return
	}

	defer func() {
		_ = ws.Close()
	}()

	hr, err := dockerService.ContainerTerminal(req)
	if err != nil {
		fmt.Println(err)
		ws.Close()
		return
	}
	defer hr.Close()

	go func() {
		socket.WsWriter(hr.Conn, ws)
	}()

	socket.WsReader(ws, hr.Conn)
}

// ContainerExecCreate Create adockerCreate a
func (d *DockerApi) ContainerExecCreate(c *gin.Context) {
	var req request.ContainerExecCreateRequest
	if errStr, err := utils.BaseValidator(&req, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	execId, err := dockerService.ContainerExecCreate(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(execId, c)
}

// ContainerExecResize changedockerchange
func (d *DockerApi) ContainerExecResize(c *gin.Context) {
	var req request.ContainerExecResizeRequest
	if errStr, err := utils.BaseValidator(&req, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	if err := dockerService.ContainerExecResize(req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// ContainerLogs obtaindockerobtain
func (d *DockerApi) ContainerLogs(c *gin.Context) {
	var req request.ContainerLogsRequest
	if errStr, err := utils.BaseValidatorQuery(&req, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	ws, err := socket.Upgrade(c)
	if err != nil {
		response.FailWithMessage("websocketinitialization failed！", c)
		return
	}

	defer func() {
		_ = ws.Close()
	}()

	logs, err := dockerService.ContainerLogs(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	go socket.WsWriterLog(ws, logs)

	socket.WsPing(ws)

	fmt.Println("websocketIt's over")
}

// ContainerInspect Obtain detailed information about the container
func (d *DockerApi) ContainerInspect(c *gin.Context) {
	var req request.ContainerLogsInspect
	if errStr, err := utils.BaseValidatorQuery(&req, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	inspect, err := dockerService.ContainerInspect(req, c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(inspect, c)
}

// ContainerInspectUpdate Modify container information
func (d *DockerApi) ContainerInspectUpdate(c *gin.Context) {
	//var req request.ContainerInspectUpdate
	//if errStr, err := utils.BaseValidator(&req, c); err != nil {
	//	response.FailWithMessage(errStr, c)
	//	return
	//}
	//
	//if err := dockerService.ContainerInspectUpdate(req); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}

	response.Ok(c)
}

// ServiceList obtaindockerobtain
func (d *DockerApi) ServiceList(c *gin.Context) {
	var req request.ServiceListRequest
	if errStr, err := utils.BaseValidatorQuery(&req, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	services, err := dockerService.ServiceList(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(services, c)
}

// ServiceLogs obtaindockerobtain
func (d *DockerApi) ServiceLogs(c *gin.Context) {
	var req request.ServiceLogsRequest
	if errStr, err := utils.BaseValidatorQuery(&req, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	logs, err := dockerService.ServiceLogs(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(logs, c)
}

// ServiceInspect obtaindockerobtain
func (d *DockerApi) ServiceInspect(c *gin.Context) {
	var req request.ServiceInspectRequest
	if errStr, err := utils.BaseValidatorQuery(&req, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	service, err := dockerService.ServiceInspect(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(service, c)
}

// UpdateService updatedockerupdate
func (d *DockerApi) UpdateService(c *gin.Context) {
	var req request.UpdateServiceRequest
	if errStr, err := utils.BaseValidator(&req, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	if err := dockerService.UpdateService(req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// ServiceRunningTime obtaindockerobtain
func (d *DockerApi) ServiceRunningTime(c *gin.Context) {
	var req request.ServiceListRequest
	if errStr, err := utils.BaseValidatorQuery(&req, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	runningTime, err := dockerService.ServiceRunningTime(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(runningTime, c)
}

// TaskList obtaindockerobtain
func (d *DockerApi) TaskList(c *gin.Context) {
	var req request.TaskListRequest
	if errStr, err := utils.BaseValidatorQuery(&req, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	tasks, err := dockerService.TaskList(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(tasks, c)
}

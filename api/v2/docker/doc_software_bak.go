// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-10 16:28
package docker

//import (
//	"cnic/fairman-gin/global"
//	"cnic/fairman-gin/model/v2/common/response"
//	"cnic/fairman-gin/model/v2/docker"
//	docRequst "cnic/fairman-gin/model/v2/docker/request"
//	"cnic/fairman-gin/utils"
//	"cnic/fairman-gin/utils/filesystem"
//	"fmt"
//
//	"github.com/gin-gonic/gin"
//)
//
//type SoftwareApi struct {
//}
//
////// GetSoftware query
////// @Tags SoftwareApi
////// @Summary Query softwareApi
////// @Security ApiKeyAuth
////// @accept application/json
////// @Produce application/json
////// @Success 200 {string} string "{"code":0,"data":{},"msg":"Operation successful"}"
////// @Router /v2/software [get]
////func (s *SoftwareApi) GetSoftware(c *gin.Context) {
////	software, err := softwareService.GetSoftware(c)
////	if err != nil {
////		response.FailWithMessage(err.Error(), c)
////		return
////	}
////
////	response.OkWithData(software, c)
////}
////
////// GetSoftwareDetails In the software detailscontainerIn the software detailstaskIn the software details
////func (s *SoftwareApi) GetSoftwareDetails(c *gin.Context) {
////	softwareId := c.Query("id")
////
////	if softwareId == "" {
////		response.FailWithMessage("softwareIDsoftware", c)
////		return
////	}
////
////	details, err := softwareService.GetSoftwareDetails(softwareId)
////
////	if err != nil {
////		response.FailWithMessage(err.Error(), c)
////		return
////	}
////
////	response.OkWithData(details, c)
////}
////
////func (s *SoftwareApi) GetSoftwareInspect(c *gin.Context) {
////	// containerId
////	var inspect request.SoftwareInspectReq
////
////	if errStr, err := utils.BaseValidatorQuery(&inspect, c); err != nil {
////		response.FailWithMessage(errStr, c)
////		return
////	}
////
////	softwareInspect, err := softwareService.GetSoftwareInspect(inspect)
////
////	if err != nil {
////		response.FailWithMessage(err.Error(), c)
////		return
////	}
////
////	response.OkWithData(softwareInspect, c)
////}
////
////// GetSoftwareLogs Obtaining softwareServiceObtaining software
////func (s *SoftwareApi) GetSoftwareLogs(c *gin.Context) {
////	var l request.GetServiceLog
////
////	if errStr, err := utils.BaseValidatorQuery(&l, c); err != nil {
////		response.FailWithMessage(errStr, c)
////		return
////	}
////	logs, err := softwareService.GetSoftwareLogs(l)
////	if err != nil {
////		response.FailWithMessage(err.Error(), c)
////		return
////	}
////
////	response.OkWithData(logs, c)
////}
//
//// CreateSoftware Create software
//// @Tags SoftwareApi
//// @Summary Create softwareApi
//// @Security ApiKeyAuth
//// @accept application/json
//// @Param object body docker.Software true "parameter"
//// @Produce application/json
//// @Success 200 {string} string "{"code":0,"data":{},"msg":"Created successfully"}"
//// @Router /v2/software [post]
//func (s *SoftwareApi) CreateSoftware(c *gin.Context) {
//	var software docker.DocSoftware
//
//	if errStr, err := utils.BaseValidator(&software, c); err != nil {
//		response.FailWithMessage(errStr, c)
//		return
//	}
//
//	err := softwareService.CreateSoftware(software, c)
//	if err != nil {
//		fmt.Println(err.Error())
//		//global.BaseChan <- global.ChanLogModal{
//		//	Level:   "error",
//		//	Message: err.Error(),
//		//}
//		//response.FailWithMessage(err.Error(), c)
//		return
//	}
//	global.BaseChan <- global.ChanLogModal{
//		Level:   "exit",
//		Message: "Successfully created software",
//	}
//	response.Ok(c)
//	//response.Ok(c)
//}
//
//func (s *SoftwareApi) UpdateSoftware(c *gin.Context) {
//	var req docker.DocSoftware
//
//	if errStr, err := utils.BaseValidator(&req, c); err != nil {
//		response.FailWithMessage(errStr, c)
//		return
//	}
//
//	if req.EntryPoint == "" {
//		req.EntryPoint = filesystem.ComposeFileDefaultName
//	}
//
//	err := softwareService.UpdateSoftware(req)
//	if err != nil {
//		global.BaseChan <- global.ChanLogModal{
//			Level:   "error",
//			Message: err.Error(),
//		}
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//
//	global.BaseChan <- global.ChanLogModal{
//		Level:   "exit",
//		Message: "Successfully upgraded software",
//	}
//	response.Ok(c)
//}
//
////func (s *SoftwareApi) UpdateSoftwareService(c *gin.Context) {
////	var service swarm.Service
////
////	softwareId := c.Param("id")
////
////	if errStr, err := utils.BaseValidator(&service, c); err != nil {
////		response.FailWithMessage(errStr, c)
////		return
////	}
////
////	if softwareId == "" {
////		response.FailWithMessage("softwareIdsoftware", c)
////		return
////	}
////
////	err := softwareService.UpdateSoftwareService(service, softwareId)
////	if err != nil {
////		response.FailWithMessage(err.Error(), c)
////		return
////	}
////
////	response.Ok(c)
////}
//
//// RemoveSoftware Remove software
//// @Tags SoftwareApi
//// @Summary Remove softwareApi
//// @Security ApiKeyAuth
//// @accept application/json
//// @Param object body request.DelSoftwareReq true "parameter"
//// @Produce application/json
//// @Success 200 {string} string "{"code":0,"data":{},"msg":"Successfully deleted"}"
//// @Router /v2/software [delete]
//func (s *SoftwareApi) RemoveSoftware(c *gin.Context) {
//	var software docRequst.DeleteDocSoftwareRequest
//
//	if errStr, err := utils.BaseValidator(&software, c); err != nil {
//		response.FailWithMessage(errStr, c)
//		return
//	}
//
//	err := softwareService.RemoveSoftware(software)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	response.Ok(c)
//}
//
////func (s *SoftwareApi) RemoveSoftwareService(c *gin.Context) {
////	var l request.DelSoftwareService
////
////	if errStr, err := utils.BaseValidator(&l, c); err != nil {
////		response.FailWithMessage(errStr, c)
////		return
////	}
////
////	err := softwareService.RemoveSoftwareService(l)
////	if err != nil {
////		response.FailWithMessage(err.Error(), c)
////		return
////	}
////
////	response.Ok(c)
////}
////
////func (s *SoftwareApi) StopSoftware(c *gin.Context) {
////	var software request.GetById
////
////	if errStr, err := utils.BaseValidator(&software, c); err != nil {
////		response.FailWithMessage(errStr, c)
////		return
////	}
////
////	fmt.Println(software.ID)
////	err := softwareService.StopSoftware(software.ID)
////	if err != nil {
////		response.FailWithMessage(err.Error(), c)
////		return
////	}
////	response.Ok(c)
////}
////
////func (s *SoftwareApi) StartSoftware(c *gin.Context) {
////	var software request.GetById
////
////	if errStr, err := utils.BaseValidator(&software, c); err != nil {
////		response.FailWithMessage(errStr, c)
////		return
////	}
////
////	err := softwareService.StartSoftware(software.ID)
////	if err != nil {
////		response.FailWithMessage(err.Error(), c)
////		return
////	}
////	response.Ok(c)
////}
////
////func (s *SoftwareApi) UpdateSoftwareReplicas(c *gin.Context) {
////	var software request.UpdateSoftwareReplicasReq
////
////	if errStr, err := utils.BaseValidator(&software, c); err != nil {
////		response.FailWithMessage(errStr, c)
////		return
////	}
////
////	updateResp, err := softwareService.UpdateSoftwareReplicas(software)
////	if err != nil {
////		response.FailWithMessage(err.Error(), c)
////		return
////	}
////
////	response.OkWithData(updateResp, c)
////}
////
////func (s *SoftwareApi) Test(c *gin.Context) {
////	//id := c.Query("id")
////	//test, s2, _ := softwareService.Test(id)
////	//
////	//response.OkWithDetailed(test, s2, c)
////	//response.OkWithData("4444", c)
////
////	// GetContainerLogsWs WebSocketGet container logs，Get container logs
////	// @Tags Container
////	// @Summary WebSocketGet container logs，Get container logs
////	// @Param containerId path string true "containerID"
////	// @Produce application/json
////	// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successfully obtained"}"
////	// @Router /v1/container/logs [get]
////	//containerId := c.Query("container")
////	//fmt.Println("containerId：", containerId)
////	fmt.Println("test")
////	//upgradegetupgradewebSocketupgrade
////	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
////	if err != nil {
////		fmt.Println("upgrade:", err)
////		response.FailWithMessage("websocketOpen failed！", c)
////	}
////	defer func() {
////		_ = ws.Close()
////	}()
////
////	fmt.Println("test1")
////
////	url := "unix:///var/run/docker.sock"
////	//if runtime.GOOS == "windows" {
////	//	url = "npipe:////./pipe/docker_engine"
////	//}
////
////	cli, err := client.NewClientWithOpts(
////		client.FromEnv,
////		client.WithAPIVersionNegotiation(),
////		client.WithHost(url),
////	)
////	fmt.Println("cli：", err)
////
////	ctx := context.Background()
////	ir, err := cli.ContainerExecCreate(ctx, "fc88161a1f50", types.ExecConfig{
////		AttachStdin:  true,
////		AttachStdout: true,
////		AttachStderr: true,
////		Cmd:          []string{"/bin/bash"},
////		Tty:          true,
////	})
////
////	if err != nil {
////		return
////	}
////
////	hr, err := cli.ContainerExecAttach(ctx, ir.ID, types.ExecStartCheck{
////		Detach: false, Tty: true,
////	})
////	fmt.Println("r：", err)
////
////	defer hr.Close()
////	defer func() {
////		hr.Conn.Write([]byte("exit\r"))
////	}()
////
////	go WsWriter(hr.Conn, ws)
////	WsReader(ws, hr.Conn)
////}
////
////func (s *SoftwareApi) Test1(c *gin.Context) {
////	var service swarm.Service
////
////	if errStr, err := utils.BaseValidator(&service, c); err != nil {
////		response.FailWithMessage(errStr, c)
////		return
////	}
////
////	test1, err := softwareService.Test1(service)
////	fmt.Println(err)
////
////	response.OkWithData(test1, c)
////}
////
////func WsWriter(reader io.Reader, writer *websocket.Conn) {
////	buf := make([]byte, 8192)
////	for {
////		nr, err := reader.Read(buf)
////		if nr > 0 {
////			err := writer.WriteMessage(websocket.BinaryMessage, buf[0:nr])
////			if err != nil {
////				return
////			}
////		}
////		if err != nil {
////			return
////		}
////	}
////}
////
////// WsPing The only purpose of this method is to determinewebsocketThe only purpose of this method is to determine，The only purpose of this method is to determine，The only purpose of this method is to determine
////func WsPing(ws *websocket.Conn) {
////	// The only purpose of this method is to determinewebsocketThe only purpose of this method is to determine，The only purpose of this method is to determine，The only purpose of this method is to determine
////	for {
////		_, _, err := ws.ReadMessage()
////		if err != nil {
////			return
////		}
////	}
////}
////
////func WsReader(reader *websocket.Conn, writer io.Writer) {
////	for {
////		messageType, p, err := reader.ReadMessage()
////		if err != nil {
////			return
////		}
////		if messageType == websocket.TextMessage {
////			_, err = writer.Write(p)
////			if err != nil {
////				return
////			}
////		}
////	}
////}
//
//// GetSoftwareStatus Obtain software status
//func (s *SoftwareApi) GetSoftwareStatus(c *gin.Context) {
//	var service []docRequst.SoftwareStatusRequest
//	if errStr, err := utils.BaseValidator(&service, c); err != nil {
//		response.FailWithMessage(errStr, c)
//		return
//	}
//	status, err := softwareService.GetSoftwareStatus(service)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	response.OkWithData(status, c)
//}

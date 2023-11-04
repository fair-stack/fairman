// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-10 16:28
package docker

import (
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/model/v2/common/response"
	"cnic/fairman-gin/model/v2/docker"
	docRequst "cnic/fairman-gin/model/v2/docker/request"
	"cnic/fairman-gin/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type SoftwareApi struct {
}

// CreateSoftware Create software
// @Tags SoftwareApi
// @Summary Create softwareApi
// @Security ApiKeyAuth
// @accept application/json
// @Param object body docker.DocSoftware true "parameter"
// @Produce application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"Created successfully"}"
// @Router /v2/software [post]
func (s *SoftwareApi) CreateSoftware(c *gin.Context) {
	var software docker.DocSoftware

	if errStr, err := utils.BaseValidator(&software, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	err := softwareService.CreateSoftware(software, c)
	if err != nil {
		fmt.Println(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	global.BaseChan <- global.ChanLogModal{
		Level:   "exit",
		Message: "Successfully created software",
	}
	response.Ok(c)
}

// UpdateSoftware Update software
// @Tags SoftwareApi
// @Summary Update softwareApi
// @Security ApiKeyAuth
// @accept application/json
// @Param object body request.UpdateSoftwareRequest true "parameter"
// @Produce application/json
func (s *SoftwareApi) UpdateSoftware(c *gin.Context) {
	var req docRequst.UpdateSoftwareRequest

	if errStr, err := utils.BaseValidator(&req, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	fmt.Println(req.Name, req.EndpointId)
	err := softwareService.UpdateSoftware(req)
	if err != nil {
		global.BaseChan <- global.ChanLogModal{
			Level:   "error",
			Message: err.Error(),
		}
		response.FailWithMessage(err.Error(), c)
		return
	}
	global.BaseChan <- global.ChanLogModal{
		Level:   "exit",
		Message: "Successfully updated the software",
	}
	response.Ok(c)
}

// RemoveSoftware Remove software
// @Tags SoftwareApi
// @Summary Remove softwareApi
// @Security ApiKeyAuth
// @accept application/json
// @Param object body request.DelSoftwareReq true "parameter"
// @Produce application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"Successfully deleted"}"
// @Router /v2/software [delete]
func (s *SoftwareApi) RemoveSoftware(c *gin.Context) {
	var software docRequst.DeleteDocSoftwareRequest

	if errStr, err := utils.BaseValidator(&software, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	err := softwareService.RemoveSoftware(software)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// StopSoftware Stop software
// @Tags SoftwareApi
// @Summary Stop softwareApi
// @Security ApiKeyAuth
// @accept application/json
// @Param object body request.StopSoftwareRequest true "parameter"
// @Produce application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"Stopped successfully"}"
// @Router /v2/software/stop [post]
func (s *SoftwareApi) StopSoftware(c *gin.Context) {
	var service docRequst.StopSoftwareRequest
	if errStr, err := utils.BaseValidator(&service, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	err := softwareService.StopSoftware(service)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// StartSoftware execute
// @Tags SoftwareApi
// @Summary executeApi
// @Security ApiKeyAuth
// @accept application/json
// @Param object body request.StartSoftwareRequest true "parameter"
// @Produce application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"Successfully started"}"
// @Router /v2/software/start [post]
func (s *SoftwareApi) StartSoftware(c *gin.Context) {
	var service docRequst.StartSoftwareRequest
	if errStr, err := utils.BaseValidator(&service, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	err := softwareService.StartSoftware(service)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// GetSoftwareStatus Obtain software status
func (s *SoftwareApi) GetSoftwareStatus(c *gin.Context) {
	var service []docRequst.SoftwareStatusRequest
	if errStr, err := utils.BaseValidator(&service, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}
	status, err := softwareService.GetSoftwareStatus(service)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(status, c)
}

// SaveEndpoint Save Software Environment
func (s *SoftwareApi) SaveEndpoint(c *gin.Context) {
	var service docRequst.SaveEndpointRequest
	if errStr, err := utils.BaseValidator(&service, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}
	err := softwareService.SaveEndpoint(service)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// UpdateSoftwareInspect Update software information
func (s *SoftwareApi) UpdateSoftwareInspect(c *gin.Context) {
	var service docRequst.UpdateSoftwareInspectRequest
	if errStr, err := utils.BaseValidator(&service, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}
	err := softwareService.UpdateSoftwareInspect(service)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// UpdateSoftwareInspect1 Update software information
func (s *SoftwareApi) UpdateSoftwareInspect1(c *gin.Context) {
	var service docRequst.UpdateSoftwareInspectRequest1
	if errStr, err := utils.BaseValidator(&service, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}
	err := softwareService.UpdateSoftwareInspect1(service)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

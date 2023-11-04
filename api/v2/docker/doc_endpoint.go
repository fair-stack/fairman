// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-04 15:20
package docker

import (
	"cnic/fairman-gin/model/v2/common/response"
	"cnic/fairman-gin/model/v2/docker"
	"cnic/fairman-gin/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type EndpointApi struct {
}

// TestEndpoint Testing the environment for local storage
// @Summary Testing the environment for local storage
// @Description Testing the environment for local storage
// @Tags docker
// @Accept json
// @Produce json
// @Param object body docker.DocEndpoint true "environmental information"
// @Success 200 {object} docker.DocEndpoint
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /v2/endpoint/test [post]
func (e *EndpointApi) TestEndpoint(c *gin.Context) {
	var endpoints []docker.DocEndpoint
	if errStr, err := utils.BaseValidator(&endpoints, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	endpoint, err := endpointService.TestEndpoint(endpoints)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(endpoint, c)
}

// TestOneEndpoint Testing a Single Environment
// @Summary Testing a Single Environment
// @Description Testing a Single Environment
// @Tags docker
// @Accept json
// @Produce json
// @Param object body docker.DocEndpoint true "environmental information "
// @Success 200 {object} docker.DocEndpoint
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /v2/endpoint/one/test [post]
func (e *EndpointApi) TestOneEndpoint(c *gin.Context) {
	var endpoint docker.DocEndpoint
	if err := c.ShouldBind(&endpoint); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(endpoint)
	resp, err := endpointService.TestOneEndpoint(endpoint)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(resp, c)
}

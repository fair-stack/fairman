// Package system
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-17 15:45
package system

import (
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/model/v2/common/response"
	"cnic/fairman-gin/model/v2/docker"
	"cnic/fairman-gin/utils"
	"fmt"
)

type EndpointService struct {
}

func (e *EndpointService) CreateEndpoint(endpoint docker.DocEndpoint) error {
	var resp response.Response
	err := utils.Post(fmt.Sprintf(AddEndpoint, global.FairConf.Service.Url), &endpoint, "", &resp)
	if err != nil {
		return utils.Errorf(err)
	}
	if resp.Code != 0 {
		return utils.Errorf("Failed to add environment information")
	}
	return nil
}

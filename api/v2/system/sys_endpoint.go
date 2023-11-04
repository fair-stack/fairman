// Package system
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-17 15:42
package system

type EndpointApi struct {
}

//func (e *EndpointApi) CreateEndpoint(c *gin.Context) {
//	var endpoint docker.DocEndpoint
//
//	if errStr, err := utils.BaseValidator(&endpoint, c); err != nil {
//		response.FailCodeMessage(http.StatusBadRequest, errStr, c)
//		return
//	}
//
//	err := endpointService.CreateEndpoint(endpoint)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	response.Ok(c)
//}

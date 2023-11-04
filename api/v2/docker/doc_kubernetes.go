// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-12-09 07:34
package docker

import (
	"cnic/fairman-gin/model/v2/common/response"
	requestV2 "cnic/fairman-gin/model/v2/docker/request"
	"cnic/fairman-gin/utils"

	"github.com/gin-gonic/gin"
)

type KubernetesApi struct {
}

// GetPodList obtaink8sPodobtain
// @Summary obtaink8sPodobtain
// @Description obtaink8sPodobtain
// @Tags kubernetes
// @Accept  json
// @Produce  json
// @Param namespace query string false "namespace"
// @Router /v2/kubernetes/pods [get]
func (k *KubernetesApi) GetPodList(ctx *gin.Context) {
	var req requestV2.PodListRequest
	if errStr, err := utils.BaseValidatorQuery(&req, ctx); err != nil {
		response.FailWithMessage(errStr, ctx)
		return
	}

	pods, err := KubernetesService.GetPodList(req, ctx)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithData(pods, ctx)
}

// GetPodListDetails obtaink8sPodobtain
// @Summary obtaink8sPodobtain
// @Description obtaink8sPodobtain
// @Tags kubernetes
// @Accept  json
// @Produce  json
// @Param namespace query string false "namespace"
// @Router /v2/kubernetes/pods [get]
func (k *KubernetesApi) GetPodListDetails(ctx *gin.Context) {
	var req requestV2.PodListRequest
	if errStr, err := utils.BaseValidatorQuery(&req, ctx); err != nil {
		response.FailWithMessage(errStr, ctx)
		return
	}

	pods, err := KubernetesService.GetPodListDetails(req, ctx)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithData(pods, ctx)
}

// PodTerminal k8sPodterminal
// @Summary k8sPodterminal
// @Description k8sPodterminal
// @Tags kubernetes
// @Accept  json
// @Produce  json
// @Param namespace query string false "namespace"
// @Param podName query string false "podName"
// @Param containerName query string false "containerName"
// @Router /v2/kubernetes/pod/terminal [get]
func (k *KubernetesApi) PodTerminal(ctx *gin.Context) {
	var req requestV2.PodTerminalRequest
	if errStr, err := utils.BaseValidatorQuery(&req, ctx); err != nil {
		response.FailWithMessage(errStr, ctx)
		return
	}

	if err := KubernetesService.PodTerminal(req, ctx); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	return
}

// PodLog k8sPodjournal
// @Summary k8sPodjournal
// @Description k8sPodjournal
// @Tags kubernetes
// @Accept  json
// @Produce  json
// @Param namespace query string false "namespace"
// @Param podName query string false "podName"
// @Param containerName query string false "containerName"
// @Param tailLines query string false "tailLines"
// @Router /v2/kubernetes/pod/log [get]
func (k *KubernetesApi) PodLog(ctx *gin.Context) {
	var req requestV2.PodLogRequest
	if errStr, err := utils.BaseValidatorQuery(&req, ctx); err != nil {
		response.FailWithMessage(errStr, ctx)
		return
	}

	if logs, err := KubernetesService.PodLog(req, ctx); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	} else {
		response.OkWithData(logs, ctx)
	}
}

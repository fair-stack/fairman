// Package response
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-12-09 20:31
package response

import (
	k8sApi "k8s.io/api/core/v1"
)

type PodListResponse struct {
	Namespace     string          `json:"namespace,omitempty" form:"namespace,omitempty" query:"namespace,omitempty"`
	PodName       string          `json:"podName,omitempty" form:"podName,omitempty" query:"podName,omitempty"`
	Status        k8sApi.PodPhase `json:"status,omitempty" form:"status,omitempty" query:"status,omitempty"`
	ContainerName string          `json:"containerName,omitempty" form:"containerName,omitempty" query:"containerName,omitempty"`
	CreatedAt     string          `json:"createdAt,omitempty" form:"createdAt,omitempty" query:"createdAt,omitempty"`
}

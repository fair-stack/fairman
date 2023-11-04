// Package request
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-12-09 07:42
package request

import "cnic/fairman-gin/model/v2/docker"

type PodListRequest struct {
	Endpoint docker.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	Name     string             `json:"name,omitempty" form:"name,omitempty" query:"name,omitempty" bson:"name"`
	Format   bool               `json:"format,omitempty" form:"format,omitempty" query:"format,omitempty"`
}

type PodTerminalRequest struct {
	Height    uint               `json:"height,omitempty" form:"height,omitempty" query:"height,omitempty" bson:"height"`
	Width     uint               `json:"width,omitempty" form:"width,omitempty" query:"width,omitempty" bson:"width"`
	Endpoint  docker.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	Namespace string             `json:"namespace,omitempty" form:"namespace,omitempty" query:"namespace,omitempty"`
	PodName   string             `json:"pod_name,omitempty" form:"pod_name,omitempty" query:"pod_name,omitempty"`
	Container string             `json:"container,omitempty" form:"container,omitempty" query:"container,omitempty"`
}

type PodLogRequest struct {
	Endpoint      docker.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	Namespace     string             `json:"namespace,omitempty" form:"namespace,omitempty" query:"namespace,omitempty"`
	PodName       string             `json:"podName,omitempty" form:"podName,omitempty" query:"podName,omitempty"`
	ContainerName string             `json:"containerName,omitempty" form:"containerName,omitempty" query:"containerName,omitempty"`
}

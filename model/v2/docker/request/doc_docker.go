// Package request
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-29 15:10
package request

import (
	"cnic/fairman-gin/model/v2/docker"

	"github.com/docker/docker/api/types/swarm"
)

type ContainerListRequest struct {
	Endpoint   docker.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	EndpointId string             `json:"endpointId,omitempty" form:"endpointId,omitempty" query:"endpointId,omitempty" bson:"endpointId"`
	Name       string             `json:"name,omitempty" form:"name,omitempty" query:"name,omitempty" bson:"name"`
}

type ContainerTerminalRequest struct {
	Height     uint               `json:"height,omitempty" form:"height,omitempty" query:"height,omitempty" bson:"height"`
	Width      uint               `json:"width,omitempty" form:"width,omitempty" query:"width,omitempty" bson:"width"`
	ExecId     string             `json:"execId,omitempty" form:"execId,omitempty" query:"execId,omitempty" bson:"execId"`
	Cmd        string             `json:"cmd,omitempty" form:"cmd,omitempty" query:"cmd,omitempty" bson:"cmd"`
	Endpoint   docker.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	EndpointId string             `json:"endpointId,omitempty" form:"endpointId,omitempty" query:"endpointId,omitempty" bson:"endpointId"`
}

type ContainerExecCreateRequest struct {
	ContainerId string             `json:"containerId,omitempty" form:"containerId,omitempty" query:"containerId,omitempty" bson:"containerId"`
	Cmd         string             `json:"cmd,omitempty" form:"cmd,omitempty" query:"cmd,omitempty" bson:"cmd"`
	Endpoint    docker.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	EndpointId  string             `json:"endpointId,omitempty" form:"endpointId,omitempty" query:"endpointId,omitempty" bson:"endpointId"`
}

type ContainerExecResizeRequest struct {
	Height     uint               `json:"height,omitempty" form:"height,omitempty" query:"height,omitempty" bson:"height"`
	Width      uint               `json:"width,omitempty" form:"width,omitempty" query:"width,omitempty" bson:"width"`
	ExecId     string             `json:"execId,omitempty" form:"execId,omitempty" query:"execId,omitempty" bson:"execId"`
	Endpoint   docker.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	EndpointId string             `json:"endpointId,omitempty" form:"endpointId,omitempty" query:"endpointId,omitempty" bson:"endpointId"`
}

type ContainerLogsRequest struct {
	Endpoint    docker.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	EndpointId  string             `json:"endpointId,omitempty" form:"endpointId,omitempty" query:"endpointId,omitempty" bson:"endpointId"`
	ContainerId string             `json:"containerId,omitempty" form:"containerId,omitempty" query:"containerId,omitempty" bson:"containerId"`
	Timestamps  bool               `json:"timestamps,omitempty" form:"timestamps,omitempty" query:"timestamps,omitempty" bson:"timestamps"`
	Tail        string             `json:"tail,omitempty" form:"tail,omitempty" query:"tail,omitempty" bson:"tail"`
}

type ContainerLogsInspect struct {
	Endpoint    docker.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	ContainerId string             `json:"containerId,omitempty" form:"containerId,omitempty" query:"containerId,omitempty" bson:"containerId"`
}

type ContainerInspectUpdate struct {
	Endpoint    docker.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	ContainerId string             `json:"containerId,omitempty" form:"containerId,omitempty" query:"containerId,omitempty" bson:"containerId"`
}

type ServiceListRequest struct {
	Endpoint docker.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	Name     string             `json:"name,omitempty" form:"name,omitempty" query:"name,omitempty" bson:"name"`
}

type ServiceLogsRequest struct {
	Endpoint   docker.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	EndpointId string             `json:"endpointId,omitempty" form:"endpointId,omitempty" query:"endpointId,omitempty" bson:"endpointId"`
	ServiceId  string             `json:"serviceId,omitempty" form:"serviceId,omitempty" query:"serviceId,omitempty" bson:"serviceId"`
	Timestamps bool               `json:"timestamps,omitempty" form:"timestamps,omitempty" query:"timestamps,omitempty" bson:"timestamps"`
	Tail       string             `json:"tail,omitempty" form:"tail,omitempty" query:"tail,omitempty" bson:"tail"`
}

type ServiceInspectRequest struct {
	Endpoint  docker.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	ServiceId string             `json:"serviceId,omitempty" form:"serviceId,omitempty" query:"serviceId,omitempty" bson:"serviceId"`
}

type UpdateServiceRequest struct {
	Endpoint docker.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	Service  swarm.Service      `json:"service,omitempty" form:"service,omitempty" query:"service,omitempty" bson:"service"`
}

type TaskListRequest struct {
	Name     string             `json:"name,omitempty" form:"name,omitempty" query:"name,omitempty" bson:"name"`
	Endpoint docker.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
}

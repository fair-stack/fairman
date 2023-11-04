// Package request
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-10 16:26
package request

import "cnic/fairman-gin/model/v2/docker"

type SoftwareReq struct {
	// build type 1. docker 2. swarm 3. compose
	//Type int `json:"type" form:"type" query:"type"`
	// stack name
	Name string `json:"name,omitempty" form:"name,omitempty" query:"name,omitempty" binding:"required"`
	// Replaced file content
	FileReplace map[string]string `json:"file_replace,omitempty" form:"file_replace,omitempty" query:"file_replace,omitempty" bson:"file_replace" binding:"required"`

	// Home page address
	HomeHost string `json:"home_host,omitempty" form:"home_host,omitempty" query:"home_host,omitempty" bson:"home_host" binding:"required"`
	// Homepage Port
	HomePort string `json:"home_port,omitempty" form:"home_port,omitempty" query:"home_port,omitempty" bson:"home_port" binding:"required"`

	// UserIp Installed user'sIP
	UserIp string `json:"user_ip,omitempty" form:"user_ip,omitempty" query:"user_ip,omitempty" bson:"user_ip"`
	UserId string `json:"user_id,omitempty" form:"user_id,omitempty" query:"user_id,omitempty" bson:"user_id"`
	// TemplateId The template called by the softwareID
	TemplateId string `json:"template_id,omitempty" form:"template_id,omitempty" query:"template_id,omitempty" bson:"template_id"`

	// environmentId
	EndpointId string             `json:"endpoint_id,omitempty" form:"endpoint_id,omitempty" query:"endpoint_id,omitempty" bson:"endpoint_id"`
	Endpoint   docker.DocEndpoint `json:"endpoint,omitempty" bson:"-"`
}

type DelSoftwareReq struct {
	Id string `json:"id,omitempty" form:"id,omitempty" query:"id,omitempty"`
}

type UpdateSoftwareReplicasReq struct {
	Id       string `json:"id,omitempty" form:"id,omitempty" query:"id,omitempty"`
	Version  uint64 `json:"version,omitempty" form:"version,omitempty" query:"version,omitempty"`
	Replicas uint64 `json:"replicas,omitempty" form:"replicas,omitempty" query:"replicas,omitempty"`
}

type SoftwareInspectReq struct {
	Id         string `json:"id,omitempty" form:"id,omitempty" query:"id,omitempty"`
	SoftwareId string `json:"software_id,omitempty" form:"software_id,omitempty" query:"software_id,omitempty"`
}

type GetServiceLog struct {
	SoftwareId string `json:"software_id,omitempty" form:"software_id,omitempty" query:"software_id,omitempty"`
	ServiceId  string `json:"service_id,omitempty" form:"service_id,omitempty" query:"service_id,omitempty"`

	Tail       string `json:"tail,omitempty" form:"tail,omitempty" query:"tail,omitempty"`
	Timestamps bool   `json:"timestamps,omitempty" form:"timestamps,omitempty" query:"timestamps,omitempty"`
	Since      string `json:"since,omitempty" form:"since,omitempty" query:"since,omitempty"`
}

type DelSoftwareService struct {
	SoftwareId string `json:"software_id,omitempty" form:"software_id,omitempty" query:"software_id,omitempty"`
	ServiceId  string `json:"service_id,omitempty" form:"service_id,omitempty" query:"service_id,omitempty"`
}

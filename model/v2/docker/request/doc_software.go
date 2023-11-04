// Package request
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-10 17:50
package request

import (
	dockerV2 "cnic/fairman-gin/model/v2/docker"
	"time"
)

type DeleteDocSoftwareRequest struct {
	Name        string               `json:"name,omitempty" form:"name,omitempty" query:"name,omitempty" bson:"name"`
	ProjectPath string               `json:"projectPath,omitempty" form:"projectPath,omitempty" query:"projectPath,omitempty" bson:"projectPath"`
	Endpoint    dockerV2.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	EndpointId  string               `json:"endpointId,omitempty" form:"endpointId,omitempty" query:"endpointId,omitempty" bson:"endpointId"`
}

type UpdateSoftwareRequest struct {
	Endpoint         dockerV2.DocEndpoint               `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	EndpointId       string                             `json:"endpointId,omitempty" form:"endpointId,omitempty" query:"endpointId,omitempty" bson:"endpointId"`
	NewContent       string                             `json:"newContent,omitempty" form:"newContent,omitempty" query:"newContent,omitempty" bson:"newContent"`
	NewConfig        map[string]dockerV2.TemplateConfig `json:"newConfig,omitempty" form:"newConfig,omitempty" query:"newConfig,omitempty" bson:"newConfig"`
	Name             string                             `json:"name,omitempty" form:"name,omitempty" query:"name,omitempty" bson:"name"`
	OldContent       string                             `json:"oldContent,omitempty" form:"oldContent,omitempty" query:"oldContent,omitempty" bson:"oldContent"`
	OldConfig        map[string]dockerV2.TemplateConfig `json:"oldConfig,omitempty" form:"oldConfig,omitempty" query:"oldConfig,omitempty" bson:"oldConfig"`
	KubernetesConfig dockerV2.KubernetesConfig          `json:"kubernetesConfig,omitempty" form:"kubernetesConfig,omitempty" query:"kubernetesConfig,omitempty"`
}

type StopSoftwareRequest struct {
	Name        string               `json:"name,omitempty" form:"name,omitempty" query:"name,omitempty" bson:"name"`
	ProjectPath string               `json:"projectPath,omitempty" form:"projectPath,omitempty" query:"projectPath,omitempty" bson:"projectPath"`
	Endpoint    dockerV2.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	EndpointId  string               `json:"endpointId,omitempty" form:"endpointId,omitempty" query:"endpointId,omitempty" bson:"endpointId"`
}

type StartSoftwareRequest struct {
	Name        string               `json:"name,omitempty" form:"name,omitempty" query:"name,omitempty" bson:"name"`
	ProjectPath string               `json:"projectPath,omitempty" form:"projectPath,omitempty" query:"projectPath,omitempty" bson:"projectPath"`
	Endpoint    dockerV2.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	EndpointId  string               `json:"endpointId,omitempty" form:"endpointId,omitempty" query:"endpointId,omitempty" bson:"endpointId"`
}

type SoftwareStatusRequest struct {
	// Default parameters _id Default parameters
	Id       string    `bson:"_id"`
	CreateAt time.Time `bson:"createAt"`
	UpdateAt time.Time `bson:"updateAt"`
	// The name of the software installation
	Name string `json:"name,omitempty" form:"name,omitempty" query:"name,omitempty" bson:"name"`
	// Home page address
	HomeHost string `json:"homeHost,omitempty" form:"homeHost,omitempty" query:"homeHost,omitempty" bson:"homeHost"`
	// Homepage Port
	HomePort string `json:"homePort,omitempty" form:"homePort,omitempty" query:"homePort,omitempty" bson:"homePort"`
	// environment
	Endpoint dockerV2.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	// Template
	Template dockerV2.DocTemplate `json:"template,omitempty" form:"template,omitempty" query:"template,omitempty" bson:"template"`
	// Update or not
	IsUpdate bool `json:"isUpdate,omitempty" form:"isUpdate,omitempty" query:"isUpdate,omitempty" bson:"-"`

	// userID
	UserId string `json:"userId,omitempty" form:"userId,omitempty" query:"userId,omitempty" bson:"userId"`
	// user name
	UserName string `json:"userName,omitempty" form:"userName,omitempty" query:"userName,omitempty" bson:"userName"`
	// userIP
	UserIp string `json:"userIp,omitempty" form:"userIp,omitempty" query:"userIp,omitempty" bson:"userIp"`
	// userIP2REGIN
	UserIp2Region string `json:"userIp2Region,omitempty" form:"userIp2Region,omitempty" query:"userIp2Region,omitempty" bson:"userIp2Region"`
	// Operation type
	OperateType string `json:"operateType,omitempty" form:"operateType,omitempty" query:"operateType,omitempty" bson:"operateType"`
	// state
	Status string `json:"status,omitempty" form:"status,omitempty" query:"status,omitempty" bson:"status"`
}

type OperationSoftwareRequest struct {
	Type string `json:"type,omitempty" form:"type,omitempty" query:"type,omitempty" bson:"type"`
	// The name of the software installation
	Name string `json:"name,omitempty" form:"name,omitempty" query:"name,omitempty" bson:"name"`
	// environmentID
	EndpointId string `json:"endpointId,omitempty" form:"endpointId,omitempty" query:"endpointId,omitempty" bson:"endpointId"`

	Endpoint *dockerV2.DocEndpoint
}

type SaveEndpointRequest struct {
	Endpoint dockerV2.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
}

type UpdateSoftwareInspectRequest struct {
	Endpoint dockerV2.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	// environmentID
	EndpointId string `json:"endpointId,omitempty" form:"endpointId,omitempty" query:"endpointId,omitempty" bson:"endpointId"`

	// Software Name
	SoftwareName string `json:"softwareName,omitempty" form:"softwareName,omitempty" query:"softwareName,omitempty" bson:"softwareName"`
	// Container Name （Container Name）
	ContainerName string `json:"containerName,omitempty" form:"containerName,omitempty" query:"containerName,omitempty" bson:"containerName"`
	// Component Configuration
	ContainerConfig dockerV2.TemplateConfig `json:"containerConfig,omitempty" form:"containerConfig,omitempty" query:"containerConfig,omitempty" bson:"containerConfig"`
}

type UpdateSoftwareInspectRequest1 struct {
	Endpoint dockerV2.DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	// environmentID
	EndpointId string `json:"endpointId,omitempty" form:"endpointId,omitempty" query:"endpointId,omitempty" bson:"endpointId"`

	// Software Name
	SoftwareName string `json:"softwareName,omitempty" form:"softwareName,omitempty" query:"softwareName,omitempty" bson:"softwareName"`
	// Component Configuration
	ContainerConfig map[string]dockerV2.TemplateConfig `json:"containerConfig,omitempty" form:"containerConfig,omitempty" query:"containerConfig,omitempty" bson:"containerConfig"`
}

//type UpdateSoftwareRequest struct {
//	Software docker.DocSoftware `json:"software,omitempty" form:"software,omitempty" query:"software,omitempty" bson:"software"`
//	Template docker.DocTemplate `json:"template,omitempty" form:"template,omitempty" query:"template,omitempty" bson:"template"`
//}

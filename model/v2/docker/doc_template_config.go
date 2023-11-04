// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-04 15:12
package docker

import (
	"cnic/fairman-gin/model/v2/system"
	"time"

	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
)

type DocTemplateConfig struct {
	// Default parameters _id Default parameters
	Id       string    `bson:"_id"`
	CreateAt time.Time `bson:"createAt"`
	UpdateAt time.Time `bson:"updateAt"`
	// TemplateID
	TemplateId string `json:"templateId,omitempty" form:"templateId,omitempty" query:"templateId,omitempty" bson:"templateId"`
	// DockerImage image
	DockerImage string `json:"dockerImage,omitempty" form:"dockerImage,omitempty" query:"dockerImage,omitempty" bson:"dockerImage"`
	// Dockerconfiguration information configuration informationdefault
	DockerConfig map[string]TemplateConfig `json:"dockerConfig,omitempty" form:"dockerConfig,omitempty" query:"dockerConfig,omitempty" bson:"dockerConfig"`
	// warehouseID
	DockerRegistry DocRegistry `json:"dockerRegistry,omitempty" form:"dockerRegistry,omitempty" query:"dockerRegistry,omitempty" bson:"dockerRegistry"`
	// SwarmFile Content
	SwarmContent string `json:"swarmContent,omitempty" form:"swarmContent,omitempty" query:"swarmContent,omitempty" bson:"swarmContent"`
	// Swarmreplace content replace content
	SwarmConfig map[string]TemplateConfig `json:"swarmConfig,omitempty" form:"swarmConfig,omitempty" query:"swarmConfig,omitempty" bson:"swarmConfig"`
	// SwarmwarehouseID
	SwarmRegistry []DocRegistry `json:"swarmRegistry,omitempty" form:"swarmRegistry,omitempty" query:"swarmRegistry,omitempty" bson:"swarmRegistry"`
	// SwarmFile Content
	K8sContent string `json:"k8SContent,omitempty" form:"k8SContent,omitempty" query:"k8SContent,omitempty" bson:"k8SContent"`
	// Swarmreplace content replace content
	K8sConfig KubernetesConfig `json:"k8SConfig,omitempty" form:"k8SConfig,omitempty" query:"k8SConfig,omitempty" bson:"k8SConfig"`
	// SwarmwarehouseID
	K8sRegistry []DocRegistry `json:"k8SRegistry,omitempty" form:"k8SRegistry,omitempty" query:"k8SRegistry,omitempty" bson:"k8SRegistry"`
	// (Reserved)Reserved
	Platform int `json:"platform,omitempty" form:"platform,omitempty" query:"platform,omitempty" bson:"platform"`
	// Software size
	FileSize float64 `json:"fileSize,omitempty" form:"fileSize,omitempty" query:"fileSize,omitempty" bson:"fileSize"`
	// userID
	UserId string `json:"userId,omitempty" form:"userId,omitempty" query:"userId,omitempty" bson:"userId"`
	// Version number
	Version string `json:"version,omitempty" form:"version,omitempty" query:"version,omitempty" bson:"version"`
	// Version Content
	VersionContent string `json:"versionContent,omitempty" form:"versionContent,omitempty" query:"versionContent,omitempty" bson:"versionContent"`
	// User Manual
	Manual system.SysBaseFiles `json:"manual,omitempty" form:"manual,omitempty" query:"manual,omitempty" bson:"manual"`
	// Delete or not
	IsDel bool `json:"isDel,omitempty" form:"isDel,omitempty" query:"isDel,omitempty" bson:"isDel"`
}

type KubernetesConfig struct {
	Deployment            appsV1.Deployment            `json:"deployment,omitempty" form:"deployment,omitempty" query:"deployment,omitempty" bson:"deployment"`
	Service               coreV1.Service               `json:"service,omitempty" form:"service,omitempty" query:"service,omitempty" bson:"service"`
	PersistentVolumeClaim coreV1.PersistentVolumeClaim `json:"persistentVolumeClaim,omitempty" form:"persistentVolumeClaim,omitempty" query:"persistentVolumeClaim,omitempty" bson:"persistentVolumeClaim"`
}

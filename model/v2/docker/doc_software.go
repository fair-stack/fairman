// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-11 16:57
package docker

import (
	"time"

	dockerYml "github.com/lliuhuan/compose-yml"
)

//type Software struct {
//	Id       string    `mapstructure:"Id" json:"Id,omitempty" yaml:"Id"`
//	CreateAt time.Time `mapstructure:"CreateAt" json:"CreateAt,omitempty" yaml:"CreateAt"`
//	UpdateAt time.Time `mapstructure:"UpdateAt" json:"UpdateAt,omitempty" yaml:"UpdateAt"`
//	// The name of the software installation
//	Name string `json:"name,omitempty" form:"name,omitempty" query:"name,omitempty" bson:"name" binding:"required"`
//	// Home page address
//	HomeHost string `json:"home_host,omitempty" form:"home_host,omitempty" query:"home_host,omitempty" bson:"home_host"`
//	// Homepage Port
//	HomePort string `json:"home_port,omitempty" form:"home_port,omitempty" query:"home_port,omitempty" bson:"home_port"`
//	// environmentId
//	EndpointId string `json:"endpoint_id,omitempty" form:"endpoint_id,omitempty" query:"endpoint_id,omitempty" bson:"endpoint_id"`
//	// TemplateId The template called by the softwareID
//	TemplateId string `json:"template_id,omitempty" form:"template_id,omitempty" query:"template_id,omitempty" bson:"template_id"`
//	// Replaced file content
//	TemplateConfig map[string]TemplateConfig `json:"template_config,omitempty" form:"template_config,omitempty" query:"template_config,omitempty"`
//	Status         string                    `json:"status,omitempty"`
//
//	// Userid Installed Users
//	UserId string `json:"user_id,omitempty" form:"user_id,omitempty" query:"user_id,omitempty" bson:"user_id"`
//	// UserIp Installed user'sIP
//	UserIp string `json:"user_ip,omitempty" form:"user_ip,omitempty" query:"user_ip,omitempty" bson:"user_ip"`
//	// Type of installed software
//	TemplateType int `json:"template_type,omitempty" form:"template_type,omitempty" query:"template_type,omitempty"`
//	// docker-compose.yaml File Content
//	FileContent string `json:"file_content,omitempty" form:"file_content,omitempty" query:"file_content,omitempty" bson:"file_content"`
//	// Path to stack file
//	EntryPoint string `json:"entry_point,omitempty" form:"entry_point,omitempty" query:"entry_point,omitempty" bson:"entry_point" example:"docker-compose.yml"`
//	// Applicable only when deploying stacks containing multiple files
//	AdditionalFiles []string `json:"additional_files,omitempty" form:"additional_files,omitempty" query:"additional_files,omitempty" bson:"additional_files"`
//	// Project Storage Path
//	ProjectPath string `json:"project_path,omitempty" form:"project_path,omitempty" query:"project_path,omitempty" bson:"project_path" example:"/data/compose/myStack_jpofkc0i9uo9wtx1zesuk649w"`
//
//	// Temporary variables
//	Endpoint DocEndpoint `json:"endpoint,omitempty" bson:"-"`
//
//	// TODO：I don't know where to put it，I don't know where to put it
//	// environment variable
//	Env []Pair `json:"env" form:"env" query:"env"`
//}

type DocSoftware struct {
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
	Endpoint DocEndpoint `json:"endpoint,omitempty" form:"endpoint,omitempty" query:"endpoint,omitempty" bson:"endpoint"`
	// Template
	Template DocTemplate `json:"template,omitempty" form:"template,omitempty" query:"template,omitempty" bson:"template"`

	// userID
	UserId string `json:"userId,omitempty" form:"userId,omitempty" query:"userId,omitempty" bson:"userId"`
	// userIP
	UserIp string `json:"userIp,omitempty" form:"userIp,omitempty" query:"userIp,omitempty" bson:"userIp"`
	// userIP2REGIN
	UserIp2Region string `json:"userIp2Region,omitempty" form:"userIp2Region,omitempty" query:"userIp2Region,omitempty" bson:"userIp2Region"`
	// running state
	Status string `json:"status,omitempty" form:"status,omitempty" query:"status,omitempty" bson:"status"`

	// The following are all temporary configurations
	ProjectPath string        `json:"projectPath,omitempty" bson:"-"`
	Env         []Pair        `json:"env,omitempty" bson:"-"`
	EntryPoint  string        `json:"entryPoint,omitempty" form:"entryPoint,omitempty" query:"entryPoint,omitempty" bson:"-"`
	DockerYml   dockerYml.Yml `bson:"-"`
}

// Pair defines a key/value string pair
type Pair struct {
	Name  string `json:"name" example:"name"`
	Value string `json:"value" example:"value"`
}

type TemplateConfig struct {
	Ports       []dockerYml.Port        `json:"ports,omitempty" form:"ports,omitempty" query:"ports,omitempty" bson:"ports"`
	Mounts      []dockerYml.VolumeMap   `json:"volume,omitempty" form:"volume,omitempty" query:"volume,omitempty" bson:"volume"`
	Environment []dockerYml.Environment `json:"environment,omitempty" form:"environmentConfig,omitempty" query:"environmentConfig,omitempty" bson:"environmentConfig"`
	Hostname    string                  `json:"hostname,omitempty" form:"hostname,omitempty" query:"hostname,omitempty" bson:"hostname"`
}

//type PublicServices []PublicService
//
//type PublicConfig struct {
//	Filename string                           `yaml:"-" json:"-"`
//	Version  string                           `json:"version"`
//	Services map[string]PublicService         `json:"services"`
//	Networks map[string]types.NetworkConfig   `yaml:",omitempty" json:"networks,omitempty"`
//	Volumes  map[string]types.VolumeConfig    `yaml:",omitempty" json:"volumes,omitempty"`
//	Secrets  map[string]types.SecretConfig    `yaml:",omitempty" json:"secrets,omitempty"`
//	Configs  map[string]types.ConfigObjConfig `yaml:",omitempty" json:"configs,omitempty"`
//	Extras   map[string]interface{}           `yaml:",inline" json:"-"`
//}
//
//type PublicService struct {
//	types.ServiceConfig
//	Build *struct{} `yaml:",omitempty" json:"build,omitempty"`
//}

// Package response
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-11 16:51
package response

import (
	"cnic/fairman-gin/model/v2/docker"

	"github.com/docker/docker/api/types/swarm"
)

type SoftwareByUserIdResp struct {
	Software []docker.DocSoftware `json:"software,omitempty"`
	Template []docker.DocTemplate `json:"template,omitempty"`
}

//type SoftwareResp struct {
//	SoftwareId string    `json:"software_id"`
//	CreateAt   time.Time `mapstructure:"CreateAt" json:"CreateAt" yaml:"CreateAt"`
//	UpdateAt   time.Time `mapstructure:"UpdateAt" json:"UpdateAt" yaml:"UpdateAt"`
//	// The name of the software installation
//	SoftwareName string `json:"software_name"`
//	// softwarelogo
//	TemplateLogo string `json:"template_logo"`
//	// Software type
//	TemplateTypes []string `json:"template_types"`
//	// Upload screenshot
//	TemplateScreenshot []docker.Screenshot `json:"template_screenshot"`
//	// Installation position
//	EndpointName string `json:"endpoint_name"`
//	// Software rating
//	TemplateScore float64 `json:"template_score"`
//	// Number of downloads
//	TemplateDownload int `json:"template_download"`
//	// System version
//	TemplateVersion string `json:"template_version"`
//	// What is software like 1. docker 2. swarm 3. compose
//	TemplateType int    `mapstructure:"template_type" json:"template_type,omitempty" yaml:"template_type"`
//	TemplateId   string `json:"template_id,omitempty"`
//	Status       string `json:"status,omitempty"`
//
//	// Home page address
//	HomeHost string `mapstructure:"home_host" json:"home_host,omitempty" yaml:"home_host"`
//	// Homepage Port
//	HomePort string `mapstructure:"home_port" json:"home_port,omitempty" yaml:"home_port"`
//
//	// Is it installed
//	IsInstall bool `json:"is_install"`
//}

type SoftwareDetailsResp struct {
	Service  swarm.Service `json:"service,omitempty"`
	Children []swarm.Task  `json:"children,omitempty"`
}

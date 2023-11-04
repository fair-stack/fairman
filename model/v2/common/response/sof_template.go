// Package response
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-28 18:48
package response

import "github.com/docker/cli/cli/compose/types"

type TemplateConfigResp struct {
	Ports       []types.ServicePortConfig   `json:"ports,omitempty"`
	Mounts      []types.ServiceVolumeConfig `json:"mounts,omitempty"`
	Environment types.MappingWithEquals     `json:"environment,omitempty"`
	Hostname    string                      `json:"hostname,omitempty"`
}

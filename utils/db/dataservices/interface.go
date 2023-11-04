// Package dataservices
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-04 15:43
package dataservices

import "cnic/fairman-gin/model/v2/docker"

type (
	// EndpointService represents a service for managing environment(endpoint) data
	EndpointService interface {
		Endpoint(ID string) (*docker.DocEndpoint, error)
		Endpoints() ([]docker.DocEndpoint, error)
		Create(endpoint *docker.DocEndpoint) error
		UpdateEndpoint(ID string, endpoint *docker.DocEndpoint) error
		DeleteEndpoint(ID string) error
		//GetNextIdentifier() int
		BucketName() string
	}
)

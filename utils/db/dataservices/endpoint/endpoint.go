// Package endpoint
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-04 11:29
package endpoint

import (
	"cnic/fairman-gin/model/v2/docker"
	"cnic/fairman-gin/typings"
	"fmt"

	"github.com/sirupsen/logrus"
)

const (
	// BucketName represents the name of the bucket where this service stores data.
	BucketName = "endpoints"
)

// Service represents a service for managing environment(endpoint) data.
type Service struct {
	connection typings.Connection
}

func (service *Service) BucketName() string {
	return BucketName
}

// NewService creates a new instance of a service.
func NewService(connection typings.Connection) (*Service, error) {
	err := connection.SetServiceName(BucketName)
	if err != nil {
		return nil, err
	}

	return &Service{
		connection: connection,
	}, nil
}

// Endpoint returns an environment(endpoint) by ID.
func (service *Service) Endpoint(ID string) (*docker.DocEndpoint, error) {
	var endpoint docker.DocEndpoint
	identifier := service.connection.ConvertToKey(ID)

	err := service.connection.GetObject(BucketName, identifier, &endpoint)
	if err != nil {
		return nil, err
	}

	return &endpoint, nil
}

// UpdateEndpoint updates an environment(endpoint).
func (service *Service) UpdateEndpoint(ID string, endpoint *docker.DocEndpoint) error {
	identifier := service.connection.ConvertToKey(ID)
	return service.connection.UpdateObject(BucketName, identifier, endpoint)
}

// DeleteEndpoint deletes an environment(endpoint).
func (service *Service) DeleteEndpoint(ID string) error {
	identifier := service.connection.ConvertToKey(ID)
	return service.connection.DeleteObject(BucketName, identifier)
}

// Endpoints return an array containing all the environments(endpoints).
func (service *Service) Endpoints() ([]docker.DocEndpoint, error) {
	var endpoints []docker.DocEndpoint

	err := service.connection.GetAllWithJsoniter(
		BucketName,
		&docker.DocEndpoint{},
		func(obj interface{}) (interface{}, error) {
			endpoint, ok := obj.(*docker.DocEndpoint)
			if !ok {
				logrus.WithField("obj", obj).Errorf("Failed to convert to Endpoint object")
				return nil, fmt.Errorf("Failed to convert to Endpoint object: %s", obj)
			}
			endpoints = append(endpoints, *endpoint)
			return &docker.DocEndpoint{}, nil
		})
	return endpoints, err
}

// Create CreateEndpoint assign an ID to a new environment(endpoint) and saves it.
func (service *Service) Create(endpoint *docker.DocEndpoint) error {
	return service.connection.CreateObjectWithId(BucketName, endpoint.Id, endpoint)
}

//// GetNextIdentifier returns the next identifier for an environment(endpoint).
//func (service *Service) GetNextIdentifier() int {
//	return service.connection.GetNextIdentifier(BucketName)
//}

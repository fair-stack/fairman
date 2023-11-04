// Package datastore
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-04 15:38
package datastore

import (
	"cnic/fairman-gin/typings"
	"cnic/fairman-gin/utils/db/dataservices"
	"cnic/fairman-gin/utils/db/dataservices/endpoint"
)

type Store struct {
	connection typings.Connection

	EndpointService *endpoint.Service
}

func (store *Store) initServices() error {

	endpointService, err := endpoint.NewService(store.connection)
	if err != nil {
		return err
	}
	store.EndpointService = endpointService

	return nil
}

func (store *Store) Endpoint() dataservices.EndpointService {
	return store.EndpointService
}

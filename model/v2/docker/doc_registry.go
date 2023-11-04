// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-11 15:33
package docker

import "time"

type RegistryType int

type DocRegistry struct {
	// Default parameters _id Default parameters
	Id       string    `bson:"_id"`
	CreateAt time.Time `bson:"createAt"`
	UpdateAt time.Time `bson:"updateAt"`
	// Type Warehouse type
	// Registry Type. Valid values are:
	//	1 (Quay.io), // Not yet developed quay.io Not yet developed
	//	2 (Azure container registry), // Not yet developed
	//	3 (custom registry), // Self built warehouse, Self built warehouse
	//	4 (Gitlab registry), // Not yet developed gitlab Not yet developed
	//	5 (ProGet registry), // Prepare for development
	//	6 (DockerHub) // docker hub
	//	7 (ECR) // Not yet developed aws ECR Not yet developed
	RegistryType RegistryType `json:"registryType" form:"registryType" query:"registryType" bson:"registryType" enums:"1,2,3,4,5,6,7"`
	// Name Warehouse alias Warehouse alias
	Name string `json:"name" form:"name" query:"name" bson:"name"`
	// URL or IP Warehouse address Warehouse address Warehouse address
	URL string `json:"url" form:"url" query:"url" bson:"url"`
	// BaseURL in order toProGetin order to
	BaseURL string `json:"baseURL,omitempty" form:"baseURL,omitempty" query:"baseURL,omitempty" bson:"baseURL"`
	// Authentication Whether to enable authentication
	Authentication bool `json:"authentication" form:"authentication" query:"authentication" bson:"authentication"`
	// Username user name
	Username string `json:"username" form:"username" query:"username" bson:"username"`
	// Password password
	Password string `json:"password" form:"password" query:"password" bson:"password"`

	// Userid Reserved Usersid，Reserved Users
	UserId string `json:"userId,omitempty" form:"userId,omitempty" query:"userId,omitempty" bson:"userId"`

	// Additional fields for other warehouse types need to be stored in the middle(Additional fields for other warehouse types need to be stored in the middle，Additional fields for other warehouse types need to be stored in the middle3，6Additional fields for other warehouse types need to be stored in the middle)
	Gitlab GitlabRegistryData `json:"gitlab" form:"gitlab" query:"gitlab" bson:"gitlab"`

	// Store temporary access token(Store temporary access token)
	AccessToken       string `json:"AccessToken,omitempty"`
	AccessTokenExpiry int64  `json:"AccessTokenExpiry,omitempty"`
}

type GitlabRegistryData struct {
	ProjectID   int    `json:"projectID,omitempty" form:"projectID,omitempty" query:"projectID,omitempty" bson:"projectID"`
	InstanceURL string `json:"instanceURL,omitempty" form:"instanceURL,omitempty" query:"instanceURL,omitempty" bson:"instanceURL"`
	ProjectPath string `json:"projectPath,omitempty" form:"projectPath,omitempty" query:"projectPath,omitempty" bson:"projectPath"`
}

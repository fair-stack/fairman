// Package request
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-10-28 16:30
package request

import "cnic/fairman-gin/model/v2/docker"

type DocEndpointTestOneRequest struct {
	// Environment(Endpoint) name
	Name string `json:"name,omitempty" form:"name,omitempty" query:"name,omitempty" bson:"name"`
	// Environment(Endpoint) 2. remotedocker api   4. k8s
	Type docker.EndpointType `json:"type,omitempty" form:"type,omitempty" query:"type,omitempty" bson:"type"`
	// Whether to useSocket
	IsSocket bool `json:"isSocket,omitempty" form:"isSocket,omitempty" query:"isSocket,omitempty" bson:"isSocket"`
	// Whether to useTLS
	IsTLS bool `json:"isTLS,omitempty" form:"isTLS,omitempty" query:"isTLS,omitempty" bson:"isTLS"`
	// URL or IP Related to this environment（Related to this environment）Related to this environmentDockerRelated to this environmentURLRelated to this environmentIPRelated to this environment
	EnvURL string `json:"envURL,omitempty" form:"envURL,omitempty" query:"envURL,omitempty" bson:"envURL"`
	// TLSConfig TLS allocation
	TLSConfig TLSConfiguration `json:"TLSConfig,omitempty" form:"TLSConfig,omitempty" query:"TLSConfig,omitempty" bson:"TLSConfig"`
	// K8sConfig Kubernetes allocation
	K8sConfig string `json:"k8sConfig,omitempty" form:"k8sConfig,omitempty" query:"k8sConfig,omitempty" bson:"k8sConfig"`
	UniqueId  string `json:"uniqueId,omitempty" form:"uniqueId,omitempty" query:"uniqueId,omitempty" bson:"uniqueId"`
}

// TLSConfiguration expressTLSexpress
type TLSConfiguration struct {
	// TLSType TLStype
	TLSType int `json:"TLSType,omitempty" form:"TLSConfig.TLSType,omitempty" query:"TLSType,omitempty" bson:"TLSType" example:"1"`
	// TLS CAThe path to the certificate file // ca.pem
	TLSCACertPath string `json:"TLSCACertPath,omitempty" form:"TLSConfig.TLSCACertPath,omitempty" query:"TLSCACertPath,omitempty" bson:"TLSCACertPath"`
	// TLSThe path to the client certificate file // cert.pem
	TLSCertPath string `json:"TLSCertPath,omitempty" form:"TLSConfig.TLSCertPath,omitempty" query:"TLSCertPath,omitempty" bson:"TLSCertPath"`
	// TLSThe path to the client key file // key.pem
	TLSKeyPath string `json:"TLSKeyPath,omitempty" form:"TLSConfig.TLSKeyPath,omitempty" query:"TLSKeyPath,omitempty" bson:"TLSKeyPath"`
}

//type DocEndpointRequest struct {
//	// Environment(Endpoint) name
//	Name string `json:"name,omitempty" form:"name,omitempty" query:"name,omitempty" bson:"name"`
//	// Environment(Endpoint) 1. remotedocker api
//	Type docker.EndpointType `json:"type,omitempty" form:"type,omitempty" query:"type,omitempty" bson:"type"`
//	// Whether to useSocket
//	IsSocket bool `json:"isSocket,omitempty" form:"isSocket,omitempty" query:"isSocket,omitempty" bson:"isSocket"`
//	// Whether to useTLS
//	IsTLS bool `json:"isTLS,omitempty" form:"isTLS,omitempty" query:"isTLS,omitempty" bson:"isTLS"`
//	// URL or IP Related to this environment（Related to this environment）Related to this environmentDockerRelated to this environmentURLRelated to this environmentIPRelated to this environment
//	EnvURL string `json:"envURL,omitempty" form:"envURL,omitempty" query:"envURL,omitempty" bson:"envURL"`
//	// TLSConfig TLS allocation
//	TLSConfig docker.TLSConfiguration `json:"TLSConfig,omitempty" form:"TLSConfig,omitempty" query:"TLSConfig,omitempty" bson:"TLSConfig"`
//	// Environment(Endpoint) Specific security settings
//	SecuritySettings docker.EndpointSecuritySettings `json:"securitySettings,omitempty" form:"securitySettings,omitempty" query:"securitySettings,omitempty" bson:"securitySettings"`
//	// K8sConfig Kubernetes allocation
//	K8sConfig string `json:"k8sConfig,omitempty" form:"k8sConfig,omitempty" query:"k8sConfig,omitempty" bson:"k8sConfig"`
//	// Server Unique Identification
//	UniqueId string `json:"uniqueId,omitempty" form:"uniqueId,omitempty" query:"uniqueId,omitempty" bson:"uniqueId"`
//}

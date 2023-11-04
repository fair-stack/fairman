// Package docker
// @program: fairman-server-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-17 15:11
package docker

import (
	"time"
)

type (
	EndpointType int

	// EndpointStatus Representing the environment（Representing the environment）Representing the environment
	EndpointStatus int

	DocEndpoint struct {
		// Default parameters _id Default parameters
		Id       string    `bson:"_id"`
		CreateAt time.Time `bson:"createAt"`
		UpdateAt time.Time `bson:"updateAt"`
		// Environment(Endpoint) name
		Name string `json:"name,omitempty" form:"name,omitempty" query:"name,omitempty" bson:"name"`
		// Environment(Endpoint) 1. remotedocker api
		Type EndpointType `json:"type,omitempty" form:"type,omitempty" query:"type,omitempty" bson:"type"`
		// Whether to useSocket
		IsSocket bool `json:"isSocket,omitempty" form:"isSocket,omitempty" query:"isSocket,omitempty" bson:"isSocket"`
		// Whether to useTLS
		IsTLS bool `json:"isTLS,omitempty" form:"isTLS,omitempty" query:"isTLS,omitempty" bson:"isTLS"`
		// URL or IP Related to this environment（Related to this environment）Related to this environmentDockerRelated to this environmentURLRelated to this environmentIPRelated to this environment
		EnvURL string `json:"envURL,omitempty" form:"envURL,omitempty" query:"envURL,omitempty" bson:"envURL"`
		// TLSConfig TLS allocation
		TLSConfig TLSConfiguration `json:"TLSConfig,omitempty" form:"TLSConfig,omitempty" query:"TLSConfig,omitempty" bson:"TLSConfig"`
		// Environmental status（Environmental status）(1 - up, 2 - down)
		Status EndpointStatus `json:"status,omitempty" form:"status,omitempty" query:"status,omitempty" bson:"status" example:"1"`
		// Environment(Endpoint) Specific security settings
		SecuritySettings EndpointSecuritySettings `json:"securitySettings,omitempty" form:"securitySettings,omitempty" query:"securitySettings,omitempty" bson:"securitySettings"`
		// userID
		UserId string `json:"userId,omitempty" form:"userId,omitempty" query:"userId,omitempty" bson:"userId"`
		// All templates // All templates All templates
		Templates []DocTemplate `json:"templates,omitempty" form:"templates,omitempty" query:"templates,omitempty" bson:"-"`
		// Connection status
		Connected bool `json:"connected,omitempty" form:"connected,omitempty" query:"connected,omitempty" bson:"connected"`
		// Server username
		UserName string `json:"username,omitempty" form:"username,omitempty" query:"username,omitempty" bson:"username"`
		// Server password
		Password string `json:"password,omitempty" form:"password,omitempty" query:"password,omitempty" bson:"password"`

		// K8sConfig Kubernetes allocation
		K8sConfig string `json:"k8sConfig,omitempty" form:"k8sConfig,omitempty" query:"k8sConfig,omitempty" bson:"k8sConfig"`
		// Server Unique Identification
		UniqueId string `json:"uniqueId,omitempty" form:"uniqueId,omitempty" query:"uniqueId,omitempty" bson:"uniqueId"`
	}

	// TLSConfiguration expressTLSexpress
	TLSConfiguration struct {
		// TLSType TLStype
		TLSType int `json:"TLSType,omitempty" form:"TLSType,omitempty" query:"TLSType,omitempty" bson:"TLSType" example:"1"`
		// TLS CAThe path to the certificate file // ca.pem
		TLSCACertPath string `json:"TLSCACertPath,omitempty" form:"TLSCACertPath,omitempty" query:"TLSCACertPath,omitempty" bson:"TLSCACertPath"`
		// TLSThe path to the client certificate file // cert.pem
		TLSCertPath string `json:"TLSCertPath,omitempty" form:"TLSCertPath,omitempty" query:"TLSCertPath,omitempty" bson:"TLSCertPath"`
		// TLSThe path to the client key file // key.pem
		TLSKeyPath string `json:"TLSKeyPath,omitempty" form:"TLSKeyPath,omitempty" query:"TLSKeyPath,omitempty" bson:"TLSKeyPath"`
	}

	// EndpointSecuritySettings represents settings for an environment(endpoint)
	// Temporarily useless
	EndpointSecuritySettings struct {
		// Whether non-administrator should be able to use bind mounts when creating containers
		// Should non administrators be able to use binding mount when creating containers
		AllowBindMountsForRegularUsers bool `json:"allowBindMountsForRegularUsers" example:"false"`
		// Whether non-administrator should be able to use privileged mode when creating containers
		// Should non administrators be able to use privileged mode when creating containers
		AllowPrivilegedModeForRegularUsers bool `json:"allowPrivilegedModeForRegularUsers" example:"false"`
		// Whether non-administrator should be able to browse volumes
		// Should non administrators be able to browse volumes
		AllowVolumeBrowserForRegularUsers bool `json:"allowVolumeBrowserForRegularUsers" example:"true"`
		// Whether non-administrator should be able to use the host pid
		// Should non administrators be able to use the hostpid
		AllowHostNamespaceForRegularUsers bool `json:"allowHostNamespaceForRegularUsers" example:"true"`
		// Whether non-administrator should be able to use device mapping
		// Should non administrators be able to use device mapping
		AllowDeviceMappingForRegularUsers bool `json:"allowDeviceMappingForRegularUsers" example:"true"`
		// Whether non-administrator should be able to manage stacks
		// Should non administrators be able to manage the stack
		AllowStackManagementForRegularUsers bool `json:"allowStackManagementForRegularUsers" example:"true"`
		// Whether non-administrator should be able to use container capabilities
		// Should non administrators be able to use container functionality
		AllowContainerCapabilitiesForRegularUsers bool `json:"allowContainerCapabilitiesForRegularUsers" example:"true"`
		// Whether non-administrator should be able to use sysctl settings
		// Should non administrators be able to usesysctlShould non administrators be able to use
		AllowSysctlSettingForRegularUsers bool `json:"allowSysctlSettingForRegularUsers" example:"true"`
		// Whether host management features are enabled
		// Enable host management function
		EnableHostManagementFeatures bool `json:"enableHostManagementFeatures" example:"true"`
	}
)

var EndpointMap = make(map[string]DocEndpoint)

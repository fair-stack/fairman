// Package global
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-17 14:27
package global

import (
	"cnic/fairman-gin/model/v2/docker"
	"context"

	"github.com/docker/docker/client"
)

type (
	// AccessPolicy represent a policy that can be associated to a user or team
	AccessPolicy struct {
		// Role identifier. Reference the role that will be associated to this access policy
		RoleID RoleID `json:"RoleId" example:"1"`
	}

	// AzureCredentials represents the credentials used to connect to an Azure
	// environment(endpoint).
	AzureCredentials struct {
		// Azure application ID
		ApplicationID string `json:"ApplicationID" example:"eag7cdo9-o09l-9i83-9dO9-f0b23oe78db4"`
		// Azure tenant ID
		TenantID string `json:"TenantID" example:"34ddc78d-4fel-2358-8cc1-df84c8o839f5"`
		// Azure authentication key
		AuthenticationKey string `json:"AuthenticationKey" example:"cOrXoK/1D35w8YQ8nH1/8ZGwzz45JIYD5jxHKXEQknk="`
	}

	// DockerSnapshot represents a snapshot of a specific Docker environment(endpoint) at a specific time
	DockerSnapshot struct {
		Time                    int64             `json:"Time"`
		DockerVersion           string            `json:"DockerVersion"`
		Swarm                   bool              `json:"Swarm"`
		TotalCPU                int               `json:"TotalCPU"`
		TotalMemory             int64             `json:"TotalMemory"`
		RunningContainerCount   int               `json:"RunningContainerCount"`
		StoppedContainerCount   int               `json:"StoppedContainerCount"`
		HealthyContainerCount   int               `json:"HealthyContainerCount"`
		UnhealthyContainerCount int               `json:"UnhealthyContainerCount"`
		VolumeCount             int               `json:"VolumeCount"`
		ImageCount              int               `json:"ImageCount"`
		ServiceCount            int               `json:"ServiceCount"`
		StackCount              int               `json:"StackCount"`
		SnapshotRaw             DockerSnapshotRaw `json:"DockerSnapshotRaw"`
		NodeCount               int               `json:"NodeCount"`
	}

	// DockerSnapshotRaw represents all the information related to a snapshot as returned by the Docker API
	DockerSnapshotRaw struct {
		Containers interface{} `json:"Containers"`
		Volumes    interface{} `json:"Volumes"`
		Networks   interface{} `json:"Networks"`
		Images     interface{} `json:"Images"`
		Info       interface{} `json:"Info"`
		Version    interface{} `json:"Version"`
	}

	// EndpointID represents an environment(endpoint) identifier
	EndpointID int

	// EndpointStatus represents the status of an environment(endpoint)
	EndpointStatus int

	// EndpointType represents the type of an environment(endpoint)
	EndpointType int

	// EndpointGroupID represents an environment(endpoint) group identifier
	EndpointGroupID int

	// EndpointExtension represents a deprecated form of Portainer extension
	// TODO: legacy extension management
	EndpointExtension struct {
		Type EndpointExtensionType `json:"Type"`
		URL  string                `json:"URL"`
	}

	// EndpointExtensionType represents the type of an environment(endpoint) extension. Only
	// one extension of each type can be associated to an environment(endpoint)
	EndpointExtensionType int

	// EndpointSecuritySettings represents settings for an environment(endpoint)
	EndpointSecuritySettings struct {
		// Whether non-administrator should be able to use bind mounts when creating containers
		AllowBindMountsForRegularUsers bool `json:"allowBindMountsForRegularUsers" example:"false"`
		// Whether non-administrator should be able to use privileged mode when creating containers
		AllowPrivilegedModeForRegularUsers bool `json:"allowPrivilegedModeForRegularUsers" example:"false"`
		// Whether non-administrator should be able to browse volumes
		AllowVolumeBrowserForRegularUsers bool `json:"allowVolumeBrowserForRegularUsers" example:"true"`
		// Whether non-administrator should be able to use the host pid
		AllowHostNamespaceForRegularUsers bool `json:"allowHostNamespaceForRegularUsers" example:"true"`
		// Whether non-administrator should be able to use device mapping
		AllowDeviceMappingForRegularUsers bool `json:"allowDeviceMappingForRegularUsers" example:"true"`
		// Whether non-administrator should be able to manage stacks
		AllowStackManagementForRegularUsers bool `json:"allowStackManagementForRegularUsers" example:"true"`
		// Whether non-administrator should be able to use container capabilities
		AllowContainerCapabilitiesForRegularUsers bool `json:"allowContainerCapabilitiesForRegularUsers" example:"true"`
		// Whether non-administrator should be able to use sysctl settings
		AllowSysctlSettingForRegularUsers bool `json:"allowSysctlSettingForRegularUsers" example:"true"`
		// Whether host management features are enabled
		EnableHostManagementFeatures bool `json:"enableHostManagementFeatures" example:"true"`
	}

	// Environment(Endpoint) represents a Docker environment(endpoint) with all the info required
	// to connect to it
	Endpoint struct {
		// Environment(Endpoint) Identifier
		ID EndpointID `json:"Id" example:"1"`
		// Environment(Endpoint) name
		Name string `json:"Name" example:"my-environment"`
		// Environment(Endpoint) environment(endpoint) type. 1 for a Docker environment(endpoint), 2 for an agent on Docker environment(endpoint) or 3 for an Azure environment(endpoint).
		Type EndpointType `json:"Type" example:"1"`
		// URL or IP address of the Docker host associated to this environment(endpoint)
		URL string `json:"URL" example:"docker.mydomain.tld:2375"`
		// Environment(Endpoint) group identifier
		GroupID EndpointGroupID `json:"GroupId" example:"1"`
		// URL or IP address where exposed containers will be reachable
		PublicURL        string              `json:"PublicURL" example:"docker.mydomain.tld:2375"`
		TLSConfig        TLSConfiguration    `json:"TLSConfig"`
		Extensions       []EndpointExtension `json:"Extensions" example:""`
		AzureCredentials AzureCredentials    `json:"AzureCredentials,omitempty" example:""`
		// List of tag identifiers to which this environment(endpoint) is associated
		TagIDs []TagID `json:"TagIds"`
		// The status of the environment(endpoint) (1 - up, 2 - down)
		Status EndpointStatus `json:"Status" example:"1"`
		// List of snapshots
		Snapshots []DockerSnapshot `json:"Snapshots" example:""`
		// List of user identifiers authorized to connect to this environment(endpoint)
		UserAccessPolicies UserAccessPolicies `json:"UserAccessPolicies"`
		// List of team identifiers authorized to connect to this environment(endpoint)
		TeamAccessPolicies TeamAccessPolicies `json:"TeamAccessPolicies" example:""`
		// The identifier of the edge agent associated with this environment(endpoint)
		EdgeID string `json:"EdgeID,omitempty" example:""`
		// The key which is used to map the agent to Portainer
		EdgeKey string `json:"EdgeKey" example:""`
		// The check in interval for edge agent (in seconds)
		EdgeCheckinInterval int `json:"EdgeCheckinInterval" example:"5"`
		// Associated Kubernetes data
		//Kubernetes KubernetesData `json:"Kubernetes" example:""`
		// Maximum version of docker-compose
		ComposeSyntaxMaxVersion string `json:"ComposeSyntaxMaxVersion" example:"3.8"`
		// Environment(Endpoint) specific security settings
		SecuritySettings EndpointSecuritySettings
		// LastCheckInDate mark last check-in date on checkin
		LastCheckInDate int64

		// Deprecated fields
		// Deprecated in DBVersion == 4
		TLS           bool   `json:"TLS,omitempty"`
		TLSCACertPath string `json:"TLSCACert,omitempty"`
		TLSCertPath   string `json:"TLSCert,omitempty"`
		TLSKeyPath    string `json:"TLSKey,omitempty"`

		// Deprecated in DBVersion == 18
		AuthorizedUsers []UserID `json:"AuthorizedUsers"`
		AuthorizedTeams []TeamID `json:"AuthorizedTeams"`

		// Deprecated in DBVersion == 22
		Tags []string `json:"Tags"`
	}

	// TLSConfiguration represents a TLS configuration
	TLSConfiguration struct {
		// Use TLS
		TLS bool `json:"TLS" example:"true"`
		// Skip the verification of the server TLS certificate
		TLSSkipVerify bool `json:"TLSSkipVerify" example:"false"`
		// Path to the TLS CA certificate file
		TLSCACertPath string `json:"TLSCACert,omitempty" example:"/data/tls/ca.pem"`
		// Path to the TLS client certificate file
		TLSCertPath string `json:"TLSCert,omitempty" example:"/data/tls/cert.pem"`
		// Path to the TLS client key file
		TLSKeyPath string `json:"TLSKey,omitempty" example:"/data/tls/key.pem"`
	}

	// TeamAccessPolicies represent the association of an access policy and a team
	TeamAccessPolicies map[TeamID]AccessPolicy

	// TeamID represents a team identifier
	TeamID int

	// RoleID represents a role identifier
	RoleID int

	// TagID represents a tag identifier
	TagID int

	// UserAccessPolicies represent the association of an access policy and a user
	UserAccessPolicies map[UserID]AccessPolicy

	// UserID represents a user identifier
	UserID int

	// StackID represents a stack identifier (it must be composed of Name + "_" + SwarmID to create a unique identifier)
	StackID int

	// StackStatus represent a status for a stack
	StackStatus int

	// StackType Docker in Docker / Docker in Compose / Docker in Swarm
	StackType int

	//StackAutoUpdate represents the git auto sync config for stack deployment
	StackAutoUpdate struct {
		// Auto update interval
		Interval string `example:"1m30s"`
		// A UUID generated from client
		Webhook string `example:"05de31a2-79fa-4644-9c12-faa67e5c49f0"`
		// Autoupdate job id
		JobID string `example:"15"`
	}

	// Pair defines a key/value string pair
	Pair struct {
		Name  string `json:"name" example:"name"`
		Value string `json:"value" example:"value"`
	}

	//// ResourceControl represent a reference to a Docker resource with specific access controls
	//ResourceControl struct {
	//	// ResourceControl Identifier
	//	ID ResourceControlID `json:"Id" example:"1"`
	//	// Docker resource identifier on which access control will be applied.\
	//	// In the case of a resource control applied to a stack, use the stack name as identifier
	//	ResourceID string `json:"ResourceId" example:"617c5f22bb9b023d6daab7cba43a57576f83492867bc767d1c59416b065e5f08"`
	//	// List of Docker resources that will inherit this access control
	//	SubResourceIDs []string `json:"SubResourceIds" example:"617c5f22bb9b023d6daab7cba43a57576f83492867bc767d1c59416b065e5f08"`
	//	// Type of Docker resource. Valid values are: 1- container, 2 -service
	//	// 3 - volume, 4 - secret, 5 - stack, 6 - config or 7 - custom template
	//	Type         ResourceControlType  `json:"Type" example:"1"`
	//	UserAccesses []UserResourceAccess `json:"UserAccesses" example:""`
	//	TeamAccesses []TeamResourceAccess `json:"TeamAccesses" example:""`
	//	// Permit access to the associated resource to any user
	//	Public bool `json:"Public" example:"true"`
	//	// Permit access to resource only to admins
	//	AdministratorsOnly bool `json:"AdministratorsOnly" example:"true"`
	//	System             bool `json:"System" example:""`
	//
	//	// Deprecated fields
	//	// Deprecated in DBVersion == 2
	//	OwnerID     UserID              `json:"OwnerId,omitempty"`
	//	AccessLevel ResourceAccessLevel `json:"AccessLevel,omitempty"`
	//}
)

const (
	_ StackType = iota
	// DockerSwarmStack represents a stack managed via docker stack
	DockerSwarmStack
	// DockerStack represents a stack managed via docker run
	DockerStack
	// DockerComposeStack represents a stack managed via docker-compose
	DockerComposeStack
	// KubernetesStack represents a stack managed via Kubernetes
	KubernetesStack
)

type (
	ClientFactory interface {
		CreateClient(endpoint docker.DocEndpoint) (*client.Client, error)
		//CreateRegistryAuthenticationHeader() (string, error)
	}

	FileSystemService interface {
		StoreStackFileFromBytes(stackIdentifier, fileName string, data []byte) (string, error)
		CreateKubernetesFile(stackIdentifier string, conf docker.KubernetesConfig) (files []string, err error)
		RemoveDirectory(directoryPath string) error
		GetFileContent(trustedRootPath, filePath string) ([]byte, error)
		GetProjectPath(fileName string) string
	}

	//// Stack represents a Docker stack created via docker stack deploy
	//Stack struct {
	//	// Stack Identifier
	//	//ID StackID `json:"Id" example:"1"`
	//	// Stack name
	//	Name string `json:"Name" example:"myStack"`
	//	// Stack type. 1 for a Swarm stack, 2 for a Compose stack
	//	Type StackType `json:"Type" example:"2"`
	//	// Environment(Endpoint) identifier. Reference the environment(endpoint) that will be used for deployment
	//	// environment（environment）environment。environment（environment）
	//	//EndpointID EndpointID `json:"EndpointId" example:"1"`
	//	// Cluster identifier of the Swarm cluster where the stack is deployed
	//	// Cluster identifier of the cluster where the stack is deployed
	//	SwarmID string `json:"SwarmId" example:"jpofkc0i9uo9wtx1zesuk649w"`
	//	// Path to the Stack file
	//	// Path to stack file
	//	EntryPoint string `json:"EntryPoint" example:"docker-compose.yml"`
	//	// A list of environment(endpoint) variables used during stack deployment
	//	// Environment used during stack deployment（Environment used during stack deployment）Environment used during stack deployment
	//	Env []Pair `json:"Env" example:""`
	//	//ResourceControl *ResourceControl `json:"ResourceControl" example:""`
	//	// Stack status (1 - active, 2 - inactive)
	//	// Stack State（1-Stack State，2-Stack State）
	//	Status StackStatus `json:"Status" example:"1"`
	//	// Path on disk to the repository hosting the Stack file
	//	// The path to the repository where stack files are stored on disk
	//	ProjectPath string `example:"/data/compose/myStack_jpofkc0i9uo9wtx1zesuk649w"`
	//	// The date in unix time when stack was created
	//	// stayunixstay
	//	CreationDate int64 `example:"1587399600"`
	//	// The username which created this stack
	//	// The username for creating this stack
	//	CreatedBy string `example:"admin"`
	//	// The date in unix time when stack was last updated
	//	// Last updated stackunixLast updated stack
	//	UpdateDate int64 `example:"1587399600"`
	//	// The username which last updated this stack
	//	// Last updated username for this stack
	//	UpdatedBy string `example:"bob"`
	//	// Only applies when deploying stack with multiple files
	//	// Applicable only when deploying stacks containing multiple files
	//	AdditionalFiles []string `json:"AdditionalFiles"`
	//	// The auto update settings of a git stack
	//	// gitAutomatic update settings for the stack
	//	AutoUpdate *StackAutoUpdate `json:"AutoUpdate"`
	//	// The git config of this stack
	//	// This stack'sgitThis stack's
	//	GitConfig *gittypes.RepoConfig
	//	// Whether the stack is from a app template
	//	// Does the stack come from an application template
	//	FromAppTemplate bool `example:"false"`
	//	// Kubernetes namespace if stack is a kube application
	//	// If the stack iskubeIf the stack is，If the stack isKubernetesIf the stack is
	//	Namespace string `example:"default"`
	//	// IsComposeFormat indicates if the Kubernetes stack is created from a Docker Compose file
	//	// IsComposeFormatindicateKubernetesindicateDocker Composeindicate
	//	IsComposeFormat bool `example:"false"`
	//}

	// SwarmStackManager represents a service to manage Swarm stacks
	SwarmStackManager interface {
		Login(registry []docker.DocRegistry, endpoint docker.DocEndpoint) error
		Logout(endpoint docker.DocEndpoint) error
		Deploy(stack docker.DocSoftware, prune bool) error
		Remove(name string, endpoint docker.DocEndpoint) error
		NormalizeStackName(name string) string
		Init(stack docker.DocSoftware) error
	}

	// ComposeStackManager represents a service to manage Compose stacks
	ComposeStackManager interface {
		ComposeSyntaxMaxVersion() string
		NormalizeStackName(name string) string
		Up(ctx context.Context, software docker.DocSoftware, isLog bool) error
		Down(ctx context.Context, software docker.DocSoftware) error
		Stop(ctx context.Context, software docker.DocSoftware) error
		Start(ctx context.Context, software docker.DocSoftware) error
		Login(registry []docker.DocRegistry, endpoint docker.DocEndpoint) error
		Logout(endpoint docker.DocEndpoint) error
	}

	ComposeDeployer interface {
		Deploy(ctx context.Context, workingDir, host, projectName string, filePaths []string, envFilePath string, isLog bool) error
		Remove(ctx context.Context, workingDir, host, projectName string, filePaths []string) error
		Stop(ctx context.Context, workingDir, host, projectName string, filePaths []string) error
		Start(ctx context.Context, workingDir, host, projectName string, filePaths []string) error
	}
)

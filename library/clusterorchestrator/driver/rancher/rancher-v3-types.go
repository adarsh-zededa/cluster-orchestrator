package rancher

// Type represents the stored type of IntOrString.
type Type int64

const (
	Int    Type = iota // The IntOrString holds an int.
	String             // The IntOrString holds a string.
)

type Resource struct {
	ID      string            `json:"id,omitempty"`
	Type    string            `json:"type,omitempty"`
	Links   map[string]string `json:"links"`
	Actions map[string]string `json:"actions"`
}

type ConfigMapKeySelector struct {
	Key      string `json:"key,omitempty" yaml:"key,omitempty"`
	Name     string `json:"name,omitempty" yaml:"name,omitempty"`
	Optional *bool  `json:"optional,omitempty" yaml:"optional,omitempty"`
}

type ObjectFieldSelector struct {
	APIVersion string `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`
	FieldPath  string `json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
}

type ResourceFieldSelector struct {
	ContainerName string `json:"containerName,omitempty" yaml:"containerName,omitempty"`
	Divisor       string `json:"divisor,omitempty" yaml:"divisor,omitempty"`
	Resource      string `json:"resource,omitempty" yaml:"resource,omitempty"`
}

type SecretKeySelector struct {
	Key      string `json:"key,omitempty" yaml:"key,omitempty"`
	Name     string `json:"name,omitempty" yaml:"name,omitempty"`
	Optional *bool  `json:"optional,omitempty" yaml:"optional,omitempty"`
}

type EnvVarSource struct {
	ConfigMapKeyRef  *ConfigMapKeySelector  `json:"configMapKeyRef,omitempty" yaml:"configMapKeyRef,omitempty"`
	FieldRef         *ObjectFieldSelector   `json:"fieldRef,omitempty" yaml:"fieldRef,omitempty"`
	ResourceFieldRef *ResourceFieldSelector `json:"resourceFieldRef,omitempty" yaml:"resourceFieldRef,omitempty"`
	SecretKeyRef     *SecretKeySelector     `json:"secretKeyRef,omitempty" yaml:"secretKeyRef,omitempty"`
}

type EnvVar struct {
	Name      string        `json:"name,omitempty" yaml:"name,omitempty"`
	Value     string        `json:"value,omitempty" yaml:"value,omitempty"`
	ValueFrom *EnvVarSource `json:"valueFrom,omitempty" yaml:"valueFrom,omitempty"`
}

type Answer struct {
	ClusterID string            `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	ProjectID string            `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	Values    map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
}

type SubQuestion struct {
	Default      string   `json:"default,omitempty" yaml:"default,omitempty"`
	Description  string   `json:"description,omitempty" yaml:"description,omitempty"`
	Group        string   `json:"group,omitempty" yaml:"group,omitempty"`
	InvalidChars string   `json:"invalidChars,omitempty" yaml:"invalidChars,omitempty"`
	Label        string   `json:"label,omitempty" yaml:"label,omitempty"`
	Max          int64    `json:"max,omitempty" yaml:"max,omitempty"`
	MaxLength    int64    `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	Min          int64    `json:"min,omitempty" yaml:"min,omitempty"`
	MinLength    int64    `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	Options      []string `json:"options,omitempty" yaml:"options,omitempty"`
	Required     bool     `json:"required,omitempty" yaml:"required,omitempty"`
	Satisfies    string   `json:"satisfies,omitempty" yaml:"satisfies,omitempty"`
	ShowIf       string   `json:"showIf,omitempty" yaml:"showIf,omitempty"`
	Type         string   `json:"type,omitempty" yaml:"type,omitempty"`
	ValidChars   string   `json:"validChars,omitempty" yaml:"validChars,omitempty"`
	Variable     string   `json:"variable,omitempty" yaml:"variable,omitempty"`
}

type Question struct {
	Default           string        `json:"default,omitempty" yaml:"default,omitempty"`
	Description       string        `json:"description,omitempty" yaml:"description,omitempty"`
	Group             string        `json:"group,omitempty" yaml:"group,omitempty"`
	InvalidChars      string        `json:"invalidChars,omitempty" yaml:"invalidChars,omitempty"`
	Label             string        `json:"label,omitempty" yaml:"label,omitempty"`
	Max               int64         `json:"max,omitempty" yaml:"max,omitempty"`
	MaxLength         int64         `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	Min               int64         `json:"min,omitempty" yaml:"min,omitempty"`
	MinLength         int64         `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	Options           []string      `json:"options,omitempty" yaml:"options,omitempty"`
	Required          bool          `json:"required,omitempty" yaml:"required,omitempty"`
	Satisfies         string        `json:"satisfies,omitempty" yaml:"satisfies,omitempty"`
	ShowIf            string        `json:"showIf,omitempty" yaml:"showIf,omitempty"`
	ShowSubquestionIf string        `json:"showSubquestionIf,omitempty" yaml:"showSubquestionIf,omitempty"`
	Subquestions      []SubQuestion `json:"subquestions,omitempty" yaml:"subquestions,omitempty"`
	Type              string        `json:"type,omitempty" yaml:"type,omitempty"`
	ValidChars        string        `json:"validChars,omitempty" yaml:"validChars,omitempty"`
	Variable          string        `json:"variable,omitempty" yaml:"variable,omitempty"`
}

type ClusterComponentStatus struct {
	Conditions []ComponentCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	Name       string               `json:"name,omitempty" yaml:"name,omitempty"`
}

type LaunchTemplate struct {
	Name    string `json:"name,omitempty" yaml:"name,omitempty"`
	Version *int64 `json:"version,omitempty" yaml:"version,omitempty"`
}

type NodeGroup struct {
	DesiredSize          *int64            `json:"desiredSize,omitempty" yaml:"desiredSize,omitempty"`
	DiskSize             *int64            `json:"diskSize,omitempty" yaml:"diskSize,omitempty"`
	Ec2SshKey            string            `json:"ec2SshKey,omitempty" yaml:"ec2SshKey,omitempty"`
	Gpu                  *bool             `json:"gpu,omitempty" yaml:"gpu,omitempty"`
	ImageID              string            `json:"imageId,omitempty" yaml:"imageId,omitempty"`
	InstanceType         string            `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`
	Labels               map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	LaunchTemplate       *LaunchTemplate   `json:"launchTemplate,omitempty" yaml:"launchTemplate,omitempty"`
	MaxSize              *int64            `json:"maxSize,omitempty" yaml:"maxSize,omitempty"`
	MinSize              *int64            `json:"minSize,omitempty" yaml:"minSize,omitempty"`
	NodegroupName        string            `json:"nodegroupName,omitempty" yaml:"nodegroupName,omitempty"`
	RequestSpotInstances *bool             `json:"requestSpotInstances,omitempty" yaml:"requestSpotInstances,omitempty"`
	ResourceTags         map[string]string `json:"resourceTags,omitempty" yaml:"resourceTags,omitempty"`
	SpotInstanceTypes    []string          `json:"spotInstanceTypes,omitempty" yaml:"spotInstanceTypes,omitempty"`
	Subnets              []string          `json:"subnets,omitempty" yaml:"subnets,omitempty"`
	Tags                 map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
	UserData             string            `json:"userData,omitempty" yaml:"userData,omitempty"`
	Version              string            `json:"version,omitempty" yaml:"version,omitempty"`
}

type EKSClusterConfigSpec struct {
	AmazonCredentialSecret string            `json:"amazonCredentialSecret,omitempty" yaml:"amazonCredentialSecret,omitempty"`
	DisplayName            string            `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	Imported               bool              `json:"imported,omitempty" yaml:"imported,omitempty"`
	KmsKey                 string            `json:"kmsKey,omitempty" yaml:"kmsKey,omitempty"`
	KubernetesVersion      string            `json:"kubernetesVersion,omitempty" yaml:"kubernetesVersion,omitempty"`
	LoggingTypes           []string          `json:"loggingTypes,omitempty" yaml:"loggingTypes,omitempty"`
	NodeGroups             []NodeGroup       `json:"nodeGroups,omitempty" yaml:"nodeGroups,omitempty"`
	PrivateAccess          *bool             `json:"privateAccess,omitempty" yaml:"privateAccess,omitempty"`
	PublicAccess           *bool             `json:"publicAccess,omitempty" yaml:"publicAccess,omitempty"`
	PublicAccessSources    []string          `json:"publicAccessSources,omitempty" yaml:"publicAccessSources,omitempty"`
	Region                 string            `json:"region,omitempty" yaml:"region,omitempty"`
	SecretsEncryption      *bool             `json:"secretsEncryption,omitempty" yaml:"secretsEncryption,omitempty"`
	SecurityGroups         []string          `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`
	ServiceRole            string            `json:"serviceRole,omitempty" yaml:"serviceRole,omitempty"`
	Subnets                []string          `json:"subnets,omitempty" yaml:"subnets,omitempty"`
	Tags                   map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

type GKEClusterAddons struct {
	HTTPLoadBalancing        bool `json:"httpLoadBalancing,omitempty" yaml:"httpLoadBalancing,omitempty"`
	HorizontalPodAutoscaling bool `json:"horizontalPodAutoscaling,omitempty" yaml:"horizontalPodAutoscaling,omitempty"`
	NetworkPolicyConfig      bool `json:"networkPolicyConfig,omitempty" yaml:"networkPolicyConfig,omitempty"`
}

type GKEIPAllocationPolicy struct {
	ClusterIpv4CidrBlock       string `json:"clusterIpv4CidrBlock,omitempty" yaml:"clusterIpv4CidrBlock,omitempty"`
	ClusterSecondaryRangeName  string `json:"clusterSecondaryRangeName,omitempty" yaml:"clusterSecondaryRangeName,omitempty"`
	CreateSubnetwork           bool   `json:"createSubnetwork,omitempty" yaml:"createSubnetwork,omitempty"`
	NodeIpv4CidrBlock          string `json:"nodeIpv4CidrBlock,omitempty" yaml:"nodeIpv4CidrBlock,omitempty"`
	ServicesIpv4CidrBlock      string `json:"servicesIpv4CidrBlock,omitempty" yaml:"servicesIpv4CidrBlock,omitempty"`
	ServicesSecondaryRangeName string `json:"servicesSecondaryRangeName,omitempty" yaml:"servicesSecondaryRangeName,omitempty"`
	SubnetworkName             string `json:"subnetworkName,omitempty" yaml:"subnetworkName,omitempty"`
	UseIPAliases               bool   `json:"useIpAliases,omitempty" yaml:"useIpAliases,omitempty"`
}

type GKECidrBlock struct {
	CidrBlock   string `json:"cidrBlock,omitempty" yaml:"cidrBlock,omitempty"`
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}

type GKEMasterAuthorizedNetworksConfig struct {
	CidrBlocks []GKECidrBlock `json:"cidrBlocks,omitempty" yaml:"cidrBlocks,omitempty"`
	Enabled    bool           `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}

type GKENodePoolAutoscaling struct {
	Enabled      bool  `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	MaxNodeCount int64 `json:"maxNodeCount,omitempty" yaml:"maxNodeCount,omitempty"`
	MinNodeCount int64 `json:"minNodeCount,omitempty" yaml:"minNodeCount,omitempty"`
}

type GKENodeTaintConfig struct {
	Effect string `json:"effect,omitempty" yaml:"effect,omitempty"`
	Key    string `json:"key,omitempty" yaml:"key,omitempty"`
	Value  string `json:"value,omitempty" yaml:"value,omitempty"`
}

type GKENodeConfig struct {
	DiskSizeGb    int64                `json:"diskSizeGb,omitempty" yaml:"diskSizeGb,omitempty"`
	DiskType      string               `json:"diskType,omitempty" yaml:"diskType,omitempty"`
	ImageType     string               `json:"imageType,omitempty" yaml:"imageType,omitempty"`
	Labels        map[string]string    `json:"labels,omitempty" yaml:"labels,omitempty"`
	LocalSsdCount int64                `json:"localSsdCount,omitempty" yaml:"localSsdCount,omitempty"`
	MachineType   string               `json:"machineType,omitempty" yaml:"machineType,omitempty"`
	OauthScopes   []string             `json:"oauthScopes,omitempty" yaml:"oauthScopes,omitempty"`
	Preemptible   bool                 `json:"preemptible,omitempty" yaml:"preemptible,omitempty"`
	Taints        []GKENodeTaintConfig `json:"taints,omitempty" yaml:"taints,omitempty"`
}

type GKENodePoolManagement struct {
	AutoRepair  bool `json:"autoRepair,omitempty" yaml:"autoRepair,omitempty"`
	AutoUpgrade bool `json:"autoUpgrade,omitempty" yaml:"autoUpgrade,omitempty"`
}

type GKENodePoolConfig struct {
	Autoscaling       *GKENodePoolAutoscaling `json:"autoscaling,omitempty" yaml:"autoscaling,omitempty"`
	Config            *GKENodeConfig          `json:"config,omitempty" yaml:"config,omitempty"`
	InitialNodeCount  *int64                  `json:"initialNodeCount,omitempty" yaml:"initialNodeCount,omitempty"`
	Management        *GKENodePoolManagement  `json:"management,omitempty" yaml:"management,omitempty"`
	MaxPodsConstraint *int64                  `json:"maxPodsConstraint,omitempty" yaml:"maxPodsConstraint,omitempty"`
	Name              string                  `json:"name,omitempty" yaml:"name,omitempty"`
	Version           string                  `json:"version,omitempty" yaml:"version,omitempty"`
}

type GKEPrivateClusterConfig struct {
	EnablePrivateEndpoint bool   `json:"enablePrivateEndpoint,omitempty" yaml:"enablePrivateEndpoint,omitempty"`
	EnablePrivateNodes    bool   `json:"enablePrivateNodes,omitempty" yaml:"enablePrivateNodes,omitempty"`
	MasterIpv4CidrBlock   string `json:"masterIpv4CidrBlock,omitempty" yaml:"masterIpv4CidrBlock,omitempty"`
}

type GKEClusterConfigSpec struct {
	ClusterAddons                  *GKEClusterAddons                  `json:"clusterAddons,omitempty" yaml:"clusterAddons,omitempty"`
	ClusterIpv4CidrBlock           string                             `json:"clusterIpv4Cidr,omitempty" yaml:"clusterIpv4Cidr,omitempty"`
	ClusterName                    string                             `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`
	Description                    string                             `json:"description,omitempty" yaml:"description,omitempty"`
	EnableKubernetesAlpha          *bool                              `json:"enableKubernetesAlpha,omitempty" yaml:"enableKubernetesAlpha,omitempty"`
	GoogleCredentialSecret         string                             `json:"googleCredentialSecret,omitempty" yaml:"googleCredentialSecret,omitempty"`
	IPAllocationPolicy             *GKEIPAllocationPolicy             `json:"ipAllocationPolicy,omitempty" yaml:"ipAllocationPolicy,omitempty"`
	Imported                       bool                               `json:"imported,omitempty" yaml:"imported,omitempty"`
	KubernetesVersion              string                             `json:"kubernetesVersion,omitempty" yaml:"kubernetesVersion,omitempty"`
	Locations                      []string                           `json:"locations,omitempty" yaml:"locations,omitempty"`
	LoggingService                 string                             `json:"loggingService,omitempty" yaml:"loggingService,omitempty"`
	MaintenanceWindow              string                             `json:"maintenanceWindow,omitempty" yaml:"maintenanceWindow,omitempty"`
	MasterAuthorizedNetworksConfig *GKEMasterAuthorizedNetworksConfig `json:"masterAuthorizedNetworks,omitempty" yaml:"masterAuthorizedNetworks,omitempty"`
	MonitoringService              string                             `json:"monitoringService,omitempty" yaml:"monitoringService,omitempty"`
	Network                        string                             `json:"network,omitempty" yaml:"network,omitempty"`
	NetworkPolicyEnabled           *bool                              `json:"networkPolicyEnabled,omitempty" yaml:"networkPolicyEnabled,omitempty"`
	NodePools                      []GKENodePoolConfig                `json:"nodePools,omitempty" yaml:"nodePools,omitempty"`
	PrivateClusterConfig           *GKEPrivateClusterConfig           `json:"privateClusterConfig,omitempty" yaml:"privateClusterConfig,omitempty"`
	ProjectID                      string                             `json:"projectID,omitempty" yaml:"projectID,omitempty"`
	Region                         string                             `json:"region,omitempty" yaml:"region,omitempty"`
	Subnetwork                     string                             `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`
	Zone                           string                             `json:"zone,omitempty" yaml:"zone,omitempty"`
}

type ImportedConfig struct {
	KubeConfig string `json:"kubeConfig,omitempty" yaml:"kubeConfig,omitempty"`
}

type ClusterUpgradeStrategy struct {
	DrainServerNodes  bool  `json:"drainServerNodes,omitempty" yaml:"drainServerNodes,omitempty"`
	DrainWorkerNodes  bool  `json:"drainWorkerNodes,omitempty" yaml:"drainWorkerNodes,omitempty"`
	ServerConcurrency int64 `json:"serverConcurrency,omitempty" yaml:"serverConcurrency,omitempty"`
	WorkerConcurrency int64 `json:"workerConcurrency,omitempty" yaml:"workerConcurrency,omitempty"`
}

type K3sConfig struct {
	ClusterUpgradeStrategy *ClusterUpgradeStrategy `json:"k3supgradeStrategy,omitempty" yaml:"k3supgradeStrategy,omitempty"`
	Version                string                  `json:"kubernetesVersion,omitempty" yaml:"kubernetesVersion,omitempty"`
}

type LocalClusterAuthEndpoint struct {
	CACerts string `json:"caCerts,omitempty" yaml:"caCerts,omitempty"`
	Enabled bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	FQDN    string `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`
}

type AuthWebhookConfig struct {
	CacheTimeout string `json:"cacheTimeout,omitempty" yaml:"cacheTimeout,omitempty"`
	ConfigFile   string `json:"configFile,omitempty" yaml:"configFile,omitempty"`
}

type AuthnConfig struct {
	SANs     []string           `json:"sans,omitempty" yaml:"sans,omitempty"`
	Strategy string             `json:"strategy,omitempty" yaml:"strategy,omitempty"`
	Webhook  *AuthWebhookConfig `json:"webhook,omitempty" yaml:"webhook,omitempty"`
}

type AuthzConfig struct {
	Mode    string            `json:"mode,omitempty" yaml:"mode,omitempty"`
	Options map[string]string `json:"options,omitempty" yaml:"options,omitempty"`
}

type BastionHost struct {
	Address      string `json:"address,omitempty" yaml:"address,omitempty"`
	Port         string `json:"port,omitempty" yaml:"port,omitempty"`
	SSHAgentAuth bool   `json:"sshAgentAuth,omitempty" yaml:"sshAgentAuth,omitempty"`
	SSHCert      string `json:"sshCert,omitempty" yaml:"sshCert,omitempty"`
	SSHCertPath  string `json:"sshCertPath,omitempty" yaml:"sshCertPath,omitempty"`
	SSHKey       string `json:"sshKey,omitempty" yaml:"sshKey,omitempty"`
	SSHKeyPath   string `json:"sshKeyPath,omitempty" yaml:"sshKeyPath,omitempty"`
	User         string `json:"user,omitempty" yaml:"user,omitempty"`
}

type GlobalAwsOpts struct {
	DisableSecurityGroupIngress bool   `json:"disable-security-group-ingress,omitempty" yaml:"disable-security-group-ingress,omitempty"`
	DisableStrictZoneCheck      bool   `json:"disable-strict-zone-check,omitempty" yaml:"disable-strict-zone-check,omitempty"`
	ElbSecurityGroup            string `json:"elb-security-group,omitempty" yaml:"elb-security-group,omitempty"`
	KubernetesClusterID         string `json:"kubernetes-cluster-id,omitempty" yaml:"kubernetes-cluster-id,omitempty"`
	KubernetesClusterTag        string `json:"kubernetes-cluster-tag,omitempty" yaml:"kubernetes-cluster-tag,omitempty"`
	RoleARN                     string `json:"role-arn,omitempty" yaml:"role-arn,omitempty"`
	RouteTableID                string `json:"routetable-id,omitempty" yaml:"routetable-id,omitempty"`
	SubnetID                    string `json:"subnet-id,omitempty" yaml:"subnet-id,omitempty"`
	VPC                         string `json:"vpc,omitempty" yaml:"vpc,omitempty"`
	Zone                        string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

type ServiceOverride struct {
	Region        string `json:"region,omitempty" yaml:"region,omitempty"`
	Service       string `json:"service,omitempty" yaml:"service,omitempty"`
	SigningMethod string `json:"signing-method,omitempty" yaml:"signing-method,omitempty"`
	SigningName   string `json:"signing-name,omitempty" yaml:"signing-name,omitempty"`
	SigningRegion string `json:"signing-region,omitempty" yaml:"signing-region,omitempty"`
	URL           string `json:"url,omitempty" yaml:"url,omitempty"`
}

type AWSCloudProvider struct {
	Global          *GlobalAwsOpts             `json:"global,omitempty" yaml:"global,omitempty"`
	ServiceOverride map[string]ServiceOverride `json:"serviceOverride,omitempty" yaml:"serviceOverride,omitempty"`
}

type AzureCloudProvider struct {
	AADClientCertPassword        string `json:"aadClientCertPassword,omitempty" yaml:"aadClientCertPassword,omitempty"`
	AADClientCertPath            string `json:"aadClientCertPath,omitempty" yaml:"aadClientCertPath,omitempty"`
	AADClientID                  string `json:"aadClientId,omitempty" yaml:"aadClientId,omitempty"`
	AADClientSecret              string `json:"aadClientSecret,omitempty" yaml:"aadClientSecret,omitempty"`
	Cloud                        string `json:"cloud,omitempty" yaml:"cloud,omitempty"`
	CloudProviderBackoff         bool   `json:"cloudProviderBackoff,omitempty" yaml:"cloudProviderBackoff,omitempty"`
	CloudProviderBackoffDuration int64  `json:"cloudProviderBackoffDuration,omitempty" yaml:"cloudProviderBackoffDuration,omitempty"`
	CloudProviderBackoffExponent int64  `json:"cloudProviderBackoffExponent,omitempty" yaml:"cloudProviderBackoffExponent,omitempty"`
	CloudProviderBackoffJitter   int64  `json:"cloudProviderBackoffJitter,omitempty" yaml:"cloudProviderBackoffJitter,omitempty"`
	CloudProviderBackoffRetries  int64  `json:"cloudProviderBackoffRetries,omitempty" yaml:"cloudProviderBackoffRetries,omitempty"`
	CloudProviderRateLimit       bool   `json:"cloudProviderRateLimit,omitempty" yaml:"cloudProviderRateLimit,omitempty"`
	CloudProviderRateLimitBucket int64  `json:"cloudProviderRateLimitBucket,omitempty" yaml:"cloudProviderRateLimitBucket,omitempty"`
	CloudProviderRateLimitQPS    int64  `json:"cloudProviderRateLimitQPS,omitempty" yaml:"cloudProviderRateLimitQPS,omitempty"`
	ExcludeMasterFromStandardLB  *bool  `json:"excludeMasterFromStandardLB,omitempty" yaml:"excludeMasterFromStandardLB,omitempty"`
	LoadBalancerSku              string `json:"loadBalancerSku,omitempty" yaml:"loadBalancerSku,omitempty"`
	Location                     string `json:"location,omitempty" yaml:"location,omitempty"`
	MaximumLoadBalancerRuleCount int64  `json:"maximumLoadBalancerRuleCount,omitempty" yaml:"maximumLoadBalancerRuleCount,omitempty"`
	PrimaryAvailabilitySetName   string `json:"primaryAvailabilitySetName,omitempty" yaml:"primaryAvailabilitySetName,omitempty"`
	PrimaryScaleSetName          string `json:"primaryScaleSetName,omitempty" yaml:"primaryScaleSetName,omitempty"`
	ResourceGroup                string `json:"resourceGroup,omitempty" yaml:"resourceGroup,omitempty"`
	RouteTableName               string `json:"routeTableName,omitempty" yaml:"routeTableName,omitempty"`
	SecurityGroupName            string `json:"securityGroupName,omitempty" yaml:"securityGroupName,omitempty"`
	SubnetName                   string `json:"subnetName,omitempty" yaml:"subnetName,omitempty"`
	SubscriptionID               string `json:"subscriptionId,omitempty" yaml:"subscriptionId,omitempty"`
	TenantID                     string `json:"tenantId,omitempty" yaml:"tenantId,omitempty"`
	UseInstanceMetadata          bool   `json:"useInstanceMetadata,omitempty" yaml:"useInstanceMetadata,omitempty"`
	UseManagedIdentityExtension  bool   `json:"useManagedIdentityExtension,omitempty" yaml:"useManagedIdentityExtension,omitempty"`
	UserAssignedIdentityID       string `json:"userAssignedIdentityID,omitempty" yaml:"userAssignedIdentityID,omitempty"`
	VMType                       string `json:"vmType,omitempty" yaml:"vmType,omitempty"`
	VnetName                     string `json:"vnetName,omitempty" yaml:"vnetName,omitempty"`
	VnetResourceGroup            string `json:"vnetResourceGroup,omitempty" yaml:"vnetResourceGroup,omitempty"`
}

type BlockStorageOpenstackOpts struct {
	BSVersion       string `json:"bs-version,omitempty" yaml:"bs-version,omitempty"`
	IgnoreVolumeAZ  bool   `json:"ignore-volume-az,omitempty" yaml:"ignore-volume-az,omitempty"`
	TrustDevicePath bool   `json:"trust-device-path,omitempty" yaml:"trust-device-path,omitempty"`
}

type GlobalOpenstackOpts struct {
	AuthURL    string `json:"auth-url,omitempty" yaml:"auth-url,omitempty"`
	CAFile     string `json:"ca-file,omitempty" yaml:"ca-file,omitempty"`
	DomainID   string `json:"domain-id,omitempty" yaml:"domain-id,omitempty"`
	DomainName string `json:"domain-name,omitempty" yaml:"domain-name,omitempty"`
	Password   string `json:"password,omitempty" yaml:"password,omitempty"`
	Region     string `json:"region,omitempty" yaml:"region,omitempty"`
	TenantID   string `json:"tenant-id,omitempty" yaml:"tenant-id,omitempty"`
	TenantName string `json:"tenant-name,omitempty" yaml:"tenant-name,omitempty"`
	TrustID    string `json:"trust-id,omitempty" yaml:"trust-id,omitempty"`
	UserID     string `json:"user-id,omitempty" yaml:"user-id,omitempty"`
	Username   string `json:"username,omitempty" yaml:"username,omitempty"`
}

type LoadBalancerOpenstackOpts struct {
	CreateMonitor        bool   `json:"create-monitor,omitempty" yaml:"create-monitor,omitempty"`
	FloatingNetworkID    string `json:"floating-network-id,omitempty" yaml:"floating-network-id,omitempty"`
	LBMethod             string `json:"lb-method,omitempty" yaml:"lb-method,omitempty"`
	LBProvider           string `json:"lb-provider,omitempty" yaml:"lb-provider,omitempty"`
	LBVersion            string `json:"lb-version,omitempty" yaml:"lb-version,omitempty"`
	ManageSecurityGroups bool   `json:"manage-security-groups,omitempty" yaml:"manage-security-groups,omitempty"`
	MonitorDelay         string `json:"monitor-delay,omitempty" yaml:"monitor-delay,omitempty"`
	MonitorMaxRetries    int64  `json:"monitor-max-retries,omitempty" yaml:"monitor-max-retries,omitempty"`
	MonitorTimeout       string `json:"monitor-timeout,omitempty" yaml:"monitor-timeout,omitempty"`
	SubnetID             string `json:"subnet-id,omitempty" yaml:"subnet-id,omitempty"`
	UseOctavia           bool   `json:"use-octavia,omitempty" yaml:"use-octavia,omitempty"`
}

type MetadataOpenstackOpts struct {
	RequestTimeout int64  `json:"request-timeout,omitempty" yaml:"request-timeout,omitempty"`
	SearchOrder    string `json:"search-order,omitempty" yaml:"search-order,omitempty"`
}

type RouteOpenstackOpts struct {
	RouterID string `json:"router-id,omitempty" yaml:"router-id,omitempty"`
}

type OpenstackCloudProvider struct {
	BlockStorage *BlockStorageOpenstackOpts `json:"blockStorage,omitempty" yaml:"blockStorage,omitempty"`
	Global       *GlobalOpenstackOpts       `json:"global,omitempty" yaml:"global,omitempty"`
	LoadBalancer *LoadBalancerOpenstackOpts `json:"loadBalancer,omitempty" yaml:"loadBalancer,omitempty"`
	Metadata     *MetadataOpenstackOpts     `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Route        *RouteOpenstackOpts        `json:"route,omitempty" yaml:"route,omitempty"`
}

type DiskVsphereOpts struct {
	SCSIControllerType string `json:"scsicontrollertype,omitempty" yaml:"scsicontrollertype,omitempty"`
}

type GlobalVsphereOpts struct {
	Datacenter        string `json:"datacenter,omitempty" yaml:"datacenter,omitempty"`
	Datacenters       string `json:"datacenters,omitempty" yaml:"datacenters,omitempty"`
	DefaultDatastore  string `json:"datastore,omitempty" yaml:"datastore,omitempty"`
	InsecureFlag      bool   `json:"insecure-flag,omitempty" yaml:"insecure-flag,omitempty"`
	Password          string `json:"password,omitempty" yaml:"password,omitempty"`
	RoundTripperCount int64  `json:"soap-roundtrip-count,omitempty" yaml:"soap-roundtrip-count,omitempty"`
	User              string `json:"user,omitempty" yaml:"user,omitempty"`
	VCenterIP         string `json:"server,omitempty" yaml:"server,omitempty"`
	VCenterPort       string `json:"port,omitempty" yaml:"port,omitempty"`
	VMName            string `json:"vm-name,omitempty" yaml:"vm-name,omitempty"`
	VMUUID            string `json:"vm-uuid,omitempty" yaml:"vm-uuid,omitempty"`
	WorkingDir        string `json:"working-dir,omitempty" yaml:"working-dir,omitempty"`
}

type NetworkVshpereOpts struct {
	PublicNetwork string `json:"public-network,omitempty" yaml:"public-network,omitempty"`
}

type VirtualCenterConfig struct {
	Datacenters       string `json:"datacenters,omitempty" yaml:"datacenters,omitempty"`
	Password          string `json:"password,omitempty" yaml:"password,omitempty"`
	RoundTripperCount int64  `json:"soap-roundtrip-count,omitempty" yaml:"soap-roundtrip-count,omitempty"`
	User              string `json:"user,omitempty" yaml:"user,omitempty"`
	VCenterPort       string `json:"port,omitempty" yaml:"port,omitempty"`
}

type WorkspaceVsphereOpts struct {
	Datacenter       string `json:"datacenter,omitempty" yaml:"datacenter,omitempty"`
	DefaultDatastore string `json:"default-datastore,omitempty" yaml:"default-datastore,omitempty"`
	Folder           string `json:"folder,omitempty" yaml:"folder,omitempty"`
	ResourcePoolPath string `json:"resourcepool-path,omitempty" yaml:"resourcepool-path,omitempty"`
	VCenterIP        string `json:"server,omitempty" yaml:"server,omitempty"`
}

type VsphereCloudProvider struct {
	Disk          *DiskVsphereOpts               `json:"disk,omitempty" yaml:"disk,omitempty"`
	Global        *GlobalVsphereOpts             `json:"global,omitempty" yaml:"global,omitempty"`
	Network       *NetworkVshpereOpts            `json:"network,omitempty" yaml:"network,omitempty"`
	VirtualCenter map[string]VirtualCenterConfig `json:"virtualCenter,omitempty" yaml:"virtualCenter,omitempty"`
	Workspace     *WorkspaceVsphereOpts          `json:"workspace,omitempty" yaml:"workspace,omitempty"`
}

type CloudProvider struct {
	AWSCloudProvider       *AWSCloudProvider       `json:"awsCloudProvider,omitempty" yaml:"awsCloudProvider,omitempty"`
	AzureCloudProvider     *AzureCloudProvider     `json:"azureCloudProvider,omitempty" yaml:"azureCloudProvider,omitempty"`
	CustomCloudProvider    string                  `json:"customCloudProvider,omitempty" yaml:"customCloudProvider,omitempty"`
	Name                   string                  `json:"name,omitempty" yaml:"name,omitempty"`
	OpenstackCloudProvider *OpenstackCloudProvider `json:"openstackCloudProvider,omitempty" yaml:"openstackCloudProvider,omitempty"`
	VsphereCloudProvider   *VsphereCloudProvider   `json:"vsphereCloudProvider,omitempty" yaml:"vsphereCloudProvider,omitempty"`
}

type LinearAutoscalerParams struct {
	CoresPerReplica           float64 `json:"coresPerReplica,omitempty" yaml:"coresPerReplica,omitempty"`
	Max                       int64   `json:"max,omitempty" yaml:"max,omitempty"`
	Min                       int64   `json:"min,omitempty" yaml:"min,omitempty"`
	NodesPerReplica           float64 `json:"nodesPerReplica,omitempty" yaml:"nodesPerReplica,omitempty"`
	PreventSinglePointFailure bool    `json:"preventSinglePointFailure,omitempty" yaml:"preventSinglePointFailure,omitempty"`
}

type IntOrString struct {
	Type   Type   `json:"type,omitempty" yaml:"type,omitempty"`
	IntVal int32  `json:"intVal,omitempty" yaml:"intVal,omitempty"`
	StrVal string `json:"strVal,omitempty" yaml:"strVal,omitempty"`
}

type RollingUpdateDaemonSet struct {
	MaxUnavailable IntOrString `json:"maxUnavailable,omitempty" yaml:"maxUnavailable,omitempty"`
}

type DaemonSetUpdateStrategy struct {
	RollingUpdate *RollingUpdateDaemonSet `json:"rollingUpdate,omitempty" yaml:"rollingUpdate,omitempty"`
	Strategy      string                  `json:"strategy,omitempty" yaml:"strategy,omitempty"`
}

type Nodelocal struct {
	IPAddress                     string                   `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`
	NodeLocalDNSPriorityClassName string                   `json:"nodeLocalDnsPriorityClassName,omitempty" yaml:"nodeLocalDnsPriorityClassName,omitempty"`
	NodeSelector                  map[string]string        `json:"nodeSelector,omitempty" yaml:"nodeSelector,omitempty"`
	UpdateStrategy                *DaemonSetUpdateStrategy `json:"updateStrategy,omitempty" yaml:"updateStrategy,omitempty"`
}

type Toleration struct {
	Effect            string `json:"effect,omitempty" yaml:"effect,omitempty"`
	Key               string `json:"key,omitempty" yaml:"key,omitempty"`
	Operator          string `json:"operator,omitempty" yaml:"operator,omitempty"`
	TolerationSeconds *int64 `json:"tolerationSeconds,omitempty" yaml:"tolerationSeconds,omitempty"`
	Value             string `json:"value,omitempty" yaml:"value,omitempty"`
}

type RollingUpdateDeployment struct {
	MaxSurge       IntOrString `json:"maxSurge,omitempty" yaml:"maxSurge,omitempty"`
	MaxUnavailable IntOrString `json:"maxUnavailable,omitempty" yaml:"maxUnavailable,omitempty"`
}

type DNSConfig struct {
	LinearAutoscalerParams *LinearAutoscalerParams `json:"linearAutoscalerParams,omitempty" yaml:"linearAutoscalerParams,omitempty"`
	NodeSelector           map[string]string       `json:"nodeSelector,omitempty" yaml:"nodeSelector,omitempty"`
	Nodelocal              *Nodelocal              `json:"nodelocal,omitempty" yaml:"nodelocal,omitempty"`
	Options                map[string]string       `json:"options,omitempty" yaml:"options,omitempty"`
	Provider               string                  `json:"provider,omitempty" yaml:"provider,omitempty"`
	ReverseCIDRs           []string                `json:"reversecidrs,omitempty" yaml:"reversecidrs,omitempty"`
	StubDomains            map[string][]string     `json:"stubdomains,omitempty" yaml:"stubdomains,omitempty"`
	Tolerations            []Toleration            `json:"tolerations,omitempty" yaml:"tolerations,omitempty"`
	UpdateStrategy         *DeploymentStrategy     `json:"updateStrategy,omitempty" yaml:"updateStrategy,omitempty"`
	UpstreamNameservers    []string                `json:"upstreamnameservers,omitempty" yaml:"upstreamnameservers,omitempty"`
}

type DeploymentStrategy struct {
	RollingUpdate *RollingUpdateDeployment `json:"rollingUpdate,omitempty" yaml:"rollingUpdate,omitempty"`
	Strategy      string                   `json:"strategy,omitempty" yaml:"strategy,omitempty"`
}

type IngressConfig struct {
	DNSPolicy                               string                   `json:"dnsPolicy,omitempty" yaml:"dnsPolicy,omitempty"`
	DefaultBackend                          *bool                    `json:"defaultBackend,omitempty" yaml:"defaultBackend,omitempty"`
	DefaultHTTPBackendPriorityClassName     string                   `json:"defaultHttpBackendPriorityClassName,omitempty" yaml:"defaultHttpBackendPriorityClassName,omitempty"`
	ExtraArgs                               map[string]string        `json:"extraArgs,omitempty" yaml:"extraArgs,omitempty"`
	ExtraEnvs                               []interface{}            `json:"extraEnvs,omitempty" yaml:"extraEnvs,omitempty"`
	ExtraVolumeMounts                       []interface{}            `json:"extraVolumeMounts,omitempty" yaml:"extraVolumeMounts,omitempty"`
	ExtraVolumes                            []interface{}            `json:"extraVolumes,omitempty" yaml:"extraVolumes,omitempty"`
	HTTPPort                                int64                    `json:"httpPort,omitempty" yaml:"httpPort,omitempty"`
	HTTPSPort                               int64                    `json:"httpsPort,omitempty" yaml:"httpsPort,omitempty"`
	NetworkMode                             string                   `json:"networkMode,omitempty" yaml:"networkMode,omitempty"`
	NginxIngressControllerPriorityClassName string                   `json:"nginxIngressControllerPriorityClassName,omitempty" yaml:"nginxIngressControllerPriorityClassName,omitempty"`
	NodeSelector                            map[string]string        `json:"nodeSelector,omitempty" yaml:"nodeSelector,omitempty"`
	Options                                 map[string]string        `json:"options,omitempty" yaml:"options,omitempty"`
	Provider                                string                   `json:"provider,omitempty" yaml:"provider,omitempty"`
	Tolerations                             []Toleration             `json:"tolerations,omitempty" yaml:"tolerations,omitempty"`
	UpdateStrategy                          *DaemonSetUpdateStrategy `json:"updateStrategy,omitempty" yaml:"updateStrategy,omitempty"`
}

type MonitoringConfig struct {
	MetricsServerPriorityClassName string              `json:"metricsServerPriorityClassName,omitempty" yaml:"metricsServerPriorityClassName,omitempty"`
	NodeSelector                   map[string]string   `json:"nodeSelector,omitempty" yaml:"nodeSelector,omitempty"`
	Options                        map[string]string   `json:"options,omitempty" yaml:"options,omitempty"`
	Provider                       string              `json:"provider,omitempty" yaml:"provider,omitempty"`
	Replicas                       *int64              `json:"replicas,omitempty" yaml:"replicas,omitempty"`
	Tolerations                    []Toleration        `json:"tolerations,omitempty" yaml:"tolerations,omitempty"`
	UpdateStrategy                 *DeploymentStrategy `json:"updateStrategy,omitempty" yaml:"updateStrategy,omitempty"`
}

type AciNetworkProvider struct {
	AEP                      string   `json:"aep,omitempty" yaml:"aep,omitempty"`
	ApicHosts                []string `json:"apicHosts,omitempty" yaml:"apicHosts,omitempty"`
	ApicRefreshTime          string   `json:"apicRefreshTime,omitempty" yaml:"apicRefreshTime,omitempty"`
	ApicUserCrt              string   `json:"apicUserCrt,omitempty" yaml:"apicUserCrt,omitempty"`
	ApicUserKey              string   `json:"apicUserKey,omitempty" yaml:"apicUserKey,omitempty"`
	ApicUserName             string   `json:"apicUserName,omitempty" yaml:"apicUserName,omitempty"`
	CApic                    string   `json:"capic,omitempty" yaml:"capic,omitempty"`
	ControllerLogLevel       string   `json:"controllerLogLevel,omitempty" yaml:"controllerLogLevel,omitempty"`
	DropLogEnable            string   `json:"dropLogEnable,omitempty" yaml:"dropLogEnable,omitempty"`
	DynamicExternalSubnet    string   `json:"externDynamic,omitempty" yaml:"externDynamic,omitempty"`
	EnableEndpointSlice      string   `json:"enableEndpointSlice,omitempty" yaml:"enableEndpointSlice,omitempty"`
	EncapType                string   `json:"encapType,omitempty" yaml:"encapType,omitempty"`
	EpRegistry               string   `json:"epRegistry,omitempty" yaml:"epRegistry,omitempty"`
	GbpPodSubnet             string   `json:"gbpPodSubnet,omitempty" yaml:"gbpPodSubnet,omitempty"`
	HostAgentLogLevel        string   `json:"hostAgentLogLevel,omitempty" yaml:"hostAgentLogLevel,omitempty"`
	ImagePullPolicy          string   `json:"imagePullPolicy,omitempty" yaml:"imagePullPolicy,omitempty"`
	ImagePullSecret          string   `json:"imagePullSecret,omitempty" yaml:"imagePullSecret,omitempty"`
	InfraVlan                string   `json:"infraVlan,omitempty" yaml:"infraVlan,omitempty"`
	InstallIstio             string   `json:"installIstio,omitempty" yaml:"installIstio,omitempty"`
	IstioProfile             string   `json:"istioProfile,omitempty" yaml:"istioProfile,omitempty"`
	KafkaBrokers             []string `json:"kafkaBrokers,omitempty" yaml:"kafkaBrokers,omitempty"`
	KafkaClientCrt           string   `json:"kafkaClientCrt,omitempty" yaml:"kafkaClientCrt,omitempty"`
	KafkaClientKey           string   `json:"kafkaClientKey,omitempty" yaml:"kafkaClientKey,omitempty"`
	KubeAPIVlan              string   `json:"kubeApiVlan,omitempty" yaml:"kubeApiVlan,omitempty"`
	L3Out                    string   `json:"l3out,omitempty" yaml:"l3out,omitempty"`
	L3OutExternalNetworks    []string `json:"l3outExternalNetworks,omitempty" yaml:"l3outExternalNetworks,omitempty"`
	MaxNodesSvcGraph         string   `json:"maxNodesSvcGraph,omitempty" yaml:"maxNodesSvcGraph,omitempty"`
	McastRangeEnd            string   `json:"mcastRangeEnd,omitempty" yaml:"mcastRangeEnd,omitempty"`
	McastRangeStart          string   `json:"mcastRangeStart,omitempty" yaml:"mcastRangeStart,omitempty"`
	NoPriorityClass          string   `json:"noPriorityClass,omitempty" yaml:"noPriorityClass,omitempty"`
	NodeSubnet               string   `json:"nodeSubnet,omitempty" yaml:"nodeSubnet,omitempty"`
	OVSMemoryLimit           string   `json:"ovsMemoryLimit,omitempty" yaml:"ovsMemoryLimit,omitempty"`
	OpflexAgentLogLevel      string   `json:"opflexLogLevel,omitempty" yaml:"opflexLogLevel,omitempty"`
	OpflexClientSSL          string   `json:"opflexClientSsl,omitempty" yaml:"opflexClientSsl,omitempty"`
	OpflexMode               string   `json:"opflexMode,omitempty" yaml:"opflexMode,omitempty"`
	OpflexServerPort         string   `json:"opflexServerPort,omitempty" yaml:"opflexServerPort,omitempty"`
	OverlayVRFName           string   `json:"overlayVrfName,omitempty" yaml:"overlayVrfName,omitempty"`
	PBRTrackingNonSnat       string   `json:"pbrTrackingNonSnat,omitempty" yaml:"pbrTrackingNonSnat,omitempty"`
	PodSubnetChunkSize       string   `json:"podSubnetChunkSize,omitempty" yaml:"podSubnetChunkSize,omitempty"`
	RunGbpContainer          string   `json:"runGbpContainer,omitempty" yaml:"runGbpContainer,omitempty"`
	RunOpflexServerContainer string   `json:"runOpflexServerContainer,omitempty" yaml:"runOpflexServerContainer,omitempty"`
	ServiceGraphSubnet       string   `json:"nodeSvcSubnet,omitempty" yaml:"nodeSvcSubnet,omitempty"`
	ServiceMonitorInterval   string   `json:"serviceMonitorInterval,omitempty" yaml:"serviceMonitorInterval,omitempty"`
	ServiceVlan              string   `json:"serviceVlan,omitempty" yaml:"serviceVlan,omitempty"`
	SnatContractScope        string   `json:"snatContractScope,omitempty" yaml:"snatContractScope,omitempty"`
	SnatNamespace            string   `json:"snatNamespace,omitempty" yaml:"snatNamespace,omitempty"`
	SnatPortRangeEnd         string   `json:"snatPortRangeEnd,omitempty" yaml:"snatPortRangeEnd,omitempty"`
	SnatPortRangeStart       string   `json:"snatPortRangeStart,omitempty" yaml:"snatPortRangeStart,omitempty"`
	SnatPortsPerNode         string   `json:"snatPortsPerNode,omitempty" yaml:"snatPortsPerNode,omitempty"`
	StaticExternalSubnet     string   `json:"externStatic,omitempty" yaml:"externStatic,omitempty"`
	SubnetDomainName         string   `json:"subnetDomainName,omitempty" yaml:"subnetDomainName,omitempty"`
	SystemIdentifier         string   `json:"systemId,omitempty" yaml:"systemId,omitempty"`
	Tenant                   string   `json:"tenant,omitempty" yaml:"tenant,omitempty"`
	Token                    string   `json:"token,omitempty" yaml:"token,omitempty"`
	UseAciAnywhereCRD        string   `json:"useAciAnywhereCrd,omitempty" yaml:"useAciAnywhereCrd,omitempty"`
	UseAciCniPriorityClass   string   `json:"useAciCniPriorityClass,omitempty" yaml:"useAciCniPriorityClass,omitempty"`
	UseHostNetnsVolume       string   `json:"useHostNetnsVolume,omitempty" yaml:"useHostNetnsVolume,omitempty"`
	UseOpflexServerVolume    string   `json:"useOpflexServerVolume,omitempty" yaml:"useOpflexServerVolume,omitempty"`
	UsePrivilegedContainer   string   `json:"usePrivilegedContainer,omitempty" yaml:"usePrivilegedContainer,omitempty"`
	VRFName                  string   `json:"vrfName,omitempty" yaml:"vrfName,omitempty"`
	VRFTenant                string   `json:"vrfTenant,omitempty" yaml:"vrfTenant,omitempty"`
	VmmController            string   `json:"vmmController,omitempty" yaml:"vmmController,omitempty"`
	VmmDomain                string   `json:"vmmDomain,omitempty" yaml:"vmmDomain,omitempty"`
}

type CalicoNetworkProvider struct {
	CloudProvider string `json:"cloudProvider,omitempty" yaml:"cloudProvider,omitempty"`
}

type CanalNetworkProvider struct {
	Iface string `json:"iface,omitempty" yaml:"iface,omitempty"`
}

type FlannelNetworkProvider struct {
	Iface string `json:"iface,omitempty" yaml:"iface,omitempty"`
}

type WeaveNetworkProvider struct {
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
}

type NetworkConfig struct {
	AciNetworkProvider     *AciNetworkProvider      `json:"aciNetworkProvider,omitempty" yaml:"aciNetworkProvider,omitempty"`
	CalicoNetworkProvider  *CalicoNetworkProvider   `json:"calicoNetworkProvider,omitempty" yaml:"calicoNetworkProvider,omitempty"`
	CanalNetworkProvider   *CanalNetworkProvider    `json:"canalNetworkProvider,omitempty" yaml:"canalNetworkProvider,omitempty"`
	FlannelNetworkProvider *FlannelNetworkProvider  `json:"flannelNetworkProvider,omitempty" yaml:"flannelNetworkProvider,omitempty"`
	MTU                    int64                    `json:"mtu,omitempty" yaml:"mtu,omitempty"`
	NodeSelector           map[string]string        `json:"nodeSelector,omitempty" yaml:"nodeSelector,omitempty"`
	Options                map[string]string        `json:"options,omitempty" yaml:"options,omitempty"`
	Plugin                 string                   `json:"plugin,omitempty" yaml:"plugin,omitempty"`
	Tolerations            []Toleration             `json:"tolerations,omitempty" yaml:"tolerations,omitempty"`
	UpdateStrategy         *DaemonSetUpdateStrategy `json:"updateStrategy,omitempty" yaml:"updateStrategy,omitempty"`
	WeaveNetworkProvider   *WeaveNetworkProvider    `json:"weaveNetworkProvider,omitempty" yaml:"weaveNetworkProvider,omitempty"`
}

type RKETaint struct {
	Effect    string `json:"effect,omitempty" yaml:"effect,omitempty"`
	Key       string `json:"key,omitempty" yaml:"key,omitempty"`
	TimeAdded string `json:"timeAdded,omitempty" yaml:"timeAdded,omitempty"`
	Value     string `json:"value,omitempty" yaml:"value,omitempty"`
}

type RKEConfigNode struct {
	Address          string            `json:"address,omitempty" yaml:"address,omitempty"`
	DockerSocket     string            `json:"dockerSocket,omitempty" yaml:"dockerSocket,omitempty"`
	HostnameOverride string            `json:"hostnameOverride,omitempty" yaml:"hostnameOverride,omitempty"`
	InternalAddress  string            `json:"internalAddress,omitempty" yaml:"internalAddress,omitempty"`
	Labels           map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	NodeID           string            `json:"nodeId,omitempty" yaml:"nodeId,omitempty"`
	Port             string            `json:"port,omitempty" yaml:"port,omitempty"`
	Role             []string          `json:"role,omitempty" yaml:"role,omitempty"`
	SSHAgentAuth     bool              `json:"sshAgentAuth,omitempty" yaml:"sshAgentAuth,omitempty"`
	SSHCert          string            `json:"sshCert,omitempty" yaml:"sshCert,omitempty"`
	SSHCertPath      string            `json:"sshCertPath,omitempty" yaml:"sshCertPath,omitempty"`
	SSHKey           string            `json:"sshKey,omitempty" yaml:"sshKey,omitempty"`
	SSHKeyPath       string            `json:"sshKeyPath,omitempty" yaml:"sshKeyPath,omitempty"`
	Taints           []RKETaint        `json:"taints,omitempty" yaml:"taints,omitempty"`
	User             string            `json:"user,omitempty" yaml:"user,omitempty"`
}

type PrivateRegistry struct {
	IsDefault bool   `json:"isDefault,omitempty" yaml:"isDefault,omitempty"`
	Password  string `json:"password,omitempty" yaml:"password,omitempty"`
	URL       string `json:"url,omitempty" yaml:"url,omitempty"`
	User      string `json:"user,omitempty" yaml:"user,omitempty"`
}

type RestoreConfig struct {
	Restore      bool   `json:"restore,omitempty" yaml:"restore,omitempty"`
	SnapshotName string `json:"snapshotName,omitempty" yaml:"snapshotName,omitempty"`
}

type RotateCertificates struct {
	CACertificates bool   `json:"caCertificates,omitempty" yaml:"caCertificates,omitempty"`
	Services       string `json:"services,omitempty" yaml:"services,omitempty"`
}

type S3BackupConfig struct {
	AccessKey  string `json:"accessKey,omitempty" yaml:"accessKey,omitempty"`
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`
	CustomCA   string `json:"customCa,omitempty" yaml:"customCa,omitempty"`
	Endpoint   string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	Folder     string `json:"folder,omitempty" yaml:"folder,omitempty"`
	Region     string `json:"region,omitempty" yaml:"region,omitempty"`
	SecretKey  string `json:"secretKey,omitempty" yaml:"secretKey,omitempty"`
}

type BackupConfig struct {
	Enabled        *bool           `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	IntervalHours  int64           `json:"intervalHours,omitempty" yaml:"intervalHours,omitempty"`
	Retention      int64           `json:"retention,omitempty" yaml:"retention,omitempty"`
	S3BackupConfig *S3BackupConfig `json:"s3BackupConfig,omitempty" yaml:"s3BackupConfig,omitempty"`
	SafeTimestamp  bool            `json:"safeTimestamp,omitempty" yaml:"safeTimestamp,omitempty"`
	Timeout        int64           `json:"timeout,omitempty" yaml:"timeout,omitempty"`
}

type ETCDService struct {
	BackupConfig      *BackupConfig     `json:"backupConfig,omitempty" yaml:"backupConfig,omitempty"`
	CACert            string            `json:"caCert,omitempty" yaml:"caCert,omitempty"`
	Cert              string            `json:"cert,omitempty" yaml:"cert,omitempty"`
	Creation          string            `json:"creation,omitempty" yaml:"creation,omitempty"`
	ExternalURLs      []string          `json:"externalUrls,omitempty" yaml:"externalUrls,omitempty"`
	ExtraArgs         map[string]string `json:"extraArgs,omitempty" yaml:"extraArgs,omitempty"`
	ExtraBinds        []string          `json:"extraBinds,omitempty" yaml:"extraBinds,omitempty"`
	ExtraEnv          []string          `json:"extraEnv,omitempty" yaml:"extraEnv,omitempty"`
	GID               int64             `json:"gid,omitempty" yaml:"gid,omitempty"`
	Image             string            `json:"image,omitempty" yaml:"image,omitempty"`
	Key               string            `json:"key,omitempty" yaml:"key,omitempty"`
	Path              string            `json:"path,omitempty" yaml:"path,omitempty"`
	Retention         string            `json:"retention,omitempty" yaml:"retention,omitempty"`
	Snapshot          *bool             `json:"snapshot,omitempty" yaml:"snapshot,omitempty"`
	UID               int64             `json:"uid,omitempty" yaml:"uid,omitempty"`
	WindowsExtraArgs  map[string]string `json:"winExtraArgs,omitempty" yaml:"winExtraArgs,omitempty"`
	WindowsExtraBinds []string          `json:"winExtraBinds,omitempty" yaml:"winExtraBinds,omitempty"`
	WindowsExtraEnv   []string          `json:"winExtraEnv,omitempty" yaml:"winExtraEnv,omitempty"`
}

type AuditLogConfig struct {
	Format    string                 `json:"format,omitempty" yaml:"format,omitempty"`
	MaxAge    int64                  `json:"maxAge,omitempty" yaml:"maxAge,omitempty"`
	MaxBackup int64                  `json:"maxBackup,omitempty" yaml:"maxBackup,omitempty"`
	MaxSize   int64                  `json:"maxSize,omitempty" yaml:"maxSize,omitempty"`
	Path      string                 `json:"path,omitempty" yaml:"path,omitempty"`
	Policy    map[string]interface{} `json:"policy,omitempty" yaml:"policy,omitempty"`
}

type AuditLog struct {
	Configuration *AuditLogConfig `json:"configuration,omitempty" yaml:"configuration,omitempty"`
	Enabled       bool            `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}

type EventRateLimit struct {
	Configuration map[string]interface{} `json:"configuration,omitempty" yaml:"configuration,omitempty"`
	Enabled       bool                   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}

type SecretsEncryptionConfig struct {
	CustomConfig map[string]interface{} `json:"customConfig,omitempty" yaml:"customConfig,omitempty"`
	Enabled      bool                   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}

type KubeAPIService struct {
	AdmissionConfiguration  map[string]interface{}   `json:"admissionConfiguration,omitempty" yaml:"admissionConfiguration,omitempty"`
	AlwaysPullImages        bool                     `json:"alwaysPullImages,omitempty" yaml:"alwaysPullImages,omitempty"`
	AuditLog                *AuditLog                `json:"auditLog,omitempty" yaml:"auditLog,omitempty"`
	EventRateLimit          *EventRateLimit          `json:"eventRateLimit,omitempty" yaml:"eventRateLimit,omitempty"`
	ExtraArgs               map[string]string        `json:"extraArgs,omitempty" yaml:"extraArgs,omitempty"`
	ExtraBinds              []string                 `json:"extraBinds,omitempty" yaml:"extraBinds,omitempty"`
	ExtraEnv                []string                 `json:"extraEnv,omitempty" yaml:"extraEnv,omitempty"`
	Image                   string                   `json:"image,omitempty" yaml:"image,omitempty"`
	PodSecurityPolicy       bool                     `json:"podSecurityPolicy,omitempty" yaml:"podSecurityPolicy,omitempty"`
	SecretsEncryptionConfig *SecretsEncryptionConfig `json:"secretsEncryptionConfig,omitempty" yaml:"secretsEncryptionConfig,omitempty"`
	ServiceClusterIPRange   string                   `json:"serviceClusterIpRange,omitempty" yaml:"serviceClusterIpRange,omitempty"`
	ServiceNodePortRange    string                   `json:"serviceNodePortRange,omitempty" yaml:"serviceNodePortRange,omitempty"`
	WindowsExtraArgs        map[string]string        `json:"winExtraArgs,omitempty" yaml:"winExtraArgs,omitempty"`
	WindowsExtraBinds       []string                 `json:"winExtraBinds,omitempty" yaml:"winExtraBinds,omitempty"`
	WindowsExtraEnv         []string                 `json:"winExtraEnv,omitempty" yaml:"winExtraEnv,omitempty"`
}

type KubeControllerService struct {
	ClusterCIDR           string            `json:"clusterCidr,omitempty" yaml:"clusterCidr,omitempty"`
	ExtraArgs             map[string]string `json:"extraArgs,omitempty" yaml:"extraArgs,omitempty"`
	ExtraBinds            []string          `json:"extraBinds,omitempty" yaml:"extraBinds,omitempty"`
	ExtraEnv              []string          `json:"extraEnv,omitempty" yaml:"extraEnv,omitempty"`
	Image                 string            `json:"image,omitempty" yaml:"image,omitempty"`
	ServiceClusterIPRange string            `json:"serviceClusterIpRange,omitempty" yaml:"serviceClusterIpRange,omitempty"`
	WindowsExtraArgs      map[string]string `json:"winExtraArgs,omitempty" yaml:"winExtraArgs,omitempty"`
	WindowsExtraBinds     []string          `json:"winExtraBinds,omitempty" yaml:"winExtraBinds,omitempty"`
	WindowsExtraEnv       []string          `json:"winExtraEnv,omitempty" yaml:"winExtraEnv,omitempty"`
}

type KubeletService struct {
	ClusterDNSServer           string            `json:"clusterDnsServer,omitempty" yaml:"clusterDnsServer,omitempty"`
	ClusterDomain              string            `json:"clusterDomain,omitempty" yaml:"clusterDomain,omitempty"`
	ExtraArgs                  map[string]string `json:"extraArgs,omitempty" yaml:"extraArgs,omitempty"`
	ExtraBinds                 []string          `json:"extraBinds,omitempty" yaml:"extraBinds,omitempty"`
	ExtraEnv                   []string          `json:"extraEnv,omitempty" yaml:"extraEnv,omitempty"`
	FailSwapOn                 bool              `json:"failSwapOn,omitempty" yaml:"failSwapOn,omitempty"`
	GenerateServingCertificate bool              `json:"generateServingCertificate,omitempty" yaml:"generateServingCertificate,omitempty"`
	Image                      string            `json:"image,omitempty" yaml:"image,omitempty"`
	InfraContainerImage        string            `json:"infraContainerImage,omitempty" yaml:"infraContainerImage,omitempty"`
	WindowsExtraArgs           map[string]string `json:"winExtraArgs,omitempty" yaml:"winExtraArgs,omitempty"`
	WindowsExtraBinds          []string          `json:"winExtraBinds,omitempty" yaml:"winExtraBinds,omitempty"`
	WindowsExtraEnv            []string          `json:"winExtraEnv,omitempty" yaml:"winExtraEnv,omitempty"`
}

type KubeproxyService struct {
	ExtraArgs         map[string]string `json:"extraArgs,omitempty" yaml:"extraArgs,omitempty"`
	ExtraBinds        []string          `json:"extraBinds,omitempty" yaml:"extraBinds,omitempty"`
	ExtraEnv          []string          `json:"extraEnv,omitempty" yaml:"extraEnv,omitempty"`
	Image             string            `json:"image,omitempty" yaml:"image,omitempty"`
	WindowsExtraArgs  map[string]string `json:"winExtraArgs,omitempty" yaml:"winExtraArgs,omitempty"`
	WindowsExtraBinds []string          `json:"winExtraBinds,omitempty" yaml:"winExtraBinds,omitempty"`
	WindowsExtraEnv   []string          `json:"winExtraEnv,omitempty" yaml:"winExtraEnv,omitempty"`
}

type SchedulerService struct {
	ExtraArgs         map[string]string `json:"extraArgs,omitempty" yaml:"extraArgs,omitempty"`
	ExtraBinds        []string          `json:"extraBinds,omitempty" yaml:"extraBinds,omitempty"`
	ExtraEnv          []string          `json:"extraEnv,omitempty" yaml:"extraEnv,omitempty"`
	Image             string            `json:"image,omitempty" yaml:"image,omitempty"`
	WindowsExtraArgs  map[string]string `json:"winExtraArgs,omitempty" yaml:"winExtraArgs,omitempty"`
	WindowsExtraBinds []string          `json:"winExtraBinds,omitempty" yaml:"winExtraBinds,omitempty"`
	WindowsExtraEnv   []string          `json:"winExtraEnv,omitempty" yaml:"winExtraEnv,omitempty"`
}

type RKEConfigServices struct {
	Etcd           *ETCDService           `json:"etcd,omitempty" yaml:"etcd,omitempty"`
	KubeAPI        *KubeAPIService        `json:"kubeApi,omitempty" yaml:"kubeApi,omitempty"`
	KubeController *KubeControllerService `json:"kubeController,omitempty" yaml:"kubeController,omitempty"`
	Kubelet        *KubeletService        `json:"kubelet,omitempty" yaml:"kubelet,omitempty"`
	Kubeproxy      *KubeproxyService      `json:"kubeproxy,omitempty" yaml:"kubeproxy,omitempty"`
	Scheduler      *SchedulerService      `json:"scheduler,omitempty" yaml:"scheduler,omitempty"`
}

type NodeDrainInput struct {
	DeleteLocalData  bool  `json:"deleteLocalData,omitempty" yaml:"deleteLocalData,omitempty"`
	Force            bool  `json:"force,omitempty" yaml:"force,omitempty"`
	GracePeriod      int64 `json:"gracePeriod,omitempty" yaml:"gracePeriod,omitempty"`
	IgnoreDaemonSets *bool `json:"ignoreDaemonSets,omitempty" yaml:"ignoreDaemonSets,omitempty"`
	Timeout          int64 `json:"timeout,omitempty" yaml:"timeout,omitempty"`
}

type NodeUpgradeStrategy struct {
	Drain                      *bool           `json:"drain,omitempty" yaml:"drain,omitempty"`
	DrainInput                 *NodeDrainInput `json:"nodeDrainInput,omitempty" yaml:"nodeDrainInput,omitempty"`
	MaxUnavailableControlplane string          `json:"maxUnavailableControlplane,omitempty" yaml:"maxUnavailableControlplane,omitempty"`
	MaxUnavailableWorker       string          `json:"maxUnavailableWorker,omitempty" yaml:"maxUnavailableWorker,omitempty"`
}

type RancherKubernetesEngineConfig struct {
	AddonJobTimeout     int64                `json:"addonJobTimeout,omitempty" yaml:"addonJobTimeout,omitempty"`
	Addons              string               `json:"addons,omitempty" yaml:"addons,omitempty"`
	AddonsInclude       []string             `json:"addonsInclude,omitempty" yaml:"addonsInclude,omitempty"`
	Authentication      *AuthnConfig         `json:"authentication,omitempty" yaml:"authentication,omitempty"`
	Authorization       *AuthzConfig         `json:"authorization,omitempty" yaml:"authorization,omitempty"`
	BastionHost         *BastionHost         `json:"bastionHost,omitempty" yaml:"bastionHost,omitempty"`
	CloudProvider       *CloudProvider       `json:"cloudProvider,omitempty" yaml:"cloudProvider,omitempty"`
	ClusterName         string               `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`
	DNS                 *DNSConfig           `json:"dns,omitempty" yaml:"dns,omitempty"`
	IgnoreDockerVersion *bool                `json:"ignoreDockerVersion,omitempty" yaml:"ignoreDockerVersion,omitempty"`
	Ingress             *IngressConfig       `json:"ingress,omitempty" yaml:"ingress,omitempty"`
	Monitoring          *MonitoringConfig    `json:"monitoring,omitempty" yaml:"monitoring,omitempty"`
	Network             *NetworkConfig       `json:"network,omitempty" yaml:"network,omitempty"`
	Nodes               []RKEConfigNode      `json:"nodes,omitempty" yaml:"nodes,omitempty"`
	PrefixPath          string               `json:"prefixPath,omitempty" yaml:"prefixPath,omitempty"`
	PrivateRegistries   []PrivateRegistry    `json:"privateRegistries,omitempty" yaml:"privateRegistries,omitempty"`
	Restore             *RestoreConfig       `json:"restore,omitempty" yaml:"restore,omitempty"`
	RotateCertificates  *RotateCertificates  `json:"rotateCertificates,omitempty" yaml:"rotateCertificates,omitempty"`
	RotateEncryptionKey bool                 `json:"rotateEncryptionKey,omitempty" yaml:"rotateEncryptionKey,omitempty"`
	SSHAgentAuth        bool                 `json:"sshAgentAuth,omitempty" yaml:"sshAgentAuth,omitempty"`
	SSHCertPath         string               `json:"sshCertPath,omitempty" yaml:"sshCertPath,omitempty"`
	SSHKeyPath          string               `json:"sshKeyPath,omitempty" yaml:"sshKeyPath,omitempty"`
	Services            *RKEConfigServices   `json:"services,omitempty" yaml:"services,omitempty"`
	UpgradeStrategy     *NodeUpgradeStrategy `json:"upgradeStrategy,omitempty" yaml:"upgradeStrategy,omitempty"`
	Version             string               `json:"kubernetesVersion,omitempty" yaml:"kubernetesVersion,omitempty"`
	WindowsPrefixPath   string               `json:"winPrefixPath,omitempty" yaml:"winPrefixPath,omitempty"`
}

type Rke2Config struct {
	ClusterUpgradeStrategy *ClusterUpgradeStrategy `json:"rke2upgradeStrategy,omitempty" yaml:"rke2upgradeStrategy,omitempty"`
	Version                string                  `json:"kubernetesVersion,omitempty" yaml:"kubernetesVersion,omitempty"`
}

type CisScanConfig struct {
	DebugMaster              bool     `json:"debugMaster,omitempty" yaml:"debugMaster,omitempty"`
	DebugWorker              bool     `json:"debugWorker,omitempty" yaml:"debugWorker,omitempty"`
	OverrideBenchmarkVersion string   `json:"overrideBenchmarkVersion,omitempty" yaml:"overrideBenchmarkVersion,omitempty"`
	OverrideSkip             []string `json:"overrideSkip,omitempty" yaml:"overrideSkip,omitempty"`
	Profile                  string   `json:"profile,omitempty" yaml:"profile,omitempty"`
}

type ClusterScanConfig struct {
	CisScanConfig *CisScanConfig `json:"cisScanConfig,omitempty" yaml:"cisScanConfig,omitempty"`
}

type ScheduledClusterScanConfig struct {
	CronSchedule string `json:"cronSchedule,omitempty" yaml:"cronSchedule,omitempty"`
	Retention    int64  `json:"retention,omitempty" yaml:"retention,omitempty"`
}

type ScheduledClusterScan struct {
	Enabled        bool                        `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	ScanConfig     *ClusterScanConfig          `json:"scanConfig,omitempty" yaml:"scanConfig,omitempty"`
	ScheduleConfig *ScheduledClusterScanConfig `json:"scheduleConfig,omitempty" yaml:"scheduleConfig,omitempty"`
}

type ClusterSpec struct {
	AgentEnvVars                        []EnvVar                       `json:"agentEnvVars,omitempty" yaml:"agentEnvVars,omitempty"`
	AgentImageOverride                  string                         `json:"agentImageOverride,omitempty" yaml:"agentImageOverride,omitempty"`
	AmazonElasticContainerServiceConfig map[string]interface{}         `json:"amazonElasticContainerServiceConfig,omitempty" yaml:"amazonElasticContainerServiceConfig,omitempty"`
	AzureKubernetesServiceConfig        map[string]interface{}         `json:"azureKubernetesServiceConfig,omitempty" yaml:"azureKubernetesServiceConfig,omitempty"`
	ClusterTemplateAnswers              *Answer                        `json:"answers,omitempty" yaml:"answers,omitempty"`
	ClusterTemplateID                   string                         `json:"clusterTemplateId,omitempty" yaml:"clusterTemplateId,omitempty"`
	ClusterTemplateQuestions            []Question                     `json:"questions,omitempty" yaml:"questions,omitempty"`
	ClusterTemplateRevisionID           string                         `json:"clusterTemplateRevisionId,omitempty" yaml:"clusterTemplateRevisionId,omitempty"`
	DefaultClusterRoleForProjectMembers string                         `json:"defaultClusterRoleForProjectMembers,omitempty" yaml:"defaultClusterRoleForProjectMembers,omitempty"`
	DefaultPodSecurityPolicyTemplateID  string                         `json:"defaultPodSecurityPolicyTemplateId,omitempty" yaml:"defaultPodSecurityPolicyTemplateId,omitempty"`
	Description                         string                         `json:"description,omitempty" yaml:"description,omitempty"`
	DesiredAgentImage                   string                         `json:"desiredAgentImage,omitempty" yaml:"desiredAgentImage,omitempty"`
	DesiredAuthImage                    string                         `json:"desiredAuthImage,omitempty" yaml:"desiredAuthImage,omitempty"`
	DisplayName                         string                         `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	DockerRootDir                       string                         `json:"dockerRootDir,omitempty" yaml:"dockerRootDir,omitempty"`
	EKSConfig                           *EKSClusterConfigSpec          `json:"eksConfig,omitempty" yaml:"eksConfig,omitempty"`
	EnableClusterAlerting               bool                           `json:"enableClusterAlerting,omitempty" yaml:"enableClusterAlerting,omitempty"`
	EnableClusterMonitoring             bool                           `json:"enableClusterMonitoring,omitempty" yaml:"enableClusterMonitoring,omitempty"`
	EnableNetworkPolicy                 *bool                          `json:"enableNetworkPolicy,omitempty" yaml:"enableNetworkPolicy,omitempty"`
	FleetWorkspaceName                  string                         `json:"fleetWorkspaceName,omitempty" yaml:"fleetWorkspaceName,omitempty"`
	GKEConfig                           *GKEClusterConfigSpec          `json:"gkeConfig,omitempty" yaml:"gkeConfig,omitempty"`
	GenericEngineConfig                 map[string]interface{}         `json:"genericEngineConfig,omitempty" yaml:"genericEngineConfig,omitempty"`
	GoogleKubernetesEngineConfig        map[string]interface{}         `json:"googleKubernetesEngineConfig,omitempty" yaml:"googleKubernetesEngineConfig,omitempty"`
	ImportedConfig                      *ImportedConfig                `json:"importedConfig,omitempty" yaml:"importedConfig,omitempty"`
	Internal                            bool                           `json:"internal,omitempty" yaml:"internal,omitempty"`
	K3sConfig                           *K3sConfig                     `json:"k3sConfig,omitempty" yaml:"k3sConfig,omitempty"`
	LocalClusterAuthEndpoint            *LocalClusterAuthEndpoint      `json:"localClusterAuthEndpoint,omitempty" yaml:"localClusterAuthEndpoint,omitempty"`
	RancherKubernetesEngineConfig       *RancherKubernetesEngineConfig `json:"rancherKubernetesEngineConfig,omitempty" yaml:"rancherKubernetesEngineConfig,omitempty"`
	Rke2Config                          *Rke2Config                    `json:"rke2Config,omitempty" yaml:"rke2Config,omitempty"`
	ScheduledClusterScan                *ScheduledClusterScan          `json:"scheduledClusterScan,omitempty" yaml:"scheduledClusterScan,omitempty"`
	WindowsPreferedCluster              bool                           `json:"windowsPreferedCluster,omitempty" yaml:"windowsPreferedCluster,omitempty"`
}

type IngressCapabilities struct {
	CustomDefaultBackend *bool  `json:"customDefaultBackend,omitempty" yaml:"customDefaultBackend,omitempty"`
	IngressProvider      string `json:"ingressProvider,omitempty" yaml:"ingressProvider,omitempty"`
}

type LoadBalancerCapabilities struct {
	Enabled              *bool    `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	HealthCheckSupported bool     `json:"healthCheckSupported,omitempty" yaml:"healthCheckSupported,omitempty"`
	ProtocolsSupported   []string `json:"protocolsSupported,omitempty" yaml:"protocolsSupported,omitempty"`
	Provider             string   `json:"provider,omitempty" yaml:"provider,omitempty"`
}

type Capabilities struct {
	IngressCapabilities      []IngressCapabilities     `json:"ingressCapabilities,omitempty" yaml:"ingressCapabilities,omitempty"`
	LoadBalancerCapabilities *LoadBalancerCapabilities `json:"loadBalancerCapabilities,omitempty" yaml:"loadBalancerCapabilities,omitempty"`
	NodePoolScalingSupported bool                      `json:"nodePoolScalingSupported,omitempty" yaml:"nodePoolScalingSupported,omitempty"`
	NodePortRange            string                    `json:"nodePortRange,omitempty" yaml:"nodePortRange,omitempty"`
	PspEnabled               bool                      `json:"pspEnabled,omitempty" yaml:"pspEnabled,omitempty"`
	TaintSupport             *bool                     `json:"taintSupport,omitempty" yaml:"taintSupport,omitempty"`
}

type CertExpiration struct {
	ExpirationDate string `json:"expirationDate,omitempty" yaml:"expirationDate,omitempty"`
}

type ComponentCondition struct {
	Error   string `json:"error,omitempty" yaml:"error,omitempty"`
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
	Status  string `json:"status,omitempty" yaml:"status,omitempty"`
	Type    string `json:"type,omitempty" yaml:"type,omitempty"`
}

type ClusterCondition struct {
	LastTransitionTime string `json:"lastTransitionTime,omitempty" yaml:"lastTransitionTime,omitempty"`
	LastUpdateTime     string `json:"lastUpdateTime,omitempty" yaml:"lastUpdateTime,omitempty"`
	Message            string `json:"message,omitempty" yaml:"message,omitempty"`
	Reason             string `json:"reason,omitempty" yaml:"reason,omitempty"`
	Status             string `json:"status,omitempty" yaml:"status,omitempty"`
	Type               string `json:"type,omitempty" yaml:"type,omitempty"`
}

type EKSStatus struct {
	ManagedLaunchTemplateID       string                `json:"managedLaunchTemplateID,omitempty" yaml:"managedLaunchTemplateID,omitempty"`
	ManagedLaunchTemplateVersions map[string]string     `json:"managedLaunchTemplateVersions,omitempty" yaml:"managedLaunchTemplateVersions,omitempty"`
	PrivateRequiresTunnel         *bool                 `json:"privateRequiresTunnel,omitempty" yaml:"privateRequiresTunnel,omitempty"`
	SecurityGroups                []string              `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`
	Subnets                       []string              `json:"subnets,omitempty" yaml:"subnets,omitempty"`
	UpstreamSpec                  *EKSClusterConfigSpec `json:"upstreamSpec,omitempty" yaml:"upstreamSpec,omitempty"`
	VirtualNetwork                string                `json:"virtualNetwork,omitempty" yaml:"virtualNetwork,omitempty"`
}

type GKEStatus struct {
	PrivateRequiresTunnel *bool                 `json:"privateRequiresTunnel,omitempty" yaml:"privateRequiresTunnel,omitempty"`
	UpstreamSpec          *GKEClusterConfigSpec `json:"upstreamSpec,omitempty" yaml:"upstreamSpec,omitempty"`
}

type MonitoringCondition struct {
	LastTransitionTime string `json:"lastTransitionTime,omitempty" yaml:"lastTransitionTime,omitempty"`
	LastUpdateTime     string `json:"lastUpdateTime,omitempty" yaml:"lastUpdateTime,omitempty"`
	Message            string `json:"message,omitempty" yaml:"message,omitempty"`
	Reason             string `json:"reason,omitempty" yaml:"reason,omitempty"`
	Status             string `json:"status,omitempty" yaml:"status,omitempty"`
	Type               string `json:"type,omitempty" yaml:"type,omitempty"`
}

type MonitoringStatus struct {
	Conditions      []MonitoringCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	GrafanaEndpoint string                `json:"grafanaEndpoint,omitempty" yaml:"grafanaEndpoint,omitempty"`
}

type OwnerReference struct {
	APIVersion         string `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`
	BlockOwnerDeletion *bool  `json:"blockOwnerDeletion,omitempty" yaml:"blockOwnerDeletion,omitempty"`
	Controller         *bool  `json:"controller,omitempty" yaml:"controller,omitempty"`
	Kind               string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name               string `json:"name,omitempty" yaml:"name,omitempty"`
	UID                string `json:"uid,omitempty" yaml:"uid,omitempty"`
}

type ScheduledClusterScanStatus struct {
	Enabled          bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	LastRunTimestamp string `json:"lastRunTimestamp,omitempty" yaml:"lastRunTimestamp,omitempty"`
}

type Info struct {
	BuildDate    string `json:"buildDate,omitempty" yaml:"buildDate,omitempty"`
	Compiler     string `json:"compiler,omitempty" yaml:"compiler,omitempty"`
	GitCommit    string `json:"gitCommit,omitempty" yaml:"gitCommit,omitempty"`
	GitTreeState string `json:"gitTreeState,omitempty" yaml:"gitTreeState,omitempty"`
	GitVersion   string `json:"gitVersion,omitempty" yaml:"gitVersion,omitempty"`
	GoVersion    string `json:"goVersion,omitempty" yaml:"goVersion,omitempty"`
	Major        string `json:"major,omitempty" yaml:"major,omitempty"`
	Minor        string `json:"minor,omitempty" yaml:"minor,omitempty"`
	Platform     string `json:"platform,omitempty" yaml:"platform,omitempty"`
}

type Cluster struct {
	Resource
	APIEndpoint                          string                         `json:"apiEndpoint,omitempty" yaml:"apiEndpoint,omitempty"`
	AgentEnvVars                         []EnvVar                       `json:"agentEnvVars,omitempty" yaml:"agentEnvVars,omitempty"`
	AgentFeatures                        map[string]bool                `json:"agentFeatures,omitempty" yaml:"agentFeatures,omitempty"`
	AgentImage                           string                         `json:"agentImage,omitempty" yaml:"agentImage,omitempty"`
	AgentImageOverride                   string                         `json:"agentImageOverride,omitempty" yaml:"agentImageOverride,omitempty"`
	Allocatable                          map[string]string              `json:"allocatable,omitempty" yaml:"allocatable,omitempty"`
	Annotations                          map[string]string              `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	AppliedAgentEnvVars                  []EnvVar                       `json:"appliedAgentEnvVars,omitempty" yaml:"appliedAgentEnvVars,omitempty"`
	AppliedEnableNetworkPolicy           bool                           `json:"appliedEnableNetworkPolicy,omitempty" yaml:"appliedEnableNetworkPolicy,omitempty"`
	AppliedPodSecurityPolicyTemplateName string                         `json:"appliedPodSecurityPolicyTemplateId,omitempty" yaml:"appliedPodSecurityPolicyTemplateId,omitempty"`
	AppliedSpec                          *ClusterSpec                   `json:"appliedSpec,omitempty" yaml:"appliedSpec,omitempty"`
	AuthImage                            string                         `json:"authImage,omitempty" yaml:"authImage,omitempty"`
	CACert                               string                         `json:"caCert,omitempty" yaml:"caCert,omitempty"`
	Capabilities                         *Capabilities                  `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`
	Capacity                             map[string]string              `json:"capacity,omitempty" yaml:"capacity,omitempty"`
	CertificatesExpiration               map[string]CertExpiration      `json:"certificatesExpiration,omitempty" yaml:"certificatesExpiration,omitempty"`
	ClusterTemplateAnswers               *Answer                        `json:"answers,omitempty" yaml:"answers,omitempty"`
	ClusterTemplateID                    string                         `json:"clusterTemplateId,omitempty" yaml:"clusterTemplateId,omitempty"`
	ClusterTemplateQuestions             []Question                     `json:"questions,omitempty" yaml:"questions,omitempty"`
	ClusterTemplateRevisionID            string                         `json:"clusterTemplateRevisionId,omitempty" yaml:"clusterTemplateRevisionId,omitempty"`
	ComponentStatuses                    []ClusterComponentStatus       `json:"componentStatuses,omitempty" yaml:"componentStatuses,omitempty"`
	Conditions                           []ClusterCondition             `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	Created                              string                         `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID                            string                         `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	CurrentCisRunName                    string                         `json:"currentCisRunName,omitempty" yaml:"currentCisRunName,omitempty"`
	DefaultClusterRoleForProjectMembers  string                         `json:"defaultClusterRoleForProjectMembers,omitempty" yaml:"defaultClusterRoleForProjectMembers,omitempty"`
	DefaultPodSecurityPolicyTemplateID   string                         `json:"defaultPodSecurityPolicyTemplateId,omitempty" yaml:"defaultPodSecurityPolicyTemplateId,omitempty"`
	Description                          string                         `json:"description,omitempty" yaml:"description,omitempty"`
	DesiredAgentImage                    string                         `json:"desiredAgentImage,omitempty" yaml:"desiredAgentImage,omitempty"`
	DesiredAuthImage                     string                         `json:"desiredAuthImage,omitempty" yaml:"desiredAuthImage,omitempty"`
	DockerRootDir                        string                         `json:"dockerRootDir,omitempty" yaml:"dockerRootDir,omitempty"`
	Driver                               string                         `json:"driver,omitempty" yaml:"driver,omitempty"`
	EKSConfig                            *EKSClusterConfigSpec          `json:"eksConfig,omitempty" yaml:"eksConfig,omitempty"`
	EKSStatus                            *EKSStatus                     `json:"eksStatus,omitempty" yaml:"eksStatus,omitempty"`
	EnableClusterAlerting                bool                           `json:"enableClusterAlerting,omitempty" yaml:"enableClusterAlerting,omitempty"`
	EnableClusterMonitoring              bool                           `json:"enableClusterMonitoring,omitempty" yaml:"enableClusterMonitoring,omitempty"`
	EnableNetworkPolicy                  *bool                          `json:"enableNetworkPolicy,omitempty" yaml:"enableNetworkPolicy,omitempty"`
	FailedSpec                           *ClusterSpec                   `json:"failedSpec,omitempty" yaml:"failedSpec,omitempty"`
	FleetWorkspaceName                   string                         `json:"fleetWorkspaceName,omitempty" yaml:"fleetWorkspaceName,omitempty"`
	GKEConfig                            *GKEClusterConfigSpec          `json:"gkeConfig,omitempty" yaml:"gkeConfig,omitempty"`
	GKEStatus                            *GKEStatus                     `json:"gkeStatus,omitempty" yaml:"gkeStatus,omitempty"`
	ImportedConfig                       *ImportedConfig                `json:"importedConfig,omitempty" yaml:"importedConfig,omitempty"`
	Internal                             bool                           `json:"internal,omitempty" yaml:"internal,omitempty"`
	IstioEnabled                         bool                           `json:"istioEnabled,omitempty" yaml:"istioEnabled,omitempty"`
	K3sConfig                            *K3sConfig                     `json:"k3sConfig,omitempty" yaml:"k3sConfig,omitempty"`
	Labels                               map[string]string              `json:"labels,omitempty" yaml:"labels,omitempty"`
	Limits                               map[string]string              `json:"limits,omitempty" yaml:"limits,omitempty"`
	LocalClusterAuthEndpoint             *LocalClusterAuthEndpoint      `json:"localClusterAuthEndpoint,omitempty" yaml:"localClusterAuthEndpoint,omitempty"`
	MonitoringStatus                     *MonitoringStatus              `json:"monitoringStatus,omitempty" yaml:"monitoringStatus,omitempty"`
	Name                                 string                         `json:"name,omitempty" yaml:"name,omitempty"`
	NodeCount                            int64                          `json:"nodeCount,omitempty" yaml:"nodeCount,omitempty"`
	NodeVersion                          int64                          `json:"nodeVersion,omitempty" yaml:"nodeVersion,omitempty"`
	OwnerReferences                      []OwnerReference               `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Provider                             string                         `json:"provider,omitempty" yaml:"provider,omitempty"`
	RancherKubernetesEngineConfig        *RancherKubernetesEngineConfig `json:"rancherKubernetesEngineConfig,omitempty" yaml:"rancherKubernetesEngineConfig,omitempty"`
	Removed                              string                         `json:"removed,omitempty" yaml:"removed,omitempty"`
	Requested                            map[string]string              `json:"requested,omitempty" yaml:"requested,omitempty"`
	Rke2Config                           *Rke2Config                    `json:"rke2Config,omitempty" yaml:"rke2Config,omitempty"`
	ScheduledClusterScan                 *ScheduledClusterScan          `json:"scheduledClusterScan,omitempty" yaml:"scheduledClusterScan,omitempty"`
	ScheduledClusterScanStatus           *ScheduledClusterScanStatus    `json:"scheduledClusterScanStatus,omitempty" yaml:"scheduledClusterScanStatus,omitempty"`
	State                                string                         `json:"state,omitempty" yaml:"state,omitempty"`
	Transitioning                        string                         `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage                 string                         `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
	UUID                                 string                         `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Version                              *Info                          `json:"version,omitempty" yaml:"version,omitempty"`
	WindowsPreferedCluster               bool                           `json:"windowsPreferedCluster,omitempty" yaml:"windowsPreferedCluster,omitempty"`
}

type Pagination struct {
	Marker   string `json:"marker,omitempty"`
	First    string `json:"first,omitempty"`
	Previous string `json:"previous,omitempty"`
	Next     string `json:"next,omitempty"`
	Last     string `json:"last,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
	Total    *int64 `json:"total,omitempty"`
	Partial  bool   `json:"partial,omitempty"`
}

type SortOrder string

var (
	ASC  = SortOrder("asc")
	DESC = SortOrder("desc")
)

type Sort struct {
	Name    string            `json:"name,omitempty"`
	Order   SortOrder         `json:"order,omitempty"`
	Reverse string            `json:"reverse,omitempty"`
	Links   map[string]string `json:"links,omitempty"`
}

var (
	ModifierEQ      ModifierType = "eq"
	ModifierNE      ModifierType = "ne"
	ModifierNull    ModifierType = "null"
	ModifierNotNull ModifierType = "notnull"
	ModifierIn      ModifierType = "in"
	ModifierNotIn   ModifierType = "notin"
)

type ModifierType string

type Condition struct {
	Modifier ModifierType `json:"modifier,omitempty"`
	Value    interface{}  `json:"value,omitempty"`
}

type Collection struct {
	Type         string                 `json:"type,omitempty"`
	Links        map[string]string      `json:"links"`
	CreateTypes  map[string]string      `json:"createTypes,omitempty"`
	Actions      map[string]string      `json:"actions"`
	Pagination   *Pagination            `json:"pagination,omitempty"`
	Sort         *Sort                  `json:"sort,omitempty"`
	Filters      map[string][]Condition `json:"filters,omitempty"`
	ResourceType string                 `json:"resourceType"`
}

type ClusterCollection struct {
	Collection
	Data []Cluster `json:"data,omitempty"`
}

type NodeCondition struct {
	LastHeartbeatTime  string `json:"lastHeartbeatTime,omitempty" yaml:"lastHeartbeatTime,omitempty"`
	LastTransitionTime string `json:"lastTransitionTime,omitempty" yaml:"lastTransitionTime,omitempty"`
	Message            string `json:"message,omitempty" yaml:"message,omitempty"`
	Reason             string `json:"reason,omitempty" yaml:"reason,omitempty"`
	Status             string `json:"status,omitempty" yaml:"status,omitempty"`
	Type               string `json:"type,omitempty" yaml:"type,omitempty"`
}

type CustomConfig struct {
	Address         string            `json:"address,omitempty" yaml:"address,omitempty"`
	DockerSocket    string            `json:"dockerSocket,omitempty" yaml:"dockerSocket,omitempty"`
	InternalAddress string            `json:"internalAddress,omitempty" yaml:"internalAddress,omitempty"`
	Label           map[string]string `json:"label,omitempty" yaml:"label,omitempty"`
	SSHCert         string            `json:"sshCert,omitempty" yaml:"sshCert,omitempty"`
	SSHKey          string            `json:"sshKey,omitempty" yaml:"sshKey,omitempty"`
	Taints          []string          `json:"taints,omitempty" yaml:"taints,omitempty"`
	User            string            `json:"user,omitempty" yaml:"user,omitempty"`
}

type DockerInfo struct {
	Architecture       string   `json:"architecture,omitempty" yaml:"architecture,omitempty"`
	CgroupDriver       string   `json:"cgroupDriver,omitempty" yaml:"cgroupDriver,omitempty"`
	Debug              bool     `json:"debug,omitempty" yaml:"debug,omitempty"`
	DockerRootDir      string   `json:"dockerRootDir,omitempty" yaml:"dockerRootDir,omitempty"`
	Driver             string   `json:"driver,omitempty" yaml:"driver,omitempty"`
	ExperimentalBuild  bool     `json:"experimentalBuild,omitempty" yaml:"experimentalBuild,omitempty"`
	HTTPProxy          string   `json:"httpProxy,omitempty" yaml:"httpProxy,omitempty"`
	HTTPSProxy         string   `json:"httpsProxy,omitempty" yaml:"httpsProxy,omitempty"`
	IndexServerAddress string   `json:"indexServerAddress,omitempty" yaml:"indexServerAddress,omitempty"`
	InitBinary         string   `json:"initBinary,omitempty" yaml:"initBinary,omitempty"`
	KernelVersion      string   `json:"kernelVersion,omitempty" yaml:"kernelVersion,omitempty"`
	Labels             []string `json:"labels,omitempty" yaml:"labels,omitempty"`
	LoggingDriver      string   `json:"loggingDriver,omitempty" yaml:"loggingDriver,omitempty"`
	Name               string   `json:"name,omitempty" yaml:"name,omitempty"`
	NoProxy            string   `json:"noProxy,omitempty" yaml:"noProxy,omitempty"`
	OSType             string   `json:"osType,omitempty" yaml:"osType,omitempty"`
	OperatingSystem    string   `json:"operatingSystem,omitempty" yaml:"operatingSystem,omitempty"`
	SecurityOptions    []string `json:"securityOptions,omitempty" yaml:"securityOptions,omitempty"`
	ServerVersion      string   `json:"serverVersion,omitempty" yaml:"serverVersion,omitempty"`
}

type CPUInfo struct {
	Count int64 `json:"count,omitempty" yaml:"count,omitempty"`
}

type KubernetesInfo struct {
	KubeProxyVersion string `json:"kubeProxyVersion,omitempty" yaml:"kubeProxyVersion,omitempty"`
	KubeletVersion   string `json:"kubeletVersion,omitempty" yaml:"kubeletVersion,omitempty"`
}

type MemoryInfo struct {
	MemTotalKiB int64 `json:"memTotalKiB,omitempty" yaml:"memTotalKiB,omitempty"`
}

type OSInfo struct {
	DockerVersion   string `json:"dockerVersion,omitempty" yaml:"dockerVersion,omitempty"`
	KernelVersion   string `json:"kernelVersion,omitempty" yaml:"kernelVersion,omitempty"`
	OperatingSystem string `json:"operatingSystem,omitempty" yaml:"operatingSystem,omitempty"`
}

type NodeInfo struct {
	CPU        *CPUInfo        `json:"cpu,omitempty" yaml:"cpu,omitempty"`
	Kubernetes *KubernetesInfo `json:"kubernetes,omitempty" yaml:"kubernetes,omitempty"`
	Memory     *MemoryInfo     `json:"memory,omitempty" yaml:"memory,omitempty"`
	OS         *OSInfo         `json:"os,omitempty" yaml:"os,omitempty"`
}

type File struct {
	Contents string `json:"contents,omitempty" yaml:"contents,omitempty"`
	Name     string `json:"name,omitempty" yaml:"name,omitempty"`
}

type PortCheck struct {
	Address  string `json:"address,omitempty" yaml:"address,omitempty"`
	Port     int64  `json:"port,omitempty" yaml:"port,omitempty"`
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
}

type HealthCheck struct {
	URL string `json:"url,omitempty" yaml:"url,omitempty"`
}

type Process struct {
	Args                    []string          `json:"args,omitempty" yaml:"args,omitempty"`
	Binds                   []string          `json:"binds,omitempty" yaml:"binds,omitempty"`
	Command                 []string          `json:"command,omitempty" yaml:"command,omitempty"`
	Env                     []string          `json:"env,omitempty" yaml:"env,omitempty"`
	HealthCheck             *HealthCheck      `json:"healthCheck,omitempty" yaml:"healthCheck,omitempty"`
	Image                   string            `json:"image,omitempty" yaml:"image,omitempty"`
	ImageRegistryAuthConfig string            `json:"imageRegistryAuthConfig,omitempty" yaml:"imageRegistryAuthConfig,omitempty"`
	Labels                  map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                    string            `json:"name,omitempty" yaml:"name,omitempty"`
	NetworkMode             string            `json:"networkMode,omitempty" yaml:"networkMode,omitempty"`
	PidMode                 string            `json:"pidMode,omitempty" yaml:"pidMode,omitempty"`
	Privileged              bool              `json:"privileged,omitempty" yaml:"privileged,omitempty"`
	Publish                 []string          `json:"publish,omitempty" yaml:"publish,omitempty"`
	RestartPolicy           string            `json:"restartPolicy,omitempty" yaml:"restartPolicy,omitempty"`
	User                    string            `json:"user,omitempty" yaml:"user,omitempty"`
	VolumesFrom             []string          `json:"volumesFrom,omitempty" yaml:"volumesFrom,omitempty"`
}

type RKEConfigNodePlan struct {
	Address     string             `json:"address,omitempty" yaml:"address,omitempty"`
	Annotations map[string]string  `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Files       []File             `json:"files,omitempty" yaml:"files,omitempty"`
	Labels      map[string]string  `json:"labels,omitempty" yaml:"labels,omitempty"`
	PortChecks  []PortCheck        `json:"portChecks,omitempty" yaml:"portChecks,omitempty"`
	Processes   map[string]Process `json:"processes,omitempty" yaml:"processes,omitempty"`
	Taints      []RKETaint         `json:"taints,omitempty" yaml:"taints,omitempty"`
}

type NodePlan struct {
	AgentCheckInterval int64              `json:"agentCheckInterval,omitempty" yaml:"agentCheckInterval,omitempty"`
	Plan               *RKEConfigNodePlan `json:"plan,omitempty" yaml:"plan,omitempty"`
	Version            int64              `json:"version,omitempty" yaml:"version,omitempty"`
}

type Taint struct {
	Effect    string `json:"effect,omitempty" yaml:"effect,omitempty"`
	Key       string `json:"key,omitempty" yaml:"key,omitempty"`
	TimeAdded string `json:"timeAdded,omitempty" yaml:"timeAdded,omitempty"`
	Value     string `json:"value,omitempty" yaml:"value,omitempty"`
}

type PublicEndpoint struct {
	Addresses []string `json:"addresses,omitempty" yaml:"addresses,omitempty"`
	AllNodes  bool     `json:"allNodes,omitempty" yaml:"allNodes,omitempty"`
	Hostname  string   `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	IngressID string   `json:"ingressId,omitempty" yaml:"ingressId,omitempty"`
	NodeID    string   `json:"nodeId,omitempty" yaml:"nodeId,omitempty"`
	Path      string   `json:"path,omitempty" yaml:"path,omitempty"`
	PodID     string   `json:"podId,omitempty" yaml:"podId,omitempty"`
	Port      int64    `json:"port,omitempty" yaml:"port,omitempty"`
	Protocol  string   `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	ServiceID string   `json:"serviceId,omitempty" yaml:"serviceId,omitempty"`
}

type AttachedVolume struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

type Node struct {
	Resource
	Allocatable          map[string]string         `json:"allocatable,omitempty" yaml:"allocatable,omitempty"`
	Annotations          map[string]string         `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	AppliedNodeVersion   int64                     `json:"appliedNodeVersion,omitempty" yaml:"appliedNodeVersion,omitempty"`
	Capacity             map[string]string         `json:"capacity,omitempty" yaml:"capacity,omitempty"`
	ClusterID            string                    `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	Conditions           []NodeCondition           `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	ControlPlane         bool                      `json:"controlPlane,omitempty" yaml:"controlPlane,omitempty"`
	Created              string                    `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID            string                    `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	CustomConfig         *CustomConfig             `json:"customConfig,omitempty" yaml:"customConfig,omitempty"`
	Description          string                    `json:"description,omitempty" yaml:"description,omitempty"`
	DockerInfo           *DockerInfo               `json:"dockerInfo,omitempty" yaml:"dockerInfo,omitempty"`
	Etcd                 bool                      `json:"etcd,omitempty" yaml:"etcd,omitempty"`
	ExternalIPAddress    string                    `json:"externalIpAddress,omitempty" yaml:"externalIpAddress,omitempty"`
	Hostname             string                    `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	IPAddress            string                    `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`
	Imported             bool                      `json:"imported,omitempty" yaml:"imported,omitempty"`
	Info                 *NodeInfo                 `json:"info,omitempty" yaml:"info,omitempty"`
	Labels               map[string]string         `json:"labels,omitempty" yaml:"labels,omitempty"`
	Limits               map[string]string         `json:"limits,omitempty" yaml:"limits,omitempty"`
	Name                 string                    `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId          string                    `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	NodeName             string                    `json:"nodeName,omitempty" yaml:"nodeName,omitempty"`
	NodePlan             *NodePlan                 `json:"nodePlan,omitempty" yaml:"nodePlan,omitempty"`
	NodePoolID           string                    `json:"nodePoolId,omitempty" yaml:"nodePoolId,omitempty"`
	NodeTaints           []Taint                   `json:"nodeTaints,omitempty" yaml:"nodeTaints,omitempty"`
	NodeTemplateID       string                    `json:"nodeTemplateId,omitempty" yaml:"nodeTemplateId,omitempty"`
	OwnerReferences      []OwnerReference          `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	PodCidr              string                    `json:"podCidr,omitempty" yaml:"podCidr,omitempty"`
	PodCidrs             []string                  `json:"podCidrs,omitempty" yaml:"podCidrs,omitempty"`
	ProviderId           string                    `json:"providerId,omitempty" yaml:"providerId,omitempty"`
	PublicEndpoints      []PublicEndpoint          `json:"publicEndpoints,omitempty" yaml:"publicEndpoints,omitempty"`
	Removed              string                    `json:"removed,omitempty" yaml:"removed,omitempty"`
	Requested            map[string]string         `json:"requested,omitempty" yaml:"requested,omitempty"`
	RequestedHostname    string                    `json:"requestedHostname,omitempty" yaml:"requestedHostname,omitempty"`
	ScaledownTime        string                    `json:"scaledownTime,omitempty" yaml:"scaledownTime,omitempty"`
	SshUser              string                    `json:"sshUser,omitempty" yaml:"sshUser,omitempty"`
	State                string                    `json:"state,omitempty" yaml:"state,omitempty"`
	Taints               []Taint                   `json:"taints,omitempty" yaml:"taints,omitempty"`
	Transitioning        string                    `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage string                    `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
	UUID                 string                    `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Unschedulable        bool                      `json:"unschedulable,omitempty" yaml:"unschedulable,omitempty"`
	VolumesAttached      map[string]AttachedVolume `json:"volumesAttached,omitempty" yaml:"volumesAttached,omitempty"`
	VolumesInUse         []string                  `json:"volumesInUse,omitempty" yaml:"volumesInUse,omitempty"`
	Worker               bool                      `json:"worker,omitempty" yaml:"worker,omitempty"`
}

type NodeCollection struct {
	Collection
	Data []Node `json:"data,omitempty"`
}

type GenericLogin struct {
	TTLMillis    int64  `json:"ttl,omitempty"`
	Description  string `json:"description,omitempty" norman:"type=string,required"`
	ResponseType string `json:"responseType,omitempty" norman:"type=string,required"` //json or cookie
}

type BasicLogin struct {
	GenericLogin `json:",inline"`
	Username     string `json:"username" norman:"type=string,required"`
	Password     string `json:"password" norman:"type=string,required"`
}

type Token struct {
	Resource
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	AuthProvider    string            `json:"authProvider,omitempty" yaml:"authProvider,omitempty"`
	ClusterID       string            `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Current         bool              `json:"current,omitempty" yaml:"current,omitempty"`
	Description     string            `json:"description,omitempty" yaml:"description,omitempty"`
	Enabled         *bool             `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Expired         bool              `json:"expired,omitempty" yaml:"expired,omitempty"`
	ExpiresAt       string            `json:"expiresAt,omitempty" yaml:"expiresAt,omitempty"`
	GroupPrincipals []string          `json:"groupPrincipals,omitempty" yaml:"groupPrincipals,omitempty"`
	IsDerived       bool              `json:"isDerived,omitempty" yaml:"isDerived,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	LastUpdateTime  string            `json:"lastUpdateTime,omitempty" yaml:"lastUpdateTime,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	ProviderInfo    map[string]string `json:"providerInfo,omitempty" yaml:"providerInfo,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	TTLMillis       int64             `json:"ttl,omitempty" yaml:"ttl,omitempty"`
	Token           string            `json:"token,omitempty" yaml:"token,omitempty"`
	UUID            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	UserID          string            `json:"userId,omitempty" yaml:"userId,omitempty"`
	UserPrincipal   string            `json:"userPrincipal,omitempty" yaml:"userPrincipal,omitempty"`
}

type TokenCollection struct {
	Collection
	Data []Token `json:"data,omitempty"`
}

type ClusterRegistrationToken struct {
	Resource
	Annotations          map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	ClusterID            string            `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	Command              string            `json:"command,omitempty" yaml:"command,omitempty"`
	Created              string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID            string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	InsecureCommand      string            `json:"insecureCommand,omitempty" yaml:"insecureCommand,omitempty"`
	Labels               map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	ManifestURL          string            `json:"manifestUrl,omitempty" yaml:"manifestUrl,omitempty"`
	Name                 string            `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId          string            `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	NodeCommand          string            `json:"nodeCommand,omitempty" yaml:"nodeCommand,omitempty"`
	OwnerReferences      []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed              string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	State                string            `json:"state,omitempty" yaml:"state,omitempty"`
	Token                string            `json:"token,omitempty" yaml:"token,omitempty"`
	Transitioning        string            `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage string            `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
	UUID                 string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	WindowsNodeCommand   string            `json:"windowsNodeCommand,omitempty" yaml:"windowsNodeCommand,omitempty"`
}

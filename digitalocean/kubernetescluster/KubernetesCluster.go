// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v11/kubernetescluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.45.0/docs/resources/kubernetes_cluster digitalocean_kubernetes_cluster}.
type KubernetesCluster interface {
	cdktf.TerraformResource
	AutoUpgrade() interface{}
	SetAutoUpgrade(val interface{})
	AutoUpgradeInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterSubnet() *string
	SetClusterSubnet(val *string)
	ClusterSubnetInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DestroyAllAssociatedResources() interface{}
	SetDestroyAllAssociatedResources(val interface{})
	DestroyAllAssociatedResourcesInput() interface{}
	Endpoint() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Ha() interface{}
	SetHa(val interface{})
	HaInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	Ipv4Address() *string
	KubeConfig() KubernetesClusterKubeConfigList
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaintenancePolicy() KubernetesClusterMaintenancePolicyOutputReference
	MaintenancePolicyInput() *KubernetesClusterMaintenancePolicy
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodePool() KubernetesClusterNodePoolOutputReference
	NodePoolInput() *KubernetesClusterNodePool
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RegistryIntegration() interface{}
	SetRegistryIntegration(val interface{})
	RegistryIntegrationInput() interface{}
	ServiceSubnet() *string
	SetServiceSubnet(val *string)
	ServiceSubnetInput() *string
	Status() *string
	SurgeUpgrade() interface{}
	SetSurgeUpgrade(val interface{})
	SurgeUpgradeInput() interface{}
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() KubernetesClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdatedAt() *string
	Urn() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	VpcUuid() *string
	SetVpcUuid(val *string)
	VpcUuidInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutMaintenancePolicy(value *KubernetesClusterMaintenancePolicy)
	PutNodePool(value *KubernetesClusterNodePool)
	PutTimeouts(value *KubernetesClusterTimeouts)
	ResetAutoUpgrade()
	ResetClusterSubnet()
	ResetDestroyAllAssociatedResources()
	ResetHa()
	ResetId()
	ResetMaintenancePolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegistryIntegration()
	ResetServiceSubnet()
	ResetSurgeUpgrade()
	ResetTags()
	ResetTimeouts()
	ResetVpcUuid()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for KubernetesCluster
type jsiiProxy_KubernetesCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KubernetesCluster) AutoUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AutoUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ClusterSubnet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSubnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ClusterSubnetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSubnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DestroyAllAssociatedResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destroyAllAssociatedResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DestroyAllAssociatedResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destroyAllAssociatedResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Ha() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) HaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"haInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Ipv4Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KubeConfig() KubernetesClusterKubeConfigList {
	var returns KubernetesClusterKubeConfigList
	_jsii_.Get(
		j,
		"kubeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) MaintenancePolicy() KubernetesClusterMaintenancePolicyOutputReference {
	var returns KubernetesClusterMaintenancePolicyOutputReference
	_jsii_.Get(
		j,
		"maintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) MaintenancePolicyInput() *KubernetesClusterMaintenancePolicy {
	var returns *KubernetesClusterMaintenancePolicy
	_jsii_.Get(
		j,
		"maintenancePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NodePool() KubernetesClusterNodePoolOutputReference {
	var returns KubernetesClusterNodePoolOutputReference
	_jsii_.Get(
		j,
		"nodePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NodePoolInput() *KubernetesClusterNodePool {
	var returns *KubernetesClusterNodePool
	_jsii_.Get(
		j,
		"nodePoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) RegistryIntegration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registryIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) RegistryIntegrationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registryIntegrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ServiceSubnet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceSubnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ServiceSubnetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceSubnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) SurgeUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"surgeUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) SurgeUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"surgeUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Timeouts() KubernetesClusterTimeoutsOutputReference {
	var returns KubernetesClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Urn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) VpcUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) VpcUuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcUuidInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.45.0/docs/resources/kubernetes_cluster digitalocean_kubernetes_cluster} Resource.
func NewKubernetesCluster(scope constructs.Construct, id *string, config *KubernetesClusterConfig) KubernetesCluster {
	_init_.Initialize()

	if err := validateNewKubernetesClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesCluster{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.kubernetesCluster.KubernetesCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.45.0/docs/resources/kubernetes_cluster digitalocean_kubernetes_cluster} Resource.
func NewKubernetesCluster_Override(k KubernetesCluster, scope constructs.Construct, id *string, config *KubernetesClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.kubernetesCluster.KubernetesCluster",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetAutoUpgrade(val interface{}) {
	if err := j.validateSetAutoUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoUpgrade",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetClusterSubnet(val *string) {
	if err := j.validateSetClusterSubnetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterSubnet",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetDestroyAllAssociatedResources(val interface{}) {
	if err := j.validateSetDestroyAllAssociatedResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destroyAllAssociatedResources",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetHa(val interface{}) {
	if err := j.validateSetHaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ha",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetRegistryIntegration(val interface{}) {
	if err := j.validateSetRegistryIntegrationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registryIntegration",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetServiceSubnet(val *string) {
	if err := j.validateSetServiceSubnetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceSubnet",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetSurgeUpgrade(val interface{}) {
	if err := j.validateSetSurgeUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"surgeUpgrade",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetVpcUuid(val *string) {
	if err := j.validateSetVpcUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcUuid",
		val,
	)
}

// Generates CDKTF code for importing a KubernetesCluster resource upon running "cdktf plan <stack-name>".
func KubernetesCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateKubernetesCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.kubernetesCluster.KubernetesCluster",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func KubernetesCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.kubernetesCluster.KubernetesCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KubernetesCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.kubernetesCluster.KubernetesCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KubernetesCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.kubernetesCluster.KubernetesCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KubernetesCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-digitalocean.kubernetesCluster.KubernetesCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubernetesCluster) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KubernetesCluster) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KubernetesCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KubernetesCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KubernetesCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KubernetesCluster) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KubernetesCluster) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutMaintenancePolicy(value *KubernetesClusterMaintenancePolicy) {
	if err := k.validatePutMaintenancePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putMaintenancePolicy",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutNodePool(value *KubernetesClusterNodePool) {
	if err := k.validatePutNodePoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putNodePool",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutTimeouts(value *KubernetesClusterTimeouts) {
	if err := k.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetAutoUpgrade() {
	_jsii_.InvokeVoid(
		k,
		"resetAutoUpgrade",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetClusterSubnet() {
	_jsii_.InvokeVoid(
		k,
		"resetClusterSubnet",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetDestroyAllAssociatedResources() {
	_jsii_.InvokeVoid(
		k,
		"resetDestroyAllAssociatedResources",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetHa() {
	_jsii_.InvokeVoid(
		k,
		"resetHa",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetMaintenancePolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetMaintenancePolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetRegistryIntegration() {
	_jsii_.InvokeVoid(
		k,
		"resetRegistryIntegration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetServiceSubnet() {
	_jsii_.InvokeVoid(
		k,
		"resetServiceSubnet",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetSurgeUpgrade() {
	_jsii_.InvokeVoid(
		k,
		"resetSurgeUpgrade",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetTags() {
	_jsii_.InvokeVoid(
		k,
		"resetTags",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetVpcUuid() {
	_jsii_.InvokeVoid(
		k,
		"resetVpcUuid",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package droplet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v10/droplet/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.34.1/docs/resources/droplet digitalocean_droplet}.
type Droplet interface {
	cdktf.TerraformResource
	Backups() interface{}
	SetBackups(val interface{})
	BackupsInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	Disk() *float64
	DropletAgent() interface{}
	SetDropletAgent(val interface{})
	DropletAgentInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GracefulShutdown() interface{}
	SetGracefulShutdown(val interface{})
	GracefulShutdownInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	Ipv4Address() *string
	Ipv4AddressPrivate() *string
	Ipv6() interface{}
	SetIpv6(val interface{})
	Ipv6Address() *string
	SetIpv6Address(val *string)
	Ipv6AddressInput() *string
	Ipv6Input() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Locked() cdktf.IResolvable
	Memory() *float64
	Monitoring() interface{}
	SetMonitoring(val interface{})
	MonitoringInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PriceHourly() *float64
	PriceMonthly() *float64
	PrivateNetworking() interface{}
	SetPrivateNetworking(val interface{})
	PrivateNetworkingInput() interface{}
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
	ResizeDisk() interface{}
	SetResizeDisk(val interface{})
	ResizeDiskInput() interface{}
	Size() *string
	SetSize(val *string)
	SizeInput() *string
	SshKeys() *[]*string
	SetSshKeys(val *[]*string)
	SshKeysInput() *[]*string
	Status() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DropletTimeoutsOutputReference
	TimeoutsInput() interface{}
	Urn() *string
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
	Vcpus() *float64
	VolumeIds() *[]*string
	SetVolumeIds(val *[]*string)
	VolumeIdsInput() *[]*string
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
	PutTimeouts(value *DropletTimeouts)
	ResetBackups()
	ResetDropletAgent()
	ResetGracefulShutdown()
	ResetId()
	ResetIpv6()
	ResetIpv6Address()
	ResetMonitoring()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateNetworking()
	ResetRegion()
	ResetResizeDisk()
	ResetSshKeys()
	ResetTags()
	ResetTimeouts()
	ResetUserData()
	ResetVolumeIds()
	ResetVpcUuid()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Droplet
type jsiiProxy_Droplet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Droplet) Backups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) BackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Disk() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"disk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) DropletAgent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropletAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) DropletAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropletAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) GracefulShutdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gracefulShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) GracefulShutdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gracefulShutdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Ipv4Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Ipv4AddressPrivate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4AddressPrivate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Ipv6() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Ipv6Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Ipv6AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6AddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Ipv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Locked() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"locked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Monitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) MonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) PriceHourly() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priceHourly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) PriceMonthly() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priceMonthly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) PrivateNetworking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateNetworking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) PrivateNetworkingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateNetworkingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) ResizeDisk() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resizeDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) ResizeDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resizeDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Size() *string {
	var returns *string
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) SizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) SshKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) SshKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Timeouts() DropletTimeoutsOutputReference {
	var returns DropletTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Urn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) Vcpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vcpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) VolumeIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"volumeIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) VolumeIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"volumeIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) VpcUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Droplet) VpcUuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcUuidInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.34.1/docs/resources/droplet digitalocean_droplet} Resource.
func NewDroplet(scope constructs.Construct, id *string, config *DropletConfig) Droplet {
	_init_.Initialize()

	if err := validateNewDropletParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Droplet{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.droplet.Droplet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.34.1/docs/resources/droplet digitalocean_droplet} Resource.
func NewDroplet_Override(d Droplet, scope constructs.Construct, id *string, config *DropletConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.droplet.Droplet",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_Droplet)SetBackups(val interface{}) {
	if err := j.validateSetBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backups",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetDropletAgent(val interface{}) {
	if err := j.validateSetDropletAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dropletAgent",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetGracefulShutdown(val interface{}) {
	if err := j.validateSetGracefulShutdownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gracefulShutdown",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetIpv6(val interface{}) {
	if err := j.validateSetIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetIpv6Address(val *string) {
	if err := j.validateSetIpv6AddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Address",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetMonitoring(val interface{}) {
	if err := j.validateSetMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoring",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetPrivateNetworking(val interface{}) {
	if err := j.validateSetPrivateNetworkingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateNetworking",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetResizeDisk(val interface{}) {
	if err := j.validateSetResizeDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resizeDisk",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetSize(val *string) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetSshKeys(val *[]*string) {
	if err := j.validateSetSshKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshKeys",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetVolumeIds(val *[]*string) {
	if err := j.validateSetVolumeIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeIds",
		val,
	)
}

func (j *jsiiProxy_Droplet)SetVpcUuid(val *string) {
	if err := j.validateSetVpcUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcUuid",
		val,
	)
}

// Generates CDKTF code for importing a Droplet resource upon running "cdktf plan <stack-name>".
func Droplet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDroplet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.droplet.Droplet",
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
func Droplet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDroplet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.droplet.Droplet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Droplet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDroplet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.droplet.Droplet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Droplet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDroplet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.droplet.Droplet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Droplet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-digitalocean.droplet.Droplet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_Droplet) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_Droplet) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_Droplet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Droplet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Droplet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Droplet) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Droplet) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Droplet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Droplet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Droplet) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Droplet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Droplet) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Droplet) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_Droplet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Droplet) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_Droplet) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_Droplet) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_Droplet) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_Droplet) PutTimeouts(value *DropletTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_Droplet) ResetBackups() {
	_jsii_.InvokeVoid(
		d,
		"resetBackups",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Droplet) ResetDropletAgent() {
	_jsii_.InvokeVoid(
		d,
		"resetDropletAgent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Droplet) ResetGracefulShutdown() {
	_jsii_.InvokeVoid(
		d,
		"resetGracefulShutdown",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Droplet) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Droplet) ResetIpv6() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv6",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Droplet) ResetIpv6Address() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv6Address",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Droplet) ResetMonitoring() {
	_jsii_.InvokeVoid(
		d,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Droplet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Droplet) ResetPrivateNetworking() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivateNetworking",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Droplet) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Droplet) ResetResizeDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetResizeDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Droplet) ResetSshKeys() {
	_jsii_.InvokeVoid(
		d,
		"resetSshKeys",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Droplet) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Droplet) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Droplet) ResetUserData() {
	_jsii_.InvokeVoid(
		d,
		"resetUserData",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Droplet) ResetVolumeIds() {
	_jsii_.InvokeVoid(
		d,
		"resetVolumeIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Droplet) ResetVpcUuid() {
	_jsii_.InvokeVoid(
		d,
		"resetVpcUuid",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Droplet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Droplet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Droplet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Droplet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


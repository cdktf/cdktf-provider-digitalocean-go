package loadbalancer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v5/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v5/loadbalancer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/digitalocean/r/loadbalancer digitalocean_loadbalancer}.
type Loadbalancer interface {
	cdktf.TerraformResource
	Algorithm() *string
	SetAlgorithm(val *string)
	AlgorithmInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableLetsEncryptDnsRecords() interface{}
	SetDisableLetsEncryptDnsRecords(val interface{})
	DisableLetsEncryptDnsRecordsInput() interface{}
	DropletIds() *[]*float64
	SetDropletIds(val *[]*float64)
	DropletIdsInput() *[]*float64
	DropletTag() *string
	SetDropletTag(val *string)
	DropletTagInput() *string
	EnableBackendKeepalive() interface{}
	SetEnableBackendKeepalive(val interface{})
	EnableBackendKeepaliveInput() interface{}
	EnableProxyProtocol() interface{}
	SetEnableProxyProtocol(val interface{})
	EnableProxyProtocolInput() interface{}
	Firewall() LoadbalancerFirewallOutputReference
	FirewallInput() *LoadbalancerFirewall
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	ForwardingRule() LoadbalancerForwardingRuleList
	ForwardingRuleInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Healthcheck() LoadbalancerHealthcheckOutputReference
	HealthcheckInput() *LoadbalancerHealthcheck
	HttpIdleTimeoutSeconds() *float64
	SetHttpIdleTimeoutSeconds(val *float64)
	HttpIdleTimeoutSecondsInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	Ip() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
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
	RedirectHttpToHttps() interface{}
	SetRedirectHttpToHttps(val interface{})
	RedirectHttpToHttpsInput() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	Size() *string
	SetSize(val *string)
	SizeInput() *string
	SizeUnit() *float64
	SetSizeUnit(val *float64)
	SizeUnitInput() *float64
	Status() *string
	StickySessions() LoadbalancerStickySessionsOutputReference
	StickySessionsInput() *LoadbalancerStickySessions
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Urn() *string
	VpcUuid() *string
	SetVpcUuid(val *string)
	VpcUuidInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutFirewall(value *LoadbalancerFirewall)
	PutForwardingRule(value interface{})
	PutHealthcheck(value *LoadbalancerHealthcheck)
	PutStickySessions(value *LoadbalancerStickySessions)
	ResetAlgorithm()
	ResetDisableLetsEncryptDnsRecords()
	ResetDropletIds()
	ResetDropletTag()
	ResetEnableBackendKeepalive()
	ResetEnableProxyProtocol()
	ResetFirewall()
	ResetHealthcheck()
	ResetHttpIdleTimeoutSeconds()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProjectId()
	ResetRedirectHttpToHttps()
	ResetSize()
	ResetSizeUnit()
	ResetStickySessions()
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

// The jsii proxy struct for Loadbalancer
type jsiiProxy_Loadbalancer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Loadbalancer) Algorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) AlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) DisableLetsEncryptDnsRecords() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableLetsEncryptDnsRecords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) DisableLetsEncryptDnsRecordsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableLetsEncryptDnsRecordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) DropletIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"dropletIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) DropletIdsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"dropletIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) DropletTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dropletTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) DropletTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dropletTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) EnableBackendKeepalive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBackendKeepalive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) EnableBackendKeepaliveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBackendKeepaliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) EnableProxyProtocol() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableProxyProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) EnableProxyProtocolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableProxyProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) Firewall() LoadbalancerFirewallOutputReference {
	var returns LoadbalancerFirewallOutputReference
	_jsii_.Get(
		j,
		"firewall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) FirewallInput() *LoadbalancerFirewall {
	var returns *LoadbalancerFirewall
	_jsii_.Get(
		j,
		"firewallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) ForwardingRule() LoadbalancerForwardingRuleList {
	var returns LoadbalancerForwardingRuleList
	_jsii_.Get(
		j,
		"forwardingRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) ForwardingRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardingRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) Healthcheck() LoadbalancerHealthcheckOutputReference {
	var returns LoadbalancerHealthcheckOutputReference
	_jsii_.Get(
		j,
		"healthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) HealthcheckInput() *LoadbalancerHealthcheck {
	var returns *LoadbalancerHealthcheck
	_jsii_.Get(
		j,
		"healthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) HttpIdleTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpIdleTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) HttpIdleTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpIdleTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) Ip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) RedirectHttpToHttps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redirectHttpToHttps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) RedirectHttpToHttpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redirectHttpToHttpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) Size() *string {
	var returns *string
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) SizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) SizeUnit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) SizeUnitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) StickySessions() LoadbalancerStickySessionsOutputReference {
	var returns LoadbalancerStickySessionsOutputReference
	_jsii_.Get(
		j,
		"stickySessions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) StickySessionsInput() *LoadbalancerStickySessions {
	var returns *LoadbalancerStickySessions
	_jsii_.Get(
		j,
		"stickySessionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) Urn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) VpcUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Loadbalancer) VpcUuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcUuidInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/digitalocean/r/loadbalancer digitalocean_loadbalancer} Resource.
func NewLoadbalancer(scope constructs.Construct, id *string, config *LoadbalancerConfig) Loadbalancer {
	_init_.Initialize()

	if err := validateNewLoadbalancerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Loadbalancer{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.loadbalancer.Loadbalancer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/digitalocean/r/loadbalancer digitalocean_loadbalancer} Resource.
func NewLoadbalancer_Override(l Loadbalancer, scope constructs.Construct, id *string, config *LoadbalancerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.loadbalancer.Loadbalancer",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_Loadbalancer)SetAlgorithm(val *string) {
	if err := j.validateSetAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"algorithm",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetDisableLetsEncryptDnsRecords(val interface{}) {
	if err := j.validateSetDisableLetsEncryptDnsRecordsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableLetsEncryptDnsRecords",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetDropletIds(val *[]*float64) {
	if err := j.validateSetDropletIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dropletIds",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetDropletTag(val *string) {
	if err := j.validateSetDropletTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dropletTag",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetEnableBackendKeepalive(val interface{}) {
	if err := j.validateSetEnableBackendKeepaliveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBackendKeepalive",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetEnableProxyProtocol(val interface{}) {
	if err := j.validateSetEnableProxyProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableProxyProtocol",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetHttpIdleTimeoutSeconds(val *float64) {
	if err := j.validateSetHttpIdleTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpIdleTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetRedirectHttpToHttps(val interface{}) {
	if err := j.validateSetRedirectHttpToHttpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectHttpToHttps",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetSize(val *string) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetSizeUnit(val *float64) {
	if err := j.validateSetSizeUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeUnit",
		val,
	)
}

func (j *jsiiProxy_Loadbalancer)SetVpcUuid(val *string) {
	if err := j.validateSetVpcUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcUuid",
		val,
	)
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
func Loadbalancer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLoadbalancer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.loadbalancer.Loadbalancer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Loadbalancer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLoadbalancer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.loadbalancer.Loadbalancer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Loadbalancer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLoadbalancer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.loadbalancer.Loadbalancer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Loadbalancer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-digitalocean.loadbalancer.Loadbalancer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_Loadbalancer) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_Loadbalancer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Loadbalancer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Loadbalancer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Loadbalancer) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Loadbalancer) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Loadbalancer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Loadbalancer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Loadbalancer) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Loadbalancer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Loadbalancer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Loadbalancer) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_Loadbalancer) PutFirewall(value *LoadbalancerFirewall) {
	if err := l.validatePutFirewallParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFirewall",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Loadbalancer) PutForwardingRule(value interface{}) {
	if err := l.validatePutForwardingRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putForwardingRule",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Loadbalancer) PutHealthcheck(value *LoadbalancerHealthcheck) {
	if err := l.validatePutHealthcheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putHealthcheck",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Loadbalancer) PutStickySessions(value *LoadbalancerStickySessions) {
	if err := l.validatePutStickySessionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putStickySessions",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Loadbalancer) ResetAlgorithm() {
	_jsii_.InvokeVoid(
		l,
		"resetAlgorithm",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Loadbalancer) ResetDisableLetsEncryptDnsRecords() {
	_jsii_.InvokeVoid(
		l,
		"resetDisableLetsEncryptDnsRecords",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Loadbalancer) ResetDropletIds() {
	_jsii_.InvokeVoid(
		l,
		"resetDropletIds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Loadbalancer) ResetDropletTag() {
	_jsii_.InvokeVoid(
		l,
		"resetDropletTag",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Loadbalancer) ResetEnableBackendKeepalive() {
	_jsii_.InvokeVoid(
		l,
		"resetEnableBackendKeepalive",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Loadbalancer) ResetEnableProxyProtocol() {
	_jsii_.InvokeVoid(
		l,
		"resetEnableProxyProtocol",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Loadbalancer) ResetFirewall() {
	_jsii_.InvokeVoid(
		l,
		"resetFirewall",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Loadbalancer) ResetHealthcheck() {
	_jsii_.InvokeVoid(
		l,
		"resetHealthcheck",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Loadbalancer) ResetHttpIdleTimeoutSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetHttpIdleTimeoutSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Loadbalancer) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Loadbalancer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Loadbalancer) ResetProjectId() {
	_jsii_.InvokeVoid(
		l,
		"resetProjectId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Loadbalancer) ResetRedirectHttpToHttps() {
	_jsii_.InvokeVoid(
		l,
		"resetRedirectHttpToHttps",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Loadbalancer) ResetSize() {
	_jsii_.InvokeVoid(
		l,
		"resetSize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Loadbalancer) ResetSizeUnit() {
	_jsii_.InvokeVoid(
		l,
		"resetSizeUnit",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Loadbalancer) ResetStickySessions() {
	_jsii_.InvokeVoid(
		l,
		"resetStickySessions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Loadbalancer) ResetVpcUuid() {
	_jsii_.InvokeVoid(
		l,
		"resetVpcUuid",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Loadbalancer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Loadbalancer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Loadbalancer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Loadbalancer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


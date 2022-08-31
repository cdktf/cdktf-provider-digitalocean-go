// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-digitalocean-go/digitalocean/v2/jsii"

	"github.com/hashicorp/cdktf-provider-digitalocean-go/digitalocean/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDigitaloceanFirewallInboundRuleOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PortRange() *string
	SetPortRange(val *string)
	PortRangeInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	SourceAddresses() *[]*string
	SetSourceAddresses(val *[]*string)
	SourceAddressesInput() *[]*string
	SourceDropletIds() *[]*float64
	SetSourceDropletIds(val *[]*float64)
	SourceDropletIdsInput() *[]*float64
	SourceKubernetesIds() *[]*string
	SetSourceKubernetesIds(val *[]*string)
	SourceKubernetesIdsInput() *[]*string
	SourceLoadBalancerUids() *[]*string
	SetSourceLoadBalancerUids(val *[]*string)
	SourceLoadBalancerUidsInput() *[]*string
	SourceTags() *[]*string
	SetSourceTags(val *[]*string)
	SourceTagsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetPortRange()
	ResetSourceAddresses()
	ResetSourceDropletIds()
	ResetSourceKubernetesIds()
	ResetSourceLoadBalancerUids()
	ResetSourceTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDigitaloceanFirewallInboundRuleOutputReference
type jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) PortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) PortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) SourceAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) SourceAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) SourceDropletIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"sourceDropletIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) SourceDropletIdsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"sourceDropletIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) SourceKubernetesIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceKubernetesIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) SourceKubernetesIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceKubernetesIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) SourceLoadBalancerUids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceLoadBalancerUids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) SourceLoadBalancerUidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceLoadBalancerUidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) SourceTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) SourceTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDigitaloceanFirewallInboundRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDigitaloceanFirewallInboundRuleOutputReference {
	_init_.Initialize()

	if err := validateNewDataDigitaloceanFirewallInboundRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.DataDigitaloceanFirewallInboundRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDigitaloceanFirewallInboundRuleOutputReference_Override(d DataDigitaloceanFirewallInboundRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.DataDigitaloceanFirewallInboundRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference)SetPortRange(val *string) {
	if err := j.validateSetPortRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portRange",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference)SetSourceAddresses(val *[]*string) {
	if err := j.validateSetSourceAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAddresses",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference)SetSourceDropletIds(val *[]*float64) {
	if err := j.validateSetSourceDropletIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDropletIds",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference)SetSourceKubernetesIds(val *[]*string) {
	if err := j.validateSetSourceKubernetesIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceKubernetesIds",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference)SetSourceLoadBalancerUids(val *[]*string) {
	if err := j.validateSetSourceLoadBalancerUidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceLoadBalancerUids",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference)SetSourceTags(val *[]*string) {
	if err := j.validateSetSourceTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceTags",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) ResetPortRange() {
	_jsii_.InvokeVoid(
		d,
		"resetPortRange",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) ResetSourceAddresses() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceAddresses",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) ResetSourceDropletIds() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceDropletIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) ResetSourceKubernetesIds() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceKubernetesIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) ResetSourceLoadBalancerUids() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceLoadBalancerUids",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) ResetSourceTags() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanFirewallInboundRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


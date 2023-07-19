package datadigitaloceanfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v8/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v8/datadigitaloceanfirewall/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDigitaloceanFirewallOutboundRuleOutputReference interface {
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
	DestinationAddresses() *[]*string
	SetDestinationAddresses(val *[]*string)
	DestinationAddressesInput() *[]*string
	DestinationDropletIds() *[]*float64
	SetDestinationDropletIds(val *[]*float64)
	DestinationDropletIdsInput() *[]*float64
	DestinationKubernetesIds() *[]*string
	SetDestinationKubernetesIds(val *[]*string)
	DestinationKubernetesIdsInput() *[]*string
	DestinationLoadBalancerUids() *[]*string
	SetDestinationLoadBalancerUids(val *[]*string)
	DestinationLoadBalancerUidsInput() *[]*string
	DestinationTags() *[]*string
	SetDestinationTags(val *[]*string)
	DestinationTagsInput() *[]*string
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
	ResetDestinationAddresses()
	ResetDestinationDropletIds()
	ResetDestinationKubernetesIds()
	ResetDestinationLoadBalancerUids()
	ResetDestinationTags()
	ResetPortRange()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDigitaloceanFirewallOutboundRuleOutputReference
type jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) DestinationAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) DestinationAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) DestinationDropletIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"destinationDropletIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) DestinationDropletIdsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"destinationDropletIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) DestinationKubernetesIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationKubernetesIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) DestinationKubernetesIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationKubernetesIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) DestinationLoadBalancerUids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationLoadBalancerUids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) DestinationLoadBalancerUidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationLoadBalancerUidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) DestinationTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) DestinationTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) PortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) PortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDigitaloceanFirewallOutboundRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDigitaloceanFirewallOutboundRuleOutputReference {
	_init_.Initialize()

	if err := validateNewDataDigitaloceanFirewallOutboundRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.dataDigitaloceanFirewall.DataDigitaloceanFirewallOutboundRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDigitaloceanFirewallOutboundRuleOutputReference_Override(d DataDigitaloceanFirewallOutboundRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.dataDigitaloceanFirewall.DataDigitaloceanFirewallOutboundRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference)SetDestinationAddresses(val *[]*string) {
	if err := j.validateSetDestinationAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddresses",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference)SetDestinationDropletIds(val *[]*float64) {
	if err := j.validateSetDestinationDropletIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationDropletIds",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference)SetDestinationKubernetesIds(val *[]*string) {
	if err := j.validateSetDestinationKubernetesIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationKubernetesIds",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference)SetDestinationLoadBalancerUids(val *[]*string) {
	if err := j.validateSetDestinationLoadBalancerUidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationLoadBalancerUids",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference)SetDestinationTags(val *[]*string) {
	if err := j.validateSetDestinationTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationTags",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference)SetPortRange(val *string) {
	if err := j.validateSetPortRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portRange",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) ResetDestinationAddresses() {
	_jsii_.InvokeVoid(
		d,
		"resetDestinationAddresses",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) ResetDestinationDropletIds() {
	_jsii_.InvokeVoid(
		d,
		"resetDestinationDropletIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) ResetDestinationKubernetesIds() {
	_jsii_.InvokeVoid(
		d,
		"resetDestinationKubernetesIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) ResetDestinationLoadBalancerUids() {
	_jsii_.InvokeVoid(
		d,
		"resetDestinationLoadBalancerUids",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) ResetDestinationTags() {
	_jsii_.InvokeVoid(
		d,
		"resetDestinationTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) ResetPortRange() {
	_jsii_.InvokeVoid(
		d,
		"resetPortRange",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDigitaloceanFirewallOutboundRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


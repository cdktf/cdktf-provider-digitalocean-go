package firewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v5/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v5/firewall/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirewallOutboundRuleOutputReference interface {
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

// The jsii proxy struct for FirewallOutboundRuleOutputReference
type jsiiProxy_FirewallOutboundRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) DestinationAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) DestinationAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) DestinationDropletIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"destinationDropletIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) DestinationDropletIdsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"destinationDropletIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) DestinationKubernetesIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationKubernetesIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) DestinationKubernetesIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationKubernetesIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) DestinationLoadBalancerUids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationLoadBalancerUids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) DestinationLoadBalancerUidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationLoadBalancerUidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) DestinationTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) DestinationTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) PortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) PortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFirewallOutboundRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FirewallOutboundRuleOutputReference {
	_init_.Initialize()

	if err := validateNewFirewallOutboundRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirewallOutboundRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.firewall.FirewallOutboundRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFirewallOutboundRuleOutputReference_Override(f FirewallOutboundRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.firewall.FirewallOutboundRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference)SetDestinationAddresses(val *[]*string) {
	if err := j.validateSetDestinationAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddresses",
		val,
	)
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference)SetDestinationDropletIds(val *[]*float64) {
	if err := j.validateSetDestinationDropletIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationDropletIds",
		val,
	)
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference)SetDestinationKubernetesIds(val *[]*string) {
	if err := j.validateSetDestinationKubernetesIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationKubernetesIds",
		val,
	)
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference)SetDestinationLoadBalancerUids(val *[]*string) {
	if err := j.validateSetDestinationLoadBalancerUidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationLoadBalancerUids",
		val,
	)
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference)SetDestinationTags(val *[]*string) {
	if err := j.validateSetDestinationTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationTags",
		val,
	)
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference)SetPortRange(val *string) {
	if err := j.validateSetPortRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portRange",
		val,
	)
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallOutboundRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FirewallOutboundRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallOutboundRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallOutboundRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallOutboundRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallOutboundRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallOutboundRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallOutboundRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallOutboundRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallOutboundRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallOutboundRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallOutboundRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallOutboundRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallOutboundRuleOutputReference) ResetDestinationAddresses() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationAddresses",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallOutboundRuleOutputReference) ResetDestinationDropletIds() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationDropletIds",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallOutboundRuleOutputReference) ResetDestinationKubernetesIds() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationKubernetesIds",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallOutboundRuleOutputReference) ResetDestinationLoadBalancerUids() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationLoadBalancerUids",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallOutboundRuleOutputReference) ResetDestinationTags() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallOutboundRuleOutputReference) ResetPortRange() {
	_jsii_.InvokeVoid(
		f,
		"resetPortRange",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallOutboundRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallOutboundRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


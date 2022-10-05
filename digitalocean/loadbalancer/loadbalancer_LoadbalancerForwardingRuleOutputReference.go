package loadbalancer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v3/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v3/loadbalancer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerForwardingRuleOutputReference interface {
	cdktf.ComplexObject
	CertificateId() *string
	SetCertificateId(val *string)
	CertificateIdInput() *string
	CertificateName() *string
	SetCertificateName(val *string)
	CertificateNameInput() *string
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
	EntryPort() *float64
	SetEntryPort(val *float64)
	EntryPortInput() *float64
	EntryProtocol() *string
	SetEntryProtocol(val *string)
	EntryProtocolInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	TargetPort() *float64
	SetTargetPort(val *float64)
	TargetPortInput() *float64
	TargetProtocol() *string
	SetTargetProtocol(val *string)
	TargetProtocolInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsPassthrough() interface{}
	SetTlsPassthrough(val interface{})
	TlsPassthroughInput() interface{}
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
	ResetCertificateId()
	ResetCertificateName()
	ResetTlsPassthrough()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LoadbalancerForwardingRuleOutputReference
type jsiiProxy_LoadbalancerForwardingRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) CertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) CertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) CertificateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) CertificateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) EntryPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"entryPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) EntryPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"entryPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) EntryProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) EntryProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) TargetPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) TargetPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) TargetProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) TargetProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) TlsPassthrough() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsPassthrough",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference) TlsPassthroughInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsPassthroughInput",
		&returns,
	)
	return returns
}


func NewLoadbalancerForwardingRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LoadbalancerForwardingRuleOutputReference {
	_init_.Initialize()

	if err := validateNewLoadbalancerForwardingRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LoadbalancerForwardingRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.loadbalancer.LoadbalancerForwardingRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLoadbalancerForwardingRuleOutputReference_Override(l LoadbalancerForwardingRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.loadbalancer.LoadbalancerForwardingRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference)SetCertificateId(val *string) {
	if err := j.validateSetCertificateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateId",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference)SetCertificateName(val *string) {
	if err := j.validateSetCertificateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateName",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference)SetEntryPort(val *float64) {
	if err := j.validateSetEntryPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryPort",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference)SetEntryProtocol(val *string) {
	if err := j.validateSetEntryProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryProtocol",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference)SetTargetPort(val *float64) {
	if err := j.validateSetTargetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetPort",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference)SetTargetProtocol(val *string) {
	if err := j.validateSetTargetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetProtocol",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerForwardingRuleOutputReference)SetTlsPassthrough(val interface{}) {
	if err := j.validateSetTlsPassthroughParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsPassthrough",
		val,
	)
}

func (l *jsiiProxy_LoadbalancerForwardingRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerForwardingRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LoadbalancerForwardingRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LoadbalancerForwardingRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LoadbalancerForwardingRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LoadbalancerForwardingRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LoadbalancerForwardingRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LoadbalancerForwardingRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LoadbalancerForwardingRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LoadbalancerForwardingRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LoadbalancerForwardingRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerForwardingRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerForwardingRuleOutputReference) ResetCertificateId() {
	_jsii_.InvokeVoid(
		l,
		"resetCertificateId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerForwardingRuleOutputReference) ResetCertificateName() {
	_jsii_.InvokeVoid(
		l,
		"resetCertificateName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerForwardingRuleOutputReference) ResetTlsPassthrough() {
	_jsii_.InvokeVoid(
		l,
		"resetTlsPassthrough",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerForwardingRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerForwardingRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package partnerattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v12/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v12/partnerattachment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PartnerAttachmentBgpOutputReference interface {
	cdktf.ComplexObject
	AuthKey() *string
	SetAuthKey(val *string)
	AuthKeyInput() *string
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
	InternalValue() *PartnerAttachmentBgp
	SetInternalValue(val *PartnerAttachmentBgp)
	LocalRouterIp() *string
	SetLocalRouterIp(val *string)
	LocalRouterIpInput() *string
	PeerRouterAsn() *float64
	SetPeerRouterAsn(val *float64)
	PeerRouterAsnInput() *float64
	PeerRouterIp() *string
	SetPeerRouterIp(val *string)
	PeerRouterIpInput() *string
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
	ResetAuthKey()
	ResetLocalRouterIp()
	ResetPeerRouterAsn()
	ResetPeerRouterIp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PartnerAttachmentBgpOutputReference
type jsiiProxy_PartnerAttachmentBgpOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference) AuthKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference) AuthKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference) InternalValue() *PartnerAttachmentBgp {
	var returns *PartnerAttachmentBgp
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference) LocalRouterIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localRouterIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference) LocalRouterIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localRouterIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference) PeerRouterAsn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peerRouterAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference) PeerRouterAsnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peerRouterAsnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference) PeerRouterIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerRouterIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference) PeerRouterIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerRouterIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPartnerAttachmentBgpOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PartnerAttachmentBgpOutputReference {
	_init_.Initialize()

	if err := validateNewPartnerAttachmentBgpOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PartnerAttachmentBgpOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.partnerAttachment.PartnerAttachmentBgpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPartnerAttachmentBgpOutputReference_Override(p PartnerAttachmentBgpOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.partnerAttachment.PartnerAttachmentBgpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference)SetAuthKey(val *string) {
	if err := j.validateSetAuthKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authKey",
		val,
	)
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference)SetInternalValue(val *PartnerAttachmentBgp) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference)SetLocalRouterIp(val *string) {
	if err := j.validateSetLocalRouterIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localRouterIp",
		val,
	)
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference)SetPeerRouterAsn(val *float64) {
	if err := j.validateSetPeerRouterAsnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerRouterAsn",
		val,
	)
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference)SetPeerRouterIp(val *string) {
	if err := j.validateSetPeerRouterIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerRouterIp",
		val,
	)
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PartnerAttachmentBgpOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PartnerAttachmentBgpOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PartnerAttachmentBgpOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PartnerAttachmentBgpOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PartnerAttachmentBgpOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PartnerAttachmentBgpOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PartnerAttachmentBgpOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PartnerAttachmentBgpOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PartnerAttachmentBgpOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PartnerAttachmentBgpOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PartnerAttachmentBgpOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PartnerAttachmentBgpOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PartnerAttachmentBgpOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PartnerAttachmentBgpOutputReference) ResetAuthKey() {
	_jsii_.InvokeVoid(
		p,
		"resetAuthKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PartnerAttachmentBgpOutputReference) ResetLocalRouterIp() {
	_jsii_.InvokeVoid(
		p,
		"resetLocalRouterIp",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PartnerAttachmentBgpOutputReference) ResetPeerRouterAsn() {
	_jsii_.InvokeVoid(
		p,
		"resetPeerRouterAsn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PartnerAttachmentBgpOutputReference) ResetPeerRouterIp() {
	_jsii_.InvokeVoid(
		p,
		"resetPeerRouterIp",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PartnerAttachmentBgpOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PartnerAttachmentBgpOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


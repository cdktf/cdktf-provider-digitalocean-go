// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiknowledgebase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/genaiknowledgebase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference interface {
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
	OriginalFileName() *string
	SetOriginalFileName(val *string)
	OriginalFileNameInput() *string
	SizeInBytes() *string
	SetSizeInBytes(val *string)
	SizeInBytesInput() *string
	StoredObjectKey() *string
	SetStoredObjectKey(val *string)
	StoredObjectKeyInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetOriginalFileName()
	ResetSizeInBytes()
	ResetStoredObjectKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference
type jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) OriginalFileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originalFileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) OriginalFileNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originalFileNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) SizeInBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) SizeInBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) StoredObjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storedObjectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) StoredObjectKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storedObjectKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference {
	_init_.Initialize()

	if err := validateNewGenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.genaiKnowledgeBase.GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference_Override(g GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.genaiKnowledgeBase.GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference)SetOriginalFileName(val *string) {
	if err := j.validateSetOriginalFileNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originalFileName",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference)SetSizeInBytes(val *string) {
	if err := j.validateSetSizeInBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeInBytes",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference)SetStoredObjectKey(val *string) {
	if err := j.validateSetStoredObjectKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storedObjectKey",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) ResetOriginalFileName() {
	_jsii_.InvokeVoid(
		g,
		"resetOriginalFileName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) ResetSizeInBytes() {
	_jsii_.InvokeVoid(
		g,
		"resetSizeInBytes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) ResetStoredObjectKey() {
	_jsii_.InvokeVoid(
		g,
		"resetStoredObjectKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesFileUploadDataSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


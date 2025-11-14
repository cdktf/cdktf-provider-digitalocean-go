// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiknowledgebase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/genaiknowledgebase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference interface {
	cdktf.ComplexObject
	CompletedDatasources() *float64
	SetCompletedDatasources(val *float64)
	CompletedDatasourcesInput() *float64
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
	CreatedAt() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataSourceUuids() *[]*string
	SetDataSourceUuids(val *[]*string)
	DataSourceUuidsInput() *[]*string
	FinishedAt() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KnowledgeBaseUuid() *string
	Phase() *string
	SetPhase(val *string)
	PhaseInput() *string
	StartedAt() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Tokens() *float64
	SetTokens(val *float64)
	TokensInput() *float64
	TotalDatasources() *float64
	SetTotalDatasources(val *float64)
	TotalDatasourcesInput() *float64
	UpdatedAt() *string
	Uuid() *string
	SetUuid(val *string)
	UuidInput() *string
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
	ResetCompletedDatasources()
	ResetDataSourceUuids()
	ResetPhase()
	ResetTokens()
	ResetTotalDatasources()
	ResetUuid()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference
type jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) CompletedDatasources() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"completedDatasources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) CompletedDatasourcesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"completedDatasourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) DataSourceUuids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataSourceUuids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) DataSourceUuidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataSourceUuidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) FinishedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finishedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) KnowledgeBaseUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) Phase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) PhaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) StartedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) Tokens() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) TokensInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) TotalDatasources() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalDatasources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) TotalDatasourcesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalDatasourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) UuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuidInput",
		&returns,
	)
	return returns
}


func NewGenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference {
	_init_.Initialize()

	if err := validateNewGenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.genaiKnowledgeBase.GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference_Override(g GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.genaiKnowledgeBase.GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference)SetCompletedDatasources(val *float64) {
	if err := j.validateSetCompletedDatasourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"completedDatasources",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference)SetDataSourceUuids(val *[]*string) {
	if err := j.validateSetDataSourceUuidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceUuids",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference)SetPhase(val *string) {
	if err := j.validateSetPhaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phase",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference)SetTokens(val *float64) {
	if err := j.validateSetTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokens",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference)SetTotalDatasources(val *float64) {
	if err := j.validateSetTotalDatasourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalDatasources",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference)SetUuid(val *string) {
	if err := j.validateSetUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uuid",
		val,
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) ResetCompletedDatasources() {
	_jsii_.InvokeVoid(
		g,
		"resetCompletedDatasources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) ResetDataSourceUuids() {
	_jsii_.InvokeVoid(
		g,
		"resetDataSourceUuids",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) ResetPhase() {
	_jsii_.InvokeVoid(
		g,
		"resetPhase",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) ResetTokens() {
	_jsii_.InvokeVoid(
		g,
		"resetTokens",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) ResetTotalDatasources() {
	_jsii_.InvokeVoid(
		g,
		"resetTotalDatasources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) ResetUuid() {
	_jsii_.InvokeVoid(
		g,
		"resetUuid",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesLastIndexingJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


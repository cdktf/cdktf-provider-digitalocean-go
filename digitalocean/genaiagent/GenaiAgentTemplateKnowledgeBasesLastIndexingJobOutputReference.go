// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/genaiagent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference interface {
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
	InternalValue() *GenaiAgentTemplateKnowledgeBasesLastIndexingJob
	SetInternalValue(val *GenaiAgentTemplateKnowledgeBasesLastIndexingJob)
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

// The jsii proxy struct for GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference
type jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) CompletedDatasources() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"completedDatasources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) CompletedDatasourcesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"completedDatasourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) DataSourceUuids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataSourceUuids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) DataSourceUuidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataSourceUuidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) FinishedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finishedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) InternalValue() *GenaiAgentTemplateKnowledgeBasesLastIndexingJob {
	var returns *GenaiAgentTemplateKnowledgeBasesLastIndexingJob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) KnowledgeBaseUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) Phase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) PhaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) StartedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) Tokens() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) TokensInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) TotalDatasources() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalDatasources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) TotalDatasourcesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalDatasourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) UuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuidInput",
		&returns,
	)
	return returns
}


func NewGenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference {
	_init_.Initialize()

	if err := validateNewGenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.genaiAgent.GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference_Override(g GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.genaiAgent.GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference)SetCompletedDatasources(val *float64) {
	if err := j.validateSetCompletedDatasourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"completedDatasources",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference)SetDataSourceUuids(val *[]*string) {
	if err := j.validateSetDataSourceUuidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceUuids",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference)SetInternalValue(val *GenaiAgentTemplateKnowledgeBasesLastIndexingJob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference)SetPhase(val *string) {
	if err := j.validateSetPhaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phase",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference)SetTokens(val *float64) {
	if err := j.validateSetTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokens",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference)SetTotalDatasources(val *float64) {
	if err := j.validateSetTotalDatasourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalDatasources",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference)SetUuid(val *string) {
	if err := j.validateSetUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uuid",
		val,
	)
}

func (g *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) ResetCompletedDatasources() {
	_jsii_.InvokeVoid(
		g,
		"resetCompletedDatasources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) ResetDataSourceUuids() {
	_jsii_.InvokeVoid(
		g,
		"resetDataSourceUuids",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) ResetPhase() {
	_jsii_.InvokeVoid(
		g,
		"resetPhase",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) ResetTokens() {
	_jsii_.InvokeVoid(
		g,
		"resetTokens",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) ResetTotalDatasources() {
	_jsii_.InvokeVoid(
		g,
		"resetTotalDatasources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) ResetUuid() {
	_jsii_.InvokeVoid(
		g,
		"resetUuid",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GenaiAgentTemplateKnowledgeBasesLastIndexingJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


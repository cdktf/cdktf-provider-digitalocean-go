// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiknowledgebase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/genaiknowledgebase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GenaiKnowledgeBaseDatasourcesOutputReference interface {
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
	CreatedAt() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FileUploadDataSource() GenaiKnowledgeBaseDatasourcesFileUploadDataSourceList
	FileUploadDataSourceInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LastIndexingJob() GenaiKnowledgeBaseDatasourcesLastIndexingJobList
	LastIndexingJobInput() interface{}
	SpacesDataSource() GenaiKnowledgeBaseDatasourcesSpacesDataSourceList
	SpacesDataSourceInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpdatedAt() *string
	Uuid() *string
	SetUuid(val *string)
	UuidInput() *string
	WebCrawlerDataSource() GenaiKnowledgeBaseDatasourcesWebCrawlerDataSourceList
	WebCrawlerDataSourceInput() interface{}
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
	PutFileUploadDataSource(value interface{})
	PutLastIndexingJob(value interface{})
	PutSpacesDataSource(value interface{})
	PutWebCrawlerDataSource(value interface{})
	ResetFileUploadDataSource()
	ResetLastIndexingJob()
	ResetSpacesDataSource()
	ResetUuid()
	ResetWebCrawlerDataSource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GenaiKnowledgeBaseDatasourcesOutputReference
type jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) FileUploadDataSource() GenaiKnowledgeBaseDatasourcesFileUploadDataSourceList {
	var returns GenaiKnowledgeBaseDatasourcesFileUploadDataSourceList
	_jsii_.Get(
		j,
		"fileUploadDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) FileUploadDataSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileUploadDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) LastIndexingJob() GenaiKnowledgeBaseDatasourcesLastIndexingJobList {
	var returns GenaiKnowledgeBaseDatasourcesLastIndexingJobList
	_jsii_.Get(
		j,
		"lastIndexingJob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) LastIndexingJobInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastIndexingJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) SpacesDataSource() GenaiKnowledgeBaseDatasourcesSpacesDataSourceList {
	var returns GenaiKnowledgeBaseDatasourcesSpacesDataSourceList
	_jsii_.Get(
		j,
		"spacesDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) SpacesDataSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spacesDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) UuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) WebCrawlerDataSource() GenaiKnowledgeBaseDatasourcesWebCrawlerDataSourceList {
	var returns GenaiKnowledgeBaseDatasourcesWebCrawlerDataSourceList
	_jsii_.Get(
		j,
		"webCrawlerDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) WebCrawlerDataSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webCrawlerDataSourceInput",
		&returns,
	)
	return returns
}


func NewGenaiKnowledgeBaseDatasourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GenaiKnowledgeBaseDatasourcesOutputReference {
	_init_.Initialize()

	if err := validateNewGenaiKnowledgeBaseDatasourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.genaiKnowledgeBase.GenaiKnowledgeBaseDatasourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGenaiKnowledgeBaseDatasourcesOutputReference_Override(g GenaiKnowledgeBaseDatasourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.genaiKnowledgeBase.GenaiKnowledgeBaseDatasourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference)SetUuid(val *string) {
	if err := j.validateSetUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uuid",
		val,
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) PutFileUploadDataSource(value interface{}) {
	if err := g.validatePutFileUploadDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFileUploadDataSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) PutLastIndexingJob(value interface{}) {
	if err := g.validatePutLastIndexingJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLastIndexingJob",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) PutSpacesDataSource(value interface{}) {
	if err := g.validatePutSpacesDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSpacesDataSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) PutWebCrawlerDataSource(value interface{}) {
	if err := g.validatePutWebCrawlerDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWebCrawlerDataSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) ResetFileUploadDataSource() {
	_jsii_.InvokeVoid(
		g,
		"resetFileUploadDataSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) ResetLastIndexingJob() {
	_jsii_.InvokeVoid(
		g,
		"resetLastIndexingJob",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) ResetSpacesDataSource() {
	_jsii_.InvokeVoid(
		g,
		"resetSpacesDataSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) ResetUuid() {
	_jsii_.InvokeVoid(
		g,
		"resetUuid",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) ResetWebCrawlerDataSource() {
	_jsii_.InvokeVoid(
		g,
		"resetWebCrawlerDataSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDatasourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


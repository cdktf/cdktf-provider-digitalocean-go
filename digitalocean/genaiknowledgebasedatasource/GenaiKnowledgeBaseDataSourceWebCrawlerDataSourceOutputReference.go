// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiknowledgebasedatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/genaiknowledgebasedatasource/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference interface {
	cdktf.ComplexObject
	BaseUrl() *string
	SetBaseUrl(val *string)
	BaseUrlInput() *string
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
	CrawlingOption() *string
	SetCrawlingOption(val *string)
	CrawlingOptionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EmbedMedia() interface{}
	SetEmbedMedia(val interface{})
	EmbedMediaInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GenaiKnowledgeBaseDataSourceWebCrawlerDataSource
	SetInternalValue(val *GenaiKnowledgeBaseDataSourceWebCrawlerDataSource)
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
	ResetBaseUrl()
	ResetCrawlingOption()
	ResetEmbedMedia()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference
type jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) BaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) BaseUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) CrawlingOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crawlingOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) CrawlingOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crawlingOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) EmbedMedia() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embedMedia",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) EmbedMediaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embedMediaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) InternalValue() *GenaiKnowledgeBaseDataSourceWebCrawlerDataSource {
	var returns *GenaiKnowledgeBaseDataSourceWebCrawlerDataSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference {
	_init_.Initialize()

	if err := validateNewGenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.genaiKnowledgeBaseDataSource.GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference_Override(g GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.genaiKnowledgeBaseDataSource.GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference)SetBaseUrl(val *string) {
	if err := j.validateSetBaseUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseUrl",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference)SetCrawlingOption(val *string) {
	if err := j.validateSetCrawlingOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crawlingOption",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference)SetEmbedMedia(val interface{}) {
	if err := j.validateSetEmbedMediaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"embedMedia",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference)SetInternalValue(val *GenaiKnowledgeBaseDataSourceWebCrawlerDataSource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) ResetBaseUrl() {
	_jsii_.InvokeVoid(
		g,
		"resetBaseUrl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) ResetCrawlingOption() {
	_jsii_.InvokeVoid(
		g,
		"resetCrawlingOption",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) ResetEmbedMedia() {
	_jsii_.InvokeVoid(
		g,
		"resetEmbedMedia",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiKnowledgeBaseDataSourceWebCrawlerDataSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


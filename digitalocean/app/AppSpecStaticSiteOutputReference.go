// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/app/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppSpecStaticSiteOutputReference interface {
	cdktf.ComplexObject
	Bitbucket() AppSpecStaticSiteBitbucketOutputReference
	BitbucketInput() *AppSpecStaticSiteBitbucket
	BuildCommand() *string
	SetBuildCommand(val *string)
	BuildCommandInput() *string
	CatchallDocument() *string
	SetCatchallDocument(val *string)
	CatchallDocumentInput() *string
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
	Cors() AppSpecStaticSiteCorsOutputReference
	CorsInput() *AppSpecStaticSiteCors
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DockerfilePath() *string
	SetDockerfilePath(val *string)
	DockerfilePathInput() *string
	Env() AppSpecStaticSiteEnvList
	EnvInput() interface{}
	EnvironmentSlug() *string
	SetEnvironmentSlug(val *string)
	EnvironmentSlugInput() *string
	ErrorDocument() *string
	SetErrorDocument(val *string)
	ErrorDocumentInput() *string
	// Experimental.
	Fqn() *string
	Git() AppSpecStaticSiteGitOutputReference
	Github() AppSpecStaticSiteGithubOutputReference
	GithubInput() *AppSpecStaticSiteGithub
	GitInput() *AppSpecStaticSiteGit
	Gitlab() AppSpecStaticSiteGitlabOutputReference
	GitlabInput() *AppSpecStaticSiteGitlab
	IndexDocument() *string
	SetIndexDocument(val *string)
	IndexDocumentInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	OutputDir() *string
	SetOutputDir(val *string)
	OutputDirInput() *string
	Routes() AppSpecStaticSiteRoutesList
	RoutesInput() interface{}
	SourceDir() *string
	SetSourceDir(val *string)
	SourceDirInput() *string
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
	PutBitbucket(value *AppSpecStaticSiteBitbucket)
	PutCors(value *AppSpecStaticSiteCors)
	PutEnv(value interface{})
	PutGit(value *AppSpecStaticSiteGit)
	PutGithub(value *AppSpecStaticSiteGithub)
	PutGitlab(value *AppSpecStaticSiteGitlab)
	PutRoutes(value interface{})
	ResetBitbucket()
	ResetBuildCommand()
	ResetCatchallDocument()
	ResetCors()
	ResetDockerfilePath()
	ResetEnv()
	ResetEnvironmentSlug()
	ResetErrorDocument()
	ResetGit()
	ResetGithub()
	ResetGitlab()
	ResetIndexDocument()
	ResetOutputDir()
	ResetRoutes()
	ResetSourceDir()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppSpecStaticSiteOutputReference
type jsiiProxy_AppSpecStaticSiteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) Bitbucket() AppSpecStaticSiteBitbucketOutputReference {
	var returns AppSpecStaticSiteBitbucketOutputReference
	_jsii_.Get(
		j,
		"bitbucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) BitbucketInput() *AppSpecStaticSiteBitbucket {
	var returns *AppSpecStaticSiteBitbucket
	_jsii_.Get(
		j,
		"bitbucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) BuildCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) BuildCommandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) CatchallDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catchallDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) CatchallDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catchallDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) Cors() AppSpecStaticSiteCorsOutputReference {
	var returns AppSpecStaticSiteCorsOutputReference
	_jsii_.Get(
		j,
		"cors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) CorsInput() *AppSpecStaticSiteCors {
	var returns *AppSpecStaticSiteCors
	_jsii_.Get(
		j,
		"corsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) DockerfilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) DockerfilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) Env() AppSpecStaticSiteEnvList {
	var returns AppSpecStaticSiteEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) EnvironmentSlug() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentSlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) EnvironmentSlugInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentSlugInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) ErrorDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) ErrorDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) Git() AppSpecStaticSiteGitOutputReference {
	var returns AppSpecStaticSiteGitOutputReference
	_jsii_.Get(
		j,
		"git",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) Github() AppSpecStaticSiteGithubOutputReference {
	var returns AppSpecStaticSiteGithubOutputReference
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) GithubInput() *AppSpecStaticSiteGithub {
	var returns *AppSpecStaticSiteGithub
	_jsii_.Get(
		j,
		"githubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) GitInput() *AppSpecStaticSiteGit {
	var returns *AppSpecStaticSiteGit
	_jsii_.Get(
		j,
		"gitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) Gitlab() AppSpecStaticSiteGitlabOutputReference {
	var returns AppSpecStaticSiteGitlabOutputReference
	_jsii_.Get(
		j,
		"gitlab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) GitlabInput() *AppSpecStaticSiteGitlab {
	var returns *AppSpecStaticSiteGitlab
	_jsii_.Get(
		j,
		"gitlabInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) IndexDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) IndexDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) OutputDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) OutputDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) Routes() AppSpecStaticSiteRoutesList {
	var returns AppSpecStaticSiteRoutesList
	_jsii_.Get(
		j,
		"routes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) RoutesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) SourceDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) SourceDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppSpecStaticSiteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppSpecStaticSiteOutputReference {
	_init_.Initialize()

	if err := validateNewAppSpecStaticSiteOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppSpecStaticSiteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.app.AppSpecStaticSiteOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppSpecStaticSiteOutputReference_Override(a AppSpecStaticSiteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.app.AppSpecStaticSiteOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference)SetBuildCommand(val *string) {
	if err := j.validateSetBuildCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildCommand",
		val,
	)
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference)SetCatchallDocument(val *string) {
	if err := j.validateSetCatchallDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catchallDocument",
		val,
	)
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference)SetDockerfilePath(val *string) {
	if err := j.validateSetDockerfilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerfilePath",
		val,
	)
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference)SetEnvironmentSlug(val *string) {
	if err := j.validateSetEnvironmentSlugParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentSlug",
		val,
	)
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference)SetErrorDocument(val *string) {
	if err := j.validateSetErrorDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorDocument",
		val,
	)
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference)SetIndexDocument(val *string) {
	if err := j.validateSetIndexDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexDocument",
		val,
	)
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference)SetOutputDir(val *string) {
	if err := j.validateSetOutputDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputDir",
		val,
	)
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference)SetSourceDir(val *string) {
	if err := j.validateSetSourceDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDir",
		val,
	)
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppSpecStaticSiteOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) PutBitbucket(value *AppSpecStaticSiteBitbucket) {
	if err := a.validatePutBitbucketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBitbucket",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) PutCors(value *AppSpecStaticSiteCors) {
	if err := a.validatePutCorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCors",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) PutEnv(value interface{}) {
	if err := a.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnv",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) PutGit(value *AppSpecStaticSiteGit) {
	if err := a.validatePutGitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGit",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) PutGithub(value *AppSpecStaticSiteGithub) {
	if err := a.validatePutGithubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGithub",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) PutGitlab(value *AppSpecStaticSiteGitlab) {
	if err := a.validatePutGitlabParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGitlab",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) PutRoutes(value interface{}) {
	if err := a.validatePutRoutesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRoutes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) ResetBitbucket() {
	_jsii_.InvokeVoid(
		a,
		"resetBitbucket",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) ResetBuildCommand() {
	_jsii_.InvokeVoid(
		a,
		"resetBuildCommand",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) ResetCatchallDocument() {
	_jsii_.InvokeVoid(
		a,
		"resetCatchallDocument",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) ResetCors() {
	_jsii_.InvokeVoid(
		a,
		"resetCors",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) ResetDockerfilePath() {
	_jsii_.InvokeVoid(
		a,
		"resetDockerfilePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		a,
		"resetEnv",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) ResetEnvironmentSlug() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironmentSlug",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) ResetErrorDocument() {
	_jsii_.InvokeVoid(
		a,
		"resetErrorDocument",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) ResetGit() {
	_jsii_.InvokeVoid(
		a,
		"resetGit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) ResetGithub() {
	_jsii_.InvokeVoid(
		a,
		"resetGithub",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) ResetGitlab() {
	_jsii_.InvokeVoid(
		a,
		"resetGitlab",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) ResetIndexDocument() {
	_jsii_.InvokeVoid(
		a,
		"resetIndexDocument",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) ResetOutputDir() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputDir",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) ResetRoutes() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) ResetSourceDir() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceDir",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecStaticSiteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


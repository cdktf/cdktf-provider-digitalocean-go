// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v11/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v11/app/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppSpecWorkerOutputReference interface {
	cdktf.ComplexObject
	Alert() AppSpecWorkerAlertList
	AlertInput() interface{}
	BuildCommand() *string
	SetBuildCommand(val *string)
	BuildCommandInput() *string
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
	DockerfilePath() *string
	SetDockerfilePath(val *string)
	DockerfilePathInput() *string
	Env() AppSpecWorkerEnvList
	EnvInput() interface{}
	EnvironmentSlug() *string
	SetEnvironmentSlug(val *string)
	EnvironmentSlugInput() *string
	// Experimental.
	Fqn() *string
	Git() AppSpecWorkerGitOutputReference
	Github() AppSpecWorkerGithubOutputReference
	GithubInput() *AppSpecWorkerGithub
	GitInput() *AppSpecWorkerGit
	Gitlab() AppSpecWorkerGitlabOutputReference
	GitlabInput() *AppSpecWorkerGitlab
	Image() AppSpecWorkerImageOutputReference
	ImageInput() *AppSpecWorkerImage
	InstanceCount() *float64
	SetInstanceCount(val *float64)
	InstanceCountInput() *float64
	InstanceSizeSlug() *string
	SetInstanceSizeSlug(val *string)
	InstanceSizeSlugInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogDestination() AppSpecWorkerLogDestinationList
	LogDestinationInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	RunCommand() *string
	SetRunCommand(val *string)
	RunCommandInput() *string
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
	PutAlert(value interface{})
	PutEnv(value interface{})
	PutGit(value *AppSpecWorkerGit)
	PutGithub(value *AppSpecWorkerGithub)
	PutGitlab(value *AppSpecWorkerGitlab)
	PutImage(value *AppSpecWorkerImage)
	PutLogDestination(value interface{})
	ResetAlert()
	ResetBuildCommand()
	ResetDockerfilePath()
	ResetEnv()
	ResetEnvironmentSlug()
	ResetGit()
	ResetGithub()
	ResetGitlab()
	ResetImage()
	ResetInstanceCount()
	ResetInstanceSizeSlug()
	ResetLogDestination()
	ResetRunCommand()
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

// The jsii proxy struct for AppSpecWorkerOutputReference
type jsiiProxy_AppSpecWorkerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) Alert() AppSpecWorkerAlertList {
	var returns AppSpecWorkerAlertList
	_jsii_.Get(
		j,
		"alert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) AlertInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) BuildCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) BuildCommandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) DockerfilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) DockerfilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) Env() AppSpecWorkerEnvList {
	var returns AppSpecWorkerEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) EnvironmentSlug() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentSlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) EnvironmentSlugInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentSlugInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) Git() AppSpecWorkerGitOutputReference {
	var returns AppSpecWorkerGitOutputReference
	_jsii_.Get(
		j,
		"git",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) Github() AppSpecWorkerGithubOutputReference {
	var returns AppSpecWorkerGithubOutputReference
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) GithubInput() *AppSpecWorkerGithub {
	var returns *AppSpecWorkerGithub
	_jsii_.Get(
		j,
		"githubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) GitInput() *AppSpecWorkerGit {
	var returns *AppSpecWorkerGit
	_jsii_.Get(
		j,
		"gitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) Gitlab() AppSpecWorkerGitlabOutputReference {
	var returns AppSpecWorkerGitlabOutputReference
	_jsii_.Get(
		j,
		"gitlab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) GitlabInput() *AppSpecWorkerGitlab {
	var returns *AppSpecWorkerGitlab
	_jsii_.Get(
		j,
		"gitlabInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) Image() AppSpecWorkerImageOutputReference {
	var returns AppSpecWorkerImageOutputReference
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) ImageInput() *AppSpecWorkerImage {
	var returns *AppSpecWorkerImage
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) InstanceSizeSlug() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceSizeSlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) InstanceSizeSlugInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceSizeSlugInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) LogDestination() AppSpecWorkerLogDestinationList {
	var returns AppSpecWorkerLogDestinationList
	_jsii_.Get(
		j,
		"logDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) LogDestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) RunCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) RunCommandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) SourceDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) SourceDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecWorkerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppSpecWorkerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppSpecWorkerOutputReference {
	_init_.Initialize()

	if err := validateNewAppSpecWorkerOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppSpecWorkerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.app.AppSpecWorkerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppSpecWorkerOutputReference_Override(a AppSpecWorkerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.app.AppSpecWorkerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppSpecWorkerOutputReference)SetBuildCommand(val *string) {
	if err := j.validateSetBuildCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildCommand",
		val,
	)
}

func (j *jsiiProxy_AppSpecWorkerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppSpecWorkerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppSpecWorkerOutputReference)SetDockerfilePath(val *string) {
	if err := j.validateSetDockerfilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerfilePath",
		val,
	)
}

func (j *jsiiProxy_AppSpecWorkerOutputReference)SetEnvironmentSlug(val *string) {
	if err := j.validateSetEnvironmentSlugParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentSlug",
		val,
	)
}

func (j *jsiiProxy_AppSpecWorkerOutputReference)SetInstanceCount(val *float64) {
	if err := j.validateSetInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_AppSpecWorkerOutputReference)SetInstanceSizeSlug(val *string) {
	if err := j.validateSetInstanceSizeSlugParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceSizeSlug",
		val,
	)
}

func (j *jsiiProxy_AppSpecWorkerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppSpecWorkerOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppSpecWorkerOutputReference)SetRunCommand(val *string) {
	if err := j.validateSetRunCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runCommand",
		val,
	)
}

func (j *jsiiProxy_AppSpecWorkerOutputReference)SetSourceDir(val *string) {
	if err := j.validateSetSourceDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDir",
		val,
	)
}

func (j *jsiiProxy_AppSpecWorkerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppSpecWorkerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppSpecWorkerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppSpecWorkerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppSpecWorkerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppSpecWorkerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppSpecWorkerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppSpecWorkerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppSpecWorkerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppSpecWorkerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppSpecWorkerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppSpecWorkerOutputReference) PutAlert(value interface{}) {
	if err := a.validatePutAlertParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAlert",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) PutEnv(value interface{}) {
	if err := a.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnv",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) PutGit(value *AppSpecWorkerGit) {
	if err := a.validatePutGitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGit",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) PutGithub(value *AppSpecWorkerGithub) {
	if err := a.validatePutGithubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGithub",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) PutGitlab(value *AppSpecWorkerGitlab) {
	if err := a.validatePutGitlabParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGitlab",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) PutImage(value *AppSpecWorkerImage) {
	if err := a.validatePutImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putImage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) PutLogDestination(value interface{}) {
	if err := a.validatePutLogDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogDestination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) ResetAlert() {
	_jsii_.InvokeVoid(
		a,
		"resetAlert",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) ResetBuildCommand() {
	_jsii_.InvokeVoid(
		a,
		"resetBuildCommand",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) ResetDockerfilePath() {
	_jsii_.InvokeVoid(
		a,
		"resetDockerfilePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		a,
		"resetEnv",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) ResetEnvironmentSlug() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironmentSlug",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) ResetGit() {
	_jsii_.InvokeVoid(
		a,
		"resetGit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) ResetGithub() {
	_jsii_.InvokeVoid(
		a,
		"resetGithub",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) ResetGitlab() {
	_jsii_.InvokeVoid(
		a,
		"resetGitlab",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		a,
		"resetImage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) ResetInstanceCount() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) ResetInstanceSizeSlug() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceSizeSlug",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) ResetLogDestination() {
	_jsii_.InvokeVoid(
		a,
		"resetLogDestination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) ResetRunCommand() {
	_jsii_.InvokeVoid(
		a,
		"resetRunCommand",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) ResetSourceDir() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceDir",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecWorkerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppSpecWorkerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


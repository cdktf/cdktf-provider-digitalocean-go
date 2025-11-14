// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/app/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppSpecServiceOutputReference interface {
	cdktf.ComplexObject
	Alert() AppSpecServiceAlertList
	AlertInput() interface{}
	Autoscaling() AppSpecServiceAutoscalingOutputReference
	AutoscalingInput() *AppSpecServiceAutoscaling
	Bitbucket() AppSpecServiceBitbucketOutputReference
	BitbucketInput() *AppSpecServiceBitbucket
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
	Cors() AppSpecServiceCorsOutputReference
	CorsInput() *AppSpecServiceCors
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DockerfilePath() *string
	SetDockerfilePath(val *string)
	DockerfilePathInput() *string
	Env() AppSpecServiceEnvList
	EnvInput() interface{}
	EnvironmentSlug() *string
	SetEnvironmentSlug(val *string)
	EnvironmentSlugInput() *string
	// Experimental.
	Fqn() *string
	Git() AppSpecServiceGitOutputReference
	Github() AppSpecServiceGithubOutputReference
	GithubInput() *AppSpecServiceGithub
	GitInput() *AppSpecServiceGit
	Gitlab() AppSpecServiceGitlabOutputReference
	GitlabInput() *AppSpecServiceGitlab
	HealthCheck() AppSpecServiceHealthCheckOutputReference
	HealthCheckInput() *AppSpecServiceHealthCheck
	HttpPort() *float64
	SetHttpPort(val *float64)
	HttpPortInput() *float64
	Image() AppSpecServiceImageOutputReference
	ImageInput() *AppSpecServiceImage
	InstanceCount() *float64
	SetInstanceCount(val *float64)
	InstanceCountInput() *float64
	InstanceSizeSlug() *string
	SetInstanceSizeSlug(val *string)
	InstanceSizeSlugInput() *string
	InternalPorts() *[]*float64
	SetInternalPorts(val *[]*float64)
	InternalPortsInput() *[]*float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogDestination() AppSpecServiceLogDestinationList
	LogDestinationInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Routes() AppSpecServiceRoutesList
	RoutesInput() interface{}
	RunCommand() *string
	SetRunCommand(val *string)
	RunCommandInput() *string
	SourceDir() *string
	SetSourceDir(val *string)
	SourceDirInput() *string
	Termination() AppSpecServiceTerminationOutputReference
	TerminationInput() *AppSpecServiceTermination
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
	PutAlert(value interface{})
	PutAutoscaling(value *AppSpecServiceAutoscaling)
	PutBitbucket(value *AppSpecServiceBitbucket)
	PutCors(value *AppSpecServiceCors)
	PutEnv(value interface{})
	PutGit(value *AppSpecServiceGit)
	PutGithub(value *AppSpecServiceGithub)
	PutGitlab(value *AppSpecServiceGitlab)
	PutHealthCheck(value *AppSpecServiceHealthCheck)
	PutImage(value *AppSpecServiceImage)
	PutLogDestination(value interface{})
	PutRoutes(value interface{})
	PutTermination(value *AppSpecServiceTermination)
	ResetAlert()
	ResetAutoscaling()
	ResetBitbucket()
	ResetBuildCommand()
	ResetCors()
	ResetDockerfilePath()
	ResetEnv()
	ResetEnvironmentSlug()
	ResetGit()
	ResetGithub()
	ResetGitlab()
	ResetHealthCheck()
	ResetHttpPort()
	ResetImage()
	ResetInstanceCount()
	ResetInstanceSizeSlug()
	ResetInternalPorts()
	ResetLogDestination()
	ResetRoutes()
	ResetRunCommand()
	ResetSourceDir()
	ResetTermination()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppSpecServiceOutputReference
type jsiiProxy_AppSpecServiceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppSpecServiceOutputReference) Alert() AppSpecServiceAlertList {
	var returns AppSpecServiceAlertList
	_jsii_.Get(
		j,
		"alert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) AlertInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) Autoscaling() AppSpecServiceAutoscalingOutputReference {
	var returns AppSpecServiceAutoscalingOutputReference
	_jsii_.Get(
		j,
		"autoscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) AutoscalingInput() *AppSpecServiceAutoscaling {
	var returns *AppSpecServiceAutoscaling
	_jsii_.Get(
		j,
		"autoscalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) Bitbucket() AppSpecServiceBitbucketOutputReference {
	var returns AppSpecServiceBitbucketOutputReference
	_jsii_.Get(
		j,
		"bitbucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) BitbucketInput() *AppSpecServiceBitbucket {
	var returns *AppSpecServiceBitbucket
	_jsii_.Get(
		j,
		"bitbucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) BuildCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) BuildCommandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) Cors() AppSpecServiceCorsOutputReference {
	var returns AppSpecServiceCorsOutputReference
	_jsii_.Get(
		j,
		"cors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) CorsInput() *AppSpecServiceCors {
	var returns *AppSpecServiceCors
	_jsii_.Get(
		j,
		"corsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) DockerfilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) DockerfilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) Env() AppSpecServiceEnvList {
	var returns AppSpecServiceEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) EnvironmentSlug() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentSlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) EnvironmentSlugInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentSlugInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) Git() AppSpecServiceGitOutputReference {
	var returns AppSpecServiceGitOutputReference
	_jsii_.Get(
		j,
		"git",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) Github() AppSpecServiceGithubOutputReference {
	var returns AppSpecServiceGithubOutputReference
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) GithubInput() *AppSpecServiceGithub {
	var returns *AppSpecServiceGithub
	_jsii_.Get(
		j,
		"githubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) GitInput() *AppSpecServiceGit {
	var returns *AppSpecServiceGit
	_jsii_.Get(
		j,
		"gitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) Gitlab() AppSpecServiceGitlabOutputReference {
	var returns AppSpecServiceGitlabOutputReference
	_jsii_.Get(
		j,
		"gitlab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) GitlabInput() *AppSpecServiceGitlab {
	var returns *AppSpecServiceGitlab
	_jsii_.Get(
		j,
		"gitlabInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) HealthCheck() AppSpecServiceHealthCheckOutputReference {
	var returns AppSpecServiceHealthCheckOutputReference
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) HealthCheckInput() *AppSpecServiceHealthCheck {
	var returns *AppSpecServiceHealthCheck
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) HttpPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) HttpPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) Image() AppSpecServiceImageOutputReference {
	var returns AppSpecServiceImageOutputReference
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) ImageInput() *AppSpecServiceImage {
	var returns *AppSpecServiceImage
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) InstanceSizeSlug() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceSizeSlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) InstanceSizeSlugInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceSizeSlugInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) InternalPorts() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"internalPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) InternalPortsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"internalPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) LogDestination() AppSpecServiceLogDestinationList {
	var returns AppSpecServiceLogDestinationList
	_jsii_.Get(
		j,
		"logDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) LogDestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) Routes() AppSpecServiceRoutesList {
	var returns AppSpecServiceRoutesList
	_jsii_.Get(
		j,
		"routes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) RoutesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) RunCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) RunCommandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) SourceDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) SourceDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) Termination() AppSpecServiceTerminationOutputReference {
	var returns AppSpecServiceTerminationOutputReference
	_jsii_.Get(
		j,
		"termination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) TerminationInput() *AppSpecServiceTermination {
	var returns *AppSpecServiceTermination
	_jsii_.Get(
		j,
		"terminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecServiceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppSpecServiceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppSpecServiceOutputReference {
	_init_.Initialize()

	if err := validateNewAppSpecServiceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppSpecServiceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.app.AppSpecServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppSpecServiceOutputReference_Override(a AppSpecServiceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.app.AppSpecServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppSpecServiceOutputReference)SetBuildCommand(val *string) {
	if err := j.validateSetBuildCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildCommand",
		val,
	)
}

func (j *jsiiProxy_AppSpecServiceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppSpecServiceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppSpecServiceOutputReference)SetDockerfilePath(val *string) {
	if err := j.validateSetDockerfilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerfilePath",
		val,
	)
}

func (j *jsiiProxy_AppSpecServiceOutputReference)SetEnvironmentSlug(val *string) {
	if err := j.validateSetEnvironmentSlugParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentSlug",
		val,
	)
}

func (j *jsiiProxy_AppSpecServiceOutputReference)SetHttpPort(val *float64) {
	if err := j.validateSetHttpPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpPort",
		val,
	)
}

func (j *jsiiProxy_AppSpecServiceOutputReference)SetInstanceCount(val *float64) {
	if err := j.validateSetInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_AppSpecServiceOutputReference)SetInstanceSizeSlug(val *string) {
	if err := j.validateSetInstanceSizeSlugParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceSizeSlug",
		val,
	)
}

func (j *jsiiProxy_AppSpecServiceOutputReference)SetInternalPorts(val *[]*float64) {
	if err := j.validateSetInternalPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalPorts",
		val,
	)
}

func (j *jsiiProxy_AppSpecServiceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppSpecServiceOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppSpecServiceOutputReference)SetRunCommand(val *string) {
	if err := j.validateSetRunCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runCommand",
		val,
	)
}

func (j *jsiiProxy_AppSpecServiceOutputReference)SetSourceDir(val *string) {
	if err := j.validateSetSourceDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDir",
		val,
	)
}

func (j *jsiiProxy_AppSpecServiceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppSpecServiceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecServiceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppSpecServiceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppSpecServiceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppSpecServiceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppSpecServiceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppSpecServiceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppSpecServiceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppSpecServiceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppSpecServiceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppSpecServiceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecServiceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecServiceOutputReference) PutAlert(value interface{}) {
	if err := a.validatePutAlertParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAlert",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) PutAutoscaling(value *AppSpecServiceAutoscaling) {
	if err := a.validatePutAutoscalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutoscaling",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) PutBitbucket(value *AppSpecServiceBitbucket) {
	if err := a.validatePutBitbucketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBitbucket",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) PutCors(value *AppSpecServiceCors) {
	if err := a.validatePutCorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCors",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) PutEnv(value interface{}) {
	if err := a.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnv",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) PutGit(value *AppSpecServiceGit) {
	if err := a.validatePutGitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGit",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) PutGithub(value *AppSpecServiceGithub) {
	if err := a.validatePutGithubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGithub",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) PutGitlab(value *AppSpecServiceGitlab) {
	if err := a.validatePutGitlabParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGitlab",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) PutHealthCheck(value *AppSpecServiceHealthCheck) {
	if err := a.validatePutHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHealthCheck",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) PutImage(value *AppSpecServiceImage) {
	if err := a.validatePutImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putImage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) PutLogDestination(value interface{}) {
	if err := a.validatePutLogDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogDestination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) PutRoutes(value interface{}) {
	if err := a.validatePutRoutesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRoutes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) PutTermination(value *AppSpecServiceTermination) {
	if err := a.validatePutTerminationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTermination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetAlert() {
	_jsii_.InvokeVoid(
		a,
		"resetAlert",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetAutoscaling() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoscaling",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetBitbucket() {
	_jsii_.InvokeVoid(
		a,
		"resetBitbucket",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetBuildCommand() {
	_jsii_.InvokeVoid(
		a,
		"resetBuildCommand",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetCors() {
	_jsii_.InvokeVoid(
		a,
		"resetCors",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetDockerfilePath() {
	_jsii_.InvokeVoid(
		a,
		"resetDockerfilePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		a,
		"resetEnv",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetEnvironmentSlug() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironmentSlug",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetGit() {
	_jsii_.InvokeVoid(
		a,
		"resetGit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetGithub() {
	_jsii_.InvokeVoid(
		a,
		"resetGithub",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetGitlab() {
	_jsii_.InvokeVoid(
		a,
		"resetGitlab",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetHttpPort() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		a,
		"resetImage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetInstanceCount() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetInstanceSizeSlug() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceSizeSlug",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetInternalPorts() {
	_jsii_.InvokeVoid(
		a,
		"resetInternalPorts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetLogDestination() {
	_jsii_.InvokeVoid(
		a,
		"resetLogDestination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetRoutes() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetRunCommand() {
	_jsii_.InvokeVoid(
		a,
		"resetRunCommand",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetSourceDir() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceDir",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ResetTermination() {
	_jsii_.InvokeVoid(
		a,
		"resetTermination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecServiceOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecServiceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


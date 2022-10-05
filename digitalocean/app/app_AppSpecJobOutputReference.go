package app

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v3/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v3/app/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppSpecJobOutputReference interface {
	cdktf.ComplexObject
	Alert() AppSpecJobAlertList
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
	Env() AppSpecJobEnvList
	EnvInput() interface{}
	EnvironmentSlug() *string
	SetEnvironmentSlug(val *string)
	EnvironmentSlugInput() *string
	// Experimental.
	Fqn() *string
	Git() AppSpecJobGitOutputReference
	Github() AppSpecJobGithubOutputReference
	GithubInput() *AppSpecJobGithub
	GitInput() *AppSpecJobGit
	Gitlab() AppSpecJobGitlabOutputReference
	GitlabInput() *AppSpecJobGitlab
	Image() AppSpecJobImageOutputReference
	ImageInput() *AppSpecJobImage
	InstanceCount() *float64
	SetInstanceCount(val *float64)
	InstanceCountInput() *float64
	InstanceSizeSlug() *string
	SetInstanceSizeSlug(val *string)
	InstanceSizeSlugInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Kind() *string
	SetKind(val *string)
	KindInput() *string
	LogDestination() AppSpecJobLogDestinationList
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
	PutGit(value *AppSpecJobGit)
	PutGithub(value *AppSpecJobGithub)
	PutGitlab(value *AppSpecJobGitlab)
	PutImage(value *AppSpecJobImage)
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
	ResetKind()
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

// The jsii proxy struct for AppSpecJobOutputReference
type jsiiProxy_AppSpecJobOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppSpecJobOutputReference) Alert() AppSpecJobAlertList {
	var returns AppSpecJobAlertList
	_jsii_.Get(
		j,
		"alert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) AlertInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) BuildCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) BuildCommandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) DockerfilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) DockerfilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) Env() AppSpecJobEnvList {
	var returns AppSpecJobEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) EnvironmentSlug() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentSlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) EnvironmentSlugInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentSlugInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) Git() AppSpecJobGitOutputReference {
	var returns AppSpecJobGitOutputReference
	_jsii_.Get(
		j,
		"git",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) Github() AppSpecJobGithubOutputReference {
	var returns AppSpecJobGithubOutputReference
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) GithubInput() *AppSpecJobGithub {
	var returns *AppSpecJobGithub
	_jsii_.Get(
		j,
		"githubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) GitInput() *AppSpecJobGit {
	var returns *AppSpecJobGit
	_jsii_.Get(
		j,
		"gitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) Gitlab() AppSpecJobGitlabOutputReference {
	var returns AppSpecJobGitlabOutputReference
	_jsii_.Get(
		j,
		"gitlab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) GitlabInput() *AppSpecJobGitlab {
	var returns *AppSpecJobGitlab
	_jsii_.Get(
		j,
		"gitlabInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) Image() AppSpecJobImageOutputReference {
	var returns AppSpecJobImageOutputReference
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) ImageInput() *AppSpecJobImage {
	var returns *AppSpecJobImage
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) InstanceSizeSlug() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceSizeSlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) InstanceSizeSlugInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceSizeSlugInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) KindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) LogDestination() AppSpecJobLogDestinationList {
	var returns AppSpecJobLogDestinationList
	_jsii_.Get(
		j,
		"logDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) LogDestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) RunCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) RunCommandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) SourceDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) SourceDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppSpecJobOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppSpecJobOutputReference {
	_init_.Initialize()

	if err := validateNewAppSpecJobOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppSpecJobOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.app.AppSpecJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppSpecJobOutputReference_Override(a AppSpecJobOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.app.AppSpecJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppSpecJobOutputReference)SetBuildCommand(val *string) {
	if err := j.validateSetBuildCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildCommand",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobOutputReference)SetDockerfilePath(val *string) {
	if err := j.validateSetDockerfilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerfilePath",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobOutputReference)SetEnvironmentSlug(val *string) {
	if err := j.validateSetEnvironmentSlugParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentSlug",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobOutputReference)SetInstanceCount(val *float64) {
	if err := j.validateSetInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobOutputReference)SetInstanceSizeSlug(val *string) {
	if err := j.validateSetInstanceSizeSlugParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceSizeSlug",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobOutputReference)SetKind(val *string) {
	if err := j.validateSetKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobOutputReference)SetRunCommand(val *string) {
	if err := j.validateSetRunCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runCommand",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobOutputReference)SetSourceDir(val *string) {
	if err := j.validateSetSourceDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDir",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppSpecJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppSpecJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppSpecJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppSpecJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppSpecJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppSpecJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppSpecJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppSpecJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppSpecJobOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecJobOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppSpecJobOutputReference) PutAlert(value interface{}) {
	if err := a.validatePutAlertParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAlert",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) PutEnv(value interface{}) {
	if err := a.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnv",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) PutGit(value *AppSpecJobGit) {
	if err := a.validatePutGitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGit",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) PutGithub(value *AppSpecJobGithub) {
	if err := a.validatePutGithubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGithub",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) PutGitlab(value *AppSpecJobGitlab) {
	if err := a.validatePutGitlabParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGitlab",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) PutImage(value *AppSpecJobImage) {
	if err := a.validatePutImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putImage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) PutLogDestination(value interface{}) {
	if err := a.validatePutLogDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogDestination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) ResetAlert() {
	_jsii_.InvokeVoid(
		a,
		"resetAlert",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) ResetBuildCommand() {
	_jsii_.InvokeVoid(
		a,
		"resetBuildCommand",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) ResetDockerfilePath() {
	_jsii_.InvokeVoid(
		a,
		"resetDockerfilePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		a,
		"resetEnv",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) ResetEnvironmentSlug() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironmentSlug",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) ResetGit() {
	_jsii_.InvokeVoid(
		a,
		"resetGit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) ResetGithub() {
	_jsii_.InvokeVoid(
		a,
		"resetGithub",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) ResetGitlab() {
	_jsii_.InvokeVoid(
		a,
		"resetGitlab",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		a,
		"resetImage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) ResetInstanceCount() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) ResetInstanceSizeSlug() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceSizeSlug",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) ResetKind() {
	_jsii_.InvokeVoid(
		a,
		"resetKind",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) ResetLogDestination() {
	_jsii_.InvokeVoid(
		a,
		"resetLogDestination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) ResetRunCommand() {
	_jsii_.InvokeVoid(
		a,
		"resetRunCommand",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) ResetSourceDir() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceDir",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecJobOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppSpecJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


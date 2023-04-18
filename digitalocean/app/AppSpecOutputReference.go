package app

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v6/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v6/app/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppSpecOutputReference interface {
	cdktf.ComplexObject
	Alert() AppSpecAlertList
	AlertInput() interface{}
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
	Database() AppSpecDatabaseList
	DatabaseInput() interface{}
	Domain() AppSpecDomainList
	DomainInput() interface{}
	Domains() *[]*string
	SetDomains(val *[]*string)
	DomainsInput() *[]*string
	Env() AppSpecEnvList
	EnvInput() interface{}
	// Experimental.
	Fqn() *string
	Function() AppSpecFunctionList
	FunctionInput() interface{}
	InternalValue() *AppSpec
	SetInternalValue(val *AppSpec)
	Job() AppSpecJobList
	JobInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	Service() AppSpecServiceList
	ServiceInput() interface{}
	StaticSite() AppSpecStaticSiteList
	StaticSiteInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Worker() AppSpecWorkerList
	WorkerInput() interface{}
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
	PutDatabase(value interface{})
	PutDomain(value interface{})
	PutEnv(value interface{})
	PutFunction(value interface{})
	PutJob(value interface{})
	PutService(value interface{})
	PutStaticSite(value interface{})
	PutWorker(value interface{})
	ResetAlert()
	ResetDatabase()
	ResetDomain()
	ResetDomains()
	ResetEnv()
	ResetFunction()
	ResetJob()
	ResetRegion()
	ResetService()
	ResetStaticSite()
	ResetWorker()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppSpecOutputReference
type jsiiProxy_AppSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppSpecOutputReference) Alert() AppSpecAlertList {
	var returns AppSpecAlertList
	_jsii_.Get(
		j,
		"alert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) AlertInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) Database() AppSpecDatabaseList {
	var returns AppSpecDatabaseList
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) DatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) Domain() AppSpecDomainList {
	var returns AppSpecDomainList
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) DomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) Domains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) DomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) Env() AppSpecEnvList {
	var returns AppSpecEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) Function() AppSpecFunctionList {
	var returns AppSpecFunctionList
	_jsii_.Get(
		j,
		"function",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) FunctionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) InternalValue() *AppSpec {
	var returns *AppSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) Job() AppSpecJobList {
	var returns AppSpecJobList
	_jsii_.Get(
		j,
		"job",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) JobInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) Service() AppSpecServiceList {
	var returns AppSpecServiceList
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) ServiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) StaticSite() AppSpecStaticSiteList {
	var returns AppSpecStaticSiteList
	_jsii_.Get(
		j,
		"staticSite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) StaticSiteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticSiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) Worker() AppSpecWorkerList {
	var returns AppSpecWorkerList
	_jsii_.Get(
		j,
		"worker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecOutputReference) WorkerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workerInput",
		&returns,
	)
	return returns
}


func NewAppSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppSpecOutputReference {
	_init_.Initialize()

	if err := validateNewAppSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.app.AppSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppSpecOutputReference_Override(a AppSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.app.AppSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppSpecOutputReference)SetDomains(val *[]*string) {
	if err := j.validateSetDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domains",
		val,
	)
}

func (j *jsiiProxy_AppSpecOutputReference)SetInternalValue(val *AppSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppSpecOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppSpecOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AppSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppSpecOutputReference) PutAlert(value interface{}) {
	if err := a.validatePutAlertParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAlert",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecOutputReference) PutDatabase(value interface{}) {
	if err := a.validatePutDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatabase",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecOutputReference) PutDomain(value interface{}) {
	if err := a.validatePutDomainParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDomain",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecOutputReference) PutEnv(value interface{}) {
	if err := a.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnv",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecOutputReference) PutFunction(value interface{}) {
	if err := a.validatePutFunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFunction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecOutputReference) PutJob(value interface{}) {
	if err := a.validatePutJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJob",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecOutputReference) PutService(value interface{}) {
	if err := a.validatePutServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putService",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecOutputReference) PutStaticSite(value interface{}) {
	if err := a.validatePutStaticSiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStaticSite",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecOutputReference) PutWorker(value interface{}) {
	if err := a.validatePutWorkerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWorker",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecOutputReference) ResetAlert() {
	_jsii_.InvokeVoid(
		a,
		"resetAlert",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		a,
		"resetDatabase",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecOutputReference) ResetDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecOutputReference) ResetDomains() {
	_jsii_.InvokeVoid(
		a,
		"resetDomains",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		a,
		"resetEnv",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecOutputReference) ResetFunction() {
	_jsii_.InvokeVoid(
		a,
		"resetFunction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecOutputReference) ResetJob() {
	_jsii_.InvokeVoid(
		a,
		"resetJob",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecOutputReference) ResetService() {
	_jsii_.InvokeVoid(
		a,
		"resetService",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecOutputReference) ResetStaticSite() {
	_jsii_.InvokeVoid(
		a,
		"resetStaticSite",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecOutputReference) ResetWorker() {
	_jsii_.InvokeVoid(
		a,
		"resetWorker",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


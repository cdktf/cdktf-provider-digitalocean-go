package datadigitaloceanapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v4/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v4/datadigitaloceanapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDigitaloceanAppSpecServiceOutputReference interface {
	cdktf.ComplexObject
	Alert() DataDigitaloceanAppSpecServiceAlertList
	BuildCommand() *string
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
	Cors() DataDigitaloceanAppSpecServiceCorsList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DockerfilePath() *string
	Env() DataDigitaloceanAppSpecServiceEnvList
	EnvironmentSlug() *string
	// Experimental.
	Fqn() *string
	Git() DataDigitaloceanAppSpecServiceGitList
	Github() DataDigitaloceanAppSpecServiceGithubList
	Gitlab() DataDigitaloceanAppSpecServiceGitlabList
	HealthCheck() DataDigitaloceanAppSpecServiceHealthCheckList
	HttpPort() *float64
	Image() DataDigitaloceanAppSpecServiceImageList
	InstanceCount() *float64
	InstanceSizeSlug() *string
	InternalPorts() *[]*float64
	InternalValue() *DataDigitaloceanAppSpecService
	SetInternalValue(val *DataDigitaloceanAppSpecService)
	LogDestination() DataDigitaloceanAppSpecServiceLogDestinationList
	Name() *string
	Routes() DataDigitaloceanAppSpecServiceRoutesList
	RunCommand() *string
	SourceDir() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDigitaloceanAppSpecServiceOutputReference
type jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) Alert() DataDigitaloceanAppSpecServiceAlertList {
	var returns DataDigitaloceanAppSpecServiceAlertList
	_jsii_.Get(
		j,
		"alert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) BuildCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) Cors() DataDigitaloceanAppSpecServiceCorsList {
	var returns DataDigitaloceanAppSpecServiceCorsList
	_jsii_.Get(
		j,
		"cors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) DockerfilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) Env() DataDigitaloceanAppSpecServiceEnvList {
	var returns DataDigitaloceanAppSpecServiceEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) EnvironmentSlug() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentSlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) Git() DataDigitaloceanAppSpecServiceGitList {
	var returns DataDigitaloceanAppSpecServiceGitList
	_jsii_.Get(
		j,
		"git",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) Github() DataDigitaloceanAppSpecServiceGithubList {
	var returns DataDigitaloceanAppSpecServiceGithubList
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) Gitlab() DataDigitaloceanAppSpecServiceGitlabList {
	var returns DataDigitaloceanAppSpecServiceGitlabList
	_jsii_.Get(
		j,
		"gitlab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) HealthCheck() DataDigitaloceanAppSpecServiceHealthCheckList {
	var returns DataDigitaloceanAppSpecServiceHealthCheckList
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) HttpPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) Image() DataDigitaloceanAppSpecServiceImageList {
	var returns DataDigitaloceanAppSpecServiceImageList
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) InstanceSizeSlug() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceSizeSlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) InternalPorts() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"internalPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) InternalValue() *DataDigitaloceanAppSpecService {
	var returns *DataDigitaloceanAppSpecService
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) LogDestination() DataDigitaloceanAppSpecServiceLogDestinationList {
	var returns DataDigitaloceanAppSpecServiceLogDestinationList
	_jsii_.Get(
		j,
		"logDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) Routes() DataDigitaloceanAppSpecServiceRoutesList {
	var returns DataDigitaloceanAppSpecServiceRoutesList
	_jsii_.Get(
		j,
		"routes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) RunCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) SourceDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDigitaloceanAppSpecServiceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDigitaloceanAppSpecServiceOutputReference {
	_init_.Initialize()

	if err := validateNewDataDigitaloceanAppSpecServiceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.dataDigitaloceanApp.DataDigitaloceanAppSpecServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDigitaloceanAppSpecServiceOutputReference_Override(d DataDigitaloceanAppSpecServiceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.dataDigitaloceanApp.DataDigitaloceanAppSpecServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference)SetInternalValue(val *DataDigitaloceanAppSpecService) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanAppSpecServiceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


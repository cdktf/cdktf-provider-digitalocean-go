//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppSpecJobList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppSpecJobList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppSpecJobList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSpecJobList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppSpecJobList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppSpecJobList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppSpecJobListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataDigitaloceanAppSpecServiceHealthCheckList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataDigitaloceanAppSpecServiceHealthCheckList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceHealthCheckList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceHealthCheckList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataDigitaloceanAppSpecServiceHealthCheckList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataDigitaloceanAppSpecServiceHealthCheckListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


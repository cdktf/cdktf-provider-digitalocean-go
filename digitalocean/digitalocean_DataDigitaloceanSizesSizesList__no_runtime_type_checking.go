//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataDigitaloceanSizesSizesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataDigitaloceanSizesSizesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataDigitaloceanSizesSizesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataDigitaloceanSizesSizesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataDigitaloceanSizesSizesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataDigitaloceanSizesSizesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DigitaloceanProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_DigitaloceanProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateDigitaloceanProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDigitaloceanProviderParameters(scope constructs.Construct, id *string, config *DigitaloceanProviderConfig) error {
	return nil
}


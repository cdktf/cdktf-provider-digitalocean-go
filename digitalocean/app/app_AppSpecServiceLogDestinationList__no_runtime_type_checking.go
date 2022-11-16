//go:build no_runtime_type_checking

package app

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppSpecServiceLogDestinationList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppSpecServiceLogDestinationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppSpecServiceLogDestinationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSpecServiceLogDestinationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppSpecServiceLogDestinationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppSpecServiceLogDestinationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppSpecServiceLogDestinationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


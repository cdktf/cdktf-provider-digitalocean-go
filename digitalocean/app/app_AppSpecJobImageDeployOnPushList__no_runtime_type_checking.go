//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package app

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppSpecJobImageDeployOnPushList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppSpecJobImageDeployOnPushList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppSpecJobImageDeployOnPushList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSpecJobImageDeployOnPushList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppSpecJobImageDeployOnPushList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppSpecJobImageDeployOnPushList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppSpecJobImageDeployOnPushListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


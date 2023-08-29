// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package app

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppSpecServiceImageDeployOnPushList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppSpecServiceImageDeployOnPushList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppSpecServiceImageDeployOnPushList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSpecServiceImageDeployOnPushList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppSpecServiceImageDeployOnPushList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppSpecServiceImageDeployOnPushList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppSpecServiceImageDeployOnPushListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


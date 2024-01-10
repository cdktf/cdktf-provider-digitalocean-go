// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package app

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppSpecJobLogDestinationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppSpecJobLogDestinationList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppSpecJobLogDestinationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppSpecJobLogDestinationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSpecJobLogDestinationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppSpecJobLogDestinationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppSpecJobLogDestinationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppSpecJobLogDestinationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


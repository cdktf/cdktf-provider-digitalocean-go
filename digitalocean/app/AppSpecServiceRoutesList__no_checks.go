// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package app

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppSpecServiceRoutesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppSpecServiceRoutesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppSpecServiceRoutesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppSpecServiceRoutesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSpecServiceRoutesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppSpecServiceRoutesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppSpecServiceRoutesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppSpecServiceRoutesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


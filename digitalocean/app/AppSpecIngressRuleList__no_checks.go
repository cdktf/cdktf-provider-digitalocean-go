// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package app

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppSpecIngressRuleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppSpecIngressRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppSpecIngressRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppSpecIngressRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSpecIngressRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppSpecIngressRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppSpecIngressRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppSpecIngressRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


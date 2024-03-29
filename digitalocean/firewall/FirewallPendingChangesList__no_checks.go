// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package firewall

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirewallPendingChangesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FirewallPendingChangesList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FirewallPendingChangesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FirewallPendingChangesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FirewallPendingChangesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FirewallPendingChangesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFirewallPendingChangesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


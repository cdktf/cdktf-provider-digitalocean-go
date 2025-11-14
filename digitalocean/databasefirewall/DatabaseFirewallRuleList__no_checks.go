// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package databasefirewall

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseFirewallRuleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseFirewallRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatabaseFirewallRuleList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatabaseFirewallRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatabaseFirewallRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatabaseFirewallRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatabaseFirewallRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatabaseFirewallRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


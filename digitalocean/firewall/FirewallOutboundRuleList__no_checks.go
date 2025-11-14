// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package firewall

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirewallOutboundRuleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FirewallOutboundRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FirewallOutboundRuleList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FirewallOutboundRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FirewallOutboundRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FirewallOutboundRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FirewallOutboundRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFirewallOutboundRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


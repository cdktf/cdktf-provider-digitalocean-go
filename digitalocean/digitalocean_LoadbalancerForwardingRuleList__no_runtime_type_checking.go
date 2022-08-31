//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadbalancerForwardingRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LoadbalancerForwardingRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerForwardingRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerForwardingRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerForwardingRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerForwardingRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLoadbalancerForwardingRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


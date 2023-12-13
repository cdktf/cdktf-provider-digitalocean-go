// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package loadbalancer

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_Loadbalancer) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (l *jsiiProxy_Loadbalancer) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (l *jsiiProxy_Loadbalancer) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Loadbalancer) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Loadbalancer) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Loadbalancer) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Loadbalancer) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Loadbalancer) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Loadbalancer) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Loadbalancer) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Loadbalancer) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Loadbalancer) validateImportFromParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_Loadbalancer) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Loadbalancer) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_Loadbalancer) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (l *jsiiProxy_Loadbalancer) validateMoveToIdParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_Loadbalancer) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (l *jsiiProxy_Loadbalancer) validatePutFirewallParameters(value *LoadbalancerFirewall) error {
	return nil
}

func (l *jsiiProxy_Loadbalancer) validatePutForwardingRuleParameters(value interface{}) error {
	return nil
}

func (l *jsiiProxy_Loadbalancer) validatePutHealthcheckParameters(value *LoadbalancerHealthcheck) error {
	return nil
}

func (l *jsiiProxy_Loadbalancer) validatePutStickySessionsParameters(value *LoadbalancerStickySessions) error {
	return nil
}

func validateLoadbalancer_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateLoadbalancer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLoadbalancer_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateLoadbalancer_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Loadbalancer) validateSetAlgorithmParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Loadbalancer) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Loadbalancer) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Loadbalancer) validateSetDisableLetsEncryptDnsRecordsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Loadbalancer) validateSetDropletIdsParameters(val *[]*float64) error {
	return nil
}

func (j *jsiiProxy_Loadbalancer) validateSetDropletTagParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Loadbalancer) validateSetEnableBackendKeepaliveParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Loadbalancer) validateSetEnableProxyProtocolParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Loadbalancer) validateSetHttpIdleTimeoutSecondsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Loadbalancer) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Loadbalancer) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Loadbalancer) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Loadbalancer) validateSetProjectIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Loadbalancer) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Loadbalancer) validateSetRedirectHttpToHttpsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Loadbalancer) validateSetRegionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Loadbalancer) validateSetSizeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Loadbalancer) validateSetSizeUnitParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Loadbalancer) validateSetTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Loadbalancer) validateSetVpcUuidParameters(val *string) error {
	return nil
}

func validateNewLoadbalancerParameters(scope constructs.Construct, id *string, config *LoadbalancerConfig) error {
	return nil
}


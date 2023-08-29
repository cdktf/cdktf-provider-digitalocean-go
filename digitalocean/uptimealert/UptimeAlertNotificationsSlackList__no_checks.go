// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package uptimealert

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UptimeAlertNotificationsSlackList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UptimeAlertNotificationsSlackList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UptimeAlertNotificationsSlackList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_UptimeAlertNotificationsSlackList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UptimeAlertNotificationsSlackList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UptimeAlertNotificationsSlackList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUptimeAlertNotificationsSlackListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


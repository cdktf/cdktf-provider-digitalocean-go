//go:build no_runtime_type_checking

package uptimealert

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UptimeAlertNotificationsList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UptimeAlertNotificationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UptimeAlertNotificationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_UptimeAlertNotificationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UptimeAlertNotificationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UptimeAlertNotificationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUptimeAlertNotificationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


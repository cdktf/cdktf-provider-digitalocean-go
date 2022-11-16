//go:build no_runtime_type_checking

package databasecluster

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseClusterMaintenanceWindowList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterMaintenanceWindowList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatabaseClusterMaintenanceWindowList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatabaseClusterMaintenanceWindowList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatabaseClusterMaintenanceWindowList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatabaseClusterMaintenanceWindowList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatabaseClusterMaintenanceWindowListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


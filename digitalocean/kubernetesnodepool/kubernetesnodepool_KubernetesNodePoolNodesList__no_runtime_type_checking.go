//go:build no_runtime_type_checking

package kubernetesnodepool

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubernetesNodePoolNodesList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KubernetesNodePoolNodesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KubernetesNodePoolNodesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KubernetesNodePoolNodesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KubernetesNodePoolNodesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKubernetesNodePoolNodesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


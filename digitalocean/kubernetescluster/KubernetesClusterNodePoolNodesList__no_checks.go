//go:build no_runtime_type_checking

package kubernetescluster

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubernetesClusterNodePoolNodesList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KubernetesClusterNodePoolNodesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KubernetesClusterNodePoolNodesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KubernetesClusterNodePoolNodesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KubernetesClusterNodePoolNodesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKubernetesClusterNodePoolNodesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


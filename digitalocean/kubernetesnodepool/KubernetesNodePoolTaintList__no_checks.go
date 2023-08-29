// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package kubernetesnodepool

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubernetesNodePoolTaintList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KubernetesNodePoolTaintList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KubernetesNodePoolTaintList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KubernetesNodePoolTaintList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KubernetesNodePoolTaintList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KubernetesNodePoolTaintList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKubernetesNodePoolTaintListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


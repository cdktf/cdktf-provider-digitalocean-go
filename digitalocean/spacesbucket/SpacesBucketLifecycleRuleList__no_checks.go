// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package spacesbucket

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SpacesBucketLifecycleRuleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SpacesBucketLifecycleRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SpacesBucketLifecycleRuleList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SpacesBucketLifecycleRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SpacesBucketLifecycleRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SpacesBucketLifecycleRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SpacesBucketLifecycleRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSpacesBucketLifecycleRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


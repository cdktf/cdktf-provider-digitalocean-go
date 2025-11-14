// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package spaceskey

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SpacesKeyGrantList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SpacesKeyGrantList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SpacesKeyGrantList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SpacesKeyGrantList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SpacesKeyGrantList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SpacesKeyGrantList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SpacesKeyGrantList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSpacesKeyGrantListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package customimage

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomImage) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (c *jsiiProxy_CustomImage) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CustomImage) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomImage) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomImage) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomImage) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomImage) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomImage) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomImage) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomImage) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomImage) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomImage) validateImportFromParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_CustomImage) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomImage) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_CustomImage) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (c *jsiiProxy_CustomImage) validateMoveToIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_CustomImage) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CustomImage) validatePutTimeoutsParameters(value *CustomImageTimeouts) error {
	return nil
}

func validateCustomImage_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateCustomImage_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCustomImage_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateCustomImage_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomImage) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomImage) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomImage) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomImage) validateSetDistributionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomImage) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomImage) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_CustomImage) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomImage) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomImage) validateSetRegionsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_CustomImage) validateSetTagsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_CustomImage) validateSetUrlParameters(val *string) error {
	return nil
}

func validateNewCustomImageParameters(scope constructs.Construct, id *string, config *CustomImageConfig) error {
	return nil
}


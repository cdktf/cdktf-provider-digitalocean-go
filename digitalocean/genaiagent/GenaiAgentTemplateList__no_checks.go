// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package genaiagent

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GenaiAgentTemplateList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GenaiAgentTemplateList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GenaiAgentTemplateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GenaiAgentTemplateList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GenaiAgentTemplateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GenaiAgentTemplateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GenaiAgentTemplateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGenaiAgentTemplateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


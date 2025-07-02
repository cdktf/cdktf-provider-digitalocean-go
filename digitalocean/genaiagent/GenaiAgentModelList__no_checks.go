// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package genaiagent

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GenaiAgentModelList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GenaiAgentModelList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GenaiAgentModelList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GenaiAgentModelList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GenaiAgentModelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GenaiAgentModelList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GenaiAgentModelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGenaiAgentModelListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


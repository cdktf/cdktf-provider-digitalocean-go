// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package genaiagent

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GenaiAgentFunctionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GenaiAgentFunctionsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GenaiAgentFunctionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GenaiAgentFunctionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GenaiAgentFunctionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GenaiAgentFunctionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GenaiAgentFunctionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGenaiAgentFunctionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


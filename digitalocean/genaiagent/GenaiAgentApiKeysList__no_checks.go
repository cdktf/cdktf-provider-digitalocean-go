// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package genaiagent

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GenaiAgentApiKeysList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GenaiAgentApiKeysList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GenaiAgentApiKeysList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GenaiAgentApiKeysList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GenaiAgentApiKeysList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GenaiAgentApiKeysList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GenaiAgentApiKeysList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGenaiAgentApiKeysListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


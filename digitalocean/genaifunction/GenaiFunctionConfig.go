// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaifunction

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GenaiFunctionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the GenAI resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/genai_function#agent_id GenaiFunction#agent_id}
	AgentId *string `field:"required" json:"agentId" yaml:"agentId"`
	// The region where the GenAI resource will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/genai_function#description GenaiFunction#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The current status of the GenAI resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/genai_function#faas_namespace GenaiFunction#faas_namespace}
	FaasNamespace *string `field:"required" json:"faasNamespace" yaml:"faasNamespace"`
	// The creation timestamp of the GenAI resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/genai_function#function_name GenaiFunction#function_name}
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// The input schema of the GenAI resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/genai_function#input_schema GenaiFunction#input_schema}
	InputSchema *string `field:"required" json:"inputSchema" yaml:"inputSchema"`
	// The model to use for the GenAI resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/genai_function#faas_name GenaiFunction#faas_name}
	FaasName *string `field:"optional" json:"faasName" yaml:"faasName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/genai_function#id GenaiFunction#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The output schema of the GenAI resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/genai_function#output_schema GenaiFunction#output_schema}
	OutputSchema *string `field:"optional" json:"outputSchema" yaml:"outputSchema"`
}


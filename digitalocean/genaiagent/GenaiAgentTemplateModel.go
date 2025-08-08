// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiagent


type GenaiAgentTemplateModel struct {
	// agreement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_agent#agreement GenaiAgent#agreement}
	Agreement interface{} `field:"optional" json:"agreement" yaml:"agreement"`
	// Inference name of the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_agent#inference_name GenaiAgent#inference_name}
	InferenceName *string `field:"optional" json:"inferenceName" yaml:"inferenceName"`
	// Infernce version of the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_agent#inference_version GenaiAgent#inference_version}
	InferenceVersion *string `field:"optional" json:"inferenceVersion" yaml:"inferenceVersion"`
	// Indicates if the Model Base is foundational.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_agent#is_foundational GenaiAgent#is_foundational}
	IsFoundational interface{} `field:"optional" json:"isFoundational" yaml:"isFoundational"`
	// Name of the Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_agent#name GenaiAgent#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Parent UUID of the Model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_agent#parent_uuid GenaiAgent#parent_uuid}
	ParentUuid *string `field:"optional" json:"parentUuid" yaml:"parentUuid"`
	// Provider of the Model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_agent#provider GenaiAgent#provider}
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// Indicates if the Model upload is complete.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_agent#upload_complete GenaiAgent#upload_complete}
	UploadComplete interface{} `field:"optional" json:"uploadComplete" yaml:"uploadComplete"`
	// URL of the Model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_agent#url GenaiAgent#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// List of Usecases for the Model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_agent#usecases GenaiAgent#usecases}
	Usecases *[]*string `field:"optional" json:"usecases" yaml:"usecases"`
	// versions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_agent#versions GenaiAgent#versions}
	Versions interface{} `field:"optional" json:"versions" yaml:"versions"`
}


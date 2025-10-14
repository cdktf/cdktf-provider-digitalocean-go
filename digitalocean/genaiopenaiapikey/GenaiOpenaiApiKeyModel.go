// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiopenaiapikey


type GenaiOpenaiApiKeyModel struct {
	// agreement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_openai_api_key#agreement GenaiOpenaiApiKey#agreement}
	Agreement interface{} `field:"optional" json:"agreement" yaml:"agreement"`
	// Inference name of the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_openai_api_key#inference_name GenaiOpenaiApiKey#inference_name}
	InferenceName *string `field:"optional" json:"inferenceName" yaml:"inferenceName"`
	// Infernce version of the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_openai_api_key#inference_version GenaiOpenaiApiKey#inference_version}
	InferenceVersion *string `field:"optional" json:"inferenceVersion" yaml:"inferenceVersion"`
	// Indicates if the Model Base is foundational.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_openai_api_key#is_foundational GenaiOpenaiApiKey#is_foundational}
	IsFoundational interface{} `field:"optional" json:"isFoundational" yaml:"isFoundational"`
	// Name of the Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_openai_api_key#name GenaiOpenaiApiKey#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Parent UUID of the Model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_openai_api_key#parent_uuid GenaiOpenaiApiKey#parent_uuid}
	ParentUuid *string `field:"optional" json:"parentUuid" yaml:"parentUuid"`
	// Provider of the Model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_openai_api_key#provider GenaiOpenaiApiKey#provider}
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// Indicates if the Model upload is complete.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_openai_api_key#upload_complete GenaiOpenaiApiKey#upload_complete}
	UploadComplete interface{} `field:"optional" json:"uploadComplete" yaml:"uploadComplete"`
	// URL of the Model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_openai_api_key#url GenaiOpenaiApiKey#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// List of Usecases for the Model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_openai_api_key#usecases GenaiOpenaiApiKey#usecases}
	Usecases *[]*string `field:"optional" json:"usecases" yaml:"usecases"`
	// versions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_openai_api_key#versions GenaiOpenaiApiKey#versions}
	Versions interface{} `field:"optional" json:"versions" yaml:"versions"`
}


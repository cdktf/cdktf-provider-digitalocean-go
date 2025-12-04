// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiagent


type DataDigitaloceanGenaiAgentModel struct {
	// agreement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/data-sources/genai_agent#agreement DataDigitaloceanGenaiAgent#agreement}
	Agreement interface{} `field:"optional" json:"agreement" yaml:"agreement"`
	// Inference name of the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/data-sources/genai_agent#inference_name DataDigitaloceanGenaiAgent#inference_name}
	InferenceName *string `field:"optional" json:"inferenceName" yaml:"inferenceName"`
	// Infernce version of the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/data-sources/genai_agent#inference_version DataDigitaloceanGenaiAgent#inference_version}
	InferenceVersion *string `field:"optional" json:"inferenceVersion" yaml:"inferenceVersion"`
	// Indicates if the Model Base is foundational.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/data-sources/genai_agent#is_foundational DataDigitaloceanGenaiAgent#is_foundational}
	IsFoundational interface{} `field:"optional" json:"isFoundational" yaml:"isFoundational"`
	// Name of the Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/data-sources/genai_agent#name DataDigitaloceanGenaiAgent#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Parent UUID of the Model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/data-sources/genai_agent#parent_uuid DataDigitaloceanGenaiAgent#parent_uuid}
	ParentUuid *string `field:"optional" json:"parentUuid" yaml:"parentUuid"`
	// Provider of the Model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/data-sources/genai_agent#provider DataDigitaloceanGenaiAgent#provider}
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// Indicates if the Model upload is complete.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/data-sources/genai_agent#upload_complete DataDigitaloceanGenaiAgent#upload_complete}
	UploadComplete interface{} `field:"optional" json:"uploadComplete" yaml:"uploadComplete"`
	// URL of the Model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/data-sources/genai_agent#url DataDigitaloceanGenaiAgent#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// List of Usecases for the Model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/data-sources/genai_agent#usecases DataDigitaloceanGenaiAgent#usecases}
	Usecases *[]*string `field:"optional" json:"usecases" yaml:"usecases"`
	// versions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/data-sources/genai_agent#versions DataDigitaloceanGenaiAgent#versions}
	Versions interface{} `field:"optional" json:"versions" yaml:"versions"`
}


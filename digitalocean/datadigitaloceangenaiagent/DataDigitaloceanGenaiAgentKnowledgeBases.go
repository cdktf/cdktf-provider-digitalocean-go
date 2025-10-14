// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiagent


type DataDigitaloceanGenaiAgentKnowledgeBases struct {
	// Database ID of the Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/data-sources/genai_agent#database_id DataDigitaloceanGenaiAgent#database_id}
	DatabaseId *string `field:"optional" json:"databaseId" yaml:"databaseId"`
	// Embedding model UUID for the Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/data-sources/genai_agent#embedding_model_uuid DataDigitaloceanGenaiAgent#embedding_model_uuid}
	EmbeddingModelUuid *string `field:"optional" json:"embeddingModelUuid" yaml:"embeddingModelUuid"`
	// Indicates if the Knowledge Base is public.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/data-sources/genai_agent#is_public DataDigitaloceanGenaiAgent#is_public}
	IsPublic interface{} `field:"optional" json:"isPublic" yaml:"isPublic"`
	// last_indexing_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/data-sources/genai_agent#last_indexing_job DataDigitaloceanGenaiAgent#last_indexing_job}
	LastIndexingJob *DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJob `field:"optional" json:"lastIndexingJob" yaml:"lastIndexingJob"`
	// Name of the Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/data-sources/genai_agent#name DataDigitaloceanGenaiAgent#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Project ID of the Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/data-sources/genai_agent#project_id DataDigitaloceanGenaiAgent#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Region of the Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/data-sources/genai_agent#region DataDigitaloceanGenaiAgent#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// List of tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/data-sources/genai_agent#tags DataDigitaloceanGenaiAgent#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// User ID of the Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/data-sources/genai_agent#user_id DataDigitaloceanGenaiAgent#user_id}
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}


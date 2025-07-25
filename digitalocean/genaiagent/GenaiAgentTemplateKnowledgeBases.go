// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiagent


type GenaiAgentTemplateKnowledgeBases struct {
	// Database ID of the Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_agent#database_id GenaiAgent#database_id}
	DatabaseId *string `field:"optional" json:"databaseId" yaml:"databaseId"`
	// Embedding model UUID for the Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_agent#embedding_model_uuid GenaiAgent#embedding_model_uuid}
	EmbeddingModelUuid *string `field:"optional" json:"embeddingModelUuid" yaml:"embeddingModelUuid"`
	// Indicates if the Knowledge Base is public.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_agent#is_public GenaiAgent#is_public}
	IsPublic interface{} `field:"optional" json:"isPublic" yaml:"isPublic"`
	// last_indexing_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_agent#last_indexing_job GenaiAgent#last_indexing_job}
	LastIndexingJob *GenaiAgentTemplateKnowledgeBasesLastIndexingJob `field:"optional" json:"lastIndexingJob" yaml:"lastIndexingJob"`
	// Name of the Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_agent#name GenaiAgent#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Project ID of the Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_agent#project_id GenaiAgent#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Region of the Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_agent#region GenaiAgent#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// List of tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_agent#tags GenaiAgent#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// User ID of the Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_agent#user_id GenaiAgent#user_id}
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}


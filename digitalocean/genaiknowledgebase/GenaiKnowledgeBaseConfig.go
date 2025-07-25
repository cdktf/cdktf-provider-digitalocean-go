// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiknowledgebase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GenaiKnowledgeBaseConfig struct {
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
	// datasources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_knowledge_base#datasources GenaiKnowledgeBase#datasources}
	Datasources interface{} `field:"required" json:"datasources" yaml:"datasources"`
	// The unique identifier of the embedding model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_knowledge_base#embedding_model_uuid GenaiKnowledgeBase#embedding_model_uuid}
	EmbeddingModelUuid *string `field:"required" json:"embeddingModelUuid" yaml:"embeddingModelUuid"`
	// The name of the knowledge base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_knowledge_base#name GenaiKnowledgeBase#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The unique identifier of the project to which the knowledge base belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_knowledge_base#project_id GenaiKnowledgeBase#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_knowledge_base#region GenaiKnowledgeBase#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The time when the knowledge base was added to the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_knowledge_base#added_to_agent_at GenaiKnowledgeBase#added_to_agent_at}
	AddedToAgentAt *string `field:"optional" json:"addedToAgentAt" yaml:"addedToAgentAt"`
	// The unique identifier of the DigitalOcean OpenSearch database this knowledge base will use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_knowledge_base#database_id GenaiKnowledgeBase#database_id}
	DatabaseId *string `field:"optional" json:"databaseId" yaml:"databaseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_knowledge_base#id GenaiKnowledgeBase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Indicates whether the knowledge base is public or private.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_knowledge_base#is_public GenaiKnowledgeBase#is_public}
	IsPublic interface{} `field:"optional" json:"isPublic" yaml:"isPublic"`
	// last_indexing_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_knowledge_base#last_indexing_job GenaiKnowledgeBase#last_indexing_job}
	LastIndexingJob interface{} `field:"optional" json:"lastIndexingJob" yaml:"lastIndexingJob"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_knowledge_base#tags GenaiKnowledgeBase#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// The unique identifier of the VPC to which the knowledge base belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/genai_knowledge_base#vpc_uuid GenaiKnowledgeBase#vpc_uuid}
	VpcUuid *string `field:"optional" json:"vpcUuid" yaml:"vpcUuid"`
}


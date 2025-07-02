// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiagent


type GenaiAgentKnowledgeBasesLastIndexingJob struct {
	// Number of completed datasources in the last indexing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/resources/genai_agent#completed_datasources GenaiAgent#completed_datasources}
	CompletedDatasources *float64 `field:"optional" json:"completedDatasources" yaml:"completedDatasources"`
	// Datasource UUIDs for the last indexing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/resources/genai_agent#datasource_uuids GenaiAgent#datasource_uuids}
	DatasourceUuids *[]*string `field:"optional" json:"datasourceUuids" yaml:"datasourceUuids"`
	// UUID	of the Knowledge Base for the last indexing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/resources/genai_agent#knowledge_base_uuid GenaiAgent#knowledge_base_uuid}
	KnowledgeBaseUuid *string `field:"optional" json:"knowledgeBaseUuid" yaml:"knowledgeBaseUuid"`
	// Phase of the last indexing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/resources/genai_agent#phase GenaiAgent#phase}
	Phase *string `field:"optional" json:"phase" yaml:"phase"`
	// Number of tokens processed in the last indexing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/resources/genai_agent#tokens GenaiAgent#tokens}
	Tokens *float64 `field:"optional" json:"tokens" yaml:"tokens"`
	// Total number of datasources in the last indexing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/resources/genai_agent#total_datasources GenaiAgent#total_datasources}
	TotalDatasources *float64 `field:"optional" json:"totalDatasources" yaml:"totalDatasources"`
	// UUID	of the last indexing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/resources/genai_agent#uuid GenaiAgent#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
}


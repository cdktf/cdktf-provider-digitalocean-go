// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiknowledgebase


type GenaiKnowledgeBaseLastIndexingJob struct {
	// Number of completed datasources in the last indexing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_knowledge_base#completed_datasources GenaiKnowledgeBase#completed_datasources}
	CompletedDatasources *float64 `field:"optional" json:"completedDatasources" yaml:"completedDatasources"`
	// Datasource UUIDs for the last indexing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_knowledge_base#data_source_uuids GenaiKnowledgeBase#data_source_uuids}
	DataSourceUuids *[]*string `field:"optional" json:"dataSourceUuids" yaml:"dataSourceUuids"`
	// Phase of the last indexing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_knowledge_base#phase GenaiKnowledgeBase#phase}
	Phase *string `field:"optional" json:"phase" yaml:"phase"`
	// Number of tokens processed in the last indexing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_knowledge_base#tokens GenaiKnowledgeBase#tokens}
	Tokens *float64 `field:"optional" json:"tokens" yaml:"tokens"`
	// Total number of datasources in the last indexing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_knowledge_base#total_datasources GenaiKnowledgeBase#total_datasources}
	TotalDatasources *float64 `field:"optional" json:"totalDatasources" yaml:"totalDatasources"`
	// UUID  of the last indexing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_knowledge_base#uuid GenaiKnowledgeBase#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
}


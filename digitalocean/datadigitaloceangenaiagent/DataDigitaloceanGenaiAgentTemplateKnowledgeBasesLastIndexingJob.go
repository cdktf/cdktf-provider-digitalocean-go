// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiagent


type DataDigitaloceanGenaiAgentTemplateKnowledgeBasesLastIndexingJob struct {
	// Number of completed datasources in the last indexing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/data-sources/genai_agent#completed_datasources DataDigitaloceanGenaiAgent#completed_datasources}
	CompletedDatasources *float64 `field:"optional" json:"completedDatasources" yaml:"completedDatasources"`
	// Datasource UUIDs for the last indexing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/data-sources/genai_agent#datasource_uuids DataDigitaloceanGenaiAgent#datasource_uuids}
	DatasourceUuids *[]*string `field:"optional" json:"datasourceUuids" yaml:"datasourceUuids"`
	// UUID	of the Knowledge Base for the last indexing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/data-sources/genai_agent#knowledge_base_uuid DataDigitaloceanGenaiAgent#knowledge_base_uuid}
	KnowledgeBaseUuid *string `field:"optional" json:"knowledgeBaseUuid" yaml:"knowledgeBaseUuid"`
	// Phase of the last indexing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/data-sources/genai_agent#phase DataDigitaloceanGenaiAgent#phase}
	Phase *string `field:"optional" json:"phase" yaml:"phase"`
	// Number of tokens processed in the last indexing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/data-sources/genai_agent#tokens DataDigitaloceanGenaiAgent#tokens}
	Tokens *float64 `field:"optional" json:"tokens" yaml:"tokens"`
	// Total number of datasources in the last indexing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/data-sources/genai_agent#total_datasources DataDigitaloceanGenaiAgent#total_datasources}
	TotalDatasources *float64 `field:"optional" json:"totalDatasources" yaml:"totalDatasources"`
	// UUID	of the last indexing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/data-sources/genai_agent#uuid DataDigitaloceanGenaiAgent#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
}


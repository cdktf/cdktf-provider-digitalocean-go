// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiknowledgebase


type GenaiKnowledgeBaseDatasources struct {
	// file_upload_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/genai_knowledge_base#file_upload_data_source GenaiKnowledgeBase#file_upload_data_source}
	FileUploadDataSource interface{} `field:"optional" json:"fileUploadDataSource" yaml:"fileUploadDataSource"`
	// last_indexing_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/genai_knowledge_base#last_indexing_job GenaiKnowledgeBase#last_indexing_job}
	LastIndexingJob interface{} `field:"optional" json:"lastIndexingJob" yaml:"lastIndexingJob"`
	// spaces_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/genai_knowledge_base#spaces_data_source GenaiKnowledgeBase#spaces_data_source}
	SpacesDataSource interface{} `field:"optional" json:"spacesDataSource" yaml:"spacesDataSource"`
	// UUID of the Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/genai_knowledge_base#uuid GenaiKnowledgeBase#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
	// web_crawler_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/genai_knowledge_base#web_crawler_data_source GenaiKnowledgeBase#web_crawler_data_source}
	WebCrawlerDataSource interface{} `field:"optional" json:"webCrawlerDataSource" yaml:"webCrawlerDataSource"`
}


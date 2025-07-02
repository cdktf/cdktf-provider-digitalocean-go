// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiagent


type DataDigitaloceanGenaiAgentKnowledgeBases struct {
	// last_indexing_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/data-sources/genai_agent#last_indexing_job DataDigitaloceanGenaiAgent#last_indexing_job}
	LastIndexingJob interface{} `field:"optional" json:"lastIndexingJob" yaml:"lastIndexingJob"`
	// Region of the Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/data-sources/genai_agent#region DataDigitaloceanGenaiAgent#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// List of tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/data-sources/genai_agent#tags DataDigitaloceanGenaiAgent#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// User ID of the Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/data-sources/genai_agent#user_id DataDigitaloceanGenaiAgent#user_id}
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
	// UUID of the Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/data-sources/genai_agent#uuid DataDigitaloceanGenaiAgent#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiknowledgebases


type DataDigitaloceanGenaiKnowledgeBasesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.59.0/docs/data-sources/genai_knowledge_bases#key DataDigitaloceanGenaiKnowledgeBases#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.59.0/docs/data-sources/genai_knowledge_bases#values DataDigitaloceanGenaiKnowledgeBases#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.59.0/docs/data-sources/genai_knowledge_bases#all DataDigitaloceanGenaiKnowledgeBases#all}.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.59.0/docs/data-sources/genai_knowledge_bases#match_by DataDigitaloceanGenaiKnowledgeBases#match_by}.
	MatchBy *string `field:"optional" json:"matchBy" yaml:"matchBy"`
}


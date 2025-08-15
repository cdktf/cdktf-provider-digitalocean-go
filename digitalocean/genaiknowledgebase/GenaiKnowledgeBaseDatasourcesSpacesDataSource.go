// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiknowledgebase


type GenaiKnowledgeBaseDatasourcesSpacesDataSource struct {
	// The name of the Spaces bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/genai_knowledge_base#bucket_name GenaiKnowledgeBase#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The path to the item in the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/genai_knowledge_base#item_path GenaiKnowledgeBase#item_path}
	ItemPath *string `field:"optional" json:"itemPath" yaml:"itemPath"`
	// The region of the Spaces bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/genai_knowledge_base#region GenaiKnowledgeBase#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}


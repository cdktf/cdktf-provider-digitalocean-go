// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiknowledgebasedatasource


type GenaiKnowledgeBaseDataSourceSpacesDataSource struct {
	// The name of the Spaces bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/genai_knowledge_base_data_source#bucket_name GenaiKnowledgeBaseDataSource#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The path to the item in the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/genai_knowledge_base_data_source#item_path GenaiKnowledgeBaseDataSource#item_path}
	ItemPath *string `field:"optional" json:"itemPath" yaml:"itemPath"`
	// The region of the Spaces bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/genai_knowledge_base_data_source#region GenaiKnowledgeBaseDataSource#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}


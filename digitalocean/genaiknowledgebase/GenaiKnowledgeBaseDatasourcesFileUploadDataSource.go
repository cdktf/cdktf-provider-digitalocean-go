// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiknowledgebase


type GenaiKnowledgeBaseDatasourcesFileUploadDataSource struct {
	// The original name of the uploaded file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/resources/genai_knowledge_base#original_file_name GenaiKnowledgeBase#original_file_name}
	OriginalFileName *string `field:"optional" json:"originalFileName" yaml:"originalFileName"`
	// The size of the file in bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/resources/genai_knowledge_base#size_in_bytes GenaiKnowledgeBase#size_in_bytes}
	SizeInBytes *string `field:"optional" json:"sizeInBytes" yaml:"sizeInBytes"`
	// The stored object key for the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/resources/genai_knowledge_base#stored_object_key GenaiKnowledgeBase#stored_object_key}
	StoredObjectKey *string `field:"optional" json:"storedObjectKey" yaml:"storedObjectKey"`
}


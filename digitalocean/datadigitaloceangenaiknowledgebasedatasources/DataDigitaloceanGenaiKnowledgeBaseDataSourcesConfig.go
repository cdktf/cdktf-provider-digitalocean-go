// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiknowledgebasedatasources

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDigitaloceanGenaiKnowledgeBaseDataSourcesConfig struct {
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
	// UUID of the Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/data-sources/genai_knowledge_base_data_sources#knowledge_base_uuid DataDigitaloceanGenaiKnowledgeBaseDataSources#knowledge_base_uuid}
	KnowledgeBaseUuid *string `field:"required" json:"knowledgeBaseUuid" yaml:"knowledgeBaseUuid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/data-sources/genai_knowledge_base_data_sources#id DataDigitaloceanGenaiKnowledgeBaseDataSources#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecJobLogDestinationOpenSearch struct {
	// basic_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.59.0/docs/resources/app#basic_auth App#basic_auth}
	BasicAuth *AppSpecJobLogDestinationOpenSearchBasicAuth `field:"required" json:"basicAuth" yaml:"basicAuth"`
	// OpenSearch cluster name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.59.0/docs/resources/app#cluster_name App#cluster_name}
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// OpenSearch endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.59.0/docs/resources/app#endpoint App#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// OpenSearch index name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.59.0/docs/resources/app#index_name App#index_name}
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
}


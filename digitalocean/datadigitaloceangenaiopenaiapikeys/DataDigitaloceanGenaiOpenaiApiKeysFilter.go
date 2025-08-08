// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiopenaiapikeys


type DataDigitaloceanGenaiOpenaiApiKeysFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/data-sources/genai_openai_api_keys#key DataDigitaloceanGenaiOpenaiApiKeys#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/data-sources/genai_openai_api_keys#values DataDigitaloceanGenaiOpenaiApiKeys#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/data-sources/genai_openai_api_keys#all DataDigitaloceanGenaiOpenaiApiKeys#all}.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/data-sources/genai_openai_api_keys#match_by DataDigitaloceanGenaiOpenaiApiKeys#match_by}.
	MatchBy *string `field:"optional" json:"matchBy" yaml:"matchBy"`
}


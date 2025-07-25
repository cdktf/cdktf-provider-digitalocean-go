// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiagentversions


type DataDigitaloceanGenaiAgentVersionsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/data-sources/genai_agent_versions#key DataDigitaloceanGenaiAgentVersions#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/data-sources/genai_agent_versions#values DataDigitaloceanGenaiAgentVersions#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/data-sources/genai_agent_versions#all DataDigitaloceanGenaiAgentVersions#all}.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/data-sources/genai_agent_versions#match_by DataDigitaloceanGenaiAgentVersions#match_by}.
	MatchBy *string `field:"optional" json:"matchBy" yaml:"matchBy"`
}


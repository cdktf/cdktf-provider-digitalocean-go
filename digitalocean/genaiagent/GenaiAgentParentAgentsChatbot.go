// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiagent


type GenaiAgentParentAgentsChatbot struct {
	// Background color for the chatbot button.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/genai_agent#button_background_color GenaiAgent#button_background_color}
	ButtonBackgroundColor *string `field:"optional" json:"buttonBackgroundColor" yaml:"buttonBackgroundColor"`
	// Logo for the chatbot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/genai_agent#logo GenaiAgent#logo}
	Logo *string `field:"optional" json:"logo" yaml:"logo"`
	// Name of the chatbot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/genai_agent#name GenaiAgent#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Primary color for the chatbot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/genai_agent#primary_color GenaiAgent#primary_color}
	PrimaryColor *string `field:"optional" json:"primaryColor" yaml:"primaryColor"`
	// Secondary color for the chatbot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/genai_agent#secondary_color GenaiAgent#secondary_color}
	SecondaryColor *string `field:"optional" json:"secondaryColor" yaml:"secondaryColor"`
	// Starting message for the chatbot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/genai_agent#starting_message GenaiAgent#starting_message}
	StartingMessage *string `field:"optional" json:"startingMessage" yaml:"startingMessage"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiagent


type DataDigitaloceanGenaiAgentChatbot struct {
	// Background color for the chatbot button.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/data-sources/genai_agent#button_background_color DataDigitaloceanGenaiAgent#button_background_color}
	ButtonBackgroundColor *string `field:"optional" json:"buttonBackgroundColor" yaml:"buttonBackgroundColor"`
	// Logo for the chatbot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/data-sources/genai_agent#logo DataDigitaloceanGenaiAgent#logo}
	Logo *string `field:"optional" json:"logo" yaml:"logo"`
	// Name of the chatbot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/data-sources/genai_agent#name DataDigitaloceanGenaiAgent#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Primary color for the chatbot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/data-sources/genai_agent#primary_color DataDigitaloceanGenaiAgent#primary_color}
	PrimaryColor *string `field:"optional" json:"primaryColor" yaml:"primaryColor"`
	// Secondary color for the chatbot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/data-sources/genai_agent#secondary_color DataDigitaloceanGenaiAgent#secondary_color}
	SecondaryColor *string `field:"optional" json:"secondaryColor" yaml:"secondaryColor"`
	// Starting message for the chatbot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/data-sources/genai_agent#starting_message DataDigitaloceanGenaiAgent#starting_message}
	StartingMessage *string `field:"optional" json:"startingMessage" yaml:"startingMessage"`
}


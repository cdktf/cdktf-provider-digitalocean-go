// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dropletautoscale


type DropletAutoscaleDropletTemplate struct {
	// Droplet image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/droplet_autoscale#image DropletAutoscale#image}
	Image *string `field:"required" json:"image" yaml:"image"`
	// Droplet region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/droplet_autoscale#region DropletAutoscale#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// Droplet size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/droplet_autoscale#size DropletAutoscale#size}
	Size *string `field:"required" json:"size" yaml:"size"`
	// Droplet SSH keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/droplet_autoscale#ssh_keys DropletAutoscale#ssh_keys}
	SshKeys *[]*string `field:"required" json:"sshKeys" yaml:"sshKeys"`
	// Enable droplet IPv6.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/droplet_autoscale#ipv6 DropletAutoscale#ipv6}
	Ipv6 interface{} `field:"optional" json:"ipv6" yaml:"ipv6"`
	// Droplet project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/droplet_autoscale#project_id DropletAutoscale#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Droplet tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/droplet_autoscale#tags DropletAutoscale#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Droplet user data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/droplet_autoscale#user_data DropletAutoscale#user_data}
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// Droplet VPC UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/droplet_autoscale#vpc_uuid DropletAutoscale#vpc_uuid}
	VpcUuid *string `field:"optional" json:"vpcUuid" yaml:"vpcUuid"`
	// Enable droplet agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/droplet_autoscale#with_droplet_agent DropletAutoscale#with_droplet_agent}
	WithDropletAgent interface{} `field:"optional" json:"withDropletAgent" yaml:"withDropletAgent"`
}


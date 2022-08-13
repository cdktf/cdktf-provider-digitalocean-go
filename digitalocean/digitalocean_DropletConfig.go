// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DropletConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/droplet#image Droplet#image}.
	Image *string `field:"required" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/droplet#name Droplet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/droplet#region Droplet#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/droplet#size Droplet#size}.
	Size *string `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/droplet#backups Droplet#backups}.
	Backups interface{} `field:"optional" json:"backups" yaml:"backups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/droplet#droplet_agent Droplet#droplet_agent}.
	DropletAgent interface{} `field:"optional" json:"dropletAgent" yaml:"dropletAgent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/droplet#graceful_shutdown Droplet#graceful_shutdown}.
	GracefulShutdown interface{} `field:"optional" json:"gracefulShutdown" yaml:"gracefulShutdown"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/droplet#id Droplet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/droplet#ipv6 Droplet#ipv6}.
	Ipv6 interface{} `field:"optional" json:"ipv6" yaml:"ipv6"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/droplet#monitoring Droplet#monitoring}.
	Monitoring interface{} `field:"optional" json:"monitoring" yaml:"monitoring"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/droplet#private_networking Droplet#private_networking}.
	PrivateNetworking interface{} `field:"optional" json:"privateNetworking" yaml:"privateNetworking"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/droplet#resize_disk Droplet#resize_disk}.
	ResizeDisk interface{} `field:"optional" json:"resizeDisk" yaml:"resizeDisk"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/droplet#ssh_keys Droplet#ssh_keys}.
	SshKeys *[]*string `field:"optional" json:"sshKeys" yaml:"sshKeys"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/droplet#tags Droplet#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/droplet#user_data Droplet#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/droplet#volume_ids Droplet#volume_ids}.
	VolumeIds *[]*string `field:"optional" json:"volumeIds" yaml:"volumeIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/droplet#vpc_uuid Droplet#vpc_uuid}.
	VpcUuid *string `field:"optional" json:"vpcUuid" yaml:"vpcUuid"`
}


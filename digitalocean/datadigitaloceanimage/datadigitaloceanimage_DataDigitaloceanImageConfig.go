package datadigitaloceanimage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDigitaloceanImageConfig struct {
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
	// id of the image.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/image#id DataDigitaloceanImage#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *float64 `field:"optional" json:"id" yaml:"id"`
	// name of the image.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/image#name DataDigitaloceanImage#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// slug of the image.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/image#slug DataDigitaloceanImage#slug}
	Slug *string `field:"optional" json:"slug" yaml:"slug"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/image#source DataDigitaloceanImage#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

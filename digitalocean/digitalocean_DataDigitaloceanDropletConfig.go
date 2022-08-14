// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDigitaloceanDropletConfig struct {
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
	// id of the Droplet.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/droplet#id DataDigitaloceanDroplet#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *float64 `field:"optional" json:"id" yaml:"id"`
	// name of the Droplet.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/droplet#name DataDigitaloceanDroplet#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// unique tag of the Droplet.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/droplet#tag DataDigitaloceanDroplet#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type DatabaseClusterMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/database_cluster#day DatabaseCluster#day}.
	Day *string `field:"required" json:"day" yaml:"day"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/database_cluster#hour DatabaseCluster#hour}.
	Hour *string `field:"required" json:"hour" yaml:"hour"`
}


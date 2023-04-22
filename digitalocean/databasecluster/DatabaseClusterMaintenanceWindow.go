package databasecluster


type DatabaseClusterMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.0/docs/resources/database_cluster#day DatabaseCluster#day}.
	Day *string `field:"required" json:"day" yaml:"day"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.0/docs/resources/database_cluster#hour DatabaseCluster#hour}.
	Hour *string `field:"required" json:"hour" yaml:"hour"`
}


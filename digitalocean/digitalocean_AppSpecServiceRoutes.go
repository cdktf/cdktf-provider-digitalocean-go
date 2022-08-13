// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type AppSpecServiceRoutes struct {
	// Path specifies an route by HTTP path prefix.
	//
	// Paths must start with / and must be unique within the app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#path App#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// An optional flag to preserve the path that is forwarded to the backend service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#preserve_path_prefix App#preserve_path_prefix}
	PreservePathPrefix interface{} `field:"optional" json:"preservePathPrefix" yaml:"preservePathPrefix"`
}


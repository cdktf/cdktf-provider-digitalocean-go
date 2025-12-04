// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package uptimealert

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UptimeAlertConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// A unique identifier for a check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/uptime_alert#check_id UptimeAlert#check_id}
	CheckId *string `field:"required" json:"checkId" yaml:"checkId"`
	// A human-friendly display name for the alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/uptime_alert#name UptimeAlert#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/uptime_alert#notifications UptimeAlert#notifications}
	Notifications interface{} `field:"required" json:"notifications" yaml:"notifications"`
	// The type of health check to perform. Enum: 'latency' 'down' 'down_global' 'ssl_expiry'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/uptime_alert#type UptimeAlert#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The comparison operator used against the alert's threshold. Enum: 'greater_than' 'less_than.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/uptime_alert#comparison UptimeAlert#comparison}
	Comparison *string `field:"optional" json:"comparison" yaml:"comparison"`
	// Period of time the threshold must be exceeded to trigger the alert.
	//
	// Enum '2m' '3m' '5m' '10m' '15m' '30m' '1h'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/uptime_alert#period UptimeAlert#period}
	Period *string `field:"optional" json:"period" yaml:"period"`
	// The threshold at which the alert will enter a trigger state.
	//
	// The specific threshold is dependent on the alert type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/uptime_alert#threshold UptimeAlert#threshold}
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
}


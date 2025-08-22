// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecServiceTermination struct {
	// The number of seconds to wait between selecting a container instance for termination and issuing the TERM signal.
	//
	// Selecting a container instance for termination begins an asynchronous drain of new requests on upstream load-balancers. Default: 15 seconds, Minimum 1, Maximum 110.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/app#drain_seconds App#drain_seconds}
	DrainSeconds *float64 `field:"optional" json:"drainSeconds" yaml:"drainSeconds"`
	// The number of seconds to wait between sending a TERM signal to a container and issuing a KILL which causes immediate shutdown.
	//
	// Default: 120, Minimum 1, Maximum 600.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/app#grace_period_seconds App#grace_period_seconds}
	GracePeriodSeconds *float64 `field:"optional" json:"gracePeriodSeconds" yaml:"gracePeriodSeconds"`
}


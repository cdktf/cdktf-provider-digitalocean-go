// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecWorkerTermination struct {
	// The number of seconds to wait between sending a TERM signal to a container and issuing a KILL which causes immediate shutdown.
	//
	// Default: 120, Minimum 1, Maximum 600.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.48.1/docs/resources/app#grace_period_seconds App#grace_period_seconds}
	GracePeriodSeconds *float64 `field:"optional" json:"gracePeriodSeconds" yaml:"gracePeriodSeconds"`
}


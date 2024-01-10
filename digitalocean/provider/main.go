// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-digitalocean.provider.DigitaloceanProvider",
		reflect.TypeOf((*DigitaloceanProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "apiEndpoint", GoGetter: "ApiEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "apiEndpointInput", GoGetter: "ApiEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "httpRetryMax", GoGetter: "HttpRetryMax"},
			_jsii_.MemberProperty{JsiiProperty: "httpRetryMaxInput", GoGetter: "HttpRetryMaxInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpRetryWaitMax", GoGetter: "HttpRetryWaitMax"},
			_jsii_.MemberProperty{JsiiProperty: "httpRetryWaitMaxInput", GoGetter: "HttpRetryWaitMaxInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpRetryWaitMin", GoGetter: "HttpRetryWaitMin"},
			_jsii_.MemberProperty{JsiiProperty: "httpRetryWaitMinInput", GoGetter: "HttpRetryWaitMinInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requestsPerSecond", GoGetter: "RequestsPerSecond"},
			_jsii_.MemberProperty{JsiiProperty: "requestsPerSecondInput", GoGetter: "RequestsPerSecondInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiEndpoint", GoMethod: "ResetApiEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpRetryMax", GoMethod: "ResetHttpRetryMax"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpRetryWaitMax", GoMethod: "ResetHttpRetryWaitMax"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpRetryWaitMin", GoMethod: "ResetHttpRetryWaitMin"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestsPerSecond", GoMethod: "ResetRequestsPerSecond"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpacesAccessId", GoMethod: "ResetSpacesAccessId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpacesEndpoint", GoMethod: "ResetSpacesEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpacesSecretKey", GoMethod: "ResetSpacesSecretKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetToken", GoMethod: "ResetToken"},
			_jsii_.MemberProperty{JsiiProperty: "spacesAccessId", GoGetter: "SpacesAccessId"},
			_jsii_.MemberProperty{JsiiProperty: "spacesAccessIdInput", GoGetter: "SpacesAccessIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "spacesEndpoint", GoGetter: "SpacesEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "spacesEndpointInput", GoGetter: "SpacesEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "spacesSecretKey", GoGetter: "SpacesSecretKey"},
			_jsii_.MemberProperty{JsiiProperty: "spacesSecretKeyInput", GoGetter: "SpacesSecretKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "token", GoGetter: "Token"},
			_jsii_.MemberProperty{JsiiProperty: "tokenInput", GoGetter: "TokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DigitaloceanProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-digitalocean.provider.DigitaloceanProviderConfig",
		reflect.TypeOf((*DigitaloceanProviderConfig)(nil)).Elem(),
	)
}

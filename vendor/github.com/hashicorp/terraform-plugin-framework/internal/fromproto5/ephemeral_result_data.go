// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fromproto5

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/internal/fwschema"
	"github.com/hashicorp/terraform-plugin-framework/internal/fwschemadata"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

// EphemeralResultData returns the *tfsdk.EphemeralResultData for a *tfprotov5.DynamicValue and
// fwschema.Schema.
func EphemeralResultData(ctx context.Context, proto5DynamicValue *tfprotov5.DynamicValue, schema fwschema.Schema) (*tfsdk.EphemeralResultData, diag.Diagnostics) {
	if proto5DynamicValue == nil {
		return nil, nil
	}

	var diags diag.Diagnostics

	// Panic prevention here to simplify the calling implementations.
	// This should not happen, but just in case.
	if schema == nil {
		diags.AddError(
			"Unable to Convert Ephemeral Result Data",
			"An unexpected error was encountered when converting the ephemeral result data from the protocol type. "+
				"This is always an issue in terraform-plugin-framework used to implement the provider and should be reported to the provider developers.\n\n"+
				"Please report this to the provider developer:\n\n"+
				"Missing schema.",
		)

		return nil, diags
	}

	data, dynamicValueDiags := DynamicValue(ctx, proto5DynamicValue, schema, fwschemadata.DataDescriptionEphemeralResultData)

	diags.Append(dynamicValueDiags...)

	if diags.HasError() {
		return nil, diags
	}

	fw := &tfsdk.EphemeralResultData{
		Raw:    data.TerraformValue,
		Schema: schema,
	}

	return fw, diags
}

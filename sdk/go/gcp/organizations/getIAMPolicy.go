// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Generates an IAM policy document that may be referenced by and applied to
// other Google Cloud Platform resources, such as the `google_project` resource.
// 
// 
// This data source is used to define IAM policies to apply to other resources.
// Currently, defining a policy through a datasource and referencing that policy
// from another resource is the only way to apply an IAM policy to a resource.
// 
// **Note:** Several restrictions apply when setting IAM policies through this API.
// See the [setIamPolicy docs](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy)
// for a list of these restrictions.
func LookupIAMPolicy(ctx *pulumi.Context, args *GetIAMPolicyArgs) (*GetIAMPolicyResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["auditConfigs"] = args.AuditConfigs
		inputs["bindings"] = args.Bindings
	}
	outputs, err := ctx.Invoke("gcp:organizations/getIAMPolicy:getIAMPolicy", inputs)
	if err != nil {
		return nil, err
	}
	return &GetIAMPolicyResult{
		PolicyData: outputs["policyData"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getIAMPolicy.
type GetIAMPolicyArgs struct {
	// A nested configuration block that defines logging additional configuration for your project.
	AuditConfigs interface{}
	// A nested configuration block (described below)
	// defining a binding to be included in the policy document. Multiple
	// `binding` arguments are supported.
	Bindings interface{}
}

// A collection of values returned by getIAMPolicy.
type GetIAMPolicyResult struct {
	// The above bindings serialized in a format suitable for
	// referencing from a resource that supports IAM.
	PolicyData interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}

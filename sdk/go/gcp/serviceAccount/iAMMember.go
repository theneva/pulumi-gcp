// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package serviceAccount

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// When managing IAM roles, you can treat a service account either as a resource or as an identity. This resource is to add iam policy bindings to a service account resource **to configure permissions for who can edit the service account**. To configure permissions for a service account to act as an identity that can manage other GCP resources, use the googleProjectIam set of resources.
// 
// Three different resources help you manage your IAM policy for a service account. Each of these resources serves a different use case:
// 
// * `serviceAccount.IAMPolicy`: Authoritative. Sets the IAM policy for the service account and replaces any existing policy already attached.
// * `serviceAccount.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service account are preserved.
// * `serviceAccount.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service account are preserved.
// 
// > **Note:** `serviceAccount.IAMPolicy` **cannot** be used in conjunction with `serviceAccount.IAMBinding` and `serviceAccount.IAMMember` or they will fight over what your policy should be.
// 
// > **Note:** `serviceAccount.IAMBinding` resources **can be** used in conjunction with `serviceAccount.IAMMember` resources **only if** they do not grant privilege to the same role.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/service_account_iam_member.html.markdown.
type IAMMember struct {
	s *pulumi.ResourceState
}

// NewIAMMember registers a new resource with the given unique name, arguments, and options.
func NewIAMMember(ctx *pulumi.Context,
	name string, args *IAMMemberArgs, opts ...pulumi.ResourceOpt) (*IAMMember, error) {
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.ServiceAccountId == nil {
		return nil, errors.New("missing required argument 'ServiceAccountId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["condition"] = nil
		inputs["member"] = nil
		inputs["role"] = nil
		inputs["serviceAccountId"] = nil
	} else {
		inputs["condition"] = args.Condition
		inputs["member"] = args.Member
		inputs["role"] = args.Role
		inputs["serviceAccountId"] = args.ServiceAccountId
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:serviceAccount/iAMMember:IAMMember", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IAMMember{s: s}, nil
}

// GetIAMMember gets an existing IAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMMember(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IAMMemberState, opts ...pulumi.ResourceOpt) (*IAMMember, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["condition"] = state.Condition
		inputs["etag"] = state.Etag
		inputs["member"] = state.Member
		inputs["role"] = state.Role
		inputs["serviceAccountId"] = state.ServiceAccountId
	}
	s, err := ctx.ReadResource("gcp:serviceAccount/iAMMember:IAMMember", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IAMMember{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *IAMMember) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *IAMMember) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
// Structure is documented below.
func (r *IAMMember) Condition() *pulumi.Output {
	return r.s.State["condition"]
}

// (Computed) The etag of the service account IAM policy.
func (r *IAMMember) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

func (r *IAMMember) Member() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["member"])
}

// The role that should be applied. Only one
// `serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *IAMMember) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// The fully-qualified name of the service account to apply policy to.
func (r *IAMMember) ServiceAccountId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serviceAccountId"])
}

// Input properties used for looking up and filtering IAMMember resources.
type IAMMemberState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition interface{}
	// (Computed) The etag of the service account IAM policy.
	Etag interface{}
	Member interface{}
	// The role that should be applied. Only one
	// `serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId interface{}
}

// The set of arguments for constructing a IAMMember resource.
type IAMMemberArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition interface{}
	Member interface{}
	// The role that should be applied. Only one
	// `serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId interface{}
}

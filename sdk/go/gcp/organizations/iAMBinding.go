// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows creation and management of a single binding within IAM policy for
// an existing Google Cloud Platform Organization.
// 
// > **Note:** This resource __must not__ be used in conjunction with
//    `google_organization_iam_member` for the __same role__ or they will fight over
//    what your policy should be.
type IAMBinding struct {
	s *pulumi.ResourceState
}

// NewIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewIAMBinding(ctx *pulumi.Context,
	name string, args *IAMBindingArgs, opts ...pulumi.ResourceOpt) (*IAMBinding, error) {
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.OrgId == nil {
		return nil, errors.New("missing required argument 'OrgId'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["members"] = nil
		inputs["orgId"] = nil
		inputs["role"] = nil
	} else {
		inputs["members"] = args.Members
		inputs["orgId"] = args.OrgId
		inputs["role"] = args.Role
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:organizations/iAMBinding:IAMBinding", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IAMBinding{s: s}, nil
}

// GetIAMBinding gets an existing IAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IAMBindingState, opts ...pulumi.ResourceOpt) (*IAMBinding, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["etag"] = state.Etag
		inputs["members"] = state.Members
		inputs["orgId"] = state.OrgId
		inputs["role"] = state.Role
	}
	s, err := ctx.ReadResource("gcp:organizations/iAMBinding:IAMBinding", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IAMBinding{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *IAMBinding) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *IAMBinding) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// (Computed) The etag of the organization's IAM policy.
func (r *IAMBinding) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// A list of users that the role should apply to.
func (r *IAMBinding) Members() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["members"])
}

// The numeric ID of the organization in which you want to create a custom role.
func (r *IAMBinding) OrgId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["orgId"])
}

// The role that should be applied. Only one
// `google_organization_iam_binding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *IAMBinding) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// Input properties used for looking up and filtering IAMBinding resources.
type IAMBindingState struct {
	// (Computed) The etag of the organization's IAM policy.
	Etag interface{}
	// A list of users that the role should apply to.
	Members interface{}
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId interface{}
	// The role that should be applied. Only one
	// `google_organization_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}

// The set of arguments for constructing a IAMBinding resource.
type IAMBindingArgs struct {
	// A list of users that the role should apply to.
	Members interface{}
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId interface{}
	// The role that should be applied. Only one
	// `google_organization_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}

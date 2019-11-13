// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Three different resources help you manage IAM policies on dataproc jobs. Each of these resources serves a different use case:
// 
// * `dataproc.JobIAMPolicy`: Authoritative. Sets the IAM policy for the job and replaces any existing policy already attached.
// * `dataproc.JobIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the job are preserved.
// * `dataproc.JobIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the job are preserved.
// 
// > **Note:** `dataproc.JobIAMPolicy` **cannot** be used in conjunction with `dataproc.JobIAMBinding` and `dataproc.JobIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the job as `dataproc.JobIAMPolicy` replaces the entire policy.
// 
// > **Note:** `dataproc.JobIAMBinding` resources **can be** used in conjunction with `dataproc.JobIAMMember` resources **only if** they do not grant privilege to the same role.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dataproc_job_iam_member.html.markdown.
type JobIAMMember struct {
	s *pulumi.ResourceState
}

// NewJobIAMMember registers a new resource with the given unique name, arguments, and options.
func NewJobIAMMember(ctx *pulumi.Context,
	name string, args *JobIAMMemberArgs, opts ...pulumi.ResourceOpt) (*JobIAMMember, error) {
	if args == nil || args.JobId == nil {
		return nil, errors.New("missing required argument 'JobId'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["condition"] = nil
		inputs["jobId"] = nil
		inputs["member"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["role"] = nil
	} else {
		inputs["condition"] = args.Condition
		inputs["jobId"] = args.JobId
		inputs["member"] = args.Member
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["role"] = args.Role
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:dataproc/jobIAMMember:JobIAMMember", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &JobIAMMember{s: s}, nil
}

// GetJobIAMMember gets an existing JobIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobIAMMember(ctx *pulumi.Context,
	name string, id pulumi.ID, state *JobIAMMemberState, opts ...pulumi.ResourceOpt) (*JobIAMMember, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["condition"] = state.Condition
		inputs["etag"] = state.Etag
		inputs["jobId"] = state.JobId
		inputs["member"] = state.Member
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["role"] = state.Role
	}
	s, err := ctx.ReadResource("gcp:dataproc/jobIAMMember:JobIAMMember", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &JobIAMMember{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *JobIAMMember) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *JobIAMMember) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *JobIAMMember) Condition() *pulumi.Output {
	return r.s.State["condition"]
}

// (Computed) The etag of the jobs's IAM policy.
func (r *JobIAMMember) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

func (r *JobIAMMember) JobId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["jobId"])
}

func (r *JobIAMMember) Member() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["member"])
}

// The project in which the job belongs. If it
// is not provided, this provider will use the provider default.
func (r *JobIAMMember) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The region in which the job belongs. If it
// is not provided, this provider will use the provider default.
func (r *JobIAMMember) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The role that should be applied. Only one
// `dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *JobIAMMember) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// Input properties used for looking up and filtering JobIAMMember resources.
type JobIAMMemberState struct {
	Condition interface{}
	// (Computed) The etag of the jobs's IAM policy.
	Etag interface{}
	JobId interface{}
	Member interface{}
	// The project in which the job belongs. If it
	// is not provided, this provider will use the provider default.
	Project interface{}
	// The region in which the job belongs. If it
	// is not provided, this provider will use the provider default.
	Region interface{}
	// The role that should be applied. Only one
	// `dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}

// The set of arguments for constructing a JobIAMMember resource.
type JobIAMMemberArgs struct {
	Condition interface{}
	JobId interface{}
	Member interface{}
	// The project in which the job belongs. If it
	// is not provided, this provider will use the provider default.
	Project interface{}
	// The region in which the job belongs. If it
	// is not provided, this provider will use the provider default.
	Region interface{}
	// The role that should be applied. Only one
	// `dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}

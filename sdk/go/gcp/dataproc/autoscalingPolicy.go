// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dataproc_autoscaling_policy.html.markdown.
type AutoscalingPolicy struct {
	s *pulumi.ResourceState
}

// NewAutoscalingPolicy registers a new resource with the given unique name, arguments, and options.
func NewAutoscalingPolicy(ctx *pulumi.Context,
	name string, args *AutoscalingPolicyArgs, opts ...pulumi.ResourceOpt) (*AutoscalingPolicy, error) {
	if args == nil || args.PolicyId == nil {
		return nil, errors.New("missing required argument 'PolicyId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["basicAlgorithm"] = nil
		inputs["location"] = nil
		inputs["policyId"] = nil
		inputs["project"] = nil
		inputs["secondaryWorkerConfig"] = nil
		inputs["workerConfig"] = nil
	} else {
		inputs["basicAlgorithm"] = args.BasicAlgorithm
		inputs["location"] = args.Location
		inputs["policyId"] = args.PolicyId
		inputs["project"] = args.Project
		inputs["secondaryWorkerConfig"] = args.SecondaryWorkerConfig
		inputs["workerConfig"] = args.WorkerConfig
	}
	inputs["name"] = nil
	s, err := ctx.RegisterResource("gcp:dataproc/autoscalingPolicy:AutoscalingPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AutoscalingPolicy{s: s}, nil
}

// GetAutoscalingPolicy gets an existing AutoscalingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoscalingPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AutoscalingPolicyState, opts ...pulumi.ResourceOpt) (*AutoscalingPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["basicAlgorithm"] = state.BasicAlgorithm
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["policyId"] = state.PolicyId
		inputs["project"] = state.Project
		inputs["secondaryWorkerConfig"] = state.SecondaryWorkerConfig
		inputs["workerConfig"] = state.WorkerConfig
	}
	s, err := ctx.ReadResource("gcp:dataproc/autoscalingPolicy:AutoscalingPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AutoscalingPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AutoscalingPolicy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AutoscalingPolicy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Basic algorithm for autoscaling.
func (r *AutoscalingPolicy) BasicAlgorithm() *pulumi.Output {
	return r.s.State["basicAlgorithm"]
}

// The location where the autoscaling poicy should reside. The default value is 'global'.
func (r *AutoscalingPolicy) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// The "resource name" of the autoscaling policy.
func (r *AutoscalingPolicy) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot
// begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
func (r *AutoscalingPolicy) PolicyId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policyId"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *AutoscalingPolicy) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// Describes how the autoscaler will operate for secondary workers.
func (r *AutoscalingPolicy) SecondaryWorkerConfig() *pulumi.Output {
	return r.s.State["secondaryWorkerConfig"]
}

// Describes how the autoscaler will operate for primary workers.
func (r *AutoscalingPolicy) WorkerConfig() *pulumi.Output {
	return r.s.State["workerConfig"]
}

// Input properties used for looking up and filtering AutoscalingPolicy resources.
type AutoscalingPolicyState struct {
	// Basic algorithm for autoscaling.
	BasicAlgorithm interface{}
	// The location where the autoscaling poicy should reside. The default value is 'global'.
	Location interface{}
	// The "resource name" of the autoscaling policy.
	Name interface{}
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot
	// begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
	PolicyId interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// Describes how the autoscaler will operate for secondary workers.
	SecondaryWorkerConfig interface{}
	// Describes how the autoscaler will operate for primary workers.
	WorkerConfig interface{}
}

// The set of arguments for constructing a AutoscalingPolicy resource.
type AutoscalingPolicyArgs struct {
	// Basic algorithm for autoscaling.
	BasicAlgorithm interface{}
	// The location where the autoscaling poicy should reside. The default value is 'global'.
	Location interface{}
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot
	// begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
	PolicyId interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// Describes how the autoscaler will operate for secondary workers.
	SecondaryWorkerConfig interface{}
	// Describes how the autoscaler will operate for primary workers.
	WorkerConfig interface{}
}

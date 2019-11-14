// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package filestore

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/filestore_instance.html.markdown.
type Instance struct {
	s *pulumi.ResourceState
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOpt) (*Instance, error) {
	if args == nil || args.FileShares == nil {
		return nil, errors.New("missing required argument 'FileShares'")
	}
	if args == nil || args.Networks == nil {
		return nil, errors.New("missing required argument 'Networks'")
	}
	if args == nil || args.Tier == nil {
		return nil, errors.New("missing required argument 'Tier'")
	}
	if args == nil || args.Zone == nil {
		return nil, errors.New("missing required argument 'Zone'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["fileShares"] = nil
		inputs["labels"] = nil
		inputs["name"] = nil
		inputs["networks"] = nil
		inputs["project"] = nil
		inputs["tier"] = nil
		inputs["zone"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["fileShares"] = args.FileShares
		inputs["labels"] = args.Labels
		inputs["name"] = args.Name
		inputs["networks"] = args.Networks
		inputs["project"] = args.Project
		inputs["tier"] = args.Tier
		inputs["zone"] = args.Zone
	}
	inputs["createTime"] = nil
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:filestore/instance:Instance", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Instance{s: s}, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.ID, state *InstanceState, opts ...pulumi.ResourceOpt) (*Instance, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["createTime"] = state.CreateTime
		inputs["description"] = state.Description
		inputs["etag"] = state.Etag
		inputs["fileShares"] = state.FileShares
		inputs["labels"] = state.Labels
		inputs["name"] = state.Name
		inputs["networks"] = state.Networks
		inputs["project"] = state.Project
		inputs["tier"] = state.Tier
		inputs["zone"] = state.Zone
	}
	s, err := ctx.ReadResource("gcp:filestore/instance:Instance", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Instance{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Instance) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Instance) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Creation timestamp in RFC3339 text format.
func (r *Instance) CreateTime() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["createTime"])
}

// A description of the instance.
func (r *Instance) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
func (r *Instance) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// File system shares on the instance. For this version, only a single file share is supported.
func (r *Instance) FileShares() *pulumi.Output {
	return r.s.State["fileShares"]
}

// Resource labels to represent user-provided metadata.
func (r *Instance) Labels() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["labels"])
}

// The resource name of the instance.
func (r *Instance) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// VPC networks to which the instance is connected. For this version, only a single network is supported.
func (r *Instance) Networks() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["networks"])
}

func (r *Instance) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The service tier of the instance.
func (r *Instance) Tier() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tier"])
}

// The name of the Filestore zone of the instance.
func (r *Instance) Zone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zone"])
}

// Input properties used for looking up and filtering Instance resources.
type InstanceState struct {
	// Creation timestamp in RFC3339 text format.
	CreateTime interface{}
	// A description of the instance.
	Description interface{}
	// Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
	Etag interface{}
	// File system shares on the instance. For this version, only a single file share is supported.
	FileShares interface{}
	// Resource labels to represent user-provided metadata.
	Labels interface{}
	// The resource name of the instance.
	Name interface{}
	// VPC networks to which the instance is connected. For this version, only a single network is supported.
	Networks interface{}
	Project interface{}
	// The service tier of the instance.
	Tier interface{}
	// The name of the Filestore zone of the instance.
	Zone interface{}
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// A description of the instance.
	Description interface{}
	// File system shares on the instance. For this version, only a single file share is supported.
	FileShares interface{}
	// Resource labels to represent user-provided metadata.
	Labels interface{}
	// The resource name of the instance.
	Name interface{}
	// VPC networks to which the instance is connected. For this version, only a single network is supported.
	Networks interface{}
	Project interface{}
	// The service tier of the instance.
	Tier interface{}
	// The name of the Filestore zone of the instance.
	Zone interface{}
}

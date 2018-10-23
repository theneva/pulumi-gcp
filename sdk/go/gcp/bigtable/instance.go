// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigtable

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a Google Bigtable instance. For more information see
// [the official documentation](https://cloud.google.com/bigtable/) and
// [API](https://cloud.google.com/bigtable/docs/go/reference).
// 
type Instance struct {
	s *pulumi.ResourceState
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOpt) (*Instance, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["cluster"] = nil
		inputs["clusterId"] = nil
		inputs["displayName"] = nil
		inputs["instanceType"] = nil
		inputs["name"] = nil
		inputs["numNodes"] = nil
		inputs["project"] = nil
		inputs["storageType"] = nil
		inputs["zone"] = nil
	} else {
		inputs["cluster"] = args.Cluster
		inputs["clusterId"] = args.ClusterId
		inputs["displayName"] = args.DisplayName
		inputs["instanceType"] = args.InstanceType
		inputs["name"] = args.Name
		inputs["numNodes"] = args.NumNodes
		inputs["project"] = args.Project
		inputs["storageType"] = args.StorageType
		inputs["zone"] = args.Zone
	}
	s, err := ctx.RegisterResource("gcp:bigtable/instance:Instance", name, true, inputs, opts...)
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
		inputs["cluster"] = state.Cluster
		inputs["clusterId"] = state.ClusterId
		inputs["displayName"] = state.DisplayName
		inputs["instanceType"] = state.InstanceType
		inputs["name"] = state.Name
		inputs["numNodes"] = state.NumNodes
		inputs["project"] = state.Project
		inputs["storageType"] = state.StorageType
		inputs["zone"] = state.Zone
	}
	s, err := ctx.ReadResource("gcp:bigtable/instance:Instance", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Instance{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Instance) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Instance) ID() *pulumi.IDOutput {
	return r.s.ID
}

// A block of cluster configuration options. Either `cluster` or `cluster_id` must be used. Only one cluster may be specified. See structure below.
func (r *Instance) Cluster() *pulumi.Output {
	return r.s.State["cluster"]
}

// The ID of the Cloud Bigtable cluster.
func (r *Instance) ClusterId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clusterId"])
}

// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
func (r *Instance) DisplayName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["displayName"])
}

// The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
func (r *Instance) InstanceType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instanceType"])
}

// The name of the Cloud Bigtable instance.
func (r *Instance) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The number of nodes in your Cloud Bigtable cluster. Minimum of `3` for a `PRODUCTION` instance. Cannot be set for a `DEVELOPMENT` instance.
func (r *Instance) NumNodes() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["numNodes"])
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *Instance) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The storage type to use. One of `"SSD"` or `"HDD"`. Defaults to `"SSD"`.
func (r *Instance) StorageType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["storageType"])
}

// The zone to create the Cloud Bigtable cluster in. Zones that support Bigtable instances are noted on the [Cloud Bigtable locations page](https://cloud.google.com/bigtable/docs/locations).
func (r *Instance) Zone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zone"])
}

// Input properties used for looking up and filtering Instance resources.
type InstanceState struct {
	// A block of cluster configuration options. Either `cluster` or `cluster_id` must be used. Only one cluster may be specified. See structure below.
	Cluster interface{}
	// The ID of the Cloud Bigtable cluster.
	ClusterId interface{}
	// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
	DisplayName interface{}
	// The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
	InstanceType interface{}
	// The name of the Cloud Bigtable instance.
	Name interface{}
	// The number of nodes in your Cloud Bigtable cluster. Minimum of `3` for a `PRODUCTION` instance. Cannot be set for a `DEVELOPMENT` instance.
	NumNodes interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The storage type to use. One of `"SSD"` or `"HDD"`. Defaults to `"SSD"`.
	StorageType interface{}
	// The zone to create the Cloud Bigtable cluster in. Zones that support Bigtable instances are noted on the [Cloud Bigtable locations page](https://cloud.google.com/bigtable/docs/locations).
	Zone interface{}
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// A block of cluster configuration options. Either `cluster` or `cluster_id` must be used. Only one cluster may be specified. See structure below.
	Cluster interface{}
	// The ID of the Cloud Bigtable cluster.
	ClusterId interface{}
	// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
	DisplayName interface{}
	// The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
	InstanceType interface{}
	// The name of the Cloud Bigtable instance.
	Name interface{}
	// The number of nodes in your Cloud Bigtable cluster. Minimum of `3` for a `PRODUCTION` instance. Cannot be set for a `DEVELOPMENT` instance.
	NumNodes interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The storage type to use. One of `"SSD"` or `"HDD"`. Defaults to `"SSD"`.
	StorageType interface{}
	// The zone to create the Cloud Bigtable cluster in. Zones that support Bigtable instances are noted on the [Cloud Bigtable locations page](https://cloud.google.com/bigtable/docs/locations).
	Zone interface{}
}

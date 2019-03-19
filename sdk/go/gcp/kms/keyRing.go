// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows creation of a Google Cloud Platform KMS KeyRing. For more information see
// [the official documentation](https://cloud.google.com/kms/docs/object-hierarchy#key_ring)
// and
// [API](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings).
// 
// A KeyRing is a grouping of CryptoKeys for organizational purposes. A KeyRing belongs to a Google Cloud Platform Project
// and resides in a specific location.
// 
// > Note: KeyRings cannot be deleted from Google Cloud Platform. Destroying a Terraform-managed KeyRing will remove it
// from state but **will not delete the resource on the server**.
type KeyRing struct {
	s *pulumi.ResourceState
}

// NewKeyRing registers a new resource with the given unique name, arguments, and options.
func NewKeyRing(ctx *pulumi.Context,
	name string, args *KeyRingArgs, opts ...pulumi.ResourceOpt) (*KeyRing, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
	} else {
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["project"] = args.Project
	}
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:kms/keyRing:KeyRing", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &KeyRing{s: s}, nil
}

// GetKeyRing gets an existing KeyRing resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyRing(ctx *pulumi.Context,
	name string, id pulumi.ID, state *KeyRingState, opts ...pulumi.ResourceOpt) (*KeyRing, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["selfLink"] = state.SelfLink
	}
	s, err := ctx.ReadResource("gcp:kms/keyRing:KeyRing", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &KeyRing{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *KeyRing) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *KeyRing) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The Google Cloud Platform location for the KeyRing.
// A full list of valid locations can be found by running `gcloud kms locations list`.
func (r *KeyRing) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// The KeyRing's name.
// A KeyRing’s name must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`
func (r *KeyRing) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *KeyRing) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The self link of the created KeyRing. Its format is `projects/{projectId}/locations/{location}/keyRings/{keyRingName}`.
func (r *KeyRing) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

// Input properties used for looking up and filtering KeyRing resources.
type KeyRingState struct {
	// The Google Cloud Platform location for the KeyRing.
	// A full list of valid locations can be found by running `gcloud kms locations list`.
	Location interface{}
	// The KeyRing's name.
	// A KeyRing’s name must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`
	Name interface{}
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The self link of the created KeyRing. Its format is `projects/{projectId}/locations/{location}/keyRings/{keyRingName}`.
	SelfLink interface{}
}

// The set of arguments for constructing a KeyRing resource.
type KeyRingArgs struct {
	// The Google Cloud Platform location for the KeyRing.
	// A full list of valid locations can be found by running `gcloud kms locations list`.
	Location interface{}
	// The KeyRing's name.
	// A KeyRing’s name must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`
	Name interface{}
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
}

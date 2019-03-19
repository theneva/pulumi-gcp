// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Authoritatively manages the default object ACLs for a Google Cloud Storage bucket
// without managing the bucket itself.
// 
// > Note that for each object, its creator will have the `"OWNER"` role in addition
// to the default ACL that has been defined.
// 
// For more information see
// [the official documentation](https://cloud.google.com/storage/docs/access-control/lists) 
// and 
// [API](https://cloud.google.com/storage/docs/json_api/v1/defaultObjectAccessControls).
// 
// > Want fine-grained control over default object ACLs? Use `google_storage_default_object_access_control`
// to control individual role entity pairs.
type DefaultObjectACL struct {
	s *pulumi.ResourceState
}

// NewDefaultObjectACL registers a new resource with the given unique name, arguments, and options.
func NewDefaultObjectACL(ctx *pulumi.Context,
	name string, args *DefaultObjectACLArgs, opts ...pulumi.ResourceOpt) (*DefaultObjectACL, error) {
	if args == nil || args.Bucket == nil {
		return nil, errors.New("missing required argument 'Bucket'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["bucket"] = nil
		inputs["roleEntities"] = nil
	} else {
		inputs["bucket"] = args.Bucket
		inputs["roleEntities"] = args.RoleEntities
	}
	s, err := ctx.RegisterResource("gcp:storage/defaultObjectACL:DefaultObjectACL", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DefaultObjectACL{s: s}, nil
}

// GetDefaultObjectACL gets an existing DefaultObjectACL resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultObjectACL(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DefaultObjectACLState, opts ...pulumi.ResourceOpt) (*DefaultObjectACL, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["bucket"] = state.Bucket
		inputs["roleEntities"] = state.RoleEntities
	}
	s, err := ctx.ReadResource("gcp:storage/defaultObjectACL:DefaultObjectACL", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DefaultObjectACL{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DefaultObjectACL) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DefaultObjectACL) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The name of the bucket it applies to.
func (r *DefaultObjectACL) Bucket() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["bucket"])
}

// List of role/entity pairs in the form `ROLE:entity`.
// See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
// Omitting the field is the same as providing an empty list.
func (r *DefaultObjectACL) RoleEntities() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["roleEntities"])
}

// Input properties used for looking up and filtering DefaultObjectACL resources.
type DefaultObjectACLState struct {
	// The name of the bucket it applies to.
	Bucket interface{}
	// List of role/entity pairs in the form `ROLE:entity`.
	// See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
	// Omitting the field is the same as providing an empty list.
	RoleEntities interface{}
}

// The set of arguments for constructing a DefaultObjectACL resource.
type DefaultObjectACLArgs struct {
	// The name of the bucket it applies to.
	Bucket interface{}
	// List of role/entity pairs in the form `ROLE:entity`.
	// See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
	// Omitting the field is the same as providing an empty list.
	RoleEntities interface{}
}

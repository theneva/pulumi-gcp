// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package storage

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The ObjectAccessControls resources represent the Access Control Lists
// (ACLs) for objects within Google Cloud Storage. ACLs let you specify
// who has access to your data and to what extent.
// 
// There are two roles that can be assigned to an entity:
// 
// READERs can get an object, though the acl property will not be revealed.
// OWNERs are READERs, and they can get the acl property, update an object,
// and call all objectAccessControls methods on the object. The owner of an
// object is always an OWNER.
// For more information, see Access Control, with the caveat that this API
// uses READER and OWNER instead of READ and FULL_CONTROL.
// 
// 
// To get more information about ObjectAccessControl, see:
// 
// * [API documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/storage/docs/access-control/create-manage-lists)
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/storage_object_access_control.html.markdown.
type ObjectAccessControl struct {
	pulumi.CustomResourceState

	// The name of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The domain associated with the entity.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The email address associated with the entity.
	Email pulumi.StringOutput `pulumi:"email"`
	// The entity holding the permission, in one of the following forms: * user-{{userId}} * user-{{email}} (such as
	// "user-liz@example.com") * group-{{groupId}} * group-{{email}} (such as "group-example@googlegroups.com") *
	// domain-{{domain}} (such as "domain-example.com") * project-team-{{projectId}} * allUsers * allAuthenticatedUsers
	Entity pulumi.StringOutput `pulumi:"entity"`
	// The ID for the entity
	EntityId pulumi.StringOutput `pulumi:"entityId"`
	// The content generation of the object, if applied to an object.
	Generation pulumi.IntOutput `pulumi:"generation"`
	// The name of the object to apply the access control to.
	Object pulumi.StringOutput `pulumi:"object"`
	// The project team associated with the entity
	ProjectTeam ObjectAccessControlProjectTeamOutput `pulumi:"projectTeam"`
	// The access permission for the entity.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewObjectAccessControl registers a new resource with the given unique name, arguments, and options.
func NewObjectAccessControl(ctx *pulumi.Context,
	name string, args *ObjectAccessControlArgs, opts ...pulumi.ResourceOption) (*ObjectAccessControl, error) {
	if args == nil || args.Bucket == nil {
		return nil, errors.New("missing required argument 'Bucket'")
	}
	if args == nil || args.Entity == nil {
		return nil, errors.New("missing required argument 'Entity'")
	}
	if args == nil || args.Object == nil {
		return nil, errors.New("missing required argument 'Object'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &ObjectAccessControlArgs{}
	}
	var resource ObjectAccessControl
	err := ctx.RegisterResource("gcp:storage/objectAccessControl:ObjectAccessControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectAccessControl gets an existing ObjectAccessControl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectAccessControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectAccessControlState, opts ...pulumi.ResourceOption) (*ObjectAccessControl, error) {
	var resource ObjectAccessControl
	err := ctx.ReadResource("gcp:storage/objectAccessControl:ObjectAccessControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectAccessControl resources.
type objectAccessControlState struct {
	// The name of the bucket.
	Bucket *string `pulumi:"bucket"`
	// The domain associated with the entity.
	Domain *string `pulumi:"domain"`
	// The email address associated with the entity.
	Email *string `pulumi:"email"`
	// The entity holding the permission, in one of the following forms: * user-{{userId}} * user-{{email}} (such as
	// "user-liz@example.com") * group-{{groupId}} * group-{{email}} (such as "group-example@googlegroups.com") *
	// domain-{{domain}} (such as "domain-example.com") * project-team-{{projectId}} * allUsers * allAuthenticatedUsers
	Entity *string `pulumi:"entity"`
	// The ID for the entity
	EntityId *string `pulumi:"entityId"`
	// The content generation of the object, if applied to an object.
	Generation *int `pulumi:"generation"`
	// The name of the object to apply the access control to.
	Object *string `pulumi:"object"`
	// The project team associated with the entity
	ProjectTeam *ObjectAccessControlProjectTeam `pulumi:"projectTeam"`
	// The access permission for the entity.
	Role *string `pulumi:"role"`
}

type ObjectAccessControlState struct {
	// The name of the bucket.
	Bucket pulumi.StringPtrInput
	// The domain associated with the entity.
	Domain pulumi.StringPtrInput
	// The email address associated with the entity.
	Email pulumi.StringPtrInput
	// The entity holding the permission, in one of the following forms: * user-{{userId}} * user-{{email}} (such as
	// "user-liz@example.com") * group-{{groupId}} * group-{{email}} (such as "group-example@googlegroups.com") *
	// domain-{{domain}} (such as "domain-example.com") * project-team-{{projectId}} * allUsers * allAuthenticatedUsers
	Entity pulumi.StringPtrInput
	// The ID for the entity
	EntityId pulumi.StringPtrInput
	// The content generation of the object, if applied to an object.
	Generation pulumi.IntPtrInput
	// The name of the object to apply the access control to.
	Object pulumi.StringPtrInput
	// The project team associated with the entity
	ProjectTeam ObjectAccessControlProjectTeamPtrInput
	// The access permission for the entity.
	Role pulumi.StringPtrInput
}

func (ObjectAccessControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectAccessControlState)(nil)).Elem()
}

type objectAccessControlArgs struct {
	// The name of the bucket.
	Bucket string `pulumi:"bucket"`
	// The entity holding the permission, in one of the following forms: * user-{{userId}} * user-{{email}} (such as
	// "user-liz@example.com") * group-{{groupId}} * group-{{email}} (such as "group-example@googlegroups.com") *
	// domain-{{domain}} (such as "domain-example.com") * project-team-{{projectId}} * allUsers * allAuthenticatedUsers
	Entity string `pulumi:"entity"`
	// The name of the object to apply the access control to.
	Object string `pulumi:"object"`
	// The access permission for the entity.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ObjectAccessControl resource.
type ObjectAccessControlArgs struct {
	// The name of the bucket.
	Bucket pulumi.StringInput
	// The entity holding the permission, in one of the following forms: * user-{{userId}} * user-{{email}} (such as
	// "user-liz@example.com") * group-{{groupId}} * group-{{email}} (such as "group-example@googlegroups.com") *
	// domain-{{domain}} (such as "domain-example.com") * project-team-{{projectId}} * allUsers * allAuthenticatedUsers
	Entity pulumi.StringInput
	// The name of the object to apply the access control to.
	Object pulumi.StringInput
	// The access permission for the entity.
	Role pulumi.StringInput
}

func (ObjectAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectAccessControlArgs)(nil)).Elem()
}


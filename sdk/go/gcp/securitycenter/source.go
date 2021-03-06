// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package securitycenter

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// A Cloud Security Command Center's (Cloud SCC) finding source. A finding
// source is an entity or a mechanism that can produce a finding. A source is
// like a container of findings that come from the same scanner, logger,
// monitor, etc.
// 
// 
// To get more information about Source, see:
// 
// * [API documentation](https://cloud.google.com/security-command-center/docs/reference/rest/v1beta1/organizations.sources)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/binary-authorization/)
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/scc_source.html.markdown.
type Source struct {
	pulumi.CustomResourceState

	// The description of the source (max of 1024 characters).
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The source’s display name. A source’s display name must be unique amongst its siblings, for example, two sources
	// with the same parent can't share the same display name. The display name must start and end with a letter or digit, may
	// contain letters, digits, spaces, hyphens, and underscores, and can be no longer than 32 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The resource name of this source, in the format 'organizations/{{organization}}/sources/{{source}}'.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization whose Cloud Security Command Center the Source lives in.
	Organization pulumi.StringOutput `pulumi:"organization"`
}

// NewSource registers a new resource with the given unique name, arguments, and options.
func NewSource(ctx *pulumi.Context,
	name string, args *SourceArgs, opts ...pulumi.ResourceOption) (*Source, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.Organization == nil {
		return nil, errors.New("missing required argument 'Organization'")
	}
	if args == nil {
		args = &SourceArgs{}
	}
	var resource Source
	err := ctx.RegisterResource("gcp:securitycenter/source:Source", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSource gets an existing Source resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceState, opts ...pulumi.ResourceOption) (*Source, error) {
	var resource Source
	err := ctx.ReadResource("gcp:securitycenter/source:Source", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Source resources.
type sourceState struct {
	// The description of the source (max of 1024 characters).
	Description *string `pulumi:"description"`
	// The source’s display name. A source’s display name must be unique amongst its siblings, for example, two sources
	// with the same parent can't share the same display name. The display name must start and end with a letter or digit, may
	// contain letters, digits, spaces, hyphens, and underscores, and can be no longer than 32 characters.
	DisplayName *string `pulumi:"displayName"`
	// The resource name of this source, in the format 'organizations/{{organization}}/sources/{{source}}'.
	Name *string `pulumi:"name"`
	// The organization whose Cloud Security Command Center the Source lives in.
	Organization *string `pulumi:"organization"`
}

type SourceState struct {
	// The description of the source (max of 1024 characters).
	Description pulumi.StringPtrInput
	// The source’s display name. A source’s display name must be unique amongst its siblings, for example, two sources
	// with the same parent can't share the same display name. The display name must start and end with a letter or digit, may
	// contain letters, digits, spaces, hyphens, and underscores, and can be no longer than 32 characters.
	DisplayName pulumi.StringPtrInput
	// The resource name of this source, in the format 'organizations/{{organization}}/sources/{{source}}'.
	Name pulumi.StringPtrInput
	// The organization whose Cloud Security Command Center the Source lives in.
	Organization pulumi.StringPtrInput
}

func (SourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceState)(nil)).Elem()
}

type sourceArgs struct {
	// The description of the source (max of 1024 characters).
	Description *string `pulumi:"description"`
	// The source’s display name. A source’s display name must be unique amongst its siblings, for example, two sources
	// with the same parent can't share the same display name. The display name must start and end with a letter or digit, may
	// contain letters, digits, spaces, hyphens, and underscores, and can be no longer than 32 characters.
	DisplayName string `pulumi:"displayName"`
	// The organization whose Cloud Security Command Center the Source lives in.
	Organization string `pulumi:"organization"`
}

// The set of arguments for constructing a Source resource.
type SourceArgs struct {
	// The description of the source (max of 1024 characters).
	Description pulumi.StringPtrInput
	// The source’s display name. A source’s display name must be unique amongst its siblings, for example, two sources
	// with the same parent can't share the same display name. The display name must start and end with a letter or digit, may
	// contain letters, digits, spaces, hyphens, and underscores, and can be no longer than 32 characters.
	DisplayName pulumi.StringInput
	// The organization whose Cloud Security Command Center the Source lives in.
	Organization pulumi.StringInput
}

func (SourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceArgs)(nil)).Elem()
}


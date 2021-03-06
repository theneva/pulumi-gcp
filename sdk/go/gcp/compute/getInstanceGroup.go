// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get a Compute Instance Group within GCE.
// For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/#unmanaged_instance_groups)
// and [API](https://cloud.google.com/compute/docs/reference/latest/instanceGroups)
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/compute_instance_group.html.markdown.
func LookupInstanceGroup(ctx *pulumi.Context, args *LookupInstanceGroupArgs, opts ...pulumi.InvokeOption) (*LookupInstanceGroupResult, error) {
	var rv LookupInstanceGroupResult
	err := ctx.Invoke("gcp:compute/getInstanceGroup:getInstanceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceGroup.
type LookupInstanceGroupArgs struct {
	// The name of the instance group. Either `name` or `selfLink` must be provided.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The self link of the instance group. Either `name` or `selfLink` must be provided.
	SelfLink *string `pulumi:"selfLink"`
	// The zone of the instance group. If referencing the instance group by name
	// and `zone` is not provided, the provider zone is used.
	Zone *string `pulumi:"zone"`
}


// A collection of values returned by getInstanceGroup.
type LookupInstanceGroupResult struct {
	// Textual description of the instance group.
	Description string `pulumi:"description"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of instances in the group.
	Instances []string `pulumi:"instances"`
	Name *string `pulumi:"name"`
	// List of named ports in the group.
	NamedPorts []GetInstanceGroupNamedPort `pulumi:"namedPorts"`
	// The URL of the network the instance group is in.
	Network string `pulumi:"network"`
	Project string `pulumi:"project"`
	// The URI of the resource.
	SelfLink string `pulumi:"selfLink"`
	// The number of instances in the group.
	Size int `pulumi:"size"`
	Zone string `pulumi:"zone"`
}


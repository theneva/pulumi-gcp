// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get information about a VM instance resource within GCE. For more information see
// [the official documentation](https://cloud.google.com/compute/docs/instances)
// and
// [API](https://cloud.google.com/compute/docs/reference/latest/instances).
// 
func LookupInstance(ctx *pulumi.Context, args *GetInstanceArgs) (*GetInstanceResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["zone"] = args.Zone
	}
	outputs, err := ctx.Invoke("gcp:compute/getInstance:getInstance", inputs)
	if err != nil {
		return nil, err
	}
	return &GetInstanceResult{
		AllowStoppingForUpdate: outputs["allowStoppingForUpdate"],
		AttachedDisks: outputs["attachedDisks"],
		BootDisks: outputs["bootDisks"],
		CanIpForward: outputs["canIpForward"],
		CpuPlatform: outputs["cpuPlatform"],
		CreateTimeout: outputs["createTimeout"],
		DeletionProtection: outputs["deletionProtection"],
		Description: outputs["description"],
		Disks: outputs["disks"],
		GuestAccelerators: outputs["guestAccelerators"],
		InstanceId: outputs["instanceId"],
		LabelFingerprint: outputs["labelFingerprint"],
		Labels: outputs["labels"],
		MachineType: outputs["machineType"],
		Metadata: outputs["metadata"],
		MetadataFingerprint: outputs["metadataFingerprint"],
		MetadataStartupScript: outputs["metadataStartupScript"],
		MinCpuPlatform: outputs["minCpuPlatform"],
		Networks: outputs["networks"],
		NetworkInterfaces: outputs["networkInterfaces"],
		Schedulings: outputs["schedulings"],
		ScratchDisks: outputs["scratchDisks"],
		SelfLink: outputs["selfLink"],
		ServiceAccounts: outputs["serviceAccounts"],
		Tags: outputs["tags"],
		TagsFingerprint: outputs["tagsFingerprint"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getInstance.
type GetInstanceArgs struct {
	// The name of the instance. One of `name` or `self_link` must be provided.
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If `self_link` is provided, this value is ignored.  If neither `self_link`
	// nor `project` are provided, the provider project is used.
	Project interface{}
	// The zone of the instance. If `self_link` is provided, this
	// value is ignored.  If neither `self_link` nor `zone` are provided, the
	// provider zone is used.
	Zone interface{}
}

// A collection of values returned by getInstance.
type GetInstanceResult struct {
	AllowStoppingForUpdate interface{}
	// List of disks attached to the instance. Structure is documented below.
	AttachedDisks interface{}
	// The boot disk for the instance. Sructure is documented below.
	BootDisks interface{}
	// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
	CanIpForward interface{}
	// The CPU platform used by this instance.
	CpuPlatform interface{}
	CreateTimeout interface{}
	// Whether deletion protection is enabled on this instance.
	DeletionProtection interface{}
	// A brief description of the resource.
	Description interface{}
	Disks interface{}
	// List of the type and count of accelerator cards attached to the instance. Structure is documented below.
	GuestAccelerators interface{}
	// The server-assigned unique identifier of this instance.
	InstanceId interface{}
	// The unique fingerprint of the labels.
	LabelFingerprint interface{}
	// A set of key/value label pairs assigned to the instance.
	Labels interface{}
	// The machine type to create.
	MachineType interface{}
	// Metadata key/value pairs made available within the instance.
	Metadata interface{}
	// The unique fingerprint of the metadata.
	MetadataFingerprint interface{}
	MetadataStartupScript interface{}
	// The minimum CPU platform specified for the VM instance.
	MinCpuPlatform interface{}
	// The name or self_link of the network attached to this interface.
	Networks interface{}
	// The networks attached to the instance. Structure is documented below.
	NetworkInterfaces interface{}
	// The scheduling strategy being used by the instance.
	Schedulings interface{}
	// The scratch disks attached to the instance. Structure is documented below.
	ScratchDisks interface{}
	// The URI of the created resource.
	SelfLink interface{}
	// The service account to attach to the instance. Structure is documented below.
	ServiceAccounts interface{}
	// The list of tags attached to the instance.
	Tags interface{}
	// The unique fingerprint of the tags.
	TagsFingerprint interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}

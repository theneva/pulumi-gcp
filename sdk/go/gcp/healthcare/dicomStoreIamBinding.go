// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dicom_store_iam_binding.html.markdown.
type DicomStoreIamBinding struct {
	s *pulumi.ResourceState
}

// NewDicomStoreIamBinding registers a new resource with the given unique name, arguments, and options.
func NewDicomStoreIamBinding(ctx *pulumi.Context,
	name string, args *DicomStoreIamBindingArgs, opts ...pulumi.ResourceOpt) (*DicomStoreIamBinding, error) {
	if args == nil || args.DicomStoreId == nil {
		return nil, errors.New("missing required argument 'DicomStoreId'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["condition"] = nil
		inputs["dicomStoreId"] = nil
		inputs["members"] = nil
		inputs["role"] = nil
	} else {
		inputs["condition"] = args.Condition
		inputs["dicomStoreId"] = args.DicomStoreId
		inputs["members"] = args.Members
		inputs["role"] = args.Role
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:healthcare/dicomStoreIamBinding:DicomStoreIamBinding", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DicomStoreIamBinding{s: s}, nil
}

// GetDicomStoreIamBinding gets an existing DicomStoreIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDicomStoreIamBinding(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DicomStoreIamBindingState, opts ...pulumi.ResourceOpt) (*DicomStoreIamBinding, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["condition"] = state.Condition
		inputs["dicomStoreId"] = state.DicomStoreId
		inputs["etag"] = state.Etag
		inputs["members"] = state.Members
		inputs["role"] = state.Role
	}
	s, err := ctx.ReadResource("gcp:healthcare/dicomStoreIamBinding:DicomStoreIamBinding", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DicomStoreIamBinding{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DicomStoreIamBinding) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DicomStoreIamBinding) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *DicomStoreIamBinding) Condition() *pulumi.Output {
	return r.s.State["condition"]
}

// The DICOM store ID, in the form
// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
// project setting will be used as a fallback.
func (r *DicomStoreIamBinding) DicomStoreId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dicomStoreId"])
}

// (Computed) The etag of the DICOM store's IAM policy.
func (r *DicomStoreIamBinding) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

func (r *DicomStoreIamBinding) Members() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["members"])
}

// The role that should be applied. Only one
// `healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *DicomStoreIamBinding) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// Input properties used for looking up and filtering DicomStoreIamBinding resources.
type DicomStoreIamBindingState struct {
	Condition interface{}
	// The DICOM store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
	// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DicomStoreId interface{}
	// (Computed) The etag of the DICOM store's IAM policy.
	Etag interface{}
	Members interface{}
	// The role that should be applied. Only one
	// `healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}

// The set of arguments for constructing a DicomStoreIamBinding resource.
type DicomStoreIamBindingArgs struct {
	Condition interface{}
	// The DICOM store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
	// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DicomStoreId interface{}
	Members interface{}
	// The role that should be applied. Only one
	// `healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}

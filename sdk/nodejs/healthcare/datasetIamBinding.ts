// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dataset_iam_binding.html.markdown.
 */
export class DatasetIamBinding extends pulumi.CustomResource {
    /**
     * Get an existing DatasetIamBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatasetIamBindingState, opts?: pulumi.CustomResourceOptions): DatasetIamBinding {
        return new DatasetIamBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:healthcare/datasetIamBinding:DatasetIamBinding';

    /**
     * Returns true if the given object is an instance of DatasetIamBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatasetIamBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatasetIamBinding.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.healthcare.DatasetIamBindingCondition | undefined>;
    /**
     * The dataset ID, in the form
     * `{project_id}/{location_name}/{dataset_name}` or
     * `{location_name}/{dataset_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    public readonly datasetId!: pulumi.Output<string>;
    /**
     * (Computed) The etag of the dataset's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The role that should be applied. Only one
     * `gcp.healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a DatasetIamBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatasetIamBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatasetIamBindingArgs | DatasetIamBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DatasetIamBindingState | undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["datasetId"] = state ? state.datasetId : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as DatasetIamBindingArgs | undefined;
            if (!args || args.datasetId === undefined) {
                throw new Error("Missing required property 'datasetId'");
            }
            if (!args || args.members === undefined) {
                throw new Error("Missing required property 'members'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["condition"] = args ? args.condition : undefined;
            inputs["datasetId"] = args ? args.datasetId : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DatasetIamBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatasetIamBinding resources.
 */
export interface DatasetIamBindingState {
    readonly condition?: pulumi.Input<inputs.healthcare.DatasetIamBindingCondition>;
    /**
     * The dataset ID, in the form
     * `{project_id}/{location_name}/{dataset_name}` or
     * `{location_name}/{dataset_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    readonly datasetId?: pulumi.Input<string>;
    /**
     * (Computed) The etag of the dataset's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The role that should be applied. Only one
     * `gcp.healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatasetIamBinding resource.
 */
export interface DatasetIamBindingArgs {
    readonly condition?: pulumi.Input<inputs.healthcare.DatasetIamBindingCondition>;
    /**
     * The dataset ID, in the form
     * `{project_id}/{location_name}/{dataset_name}` or
     * `{location_name}/{dataset_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    readonly datasetId: pulumi.Input<string>;
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The role that should be applied. Only one
     * `gcp.healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}

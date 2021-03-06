// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * An AccessLevel is a label that can be applied to requests to GCP services,
 * along with a list of requirements necessary for the label to be applied.
 * 
 * 
 * To get more information about AccessLevel, see:
 * 
 * * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.accessLevels)
 * * How-to Guides
 *     * [Access Policy Quickstart](https://cloud.google.com/access-context-manager/docs/quickstart)
 * 
 * ## Example Usage - Access Context Manager Access Level Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const accessLevel = new gcp.accesscontextmanager.AccessLevel("access-level", {
 *     basic: {
 *         conditions: [{
 *             devicePolicy: {
 *                 osConstraints: [{
 *                     osType: "DESKTOP_CHROME_OS",
 *                 }],
 *                 requireScreenLock: false,
 *             },
 *         }],
 *     },
 *     parent: pulumi.interpolate`accessPolicies/${google_access_context_manager_access_policy_test_access.name}`,
 *     title: "chromeosNoLock",
 * });
 * const accessPolicy = new gcp.accesscontextmanager.AccessPolicy("access-policy", {
 *     parent: "organizations/123456789",
 *     title: "my policy",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/access_context_manager_access_level.html.markdown.
 */
export class AccessLevel extends pulumi.CustomResource {
    /**
     * Get an existing AccessLevel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessLevelState, opts?: pulumi.CustomResourceOptions): AccessLevel {
        return new AccessLevel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:accesscontextmanager/accessLevel:AccessLevel';

    /**
     * Returns true if the given object is an instance of AccessLevel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessLevel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessLevel.__pulumiType;
    }

    /**
     * A set of predefined conditions for the access level and a combining function.
     */
    public readonly basic!: pulumi.Output<outputs.accesscontextmanager.AccessLevelBasic | undefined>;
    /**
     * Description of the AccessLevel and its use. Does not affect behavior.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Resource name for the Access Level. The short_name component must begin with a letter and only include alphanumeric
     * and '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name}
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}
     */
    public readonly parent!: pulumi.Output<string>;
    /**
     * Human readable title. Must be unique within the Policy.
     */
    public readonly title!: pulumi.Output<string>;

    /**
     * Create a AccessLevel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessLevelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessLevelArgs | AccessLevelState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AccessLevelState | undefined;
            inputs["basic"] = state ? state.basic : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parent"] = state ? state.parent : undefined;
            inputs["title"] = state ? state.title : undefined;
        } else {
            const args = argsOrState as AccessLevelArgs | undefined;
            if (!args || args.parent === undefined) {
                throw new Error("Missing required property 'parent'");
            }
            if (!args || args.title === undefined) {
                throw new Error("Missing required property 'title'");
            }
            inputs["basic"] = args ? args.basic : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parent"] = args ? args.parent : undefined;
            inputs["title"] = args ? args.title : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AccessLevel.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessLevel resources.
 */
export interface AccessLevelState {
    /**
     * A set of predefined conditions for the access level and a combining function.
     */
    readonly basic?: pulumi.Input<inputs.accesscontextmanager.AccessLevelBasic>;
    /**
     * Description of the AccessLevel and its use. Does not affect behavior.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Resource name for the Access Level. The short_name component must begin with a letter and only include alphanumeric
     * and '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name}
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}
     */
    readonly parent?: pulumi.Input<string>;
    /**
     * Human readable title. Must be unique within the Policy.
     */
    readonly title?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccessLevel resource.
 */
export interface AccessLevelArgs {
    /**
     * A set of predefined conditions for the access level and a combining function.
     */
    readonly basic?: pulumi.Input<inputs.accesscontextmanager.AccessLevelBasic>;
    /**
     * Description of the AccessLevel and its use. Does not affect behavior.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Resource name for the Access Level. The short_name component must begin with a letter and only include alphanumeric
     * and '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name}
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}
     */
    readonly parent: pulumi.Input<string>;
    /**
     * Human readable title. Must be unique within the Policy.
     */
    readonly title: pulumi.Input<string>;
}

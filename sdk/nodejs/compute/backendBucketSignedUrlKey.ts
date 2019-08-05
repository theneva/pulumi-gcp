// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A key for signing Cloud CDN signed URLs for BackendBuckets.
 * 
 * 
 * To get more information about BackendBucketSignedUrlKey, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/backendBuckets)
 * * How-to Guides
 *     * [Using Signed URLs](https://cloud.google.com/cdn/docs/using-signed-urls/)
 * 
 * > **Warning:** All arguments including the key's value will be stored in the raw
 * state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
 * Because the API does not return the sensitive key value,
 * we cannot confirm or reverse changes to a key outside of this provider.
 * 
 * ## Example Usage - Backend Bucket Signed Url Key
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const bucket = new gcp.storage.Bucket("bucket", {
 *     location: "EU",
 * });
 * const testBackend = new gcp.compute.BackendBucket("test_backend", {
 *     bucketName: bucket.name,
 *     description: "Contains beautiful images",
 *     enableCdn: true,
 * });
 * const backendKey = new gcp.compute.BackendBucketSignedUrlKey("backend_key", {
 *     backendBucket: testBackend.name,
 *     keyValue: "pPsVemX8GM46QVeezid6Rw==",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_backend_bucket_signed_url_key.html.markdown.
 */
export class BackendBucketSignedUrlKey extends pulumi.CustomResource {
    /**
     * Get an existing BackendBucketSignedUrlKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackendBucketSignedUrlKeyState, opts?: pulumi.CustomResourceOptions): BackendBucketSignedUrlKey {
        return new BackendBucketSignedUrlKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/backendBucketSignedUrlKey:BackendBucketSignedUrlKey';

    /**
     * Returns true if the given object is an instance of BackendBucketSignedUrlKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackendBucketSignedUrlKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackendBucketSignedUrlKey.__pulumiType;
    }

    public readonly backendBucket!: pulumi.Output<string>;
    public readonly keyValue!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a BackendBucketSignedUrlKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackendBucketSignedUrlKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackendBucketSignedUrlKeyArgs | BackendBucketSignedUrlKeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BackendBucketSignedUrlKeyState | undefined;
            inputs["backendBucket"] = state ? state.backendBucket : undefined;
            inputs["keyValue"] = state ? state.keyValue : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as BackendBucketSignedUrlKeyArgs | undefined;
            if (!args || args.backendBucket === undefined) {
                throw new Error("Missing required property 'backendBucket'");
            }
            if (!args || args.keyValue === undefined) {
                throw new Error("Missing required property 'keyValue'");
            }
            inputs["backendBucket"] = args ? args.backendBucket : undefined;
            inputs["keyValue"] = args ? args.keyValue : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BackendBucketSignedUrlKey.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackendBucketSignedUrlKey resources.
 */
export interface BackendBucketSignedUrlKeyState {
    readonly backendBucket?: pulumi.Input<string>;
    readonly keyValue?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BackendBucketSignedUrlKey resource.
 */
export interface BackendBucketSignedUrlKeyArgs {
    readonly backendBucket: pulumi.Input<string>;
    readonly keyValue: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
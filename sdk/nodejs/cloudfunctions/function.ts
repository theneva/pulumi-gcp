// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a new Cloud Function. For more information see
 * [the official documentation](https://cloud.google.com/functions/docs/)
 * and
 * [API](https://cloud.google.com/functions/docs/apis).
 * 
 * > **Warning:** As of November 1, 2019, newly created Functions are
 * private-by-default and will require [appropriate IAM permissions](https://cloud.google.com/functions/docs/reference/iam/roles)
 * to be invoked. See below examples for how to set up the appropriate permissions,
 * or view the [Cloud Functions IAM resources](https://www.terraform.io/docs/providers/google/r/cloudfunctions_cloud_function_iam.html)
 * for Cloud Functions.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloudfunctions_function.html.markdown.
 */
export class Function extends pulumi.CustomResource {
    /**
     * Get an existing Function resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionState, opts?: pulumi.CustomResourceOptions): Function {
        return new Function(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:cloudfunctions/function:Function';

    /**
     * Returns true if the given object is an instance of Function.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Function {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Function.__pulumiType;
    }

    /**
     * Memory (in MB), available to the function. Default value is 256MB. Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.
     */
    public readonly availableMemoryMb!: pulumi.Output<number | undefined>;
    /**
     * Description of the function.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the function that will be executed when the Google Cloud Function is triggered.
     */
    public readonly entryPoint!: pulumi.Output<string | undefined>;
    /**
     * A set of key/value environment variable pairs to assign to the function.
     */
    public readonly environmentVariables!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `triggerHttp`.
     */
    public readonly eventTrigger!: pulumi.Output<outputs.cloudfunctions.FunctionEventTrigger>;
    /**
     * URL which triggers function execution. Returned only if `triggerHttp` is used.
     */
    public readonly httpsTriggerUrl!: pulumi.Output<string>;
    /**
     * A set of key/value label pairs to assign to the function.
     */
    public readonly labels!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The limit on the maximum number of function instances that may coexist at a given time.
     */
    public readonly maxInstances!: pulumi.Output<number | undefined>;
    /**
     * A user-defined name of the function. Function names must be unique globally.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Project of the function. If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Region of function. Currently can be only "us-central1". If it is not provided, the provider region is used.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The runtime in which the function is going to run.
     * Eg. `"nodejs8"`, `"nodejs10"`, `"python37"`, `"go111"`.
     */
    public readonly runtime!: pulumi.Output<string>;
    /**
     * If provided, the self-provided service account to run the function with.
     */
    public readonly serviceAccountEmail!: pulumi.Output<string>;
    /**
     * The GCS bucket containing the zip archive which contains the function.
     */
    public readonly sourceArchiveBucket!: pulumi.Output<string | undefined>;
    /**
     * The source archive object (file) in archive bucket.
     */
    public readonly sourceArchiveObject!: pulumi.Output<string | undefined>;
    /**
     * Represents parameters related to source repository where a function is hosted.
     * Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below.
     */
    public readonly sourceRepository!: pulumi.Output<outputs.cloudfunctions.FunctionSourceRepository | undefined>;
    /**
     * Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
     */
    public readonly timeout!: pulumi.Output<number | undefined>;
    /**
     * Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `httpsTriggerUrl`. Cannot be used with `triggerBucket` and `triggerTopic`.
     */
    public readonly triggerHttp!: pulumi.Output<boolean | undefined>;
    /**
     * The VPC Network Connector that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is `projects/*&#47;locations/*&#47;connectors/*`.
     */
    public readonly vpcConnector!: pulumi.Output<string | undefined>;

    /**
     * Create a Function resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionArgs | FunctionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FunctionState | undefined;
            inputs["availableMemoryMb"] = state ? state.availableMemoryMb : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["entryPoint"] = state ? state.entryPoint : undefined;
            inputs["environmentVariables"] = state ? state.environmentVariables : undefined;
            inputs["eventTrigger"] = state ? state.eventTrigger : undefined;
            inputs["httpsTriggerUrl"] = state ? state.httpsTriggerUrl : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["maxInstances"] = state ? state.maxInstances : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["runtime"] = state ? state.runtime : undefined;
            inputs["serviceAccountEmail"] = state ? state.serviceAccountEmail : undefined;
            inputs["sourceArchiveBucket"] = state ? state.sourceArchiveBucket : undefined;
            inputs["sourceArchiveObject"] = state ? state.sourceArchiveObject : undefined;
            inputs["sourceRepository"] = state ? state.sourceRepository : undefined;
            inputs["timeout"] = state ? state.timeout : undefined;
            inputs["triggerHttp"] = state ? state.triggerHttp : undefined;
            inputs["vpcConnector"] = state ? state.vpcConnector : undefined;
        } else {
            const args = argsOrState as FunctionArgs | undefined;
            if (!args || args.runtime === undefined) {
                throw new Error("Missing required property 'runtime'");
            }
            inputs["availableMemoryMb"] = args ? args.availableMemoryMb : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["entryPoint"] = args ? args.entryPoint : undefined;
            inputs["environmentVariables"] = args ? args.environmentVariables : undefined;
            inputs["eventTrigger"] = args ? args.eventTrigger : undefined;
            inputs["httpsTriggerUrl"] = args ? args.httpsTriggerUrl : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["maxInstances"] = args ? args.maxInstances : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["runtime"] = args ? args.runtime : undefined;
            inputs["serviceAccountEmail"] = args ? args.serviceAccountEmail : undefined;
            inputs["sourceArchiveBucket"] = args ? args.sourceArchiveBucket : undefined;
            inputs["sourceArchiveObject"] = args ? args.sourceArchiveObject : undefined;
            inputs["sourceRepository"] = args ? args.sourceRepository : undefined;
            inputs["timeout"] = args ? args.timeout : undefined;
            inputs["triggerHttp"] = args ? args.triggerHttp : undefined;
            inputs["vpcConnector"] = args ? args.vpcConnector : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Function.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Function resources.
 */
export interface FunctionState {
    /**
     * Memory (in MB), available to the function. Default value is 256MB. Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.
     */
    readonly availableMemoryMb?: pulumi.Input<number>;
    /**
     * Description of the function.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the function that will be executed when the Google Cloud Function is triggered.
     */
    readonly entryPoint?: pulumi.Input<string>;
    /**
     * A set of key/value environment variable pairs to assign to the function.
     */
    readonly environmentVariables?: pulumi.Input<{[key: string]: any}>;
    /**
     * A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `triggerHttp`.
     */
    readonly eventTrigger?: pulumi.Input<inputs.cloudfunctions.FunctionEventTrigger>;
    /**
     * URL which triggers function execution. Returned only if `triggerHttp` is used.
     */
    readonly httpsTriggerUrl?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the function.
     */
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * The limit on the maximum number of function instances that may coexist at a given time.
     */
    readonly maxInstances?: pulumi.Input<number>;
    /**
     * A user-defined name of the function. Function names must be unique globally.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Project of the function. If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Region of function. Currently can be only "us-central1". If it is not provided, the provider region is used.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The runtime in which the function is going to run.
     * Eg. `"nodejs8"`, `"nodejs10"`, `"python37"`, `"go111"`.
     */
    readonly runtime?: pulumi.Input<string>;
    /**
     * If provided, the self-provided service account to run the function with.
     */
    readonly serviceAccountEmail?: pulumi.Input<string>;
    /**
     * The GCS bucket containing the zip archive which contains the function.
     */
    readonly sourceArchiveBucket?: pulumi.Input<string>;
    /**
     * The source archive object (file) in archive bucket.
     */
    readonly sourceArchiveObject?: pulumi.Input<string>;
    /**
     * Represents parameters related to source repository where a function is hosted.
     * Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below.
     */
    readonly sourceRepository?: pulumi.Input<inputs.cloudfunctions.FunctionSourceRepository>;
    /**
     * Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
     */
    readonly timeout?: pulumi.Input<number>;
    /**
     * Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `httpsTriggerUrl`. Cannot be used with `triggerBucket` and `triggerTopic`.
     */
    readonly triggerHttp?: pulumi.Input<boolean>;
    /**
     * The VPC Network Connector that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is `projects/*&#47;locations/*&#47;connectors/*`.
     */
    readonly vpcConnector?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Function resource.
 */
export interface FunctionArgs {
    /**
     * Memory (in MB), available to the function. Default value is 256MB. Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.
     */
    readonly availableMemoryMb?: pulumi.Input<number>;
    /**
     * Description of the function.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the function that will be executed when the Google Cloud Function is triggered.
     */
    readonly entryPoint?: pulumi.Input<string>;
    /**
     * A set of key/value environment variable pairs to assign to the function.
     */
    readonly environmentVariables?: pulumi.Input<{[key: string]: any}>;
    /**
     * A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `triggerHttp`.
     */
    readonly eventTrigger?: pulumi.Input<inputs.cloudfunctions.FunctionEventTrigger>;
    /**
     * URL which triggers function execution. Returned only if `triggerHttp` is used.
     */
    readonly httpsTriggerUrl?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the function.
     */
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * The limit on the maximum number of function instances that may coexist at a given time.
     */
    readonly maxInstances?: pulumi.Input<number>;
    /**
     * A user-defined name of the function. Function names must be unique globally.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Project of the function. If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Region of function. Currently can be only "us-central1". If it is not provided, the provider region is used.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The runtime in which the function is going to run.
     * Eg. `"nodejs8"`, `"nodejs10"`, `"python37"`, `"go111"`.
     */
    readonly runtime: pulumi.Input<string>;
    /**
     * If provided, the self-provided service account to run the function with.
     */
    readonly serviceAccountEmail?: pulumi.Input<string>;
    /**
     * The GCS bucket containing the zip archive which contains the function.
     */
    readonly sourceArchiveBucket?: pulumi.Input<string>;
    /**
     * The source archive object (file) in archive bucket.
     */
    readonly sourceArchiveObject?: pulumi.Input<string>;
    /**
     * Represents parameters related to source repository where a function is hosted.
     * Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below.
     */
    readonly sourceRepository?: pulumi.Input<inputs.cloudfunctions.FunctionSourceRepository>;
    /**
     * Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
     */
    readonly timeout?: pulumi.Input<number>;
    /**
     * Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `httpsTriggerUrl`. Cannot be used with `triggerBucket` and `triggerTopic`.
     */
    readonly triggerHttp?: pulumi.Input<boolean>;
    /**
     * The VPC Network Connector that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is `projects/*&#47;locations/*&#47;connectors/*`.
     */
    readonly vpcConnector?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a Google Bigtable instance. For more information see
 * [the official documentation](https://cloud.google.com/bigtable/) and
 * [API](https://cloud.google.com/bigtable/docs/go/reference).
 * 
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState): Instance {
        return new Instance(name, <any>state, { id });
    }

    /**
     * A block of cluster configuration options. Either `cluster` or `cluster_id` must be used. Only one cluster may be specified. See structure below.
     */
    public readonly cluster: pulumi.Output<{ clusterId?: string, numNodes?: number, storageType?: string, zone: string } | undefined>;
    /**
     * The ID of the Cloud Bigtable cluster.
     */
    public readonly clusterId: pulumi.Output<string | undefined>;
    /**
     * The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
     */
    public readonly displayName: pulumi.Output<string>;
    /**
     * The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
     */
    public readonly instanceType: pulumi.Output<string | undefined>;
    /**
     * The name of the Cloud Bigtable instance.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The number of nodes in your Cloud Bigtable cluster. Minimum of `3` for a `PRODUCTION` instance. Cannot be set for a `DEVELOPMENT` instance.
     */
    public readonly numNodes: pulumi.Output<number | undefined>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    /**
     * The storage type to use. One of `"SSD"` or `"HDD"`. Defaults to `"SSD"`.
     */
    public readonly storageType: pulumi.Output<string | undefined>;
    /**
     * The zone to create the Cloud Bigtable cluster in. Zones that support Bigtable instances are noted on the [Cloud Bigtable locations page](https://cloud.google.com/bigtable/docs/locations).
     */
    public readonly zone: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: InstanceState = argsOrState as InstanceState | undefined;
            inputs["cluster"] = state ? state.cluster : undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["numNodes"] = state ? state.numNodes : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["storageType"] = state ? state.storageType : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            inputs["cluster"] = args ? args.cluster : undefined;
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["numNodes"] = args ? args.numNodes : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["storageType"] = args ? args.storageType : undefined;
            inputs["zone"] = args ? args.zone : undefined;
        }
        super("gcp:bigtable/instance:Instance", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * A block of cluster configuration options. Either `cluster` or `cluster_id` must be used. Only one cluster may be specified. See structure below.
     */
    readonly cluster?: pulumi.Input<{ clusterId?: pulumi.Input<string>, numNodes?: pulumi.Input<number>, storageType?: pulumi.Input<string>, zone?: pulumi.Input<string> }>;
    /**
     * The ID of the Cloud Bigtable cluster.
     */
    readonly clusterId?: pulumi.Input<string>;
    /**
     * The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
     */
    readonly instanceType?: pulumi.Input<string>;
    /**
     * The name of the Cloud Bigtable instance.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The number of nodes in your Cloud Bigtable cluster. Minimum of `3` for a `PRODUCTION` instance. Cannot be set for a `DEVELOPMENT` instance.
     */
    readonly numNodes?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The storage type to use. One of `"SSD"` or `"HDD"`. Defaults to `"SSD"`.
     */
    readonly storageType?: pulumi.Input<string>;
    /**
     * The zone to create the Cloud Bigtable cluster in. Zones that support Bigtable instances are noted on the [Cloud Bigtable locations page](https://cloud.google.com/bigtable/docs/locations).
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * A block of cluster configuration options. Either `cluster` or `cluster_id` must be used. Only one cluster may be specified. See structure below.
     */
    readonly cluster?: pulumi.Input<{ clusterId?: pulumi.Input<string>, numNodes?: pulumi.Input<number>, storageType?: pulumi.Input<string>, zone?: pulumi.Input<string> }>;
    /**
     * The ID of the Cloud Bigtable cluster.
     */
    readonly clusterId?: pulumi.Input<string>;
    /**
     * The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
     */
    readonly instanceType?: pulumi.Input<string>;
    /**
     * The name of the Cloud Bigtable instance.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The number of nodes in your Cloud Bigtable cluster. Minimum of `3` for a `PRODUCTION` instance. Cannot be set for a `DEVELOPMENT` instance.
     */
    readonly numNodes?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The storage type to use. One of `"SSD"` or `"HDD"`. Defaults to `"SSD"`.
     */
    readonly storageType?: pulumi.Input<string>;
    /**
     * The zone to create the Cloud Bigtable cluster in. Zones that support Bigtable instances are noted on the [Cloud Bigtable locations page](https://cloud.google.com/bigtable/docs/locations).
     */
    readonly zone?: pulumi.Input<string>;
}

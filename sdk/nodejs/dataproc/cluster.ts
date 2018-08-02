// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Manages a Cloud Dataproc cluster resource within GCP. For more information see
 * [the official dataproc documentation](https://cloud.google.com/dataproc/).
 * 
 * 
 * !> **Warning:** Due to limitations of the API, all arguments except
 * `labels`,`cluster_config.worker_config.num_instances` and `cluster_config.preemptible_worker_config.num_instances` are non-updateable. Changing others will cause recreation of the
 * whole cluster!
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState): Cluster {
        return new Cluster(name, <any>state, { id });
    }

    /**
     * Allows you to configure various aspects of the cluster.
     * Structure defined below.
     */
    public readonly clusterConfig: pulumi.Output<{ bucket: string, deleteAutogenBucket?: boolean, gceClusterConfig: { internalIpOnly?: boolean, metadata?: {[key: string]: string}, network: string, serviceAccount?: string, serviceAccountScopes: string[], subnetwork?: string, tags?: string[], zone: string }, initializationActions?: { script: string, timeoutSec?: number }[], masterConfig: { diskConfig: { bootDiskSizeGb: number, numLocalSsds: number }, instanceNames: string[], machineType: string, numInstances: number }, preemptibleWorkerConfig: { diskConfig: { bootDiskSizeGb: number }, instanceNames: string[], numInstances: number }, softwareConfig: { imageVersion: string, overrideProperties?: {[key: string]: string}, properties: {[key: string]: any} }, stagingBucket?: string, workerConfig: { diskConfig: { bootDiskSizeGb: number, numLocalSsds: number }, instanceNames: string[], machineType: string, numInstances: number } }>;
    /**
     * The list of labels (key/value pairs) to be applied to
     * instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
     * which is the name of the cluster.
     */
    public readonly labels: pulumi.Output<{[key: string]: string}>;
    /**
     * The name of the cluster, unique within the project and
     * zone.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The ID of the project in which the `cluster` will exist. If it
     * is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    /**
     * The region in which the cluster and associated nodes will be created in.
     * Defaults to `global`.
     */
    public readonly region: pulumi.Output<string | undefined>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ClusterArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ClusterState = argsOrState as ClusterState | undefined;
            inputs["clusterConfig"] = state ? state.clusterConfig : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            inputs["clusterConfig"] = args ? args.clusterConfig : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
        }
        super("gcp:dataproc/cluster:Cluster", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * Allows you to configure various aspects of the cluster.
     * Structure defined below.
     */
    readonly clusterConfig?: pulumi.Input<{ bucket?: pulumi.Input<string>, deleteAutogenBucket?: pulumi.Input<boolean>, gceClusterConfig?: pulumi.Input<{ internalIpOnly?: pulumi.Input<boolean>, metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, network?: pulumi.Input<string>, serviceAccount?: pulumi.Input<string>, serviceAccountScopes?: pulumi.Input<pulumi.Input<string>[]>, subnetwork?: pulumi.Input<string>, tags?: pulumi.Input<pulumi.Input<string>[]>, zone?: pulumi.Input<string> }>, initializationActions?: pulumi.Input<pulumi.Input<{ script: pulumi.Input<string>, timeoutSec?: pulumi.Input<number> }>[]>, masterConfig?: pulumi.Input<{ diskConfig?: pulumi.Input<{ bootDiskSizeGb?: pulumi.Input<number>, numLocalSsds?: pulumi.Input<number> }>, instanceNames?: pulumi.Input<pulumi.Input<string>[]>, machineType?: pulumi.Input<string>, numInstances?: pulumi.Input<number> }>, preemptibleWorkerConfig?: pulumi.Input<{ diskConfig?: pulumi.Input<{ bootDiskSizeGb?: pulumi.Input<number> }>, instanceNames?: pulumi.Input<pulumi.Input<string>[]>, numInstances?: pulumi.Input<number> }>, softwareConfig?: pulumi.Input<{ imageVersion?: pulumi.Input<string>, overrideProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, properties?: pulumi.Input<{[key: string]: any}> }>, stagingBucket?: pulumi.Input<string>, workerConfig?: pulumi.Input<{ diskConfig?: pulumi.Input<{ bootDiskSizeGb?: pulumi.Input<number>, numLocalSsds?: pulumi.Input<number> }>, instanceNames?: pulumi.Input<pulumi.Input<string>[]>, machineType?: pulumi.Input<string>, numInstances?: pulumi.Input<number> }> }>;
    /**
     * The list of labels (key/value pairs) to be applied to
     * instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
     * which is the name of the cluster.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the cluster, unique within the project and
     * zone.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the `cluster` will exist. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region in which the cluster and associated nodes will be created in.
     * Defaults to `global`.
     */
    readonly region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * Allows you to configure various aspects of the cluster.
     * Structure defined below.
     */
    readonly clusterConfig?: pulumi.Input<{ bucket?: pulumi.Input<string>, deleteAutogenBucket?: pulumi.Input<boolean>, gceClusterConfig?: pulumi.Input<{ internalIpOnly?: pulumi.Input<boolean>, metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, network?: pulumi.Input<string>, serviceAccount?: pulumi.Input<string>, serviceAccountScopes?: pulumi.Input<pulumi.Input<string>[]>, subnetwork?: pulumi.Input<string>, tags?: pulumi.Input<pulumi.Input<string>[]>, zone?: pulumi.Input<string> }>, initializationActions?: pulumi.Input<pulumi.Input<{ script: pulumi.Input<string>, timeoutSec?: pulumi.Input<number> }>[]>, masterConfig?: pulumi.Input<{ diskConfig?: pulumi.Input<{ bootDiskSizeGb?: pulumi.Input<number>, numLocalSsds?: pulumi.Input<number> }>, instanceNames?: pulumi.Input<pulumi.Input<string>[]>, machineType?: pulumi.Input<string>, numInstances?: pulumi.Input<number> }>, preemptibleWorkerConfig?: pulumi.Input<{ diskConfig?: pulumi.Input<{ bootDiskSizeGb?: pulumi.Input<number> }>, instanceNames?: pulumi.Input<pulumi.Input<string>[]>, numInstances?: pulumi.Input<number> }>, softwareConfig?: pulumi.Input<{ imageVersion?: pulumi.Input<string>, overrideProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, properties?: pulumi.Input<{[key: string]: any}> }>, stagingBucket?: pulumi.Input<string>, workerConfig?: pulumi.Input<{ diskConfig?: pulumi.Input<{ bootDiskSizeGb?: pulumi.Input<number>, numLocalSsds?: pulumi.Input<number> }>, instanceNames?: pulumi.Input<pulumi.Input<string>[]>, machineType?: pulumi.Input<string>, numInstances?: pulumi.Input<number> }> }>;
    /**
     * The list of labels (key/value pairs) to be applied to
     * instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
     * which is the name of the cluster.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the cluster, unique within the project and
     * zone.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the `cluster` will exist. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region in which the cluster and associated nodes will be created in.
     * Defaults to `global`.
     */
    readonly region?: pulumi.Input<string>;
}

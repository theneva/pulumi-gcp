// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Represents an Autoscaler resource.
 * 
 * Autoscalers allow you to automatically scale virtual machine instances in
 * managed instance groups according to an autoscaling policy that you
 * define.
 * 
 * 
 * To get more information about Autoscaler, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/autoscalers)
 * * How-to Guides
 *     * [Autoscaling Groups of Instances](https://cloud.google.com/compute/docs/autoscaler/)
 * 
 * <div class = "oics-button" style="float: right; margin: 0 0 -15px">
 *   <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=autoscaler_beta&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
 *     <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
 *   </a>
 * </div>
 * ## Example Usage - Autoscaler Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const debian9 = pulumi.output(gcp.compute.getImage({
 *     family: "debian-9",
 *     project: "debian-cloud",
 * }));
 * const foobarTargetPool = new gcp.compute.TargetPool("foobar", {});
 * const foobarInstanceTemplate = new gcp.compute.InstanceTemplate("foobar", {
 *     canIpForward: false,
 *     disks: [{
 *         sourceImage: debian9.apply(debian9 => debian9.selfLink),
 *     }],
 *     machineType: "n1-standard-1",
 *     metadata: {
 *         foo: "bar",
 *     },
 *     networkInterfaces: [{
 *         network: "default",
 *     }],
 *     serviceAccount: {
 *         scopes: [
 *             "userinfo-email",
 *             "compute-ro",
 *             "storage-ro",
 *         ],
 *     },
 *     tags: [
 *         "foo",
 *         "bar",
 *     ],
 * });
 * const foobarInstanceGroupManager = new gcp.compute.InstanceGroupManager("foobar", {
 *     baseInstanceName: "foobar",
 *     instanceTemplate: foobarInstanceTemplate.selfLink,
 *     targetPools: [foobarTargetPool.selfLink],
 *     zone: "us-central1-f",
 * });
 * const foobarAutoscalar = new gcp.compute.Autoscalar("foobar", {
 *     autoscalingPolicy: {
 *         cooldownPeriod: 60,
 *         cpuUtilization: {
 *             target: 0.5,
 *         },
 *         maxReplicas: 5,
 *         minReplicas: 1,
 *     },
 *     target: foobarInstanceGroupManager.selfLink,
 *     zone: "us-central1-f",
 * });
 * ```
 */
export class Autoscalar extends pulumi.CustomResource {
    /**
     * Get an existing Autoscalar resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AutoscalarState, opts?: pulumi.CustomResourceOptions): Autoscalar {
        return new Autoscalar(name, <any>state, { ...opts, id: id });
    }

    public readonly autoscalingPolicy: pulumi.Output<{ cooldownPeriod?: number, cpuUtilization: { target: number }, loadBalancingUtilization?: { target: number }, maxReplicas: number, metrics?: { filter?: string, name: string, target: number, type: string }[], minReplicas: number }>;
    public /*out*/ readonly creationTimestamp: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly project: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    public readonly target: pulumi.Output<string>;
    public readonly zone: pulumi.Output<string>;

    /**
     * Create a Autoscalar resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AutoscalarArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AutoscalarArgs | AutoscalarState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: AutoscalarState = argsOrState as AutoscalarState | undefined;
            inputs["autoscalingPolicy"] = state ? state.autoscalingPolicy : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["target"] = state ? state.target : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as AutoscalarArgs | undefined;
            if (!args || args.autoscalingPolicy === undefined) {
                throw new Error("Missing required property 'autoscalingPolicy'");
            }
            if (!args || args.target === undefined) {
                throw new Error("Missing required property 'target'");
            }
            inputs["autoscalingPolicy"] = args ? args.autoscalingPolicy : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["target"] = args ? args.target : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        super("gcp:compute/autoscalar:Autoscalar", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Autoscalar resources.
 */
export interface AutoscalarState {
    readonly autoscalingPolicy?: pulumi.Input<{ cooldownPeriod?: pulumi.Input<number>, cpuUtilization?: pulumi.Input<{ target: pulumi.Input<number> }>, loadBalancingUtilization?: pulumi.Input<{ target: pulumi.Input<number> }>, maxReplicas: pulumi.Input<number>, metrics?: pulumi.Input<pulumi.Input<{ filter?: pulumi.Input<string>, name: pulumi.Input<string>, target: pulumi.Input<number>, type: pulumi.Input<string> }>[]>, minReplicas: pulumi.Input<number> }>;
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly target?: pulumi.Input<string>;
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Autoscalar resource.
 */
export interface AutoscalarArgs {
    readonly autoscalingPolicy: pulumi.Input<{ cooldownPeriod?: pulumi.Input<number>, cpuUtilization?: pulumi.Input<{ target: pulumi.Input<number> }>, loadBalancingUtilization?: pulumi.Input<{ target: pulumi.Input<number> }>, maxReplicas: pulumi.Input<number>, metrics?: pulumi.Input<pulumi.Input<{ filter?: pulumi.Input<string>, name: pulumi.Input<string>, target: pulumi.Input<number>, type: pulumi.Input<string> }>[]>, minReplicas: pulumi.Input<number> }>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly target: pulumi.Input<string>;
    readonly zone?: pulumi.Input<string>;
}

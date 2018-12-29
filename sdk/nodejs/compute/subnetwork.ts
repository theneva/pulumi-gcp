// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A VPC network is a virtual version of the traditional physical networks
 * that exist within and between physical data centers. A VPC network
 * provides connectivity for your Compute Engine virtual machine (VM)
 * instances, Container Engine containers, App Engine Flex services, and
 * other network-related resources.
 * 
 * Each GCP project contains one or more VPC networks. Each VPC network is a
 * global entity spanning all GCP regions. This global VPC network allows VM
 * instances and other resources to communicate with each other via internal,
 * private IP addresses.
 * 
 * Each VPC network is subdivided into subnets, and each subnet is contained
 * within a single region. You can have more than one subnet in a region for
 * a given VPC network. Each subnet has a contiguous private RFC1918 IP
 * space. You create instances, containers, and the like in these subnets.
 * When you create an instance, you must create it in a subnet, and the
 * instance draws its internal IP address from that subnet.
 * 
 * Virtual machine (VM) instances in a VPC network can communicate with
 * instances in all other subnets of the same VPC network, regardless of
 * region, using their RFC1918 private IP addresses. You can isolate portions
 * of the network, even entire subnets, using firewall rules.
 * 
 * 
 * To get more information about Subnetwork, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/subnetworks)
 * * How-to Guides
 *     * [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
 *     * [Cloud Networking](https://cloud.google.com/vpc/docs/using-vpc)
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const google_compute_network_custom_test = new gcp.compute.Network("custom-test", {
 *     autoCreateSubnetworks: false,
 *     name: "test-network",
 * });
 * const google_compute_subnetwork_network_with_private_secondary_ip_ranges = new gcp.compute.Subnetwork("network-with-private-secondary-ip-ranges", {
 *     ipCidrRange: "10.2.0.0/16",
 *     name: "test-subnetwork",
 *     network: google_compute_network_custom_test.selfLink,
 *     region: "us-central1",
 *     secondaryIpRanges: [{
 *         ipCidrRange: "192.168.10.0/24",
 *         rangeName: "tf-test-secondary-range-update1",
 *     }],
 * });
 * ```
 */
export class Subnetwork extends pulumi.CustomResource {
    /**
     * Get an existing Subnetwork resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubnetworkState, opts?: pulumi.CustomResourceOptions): Subnetwork {
        return new Subnetwork(name, <any>state, { ...opts, id: id });
    }

    public /*out*/ readonly creationTimestamp: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly enableFlowLogs: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly fingerprint: pulumi.Output<string>;
    public /*out*/ readonly gatewayAddress: pulumi.Output<string>;
    public readonly ipCidrRange: pulumi.Output<string>;
    public readonly name: pulumi.Output<string>;
    public readonly network: pulumi.Output<string>;
    public readonly privateIpGoogleAccess: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    public readonly region: pulumi.Output<string>;
    public readonly secondaryIpRanges: pulumi.Output<{ ipCidrRange: string, rangeName: string }[]>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink: pulumi.Output<string>;

    /**
     * Create a Subnetwork resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubnetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubnetworkArgs | SubnetworkState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: SubnetworkState = argsOrState as SubnetworkState | undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enableFlowLogs"] = state ? state.enableFlowLogs : undefined;
            inputs["fingerprint"] = state ? state.fingerprint : undefined;
            inputs["gatewayAddress"] = state ? state.gatewayAddress : undefined;
            inputs["ipCidrRange"] = state ? state.ipCidrRange : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["privateIpGoogleAccess"] = state ? state.privateIpGoogleAccess : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["secondaryIpRanges"] = state ? state.secondaryIpRanges : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as SubnetworkArgs | undefined;
            if (!args || args.ipCidrRange === undefined) {
                throw new Error("Missing required property 'ipCidrRange'");
            }
            if (!args || args.network === undefined) {
                throw new Error("Missing required property 'network'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["enableFlowLogs"] = args ? args.enableFlowLogs : undefined;
            inputs["ipCidrRange"] = args ? args.ipCidrRange : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["privateIpGoogleAccess"] = args ? args.privateIpGoogleAccess : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["secondaryIpRanges"] = args ? args.secondaryIpRanges : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["gatewayAddress"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        super("gcp:compute/subnetwork:Subnetwork", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Subnetwork resources.
 */
export interface SubnetworkState {
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly enableFlowLogs?: pulumi.Input<boolean>;
    readonly fingerprint?: pulumi.Input<string>;
    readonly gatewayAddress?: pulumi.Input<string>;
    readonly ipCidrRange?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly network?: pulumi.Input<string>;
    readonly privateIpGoogleAccess?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly secondaryIpRanges?: pulumi.Input<pulumi.Input<{ ipCidrRange: pulumi.Input<string>, rangeName: pulumi.Input<string> }>[]>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Subnetwork resource.
 */
export interface SubnetworkArgs {
    readonly description?: pulumi.Input<string>;
    readonly enableFlowLogs?: pulumi.Input<boolean>;
    readonly ipCidrRange: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly network: pulumi.Input<string>;
    readonly privateIpGoogleAccess?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly secondaryIpRanges?: pulumi.Input<pulumi.Input<{ ipCidrRange: pulumi.Input<string>, rangeName: pulumi.Input<string> }>[]>;
}

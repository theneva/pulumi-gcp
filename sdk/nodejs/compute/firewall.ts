// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Manages a firewall resource within GCE. For more information see
 * [the official documentation](https://cloud.google.com/compute/docs/vpc/firewalls)
 * and
 * [API](https://cloud.google.com/compute/docs/reference/latest/firewalls).
 */
export class Firewall extends pulumi.CustomResource {
    /**
     * Get an existing Firewall resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallState): Firewall {
        return new Firewall(name, <any>state, { id });
    }

    /**
     * Can be specified multiple times for each allow
     * rule. Each allow block supports fields documented below.
     */
    public readonly allows: pulumi.Output<{ ports?: string[], protocol: string }[] | undefined>;
    /**
     * Can be specified multiple times for each deny
     * rule. Each deny block supports fields documented below. Can be specified
     * instead of allow.
     */
    public readonly denies: pulumi.Output<{ ports?: string[], protocol: string }[] | undefined>;
    /**
     * Textual description field.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * A list of destination CIDR ranges that this
     * firewall applies to. Can't be used for `INGRESS`.
     */
    public readonly destinationRanges: pulumi.Output<string[]>;
    /**
     * Direction of traffic to which this firewall applies;
     * One of `INGRESS` or `EGRESS`. Defaults to `INGRESS`.
     */
    public readonly direction: pulumi.Output<string>;
    /**
     * Denotes whether the firewall rule is disabled, i.e not applied to the network it is associated with.
     * When set to true, the firewall rule is not enforced and the network behaves as if it did not exist.
     */
    public readonly disabled: pulumi.Output<boolean | undefined>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The name or self_link of the network to attach this firewall to.
     */
    public readonly network: pulumi.Output<string>;
    /**
     * The priority for this firewall. Ranges from 0-65535, inclusive. Defaults to 1000. Firewall
     * resources with lower priority values have higher precedence (e.g. a firewall resource with a priority value of 0
     * takes effect over all other firewall rules with a non-zero priority).
     */
    public readonly priority: pulumi.Output<number | undefined>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    /**
     * A list of source CIDR ranges that this
     * firewall applies to. Can't be used for `EGRESS`.
     */
    public readonly sourceRanges: pulumi.Output<string[]>;
    /**
     * A list of service accounts such that
     * the firewall will apply only to traffic originating from an instance with a service account in this list.  Note that as of May 2018,
     * this list can contain only one item, due to a change in the way that these firewall rules are handled.  Source service accounts
     * cannot be used to control traffic to an instance's external IP address because service accounts are associated with an instance, not
     * an IP address. `source_ranges` can be set at the same time as `source_service_accounts`. If both are set, the firewall will apply to
     * traffic that has source IP address within `source_ranges` OR the source IP belongs to an instance with service account listed in
     * `source_service_accounts`. The connection does not need to match both properties for the firewall to apply. `source_service_accounts`
     * cannot be used at the same time as `source_tags` or `target_tags`.
     */
    public readonly sourceServiceAccounts: pulumi.Output<string | undefined>;
    /**
     * A list of source tags for this firewall. Can't be used for `EGRESS`.
     */
    public readonly sourceTags: pulumi.Output<string[] | undefined>;
    /**
     * A list of service accounts indicating
     * sets of instances located in the network that may make network connections as specified in `allow`. `target_service_accounts` cannot
     * be used at the same time as `source_tags` or `target_tags`. If neither `target_service_accounts` nor `target_tags` are specified, the
     * firewall rule applies to all instances on the specified network.  Note that as of May 2018, this list can contain only one item, due
     * to a change in the way that these firewall rules are handled.
     */
    public readonly targetServiceAccounts: pulumi.Output<string | undefined>;
    /**
     * A list of target tags for this firewall.
     */
    public readonly targetTags: pulumi.Output<string[] | undefined>;

    /**
     * Create a Firewall resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: FirewallArgs | FirewallState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: FirewallState = argsOrState as FirewallState | undefined;
            inputs["allows"] = state ? state.allows : undefined;
            inputs["denies"] = state ? state.denies : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["destinationRanges"] = state ? state.destinationRanges : undefined;
            inputs["direction"] = state ? state.direction : undefined;
            inputs["disabled"] = state ? state.disabled : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["sourceRanges"] = state ? state.sourceRanges : undefined;
            inputs["sourceServiceAccounts"] = state ? state.sourceServiceAccounts : undefined;
            inputs["sourceTags"] = state ? state.sourceTags : undefined;
            inputs["targetServiceAccounts"] = state ? state.targetServiceAccounts : undefined;
            inputs["targetTags"] = state ? state.targetTags : undefined;
        } else {
            const args = argsOrState as FirewallArgs | undefined;
            if (!args || args.network === undefined) {
                throw new Error("Missing required property 'network'");
            }
            inputs["allows"] = args ? args.allows : undefined;
            inputs["denies"] = args ? args.denies : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["destinationRanges"] = args ? args.destinationRanges : undefined;
            inputs["direction"] = args ? args.direction : undefined;
            inputs["disabled"] = args ? args.disabled : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["sourceRanges"] = args ? args.sourceRanges : undefined;
            inputs["sourceServiceAccounts"] = args ? args.sourceServiceAccounts : undefined;
            inputs["sourceTags"] = args ? args.sourceTags : undefined;
            inputs["targetServiceAccounts"] = args ? args.targetServiceAccounts : undefined;
            inputs["targetTags"] = args ? args.targetTags : undefined;
            inputs["selfLink"] = undefined /*out*/;
        }
        super("gcp:compute/firewall:Firewall", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Firewall resources.
 */
export interface FirewallState {
    /**
     * Can be specified multiple times for each allow
     * rule. Each allow block supports fields documented below.
     */
    readonly allows?: pulumi.Input<pulumi.Input<{ ports?: pulumi.Input<pulumi.Input<string>[]>, protocol: pulumi.Input<string> }>[]>;
    /**
     * Can be specified multiple times for each deny
     * rule. Each deny block supports fields documented below. Can be specified
     * instead of allow.
     */
    readonly denies?: pulumi.Input<pulumi.Input<{ ports?: pulumi.Input<pulumi.Input<string>[]>, protocol: pulumi.Input<string> }>[]>;
    /**
     * Textual description field.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A list of destination CIDR ranges that this
     * firewall applies to. Can't be used for `INGRESS`.
     */
    readonly destinationRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Direction of traffic to which this firewall applies;
     * One of `INGRESS` or `EGRESS`. Defaults to `INGRESS`.
     */
    readonly direction?: pulumi.Input<string>;
    /**
     * Denotes whether the firewall rule is disabled, i.e not applied to the network it is associated with.
     * When set to true, the firewall rule is not enforced and the network behaves as if it did not exist.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name or self_link of the network to attach this firewall to.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * The priority for this firewall. Ranges from 0-65535, inclusive. Defaults to 1000. Firewall
     * resources with lower priority values have higher precedence (e.g. a firewall resource with a priority value of 0
     * takes effect over all other firewall rules with a non-zero priority).
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * A list of source CIDR ranges that this
     * firewall applies to. Can't be used for `EGRESS`.
     */
    readonly sourceRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of service accounts such that
     * the firewall will apply only to traffic originating from an instance with a service account in this list.  Note that as of May 2018,
     * this list can contain only one item, due to a change in the way that these firewall rules are handled.  Source service accounts
     * cannot be used to control traffic to an instance's external IP address because service accounts are associated with an instance, not
     * an IP address. `source_ranges` can be set at the same time as `source_service_accounts`. If both are set, the firewall will apply to
     * traffic that has source IP address within `source_ranges` OR the source IP belongs to an instance with service account listed in
     * `source_service_accounts`. The connection does not need to match both properties for the firewall to apply. `source_service_accounts`
     * cannot be used at the same time as `source_tags` or `target_tags`.
     */
    readonly sourceServiceAccounts?: pulumi.Input<string>;
    /**
     * A list of source tags for this firewall. Can't be used for `EGRESS`.
     */
    readonly sourceTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of service accounts indicating
     * sets of instances located in the network that may make network connections as specified in `allow`. `target_service_accounts` cannot
     * be used at the same time as `source_tags` or `target_tags`. If neither `target_service_accounts` nor `target_tags` are specified, the
     * firewall rule applies to all instances on the specified network.  Note that as of May 2018, this list can contain only one item, due
     * to a change in the way that these firewall rules are handled.
     */
    readonly targetServiceAccounts?: pulumi.Input<string>;
    /**
     * A list of target tags for this firewall.
     */
    readonly targetTags?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Firewall resource.
 */
export interface FirewallArgs {
    /**
     * Can be specified multiple times for each allow
     * rule. Each allow block supports fields documented below.
     */
    readonly allows?: pulumi.Input<pulumi.Input<{ ports?: pulumi.Input<pulumi.Input<string>[]>, protocol: pulumi.Input<string> }>[]>;
    /**
     * Can be specified multiple times for each deny
     * rule. Each deny block supports fields documented below. Can be specified
     * instead of allow.
     */
    readonly denies?: pulumi.Input<pulumi.Input<{ ports?: pulumi.Input<pulumi.Input<string>[]>, protocol: pulumi.Input<string> }>[]>;
    /**
     * Textual description field.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A list of destination CIDR ranges that this
     * firewall applies to. Can't be used for `INGRESS`.
     */
    readonly destinationRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Direction of traffic to which this firewall applies;
     * One of `INGRESS` or `EGRESS`. Defaults to `INGRESS`.
     */
    readonly direction?: pulumi.Input<string>;
    /**
     * Denotes whether the firewall rule is disabled, i.e not applied to the network it is associated with.
     * When set to true, the firewall rule is not enforced and the network behaves as if it did not exist.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name or self_link of the network to attach this firewall to.
     */
    readonly network: pulumi.Input<string>;
    /**
     * The priority for this firewall. Ranges from 0-65535, inclusive. Defaults to 1000. Firewall
     * resources with lower priority values have higher precedence (e.g. a firewall resource with a priority value of 0
     * takes effect over all other firewall rules with a non-zero priority).
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * A list of source CIDR ranges that this
     * firewall applies to. Can't be used for `EGRESS`.
     */
    readonly sourceRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of service accounts such that
     * the firewall will apply only to traffic originating from an instance with a service account in this list.  Note that as of May 2018,
     * this list can contain only one item, due to a change in the way that these firewall rules are handled.  Source service accounts
     * cannot be used to control traffic to an instance's external IP address because service accounts are associated with an instance, not
     * an IP address. `source_ranges` can be set at the same time as `source_service_accounts`. If both are set, the firewall will apply to
     * traffic that has source IP address within `source_ranges` OR the source IP belongs to an instance with service account listed in
     * `source_service_accounts`. The connection does not need to match both properties for the firewall to apply. `source_service_accounts`
     * cannot be used at the same time as `source_tags` or `target_tags`.
     */
    readonly sourceServiceAccounts?: pulumi.Input<string>;
    /**
     * A list of source tags for this firewall. Can't be used for `EGRESS`.
     */
    readonly sourceTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of service accounts indicating
     * sets of instances located in the network that may make network connections as specified in `allow`. `target_service_accounts` cannot
     * be used at the same time as `source_tags` or `target_tags`. If neither `target_service_accounts` nor `target_tags` are specified, the
     * firewall rule applies to all instances on the specified network.  Note that as of May 2018, this list can contain only one item, due
     * to a change in the way that these firewall rules are handled.
     */
    readonly targetServiceAccounts?: pulumi.Input<string>;
    /**
     * A list of target tags for this firewall.
     */
    readonly targetTags?: pulumi.Input<pulumi.Input<string>[]>;
}

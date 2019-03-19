// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a VM instance resource within GCE. For more information see
 * [the official documentation](https://cloud.google.com/compute/docs/instances)
 * and
 * [API](https://cloud.google.com/compute/docs/reference/latest/instances).
 * 
 * This resource is specifically to create a compute instance from a given
 * `source_instance_template`. To create an instance without a template, use the
 * `google_compute_instance` resource.
 * 
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const tplInstanceTemplate = new gcp.compute.InstanceTemplate("tpl", {
 *     canIpForward: true,
 *     disks: [{
 *         autoDelete: true,
 *         boot: true,
 *         diskSizeGb: 100,
 *         sourceImage: "debian-cloud/debian-9",
 *     }],
 *     machineType: "n1-standard-1",
 *     metadata: {
 *         foo: "bar",
 *     },
 *     networkInterfaces: [{
 *         network: "default",
 *     }],
 * });
 * const tplInstanceFromTemplate = new gcp.compute.InstanceFromTemplate("tpl", {
 *     // Override fields from instance template
 *     canIpForward: false,
 *     labels: {
 *         my_key: "my_value",
 *     },
 *     sourceInstanceTemplate: tplInstanceTemplate.selfLink,
 *     zone: "us-central1-a",
 * });
 * ```
 */
export class InstanceFromTemplate extends pulumi.CustomResource {
    /**
     * Get an existing InstanceFromTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceFromTemplateState, opts?: pulumi.CustomResourceOptions): InstanceFromTemplate {
        return new InstanceFromTemplate(name, <any>state, { ...opts, id: id });
    }

    public readonly allowStoppingForUpdate: pulumi.Output<boolean>;
    public readonly attachedDisks: pulumi.Output<{ deviceName: string, diskEncryptionKeyRaw: string, diskEncryptionKeySha256: string, mode: string, source: string }[]>;
    public readonly bootDisk: pulumi.Output<{ autoDelete: boolean, deviceName: string, diskEncryptionKeyRaw: string, diskEncryptionKeySha256: string, initializeParams: { image: string, size: number, type: string }, source: string }>;
    public readonly canIpForward: pulumi.Output<boolean>;
    public /*out*/ readonly cpuPlatform: pulumi.Output<string>;
    public readonly deletionProtection: pulumi.Output<boolean>;
    public readonly description: pulumi.Output<string>;
    public readonly guestAccelerators: pulumi.Output<{ count: number, type: string }[]>;
    public readonly hostname: pulumi.Output<string>;
    public /*out*/ readonly instanceId: pulumi.Output<string>;
    public /*out*/ readonly labelFingerprint: pulumi.Output<string>;
    public readonly labels: pulumi.Output<{[key: string]: string}>;
    public readonly machineType: pulumi.Output<string>;
    public readonly metadata: pulumi.Output<{[key: string]: string}>;
    public /*out*/ readonly metadataFingerprint: pulumi.Output<string>;
    public readonly metadataStartupScript: pulumi.Output<string>;
    public readonly minCpuPlatform: pulumi.Output<string>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    public readonly name: pulumi.Output<string>;
    public readonly networkInterfaces: pulumi.Output<{ accessConfigs: { natIp: string, networkTier: string, publicPtrDomainName: string }[], aliasIpRanges: { ipCidrRange: string, subnetworkRangeName: string }[], name: string, network: string, networkIp: string, subnetwork: string, subnetworkProject: string }[]>;
    public readonly project: pulumi.Output<string>;
    public readonly scheduling: pulumi.Output<{ automaticRestart: boolean, onHostMaintenance: string, preemptible: boolean }>;
    public readonly scratchDisks: pulumi.Output<{ interface: string }[]>;
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    public readonly serviceAccount: pulumi.Output<{ email: string, scopes: string[] }>;
    /**
     * Name or self link of an instance
     * template to create the instance based on.
     */
    public readonly sourceInstanceTemplate: pulumi.Output<string>;
    public readonly tags: pulumi.Output<string[]>;
    public /*out*/ readonly tagsFingerprint: pulumi.Output<string>;
    /**
     * The zone that the machine should be created in. If not
     * set, the provider zone is used.
     */
    public readonly zone: pulumi.Output<string>;

    /**
     * Create a InstanceFromTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceFromTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceFromTemplateArgs | InstanceFromTemplateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: InstanceFromTemplateState = argsOrState as InstanceFromTemplateState | undefined;
            inputs["allowStoppingForUpdate"] = state ? state.allowStoppingForUpdate : undefined;
            inputs["attachedDisks"] = state ? state.attachedDisks : undefined;
            inputs["bootDisk"] = state ? state.bootDisk : undefined;
            inputs["canIpForward"] = state ? state.canIpForward : undefined;
            inputs["cpuPlatform"] = state ? state.cpuPlatform : undefined;
            inputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["guestAccelerators"] = state ? state.guestAccelerators : undefined;
            inputs["hostname"] = state ? state.hostname : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["labelFingerprint"] = state ? state.labelFingerprint : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["machineType"] = state ? state.machineType : undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["metadataFingerprint"] = state ? state.metadataFingerprint : undefined;
            inputs["metadataStartupScript"] = state ? state.metadataStartupScript : undefined;
            inputs["minCpuPlatform"] = state ? state.minCpuPlatform : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkInterfaces"] = state ? state.networkInterfaces : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["scheduling"] = state ? state.scheduling : undefined;
            inputs["scratchDisks"] = state ? state.scratchDisks : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["serviceAccount"] = state ? state.serviceAccount : undefined;
            inputs["sourceInstanceTemplate"] = state ? state.sourceInstanceTemplate : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsFingerprint"] = state ? state.tagsFingerprint : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceFromTemplateArgs | undefined;
            if (!args || args.sourceInstanceTemplate === undefined) {
                throw new Error("Missing required property 'sourceInstanceTemplate'");
            }
            inputs["allowStoppingForUpdate"] = args ? args.allowStoppingForUpdate : undefined;
            inputs["attachedDisks"] = args ? args.attachedDisks : undefined;
            inputs["bootDisk"] = args ? args.bootDisk : undefined;
            inputs["canIpForward"] = args ? args.canIpForward : undefined;
            inputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["guestAccelerators"] = args ? args.guestAccelerators : undefined;
            inputs["hostname"] = args ? args.hostname : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["machineType"] = args ? args.machineType : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["metadataStartupScript"] = args ? args.metadataStartupScript : undefined;
            inputs["minCpuPlatform"] = args ? args.minCpuPlatform : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkInterfaces"] = args ? args.networkInterfaces : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["scheduling"] = args ? args.scheduling : undefined;
            inputs["scratchDisks"] = args ? args.scratchDisks : undefined;
            inputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            inputs["sourceInstanceTemplate"] = args ? args.sourceInstanceTemplate : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["cpuPlatform"] = undefined /*out*/;
            inputs["instanceId"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["metadataFingerprint"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["tagsFingerprint"] = undefined /*out*/;
        }
        super("gcp:compute/instanceFromTemplate:InstanceFromTemplate", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceFromTemplate resources.
 */
export interface InstanceFromTemplateState {
    readonly allowStoppingForUpdate?: pulumi.Input<boolean>;
    readonly attachedDisks?: pulumi.Input<pulumi.Input<{ deviceName?: pulumi.Input<string>, diskEncryptionKeyRaw?: pulumi.Input<string>, diskEncryptionKeySha256?: pulumi.Input<string>, mode?: pulumi.Input<string>, source: pulumi.Input<string> }>[]>;
    readonly bootDisk?: pulumi.Input<{ autoDelete?: pulumi.Input<boolean>, deviceName?: pulumi.Input<string>, diskEncryptionKeyRaw?: pulumi.Input<string>, diskEncryptionKeySha256?: pulumi.Input<string>, initializeParams?: pulumi.Input<{ image?: pulumi.Input<string>, size?: pulumi.Input<number>, type?: pulumi.Input<string> }>, source?: pulumi.Input<string> }>;
    readonly canIpForward?: pulumi.Input<boolean>;
    readonly cpuPlatform?: pulumi.Input<string>;
    readonly deletionProtection?: pulumi.Input<boolean>;
    readonly description?: pulumi.Input<string>;
    readonly guestAccelerators?: pulumi.Input<pulumi.Input<{ count: pulumi.Input<number>, type: pulumi.Input<string> }>[]>;
    readonly hostname?: pulumi.Input<string>;
    readonly instanceId?: pulumi.Input<string>;
    readonly labelFingerprint?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly machineType?: pulumi.Input<string>;
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly metadataFingerprint?: pulumi.Input<string>;
    readonly metadataStartupScript?: pulumi.Input<string>;
    readonly minCpuPlatform?: pulumi.Input<string>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    readonly networkInterfaces?: pulumi.Input<pulumi.Input<{ accessConfigs?: pulumi.Input<pulumi.Input<{ natIp?: pulumi.Input<string>, networkTier?: pulumi.Input<string>, publicPtrDomainName?: pulumi.Input<string> }>[]>, aliasIpRanges?: pulumi.Input<pulumi.Input<{ ipCidrRange: pulumi.Input<string>, subnetworkRangeName?: pulumi.Input<string> }>[]>, name?: pulumi.Input<string>, network?: pulumi.Input<string>, networkIp?: pulumi.Input<string>, subnetwork?: pulumi.Input<string>, subnetworkProject?: pulumi.Input<string> }>[]>;
    readonly project?: pulumi.Input<string>;
    readonly scheduling?: pulumi.Input<{ automaticRestart?: pulumi.Input<boolean>, onHostMaintenance?: pulumi.Input<string>, preemptible?: pulumi.Input<boolean> }>;
    readonly scratchDisks?: pulumi.Input<pulumi.Input<{ interface?: pulumi.Input<string> }>[]>;
    readonly selfLink?: pulumi.Input<string>;
    readonly serviceAccount?: pulumi.Input<{ email?: pulumi.Input<string>, scopes: pulumi.Input<pulumi.Input<string>[]> }>;
    /**
     * Name or self link of an instance
     * template to create the instance based on.
     */
    readonly sourceInstanceTemplate?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    readonly tagsFingerprint?: pulumi.Input<string>;
    /**
     * The zone that the machine should be created in. If not
     * set, the provider zone is used.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceFromTemplate resource.
 */
export interface InstanceFromTemplateArgs {
    readonly allowStoppingForUpdate?: pulumi.Input<boolean>;
    readonly attachedDisks?: pulumi.Input<pulumi.Input<{ deviceName?: pulumi.Input<string>, diskEncryptionKeyRaw?: pulumi.Input<string>, diskEncryptionKeySha256?: pulumi.Input<string>, mode?: pulumi.Input<string>, source: pulumi.Input<string> }>[]>;
    readonly bootDisk?: pulumi.Input<{ autoDelete?: pulumi.Input<boolean>, deviceName?: pulumi.Input<string>, diskEncryptionKeyRaw?: pulumi.Input<string>, diskEncryptionKeySha256?: pulumi.Input<string>, initializeParams?: pulumi.Input<{ image?: pulumi.Input<string>, size?: pulumi.Input<number>, type?: pulumi.Input<string> }>, source?: pulumi.Input<string> }>;
    readonly canIpForward?: pulumi.Input<boolean>;
    readonly deletionProtection?: pulumi.Input<boolean>;
    readonly description?: pulumi.Input<string>;
    readonly guestAccelerators?: pulumi.Input<pulumi.Input<{ count: pulumi.Input<number>, type: pulumi.Input<string> }>[]>;
    readonly hostname?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly machineType?: pulumi.Input<string>;
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly metadataStartupScript?: pulumi.Input<string>;
    readonly minCpuPlatform?: pulumi.Input<string>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    readonly networkInterfaces?: pulumi.Input<pulumi.Input<{ accessConfigs?: pulumi.Input<pulumi.Input<{ natIp?: pulumi.Input<string>, networkTier?: pulumi.Input<string>, publicPtrDomainName?: pulumi.Input<string> }>[]>, aliasIpRanges?: pulumi.Input<pulumi.Input<{ ipCidrRange: pulumi.Input<string>, subnetworkRangeName?: pulumi.Input<string> }>[]>, name?: pulumi.Input<string>, network?: pulumi.Input<string>, networkIp?: pulumi.Input<string>, subnetwork?: pulumi.Input<string>, subnetworkProject?: pulumi.Input<string> }>[]>;
    readonly project?: pulumi.Input<string>;
    readonly scheduling?: pulumi.Input<{ automaticRestart?: pulumi.Input<boolean>, onHostMaintenance?: pulumi.Input<string>, preemptible?: pulumi.Input<boolean> }>;
    readonly scratchDisks?: pulumi.Input<pulumi.Input<{ interface?: pulumi.Input<string> }>[]>;
    readonly serviceAccount?: pulumi.Input<{ email?: pulumi.Input<string>, scopes: pulumi.Input<pulumi.Input<string>[]> }>;
    /**
     * Name or self link of an instance
     * template to create the instance based on.
     */
    readonly sourceInstanceTemplate: pulumi.Input<string>;
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The zone that the machine should be created in. If not
     * set, the provider zone is used.
     */
    readonly zone?: pulumi.Input<string>;
}

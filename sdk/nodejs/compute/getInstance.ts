// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getInstance(args: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {
    return pulumi.runtime.invoke("gcp:compute/getInstance:getInstance", {
        "name": args.name,
        "project": args.project,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceArgs {
    readonly name: string;
    readonly project?: string;
    readonly zone?: string;
}

/**
 * A collection of values returned by getInstance.
 */
export interface GetInstanceResult {
    readonly allowStoppingForUpdate: boolean;
    readonly attachedDisks: { deviceName: string, diskEncryptionKeyRaw: string, diskEncryptionKeySha256: string, mode: string, source: string }[];
    readonly bootDisks: { autoDelete: boolean, deviceName: string, diskEncryptionKeyRaw: string, diskEncryptionKeySha256: string, initializeParams: { image: string, size: number, type: string }[], source: string }[];
    readonly canIpForward: boolean;
    readonly cpuPlatform: string;
    readonly createTimeout: number;
    readonly deletionProtection: boolean;
    readonly description: string;
    readonly disks: { autoDelete: boolean, deviceName: string, disk: string, diskEncryptionKeyRaw: string, diskEncryptionKeySha256: string, image: string, scratch: boolean, size: number, type: string }[];
    readonly guestAccelerators: { count: number, type: string }[];
    readonly instanceId: string;
    readonly labelFingerprint: string;
    readonly labels: {[key: string]: string};
    readonly machineType: string;
    readonly metadata: {[key: string]: string};
    readonly metadataFingerprint: string;
    readonly metadataStartupScript: string;
    readonly minCpuPlatform: string;
    readonly networks: { address: string, externalAddress: string, internalAddress: string, name: string, source: string }[];
    readonly networkInterfaces: { accessConfigs: { assignedNatIp: string, natIp: string, networkTier: string, publicPtrDomainName: string }[], address: string, aliasIpRanges: { ipCidrRange: string, subnetworkRangeName: string }[], name: string, network: string, networkIp: string, subnetwork: string, subnetworkProject: string }[];
    readonly schedulings: { automaticRestart: boolean, onHostMaintenance: string, preemptible: boolean }[];
    readonly scratchDisks: { interface: string }[];
    readonly selfLink: string;
    readonly serviceAccounts: { email: string, scopes: string[] }[];
    readonly tags: string[];
    readonly tagsFingerprint: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
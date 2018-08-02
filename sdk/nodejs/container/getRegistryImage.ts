// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * This data source fetches the project name, and provides the appropriate URLs to use for container registry for this project.
 * 
 * The URLs are computed entirely offline - as long as the project exists, they will be valid, but this data source does not contact Google Container Registry (GCR) at any point.
 */
export function getRegistryImage(args: GetRegistryImageArgs): Promise<GetRegistryImageResult> {
    return pulumi.runtime.invoke("gcp:container/getRegistryImage:getRegistryImage", {
        "digest": args.digest,
        "name": args.name,
        "project": args.project,
        "region": args.region,
        "tag": args.tag,
    });
}

/**
 * A collection of arguments for invoking getRegistryImage.
 */
export interface GetRegistryImageArgs {
    readonly digest?: string;
    readonly name: string;
    readonly project?: string;
    readonly region?: string;
    readonly tag?: string;
}

/**
 * A collection of values returned by getRegistryImage.
 */
export interface GetRegistryImageResult {
    readonly imageUrl: string;
    readonly project: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}

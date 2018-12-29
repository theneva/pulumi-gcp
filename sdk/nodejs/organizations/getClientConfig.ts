// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access the configuration of the Google Cloud provider.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const google_client_config_current = pulumi.output(gcp.organizations.getClientConfig({}));
 * 
 * export const project = google_client_config_current.apply(__arg0 => __arg0.project);
 * ```
 */
export function getClientConfig(opts?: pulumi.InvokeOptions): Promise<GetClientConfigResult> {
    return pulumi.runtime.invoke("gcp:organizations/getClientConfig:getClientConfig", {
    }, opts);
}

/**
 * A collection of values returned by getClientConfig.
 */
export interface GetClientConfigResult {
    /**
     * The OAuth2 access token used by the client to authenticate against the Google Cloud API.
     */
    readonly accessToken: string;
    /**
     * The ID of the project to apply any resources to.
     */
    readonly project: string;
    /**
     * The region to operate under.
     */
    readonly region: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}

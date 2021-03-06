// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Organizations
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get information about a Google Cloud Organization.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/organization.html.markdown.
        /// </summary>
        public static Task<GetOrganizationResult> GetOrganization(GetOrganizationArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationResult>("gcp:organizations/getOrganization:getOrganization", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetOrganizationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The domain name of the Organization.
        /// </summary>
        [Input("domain")]
        public string? Domain { get; set; }

        /// <summary>
        /// The name of the Organization in the form `{organization_id}` or `organizations/{organization_id}`.
        /// </summary>
        [Input("organization")]
        public string? Organization { get; set; }

        public GetOrganizationArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetOrganizationResult
    {
        /// <summary>
        /// Timestamp when the Organization was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The Google for Work customer ID of the Organization.
        /// </summary>
        public readonly string DirectoryCustomerId;
        public readonly string Domain;
        /// <summary>
        /// The Organization's current lifecycle state.
        /// </summary>
        public readonly string LifecycleState;
        /// <summary>
        /// The resource name of the Organization in the form `organizations/{organization_id}`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Organization ID.
        /// </summary>
        public readonly string OrgId;
        public readonly string? Organization;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetOrganizationResult(
            string createTime,
            string directoryCustomerId,
            string domain,
            string lifecycleState,
            string name,
            string orgId,
            string? organization,
            string id)
        {
            CreateTime = createTime;
            DirectoryCustomerId = directoryCustomerId;
            Domain = domain;
            LifecycleState = lifecycleState;
            Name = name;
            OrgId = orgId;
            Organization = organization;
            Id = id;
        }
    }
}

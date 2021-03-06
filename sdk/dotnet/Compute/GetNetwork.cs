// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static partial class Invokes
    {
        /// <summary>
        /// Get a network within GCE from its name.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/compute_network.html.markdown.
        /// </summary>
        public static Task<GetNetworkResult> GetNetwork(GetNetworkArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkResult>("gcp:compute/getNetwork:getNetwork", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetNetworkArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the network.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetNetworkArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetNetworkResult
    {
        /// <summary>
        /// Description of this network.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The IP address of the gateway.
        /// </summary>
        public readonly string GatewayIpv4;
        public readonly string Name;
        public readonly string? Project;
        /// <summary>
        /// The URI of the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// the list of subnetworks which belong to the network
        /// </summary>
        public readonly ImmutableArray<string> SubnetworksSelfLinks;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetNetworkResult(
            string description,
            string gatewayIpv4,
            string name,
            string? project,
            string selfLink,
            ImmutableArray<string> subnetworksSelfLinks,
            string id)
        {
            Description = description;
            GatewayIpv4 = gatewayIpv4;
            Name = name;
            Project = project;
            SelfLink = selfLink;
            SubnetworksSelfLinks = subnetworksSelfLinks;
            Id = id;
        }
    }
}

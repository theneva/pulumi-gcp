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
        /// Provides access to available Google Compute regions for a given project.
        /// See more about [regions and regions](https://cloud.google.com/compute/docs/regions-zones/) in the upstream docs.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/compute_regions.html.markdown.
        /// </summary>
        public static Task<GetRegionsResult> GetRegions(GetRegionsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegionsResult>("gcp:compute/getRegions:getRegions", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetRegionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Project from which to list available regions. Defaults to project declared in the provider.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// Allows to filter list of regions based on their current status. Status can be either `UP` or `DOWN`.
        /// Defaults to no filtering (all available regions - both `UP` and `DOWN`).
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetRegionsArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetRegionsResult
    {
        /// <summary>
        /// A list of regions available in the given project
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string Project;
        public readonly string? Status;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetRegionsResult(
            ImmutableArray<string> names,
            string project,
            string? status,
            string id)
        {
            Names = names;
            Project = project;
            Status = status;
            Id = id;
        }
    }
}

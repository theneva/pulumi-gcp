// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Appengine
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/app_engine_application_url_dispatch_rules.html.markdown.
    /// </summary>
    public partial class ApplicationUrlDispatchRules : Pulumi.CustomResource
    {
        /// <summary>
        /// Rules to match an HTTP request and dispatch that request to a service.
        /// </summary>
        [Output("dispatchRules")]
        public Output<ImmutableArray<Outputs.ApplicationUrlDispatchRulesDispatchRules>> DispatchRules { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationUrlDispatchRules resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationUrlDispatchRules(string name, ApplicationUrlDispatchRulesArgs args, CustomResourceOptions? options = null)
            : base("gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules", name, args, MakeResourceOptions(options, ""))
        {
        }

        private ApplicationUrlDispatchRules(string name, Input<string> id, ApplicationUrlDispatchRulesState? state = null, CustomResourceOptions? options = null)
            : base("gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ApplicationUrlDispatchRules resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationUrlDispatchRules Get(string name, Input<string> id, ApplicationUrlDispatchRulesState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationUrlDispatchRules(name, id, state, options);
        }
    }

    public sealed class ApplicationUrlDispatchRulesArgs : Pulumi.ResourceArgs
    {
        [Input("dispatchRules", required: true)]
        private InputList<Inputs.ApplicationUrlDispatchRulesDispatchRulesArgs>? _dispatchRules;

        /// <summary>
        /// Rules to match an HTTP request and dispatch that request to a service.
        /// </summary>
        public InputList<Inputs.ApplicationUrlDispatchRulesDispatchRulesArgs> DispatchRules
        {
            get => _dispatchRules ?? (_dispatchRules = new InputList<Inputs.ApplicationUrlDispatchRulesDispatchRulesArgs>());
            set => _dispatchRules = value;
        }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public ApplicationUrlDispatchRulesArgs()
        {
        }
    }

    public sealed class ApplicationUrlDispatchRulesState : Pulumi.ResourceArgs
    {
        [Input("dispatchRules")]
        private InputList<Inputs.ApplicationUrlDispatchRulesDispatchRulesGetArgs>? _dispatchRules;

        /// <summary>
        /// Rules to match an HTTP request and dispatch that request to a service.
        /// </summary>
        public InputList<Inputs.ApplicationUrlDispatchRulesDispatchRulesGetArgs> DispatchRules
        {
            get => _dispatchRules ?? (_dispatchRules = new InputList<Inputs.ApplicationUrlDispatchRulesDispatchRulesGetArgs>());
            set => _dispatchRules = value;
        }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public ApplicationUrlDispatchRulesState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ApplicationUrlDispatchRulesDispatchRulesArgs : Pulumi.ResourceArgs
    {
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public ApplicationUrlDispatchRulesDispatchRulesArgs()
        {
        }
    }

    public sealed class ApplicationUrlDispatchRulesDispatchRulesGetArgs : Pulumi.ResourceArgs
    {
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public ApplicationUrlDispatchRulesDispatchRulesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ApplicationUrlDispatchRulesDispatchRules
    {
        public readonly string? Domain;
        public readonly string Path;
        public readonly string Service;

        [OutputConstructor]
        private ApplicationUrlDispatchRulesDispatchRules(
            string? domain,
            string path,
            string service)
        {
            Domain = domain;
            Path = path;
            Service = service;
        }
    }
    }
}

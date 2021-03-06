// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AccessContextManager
{
    /// <summary>
    /// An AccessLevel is a label that can be applied to requests to GCP services,
    /// along with a list of requirements necessary for the label to be applied.
    /// 
    /// 
    /// To get more information about AccessLevel, see:
    /// 
    /// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.accessLevels)
    /// * How-to Guides
    ///     * [Access Policy Quickstart](https://cloud.google.com/access-context-manager/docs/quickstart)
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/access_context_manager_access_level.html.markdown.
    /// </summary>
    public partial class AccessLevel : Pulumi.CustomResource
    {
        /// <summary>
        /// A set of predefined conditions for the access level and a combining function.
        /// </summary>
        [Output("basic")]
        public Output<Outputs.AccessLevelBasic?> Basic { get; private set; } = null!;

        /// <summary>
        /// Description of the AccessLevel and its use. Does not affect behavior.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Resource name for the Access Level. The short_name component must begin with a letter and only include
        /// alphanumeric and '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}
        /// </summary>
        [Output("parent")]
        public Output<string> Parent { get; private set; } = null!;

        /// <summary>
        /// Human readable title. Must be unique within the Policy.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;


        /// <summary>
        /// Create a AccessLevel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessLevel(string name, AccessLevelArgs args, CustomResourceOptions? options = null)
            : base("gcp:accesscontextmanager/accessLevel:AccessLevel", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private AccessLevel(string name, Input<string> id, AccessLevelState? state = null, CustomResourceOptions? options = null)
            : base("gcp:accesscontextmanager/accessLevel:AccessLevel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccessLevel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessLevel Get(string name, Input<string> id, AccessLevelState? state = null, CustomResourceOptions? options = null)
        {
            return new AccessLevel(name, id, state, options);
        }
    }

    public sealed class AccessLevelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A set of predefined conditions for the access level and a combining function.
        /// </summary>
        [Input("basic")]
        public Input<Inputs.AccessLevelBasicArgs>? Basic { get; set; }

        /// <summary>
        /// Description of the AccessLevel and its use. Does not affect behavior.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Resource name for the Access Level. The short_name component must begin with a letter and only include
        /// alphanumeric and '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name}
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}
        /// </summary>
        [Input("parent", required: true)]
        public Input<string> Parent { get; set; } = null!;

        /// <summary>
        /// Human readable title. Must be unique within the Policy.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public AccessLevelArgs()
        {
        }
    }

    public sealed class AccessLevelState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A set of predefined conditions for the access level and a combining function.
        /// </summary>
        [Input("basic")]
        public Input<Inputs.AccessLevelBasicGetArgs>? Basic { get; set; }

        /// <summary>
        /// Description of the AccessLevel and its use. Does not affect behavior.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Resource name for the Access Level. The short_name component must begin with a letter and only include
        /// alphanumeric and '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name}
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        /// <summary>
        /// Human readable title. Must be unique within the Policy.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public AccessLevelState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class AccessLevelBasicArgs : Pulumi.ResourceArgs
    {
        [Input("combiningFunction")]
        public Input<string>? CombiningFunction { get; set; }

        [Input("conditions", required: true)]
        private InputList<AccessLevelBasicConditionsArgs>? _conditions;
        public InputList<AccessLevelBasicConditionsArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<AccessLevelBasicConditionsArgs>());
            set => _conditions = value;
        }

        public AccessLevelBasicArgs()
        {
        }
    }

    public sealed class AccessLevelBasicConditionsArgs : Pulumi.ResourceArgs
    {
        [Input("devicePolicy")]
        public Input<AccessLevelBasicConditionsDevicePolicyArgs>? DevicePolicy { get; set; }

        [Input("ipSubnetworks")]
        private InputList<string>? _ipSubnetworks;
        public InputList<string> IpSubnetworks
        {
            get => _ipSubnetworks ?? (_ipSubnetworks = new InputList<string>());
            set => _ipSubnetworks = value;
        }

        [Input("members")]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        [Input("negate")]
        public Input<bool>? Negate { get; set; }

        [Input("requiredAccessLevels")]
        private InputList<string>? _requiredAccessLevels;
        public InputList<string> RequiredAccessLevels
        {
            get => _requiredAccessLevels ?? (_requiredAccessLevels = new InputList<string>());
            set => _requiredAccessLevels = value;
        }

        public AccessLevelBasicConditionsArgs()
        {
        }
    }

    public sealed class AccessLevelBasicConditionsDevicePolicyArgs : Pulumi.ResourceArgs
    {
        [Input("allowedDeviceManagementLevels")]
        private InputList<string>? _allowedDeviceManagementLevels;
        public InputList<string> AllowedDeviceManagementLevels
        {
            get => _allowedDeviceManagementLevels ?? (_allowedDeviceManagementLevels = new InputList<string>());
            set => _allowedDeviceManagementLevels = value;
        }

        [Input("allowedEncryptionStatuses")]
        private InputList<string>? _allowedEncryptionStatuses;
        public InputList<string> AllowedEncryptionStatuses
        {
            get => _allowedEncryptionStatuses ?? (_allowedEncryptionStatuses = new InputList<string>());
            set => _allowedEncryptionStatuses = value;
        }

        [Input("osConstraints")]
        private InputList<AccessLevelBasicConditionsDevicePolicyOsConstraintsArgs>? _osConstraints;
        public InputList<AccessLevelBasicConditionsDevicePolicyOsConstraintsArgs> OsConstraints
        {
            get => _osConstraints ?? (_osConstraints = new InputList<AccessLevelBasicConditionsDevicePolicyOsConstraintsArgs>());
            set => _osConstraints = value;
        }

        [Input("requireAdminApproval")]
        public Input<bool>? RequireAdminApproval { get; set; }

        [Input("requireCorpOwned")]
        public Input<bool>? RequireCorpOwned { get; set; }

        [Input("requireScreenLock")]
        public Input<bool>? RequireScreenLock { get; set; }

        public AccessLevelBasicConditionsDevicePolicyArgs()
        {
        }
    }

    public sealed class AccessLevelBasicConditionsDevicePolicyGetArgs : Pulumi.ResourceArgs
    {
        [Input("allowedDeviceManagementLevels")]
        private InputList<string>? _allowedDeviceManagementLevels;
        public InputList<string> AllowedDeviceManagementLevels
        {
            get => _allowedDeviceManagementLevels ?? (_allowedDeviceManagementLevels = new InputList<string>());
            set => _allowedDeviceManagementLevels = value;
        }

        [Input("allowedEncryptionStatuses")]
        private InputList<string>? _allowedEncryptionStatuses;
        public InputList<string> AllowedEncryptionStatuses
        {
            get => _allowedEncryptionStatuses ?? (_allowedEncryptionStatuses = new InputList<string>());
            set => _allowedEncryptionStatuses = value;
        }

        [Input("osConstraints")]
        private InputList<AccessLevelBasicConditionsDevicePolicyOsConstraintsGetArgs>? _osConstraints;
        public InputList<AccessLevelBasicConditionsDevicePolicyOsConstraintsGetArgs> OsConstraints
        {
            get => _osConstraints ?? (_osConstraints = new InputList<AccessLevelBasicConditionsDevicePolicyOsConstraintsGetArgs>());
            set => _osConstraints = value;
        }

        [Input("requireAdminApproval")]
        public Input<bool>? RequireAdminApproval { get; set; }

        [Input("requireCorpOwned")]
        public Input<bool>? RequireCorpOwned { get; set; }

        [Input("requireScreenLock")]
        public Input<bool>? RequireScreenLock { get; set; }

        public AccessLevelBasicConditionsDevicePolicyGetArgs()
        {
        }
    }

    public sealed class AccessLevelBasicConditionsDevicePolicyOsConstraintsArgs : Pulumi.ResourceArgs
    {
        [Input("minimumVersion")]
        public Input<string>? MinimumVersion { get; set; }

        [Input("osType", required: true)]
        public Input<string> OsType { get; set; } = null!;

        public AccessLevelBasicConditionsDevicePolicyOsConstraintsArgs()
        {
        }
    }

    public sealed class AccessLevelBasicConditionsDevicePolicyOsConstraintsGetArgs : Pulumi.ResourceArgs
    {
        [Input("minimumVersion")]
        public Input<string>? MinimumVersion { get; set; }

        [Input("osType", required: true)]
        public Input<string> OsType { get; set; } = null!;

        public AccessLevelBasicConditionsDevicePolicyOsConstraintsGetArgs()
        {
        }
    }

    public sealed class AccessLevelBasicConditionsGetArgs : Pulumi.ResourceArgs
    {
        [Input("devicePolicy")]
        public Input<AccessLevelBasicConditionsDevicePolicyGetArgs>? DevicePolicy { get; set; }

        [Input("ipSubnetworks")]
        private InputList<string>? _ipSubnetworks;
        public InputList<string> IpSubnetworks
        {
            get => _ipSubnetworks ?? (_ipSubnetworks = new InputList<string>());
            set => _ipSubnetworks = value;
        }

        [Input("members")]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        [Input("negate")]
        public Input<bool>? Negate { get; set; }

        [Input("requiredAccessLevels")]
        private InputList<string>? _requiredAccessLevels;
        public InputList<string> RequiredAccessLevels
        {
            get => _requiredAccessLevels ?? (_requiredAccessLevels = new InputList<string>());
            set => _requiredAccessLevels = value;
        }

        public AccessLevelBasicConditionsGetArgs()
        {
        }
    }

    public sealed class AccessLevelBasicGetArgs : Pulumi.ResourceArgs
    {
        [Input("combiningFunction")]
        public Input<string>? CombiningFunction { get; set; }

        [Input("conditions", required: true)]
        private InputList<AccessLevelBasicConditionsGetArgs>? _conditions;
        public InputList<AccessLevelBasicConditionsGetArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<AccessLevelBasicConditionsGetArgs>());
            set => _conditions = value;
        }

        public AccessLevelBasicGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class AccessLevelBasic
    {
        public readonly string? CombiningFunction;
        public readonly ImmutableArray<AccessLevelBasicConditions> Conditions;

        [OutputConstructor]
        private AccessLevelBasic(
            string? combiningFunction,
            ImmutableArray<AccessLevelBasicConditions> conditions)
        {
            CombiningFunction = combiningFunction;
            Conditions = conditions;
        }
    }

    [OutputType]
    public sealed class AccessLevelBasicConditions
    {
        public readonly AccessLevelBasicConditionsDevicePolicy? DevicePolicy;
        public readonly ImmutableArray<string> IpSubnetworks;
        public readonly ImmutableArray<string> Members;
        public readonly bool? Negate;
        public readonly ImmutableArray<string> RequiredAccessLevels;

        [OutputConstructor]
        private AccessLevelBasicConditions(
            AccessLevelBasicConditionsDevicePolicy? devicePolicy,
            ImmutableArray<string> ipSubnetworks,
            ImmutableArray<string> members,
            bool? negate,
            ImmutableArray<string> requiredAccessLevels)
        {
            DevicePolicy = devicePolicy;
            IpSubnetworks = ipSubnetworks;
            Members = members;
            Negate = negate;
            RequiredAccessLevels = requiredAccessLevels;
        }
    }

    [OutputType]
    public sealed class AccessLevelBasicConditionsDevicePolicy
    {
        public readonly ImmutableArray<string> AllowedDeviceManagementLevels;
        public readonly ImmutableArray<string> AllowedEncryptionStatuses;
        public readonly ImmutableArray<AccessLevelBasicConditionsDevicePolicyOsConstraints> OsConstraints;
        public readonly bool? RequireAdminApproval;
        public readonly bool? RequireCorpOwned;
        public readonly bool? RequireScreenLock;

        [OutputConstructor]
        private AccessLevelBasicConditionsDevicePolicy(
            ImmutableArray<string> allowedDeviceManagementLevels,
            ImmutableArray<string> allowedEncryptionStatuses,
            ImmutableArray<AccessLevelBasicConditionsDevicePolicyOsConstraints> osConstraints,
            bool? requireAdminApproval,
            bool? requireCorpOwned,
            bool? requireScreenLock)
        {
            AllowedDeviceManagementLevels = allowedDeviceManagementLevels;
            AllowedEncryptionStatuses = allowedEncryptionStatuses;
            OsConstraints = osConstraints;
            RequireAdminApproval = requireAdminApproval;
            RequireCorpOwned = requireCorpOwned;
            RequireScreenLock = requireScreenLock;
        }
    }

    [OutputType]
    public sealed class AccessLevelBasicConditionsDevicePolicyOsConstraints
    {
        public readonly string? MinimumVersion;
        public readonly string OsType;

        [OutputConstructor]
        private AccessLevelBasicConditionsDevicePolicyOsConstraints(
            string? minimumVersion,
            string osType)
        {
            MinimumVersion = minimumVersion;
            OsType = osType;
        }
    }
    }
}

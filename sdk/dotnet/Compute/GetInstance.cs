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
        /// Get information about a VM instance resource within GCE. For more information see
        /// [the official documentation](https://cloud.google.com/compute/docs/instances)
        /// and
        /// [API](https://cloud.google.com/compute/docs/reference/latest/instances).
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/compute_instance.html.markdown.
        /// </summary>
        public static Task<GetInstanceResult> GetInstance(GetInstanceArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("gcp:compute/getInstance:getInstance", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetInstanceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the instance. One of `name` or `self_link` must be provided.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If `self_link` is provided, this value is ignored.  If neither `self_link`
        /// nor `project` are provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// The self link of the instance. One of `name` or `self_link` must be provided.
        /// </summary>
        [Input("selfLink")]
        public string? SelfLink { get; set; }

        /// <summary>
        /// The zone of the instance. If `self_link` is provided, this
        /// value is ignored.  If neither `self_link` nor `zone` are provided, the
        /// provider zone is used.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetInstanceArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetInstanceResult
    {
        public readonly bool AllowStoppingForUpdate;
        /// <summary>
        /// List of disks attached to the instance. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceAttachedDisksResult> AttachedDisks;
        /// <summary>
        /// The boot disk for the instance. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceBootDisksResult> BootDisks;
        /// <summary>
        /// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
        /// </summary>
        public readonly bool CanIpForward;
        /// <summary>
        /// The CPU platform used by this instance.
        /// </summary>
        public readonly string CpuPlatform;
        /// <summary>
        /// Whether deletion protection is enabled on this instance.
        /// </summary>
        public readonly bool DeletionProtection;
        /// <summary>
        /// A brief description of the resource.
        /// </summary>
        public readonly string Description;
        public readonly bool EnableDisplay;
        /// <summary>
        /// List of the type and count of accelerator cards attached to the instance. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceGuestAcceleratorsResult> GuestAccelerators;
        public readonly string Hostname;
        /// <summary>
        /// The server-assigned unique identifier of this instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The unique fingerprint of the labels.
        /// </summary>
        public readonly string LabelFingerprint;
        /// <summary>
        /// A set of key/value label pairs assigned to the instance.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The machine type to create.
        /// </summary>
        public readonly string MachineType;
        /// <summary>
        /// Metadata key/value pairs made available within the instance.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// The unique fingerprint of the metadata.
        /// </summary>
        public readonly string MetadataFingerprint;
        public readonly string MetadataStartupScript;
        /// <summary>
        /// The minimum CPU platform specified for the VM instance.
        /// </summary>
        public readonly string MinCpuPlatform;
        public readonly string? Name;
        /// <summary>
        /// The networks attached to the instance. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceNetworkInterfacesResult> NetworkInterfaces;
        public readonly string? Project;
        /// <summary>
        /// The scheduling strategy being used by the instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceSchedulingsResult> Schedulings;
        /// <summary>
        /// The scratch disks attached to the instance. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceScratchDisksResult> ScratchDisks;
        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        public readonly string? SelfLink;
        /// <summary>
        /// The service account to attach to the instance. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceServiceAccountsResult> ServiceAccounts;
        /// <summary>
        /// The shielded vm config being used by the instance. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceShieldedInstanceConfigsResult> ShieldedInstanceConfigs;
        /// <summary>
        /// The list of tags attached to the instance.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The unique fingerprint of the tags.
        /// </summary>
        public readonly string TagsFingerprint;
        public readonly string? Zone;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetInstanceResult(
            bool allowStoppingForUpdate,
            ImmutableArray<Outputs.GetInstanceAttachedDisksResult> attachedDisks,
            ImmutableArray<Outputs.GetInstanceBootDisksResult> bootDisks,
            bool canIpForward,
            string cpuPlatform,
            bool deletionProtection,
            string description,
            bool enableDisplay,
            ImmutableArray<Outputs.GetInstanceGuestAcceleratorsResult> guestAccelerators,
            string hostname,
            string instanceId,
            string labelFingerprint,
            ImmutableDictionary<string, string> labels,
            string machineType,
            ImmutableDictionary<string, string> metadata,
            string metadataFingerprint,
            string metadataStartupScript,
            string minCpuPlatform,
            string? name,
            ImmutableArray<Outputs.GetInstanceNetworkInterfacesResult> networkInterfaces,
            string? project,
            ImmutableArray<Outputs.GetInstanceSchedulingsResult> schedulings,
            ImmutableArray<Outputs.GetInstanceScratchDisksResult> scratchDisks,
            string? selfLink,
            ImmutableArray<Outputs.GetInstanceServiceAccountsResult> serviceAccounts,
            ImmutableArray<Outputs.GetInstanceShieldedInstanceConfigsResult> shieldedInstanceConfigs,
            ImmutableArray<string> tags,
            string tagsFingerprint,
            string? zone,
            string id)
        {
            AllowStoppingForUpdate = allowStoppingForUpdate;
            AttachedDisks = attachedDisks;
            BootDisks = bootDisks;
            CanIpForward = canIpForward;
            CpuPlatform = cpuPlatform;
            DeletionProtection = deletionProtection;
            Description = description;
            EnableDisplay = enableDisplay;
            GuestAccelerators = guestAccelerators;
            Hostname = hostname;
            InstanceId = instanceId;
            LabelFingerprint = labelFingerprint;
            Labels = labels;
            MachineType = machineType;
            Metadata = metadata;
            MetadataFingerprint = metadataFingerprint;
            MetadataStartupScript = metadataStartupScript;
            MinCpuPlatform = minCpuPlatform;
            Name = name;
            NetworkInterfaces = networkInterfaces;
            Project = project;
            Schedulings = schedulings;
            ScratchDisks = scratchDisks;
            SelfLink = selfLink;
            ServiceAccounts = serviceAccounts;
            ShieldedInstanceConfigs = shieldedInstanceConfigs;
            Tags = tags;
            TagsFingerprint = tagsFingerprint;
            Zone = zone;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetInstanceAttachedDisksResult
    {
        /// <summary>
        /// Name with which the attached disk is accessible
        /// under `/dev/disk/by-id/`
        /// </summary>
        public readonly string DeviceName;
        public readonly string DiskEncryptionKeyRaw;
        public readonly string DiskEncryptionKeySha256;
        public readonly string KmsKeySelfLink;
        /// <summary>
        /// Read/write mode for the disk. One of `"READ_ONLY"` or `"READ_WRITE"`.
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// The name or self_link of the disk attached to this instance.
        /// </summary>
        public readonly string Source;

        [OutputConstructor]
        private GetInstanceAttachedDisksResult(
            string deviceName,
            string diskEncryptionKeyRaw,
            string diskEncryptionKeySha256,
            string kmsKeySelfLink,
            string mode,
            string source)
        {
            DeviceName = deviceName;
            DiskEncryptionKeyRaw = diskEncryptionKeyRaw;
            DiskEncryptionKeySha256 = diskEncryptionKeySha256;
            KmsKeySelfLink = kmsKeySelfLink;
            Mode = mode;
            Source = source;
        }
    }

    [OutputType]
    public sealed class GetInstanceBootDisksInitializeParamsResult
    {
        /// <summary>
        /// The image from which this disk was initialised.
        /// </summary>
        public readonly string Image;
        /// <summary>
        /// A set of key/value label pairs assigned to the instance.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Labels;
        /// <summary>
        /// The size of the image in gigabytes.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// The accelerator type resource exposed to this instance. E.g. `nvidia-tesla-k80`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetInstanceBootDisksInitializeParamsResult(
            string image,
            ImmutableDictionary<string, object> labels,
            int size,
            string type)
        {
            Image = image;
            Labels = labels;
            Size = size;
            Type = type;
        }
    }

    [OutputType]
    public sealed class GetInstanceBootDisksResult
    {
        /// <summary>
        /// Whether the disk will be auto-deleted when the instance is deleted.
        /// </summary>
        public readonly bool AutoDelete;
        /// <summary>
        /// Name with which the attached disk is accessible
        /// under `/dev/disk/by-id/`
        /// </summary>
        public readonly string DeviceName;
        public readonly string DiskEncryptionKeyRaw;
        public readonly string DiskEncryptionKeySha256;
        /// <summary>
        /// Parameters with which a disk was created alongside the instance.
        /// Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<GetInstanceBootDisksInitializeParamsResult> InitializeParams;
        public readonly string KmsKeySelfLink;
        /// <summary>
        /// Read/write mode for the disk. One of `"READ_ONLY"` or `"READ_WRITE"`.
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// The name or self_link of the disk attached to this instance.
        /// </summary>
        public readonly string Source;

        [OutputConstructor]
        private GetInstanceBootDisksResult(
            bool autoDelete,
            string deviceName,
            string diskEncryptionKeyRaw,
            string diskEncryptionKeySha256,
            ImmutableArray<GetInstanceBootDisksInitializeParamsResult> initializeParams,
            string kmsKeySelfLink,
            string mode,
            string source)
        {
            AutoDelete = autoDelete;
            DeviceName = deviceName;
            DiskEncryptionKeyRaw = diskEncryptionKeyRaw;
            DiskEncryptionKeySha256 = diskEncryptionKeySha256;
            InitializeParams = initializeParams;
            KmsKeySelfLink = kmsKeySelfLink;
            Mode = mode;
            Source = source;
        }
    }

    [OutputType]
    public sealed class GetInstanceGuestAcceleratorsResult
    {
        /// <summary>
        /// The number of the guest accelerator cards exposed to this instance.
        /// </summary>
        public readonly int Count;
        /// <summary>
        /// The accelerator type resource exposed to this instance. E.g. `nvidia-tesla-k80`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetInstanceGuestAcceleratorsResult(
            int count,
            string type)
        {
            Count = count;
            Type = type;
        }
    }

    [OutputType]
    public sealed class GetInstanceNetworkInterfacesAccessConfigsResult
    {
        /// <summary>
        /// The IP address that is be 1:1 mapped to the instance's
        /// network ip.
        /// </summary>
        public readonly string NatIp;
        /// <summary>
        /// The [networking tier][network-tier] used for configuring this instance. One of `PREMIUM` or `STANDARD`.
        /// </summary>
        public readonly string NetworkTier;
        /// <summary>
        /// The DNS domain name for the public PTR record.
        /// </summary>
        public readonly string PublicPtrDomainName;

        [OutputConstructor]
        private GetInstanceNetworkInterfacesAccessConfigsResult(
            string natIp,
            string networkTier,
            string publicPtrDomainName)
        {
            NatIp = natIp;
            NetworkTier = networkTier;
            PublicPtrDomainName = publicPtrDomainName;
        }
    }

    [OutputType]
    public sealed class GetInstanceNetworkInterfacesAliasIpRangesResult
    {
        /// <summary>
        /// The IP CIDR range represented by this alias IP range.
        /// </summary>
        public readonly string IpCidrRange;
        /// <summary>
        /// The subnetwork secondary range name specifying
        /// the secondary range from which to allocate the IP CIDR range for this alias IP
        /// range.
        /// </summary>
        public readonly string SubnetworkRangeName;

        [OutputConstructor]
        private GetInstanceNetworkInterfacesAliasIpRangesResult(
            string ipCidrRange,
            string subnetworkRangeName)
        {
            IpCidrRange = ipCidrRange;
            SubnetworkRangeName = subnetworkRangeName;
        }
    }

    [OutputType]
    public sealed class GetInstanceNetworkInterfacesResult
    {
        /// <summary>
        /// Access configurations, i.e. IPs via which this
        /// instance can be accessed via the Internet. Structure documented below.
        /// </summary>
        public readonly ImmutableArray<GetInstanceNetworkInterfacesAccessConfigsResult> AccessConfigs;
        /// <summary>
        /// An array of alias IP ranges for this network interface. Structure documented below.
        /// </summary>
        public readonly ImmutableArray<GetInstanceNetworkInterfacesAliasIpRangesResult> AliasIpRanges;
        /// <summary>
        /// The name of the instance. One of `name` or `self_link` must be provided.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name or self_link of the network attached to this interface.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// The private IP address assigned to the instance.
        /// </summary>
        public readonly string NetworkIp;
        /// <summary>
        /// The name or self_link of the subnetwork attached to this interface.
        /// </summary>
        public readonly string Subnetwork;
        /// <summary>
        /// The project in which the subnetwork belongs.
        /// </summary>
        public readonly string SubnetworkProject;

        [OutputConstructor]
        private GetInstanceNetworkInterfacesResult(
            ImmutableArray<GetInstanceNetworkInterfacesAccessConfigsResult> accessConfigs,
            ImmutableArray<GetInstanceNetworkInterfacesAliasIpRangesResult> aliasIpRanges,
            string name,
            string network,
            string networkIp,
            string subnetwork,
            string subnetworkProject)
        {
            AccessConfigs = accessConfigs;
            AliasIpRanges = aliasIpRanges;
            Name = name;
            Network = network;
            NetworkIp = networkIp;
            Subnetwork = subnetwork;
            SubnetworkProject = subnetworkProject;
        }
    }

    [OutputType]
    public sealed class GetInstanceSchedulingsNodeAffinitiesResult
    {
        public readonly string Key;
        public readonly string Operator;
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetInstanceSchedulingsNodeAffinitiesResult(
            string key,
            string @operator,
            ImmutableArray<string> values)
        {
            Key = key;
            Operator = @operator;
            Values = values;
        }
    }

    [OutputType]
    public sealed class GetInstanceSchedulingsResult
    {
        /// <summary>
        /// Specifies if the instance should be
        /// restarted if it was terminated by Compute Engine (not a user).
        /// </summary>
        public readonly bool AutomaticRestart;
        public readonly ImmutableArray<GetInstanceSchedulingsNodeAffinitiesResult> NodeAffinities;
        /// <summary>
        /// Describes maintenance behavior for the
        /// instance. One of `MIGRATE` or `TERMINATE`, for more info, read
        /// [here](https://cloud.google.com/compute/docs/instances/setting-instance-scheduling-options)
        /// </summary>
        public readonly string OnHostMaintenance;
        /// <summary>
        /// Whether the instance is preemptible.
        /// </summary>
        public readonly bool Preemptible;

        [OutputConstructor]
        private GetInstanceSchedulingsResult(
            bool automaticRestart,
            ImmutableArray<GetInstanceSchedulingsNodeAffinitiesResult> nodeAffinities,
            string onHostMaintenance,
            bool preemptible)
        {
            AutomaticRestart = automaticRestart;
            NodeAffinities = nodeAffinities;
            OnHostMaintenance = onHostMaintenance;
            Preemptible = preemptible;
        }
    }

    [OutputType]
    public sealed class GetInstanceScratchDisksResult
    {
        /// <summary>
        /// The disk interface used for attaching this disk. One of `SCSI` or `NVME`.
        /// </summary>
        public readonly string Interface;

        [OutputConstructor]
        private GetInstanceScratchDisksResult(string @interface)
        {
            Interface = @interface;
        }
    }

    [OutputType]
    public sealed class GetInstanceServiceAccountsResult
    {
        /// <summary>
        /// The service account e-mail address.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// A list of service scopes.
        /// </summary>
        public readonly ImmutableArray<string> Scopes;

        [OutputConstructor]
        private GetInstanceServiceAccountsResult(
            string email,
            ImmutableArray<string> scopes)
        {
            Email = email;
            Scopes = scopes;
        }
    }

    [OutputType]
    public sealed class GetInstanceShieldedInstanceConfigsResult
    {
        public readonly bool EnableIntegrityMonitoring;
        public readonly bool EnableSecureBoot;
        public readonly bool EnableVtpm;

        [OutputConstructor]
        private GetInstanceShieldedInstanceConfigsResult(
            bool enableIntegrityMonitoring,
            bool enableSecureBoot,
            bool enableVtpm)
        {
            EnableIntegrityMonitoring = enableIntegrityMonitoring;
            EnableSecureBoot = enableSecureBoot;
            EnableVtpm = enableVtpm;
        }
    }
    }
}

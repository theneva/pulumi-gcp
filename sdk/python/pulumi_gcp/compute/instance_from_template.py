# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class InstanceFromTemplate(pulumi.CustomResource):
    """
    Manages a VM instance resource within GCE. For more information see
    [the official documentation](https://cloud.google.com/compute/docs/instances)
    and
    [API](https://cloud.google.com/compute/docs/reference/latest/instances).
    
    This resource is specifically to create a compute instance from a given
    `source_instance_template`. To create an instance without a template, use the
    `google_compute_instance` resource.
    
    """
    def __init__(__self__, __name__, __opts__=None, allow_stopping_for_update=None, attached_disks=None, boot_disk=None, can_ip_forward=None, deletion_protection=None, description=None, guest_accelerators=None, labels=None, machine_type=None, metadata=None, metadata_startup_script=None, min_cpu_platform=None, name=None, network_interfaces=None, project=None, scheduling=None, scratch_disks=None, service_account=None, source_instance_template=None, tags=None, zone=None):
        """Create a InstanceFromTemplate resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if allow_stopping_for_update and not isinstance(allow_stopping_for_update, bool):
            raise TypeError('Expected property allow_stopping_for_update to be a bool')
        __self__.allow_stopping_for_update = allow_stopping_for_update
        __props__['allowStoppingForUpdate'] = allow_stopping_for_update

        if attached_disks and not isinstance(attached_disks, list):
            raise TypeError('Expected property attached_disks to be a list')
        __self__.attached_disks = attached_disks
        __props__['attachedDisks'] = attached_disks

        if boot_disk and not isinstance(boot_disk, dict):
            raise TypeError('Expected property boot_disk to be a dict')
        __self__.boot_disk = boot_disk
        __props__['bootDisk'] = boot_disk

        if can_ip_forward and not isinstance(can_ip_forward, bool):
            raise TypeError('Expected property can_ip_forward to be a bool')
        __self__.can_ip_forward = can_ip_forward
        __props__['canIpForward'] = can_ip_forward

        if deletion_protection and not isinstance(deletion_protection, bool):
            raise TypeError('Expected property deletion_protection to be a bool')
        __self__.deletion_protection = deletion_protection
        __props__['deletionProtection'] = deletion_protection

        if description and not isinstance(description, basestring):
            raise TypeError('Expected property description to be a basestring')
        __self__.description = description
        __props__['description'] = description

        if guest_accelerators and not isinstance(guest_accelerators, list):
            raise TypeError('Expected property guest_accelerators to be a list')
        __self__.guest_accelerators = guest_accelerators
        __props__['guestAccelerators'] = guest_accelerators

        if labels and not isinstance(labels, dict):
            raise TypeError('Expected property labels to be a dict')
        __self__.labels = labels
        __props__['labels'] = labels

        if machine_type and not isinstance(machine_type, basestring):
            raise TypeError('Expected property machine_type to be a basestring')
        __self__.machine_type = machine_type
        __props__['machineType'] = machine_type

        if metadata and not isinstance(metadata, dict):
            raise TypeError('Expected property metadata to be a dict')
        __self__.metadata = metadata
        __props__['metadata'] = metadata

        if metadata_startup_script and not isinstance(metadata_startup_script, basestring):
            raise TypeError('Expected property metadata_startup_script to be a basestring')
        __self__.metadata_startup_script = metadata_startup_script
        __props__['metadataStartupScript'] = metadata_startup_script

        if min_cpu_platform and not isinstance(min_cpu_platform, basestring):
            raise TypeError('Expected property min_cpu_platform to be a basestring')
        __self__.min_cpu_platform = min_cpu_platform
        __props__['minCpuPlatform'] = min_cpu_platform

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        A unique name for the resource, required by GCE.
        Changing this forces a new resource to be created.
        """
        __props__['name'] = name

        if network_interfaces and not isinstance(network_interfaces, list):
            raise TypeError('Expected property network_interfaces to be a list')
        __self__.network_interfaces = network_interfaces
        __props__['networkInterfaces'] = network_interfaces

        if project and not isinstance(project, basestring):
            raise TypeError('Expected property project to be a basestring')
        __self__.project = project
        __props__['project'] = project

        if scheduling and not isinstance(scheduling, dict):
            raise TypeError('Expected property scheduling to be a dict')
        __self__.scheduling = scheduling
        __props__['scheduling'] = scheduling

        if scratch_disks and not isinstance(scratch_disks, list):
            raise TypeError('Expected property scratch_disks to be a list')
        __self__.scratch_disks = scratch_disks
        __props__['scratchDisks'] = scratch_disks

        if service_account and not isinstance(service_account, dict):
            raise TypeError('Expected property service_account to be a dict')
        __self__.service_account = service_account
        __props__['serviceAccount'] = service_account

        if not source_instance_template:
            raise TypeError('Missing required property source_instance_template')
        elif not isinstance(source_instance_template, basestring):
            raise TypeError('Expected property source_instance_template to be a basestring')
        __self__.source_instance_template = source_instance_template
        """
        Name or self link of an instance
        template to create the instance based on.
        """
        __props__['sourceInstanceTemplate'] = source_instance_template

        if tags and not isinstance(tags, list):
            raise TypeError('Expected property tags to be a list')
        __self__.tags = tags
        __props__['tags'] = tags

        if zone and not isinstance(zone, basestring):
            raise TypeError('Expected property zone to be a basestring')
        __self__.zone = zone
        """
        The zone that the machine should be created in. If not
        set, the provider zone is used.
        """
        __props__['zone'] = zone

        __self__.cpu_platform = pulumi.runtime.UNKNOWN
        __self__.instance_id = pulumi.runtime.UNKNOWN
        __self__.label_fingerprint = pulumi.runtime.UNKNOWN
        __self__.metadata_fingerprint = pulumi.runtime.UNKNOWN
        __self__.self_link = pulumi.runtime.UNKNOWN
        __self__.tags_fingerprint = pulumi.runtime.UNKNOWN

        super(InstanceFromTemplate, __self__).__init__(
            'gcp:compute/instanceFromTemplate:InstanceFromTemplate',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'allowStoppingForUpdate' in outs:
            self.allow_stopping_for_update = outs['allowStoppingForUpdate']
        if 'attachedDisks' in outs:
            self.attached_disks = outs['attachedDisks']
        if 'bootDisk' in outs:
            self.boot_disk = outs['bootDisk']
        if 'canIpForward' in outs:
            self.can_ip_forward = outs['canIpForward']
        if 'cpuPlatform' in outs:
            self.cpu_platform = outs['cpuPlatform']
        if 'deletionProtection' in outs:
            self.deletion_protection = outs['deletionProtection']
        if 'description' in outs:
            self.description = outs['description']
        if 'guestAccelerators' in outs:
            self.guest_accelerators = outs['guestAccelerators']
        if 'instanceId' in outs:
            self.instance_id = outs['instanceId']
        if 'labelFingerprint' in outs:
            self.label_fingerprint = outs['labelFingerprint']
        if 'labels' in outs:
            self.labels = outs['labels']
        if 'machineType' in outs:
            self.machine_type = outs['machineType']
        if 'metadata' in outs:
            self.metadata = outs['metadata']
        if 'metadataFingerprint' in outs:
            self.metadata_fingerprint = outs['metadataFingerprint']
        if 'metadataStartupScript' in outs:
            self.metadata_startup_script = outs['metadataStartupScript']
        if 'minCpuPlatform' in outs:
            self.min_cpu_platform = outs['minCpuPlatform']
        if 'name' in outs:
            self.name = outs['name']
        if 'networkInterfaces' in outs:
            self.network_interfaces = outs['networkInterfaces']
        if 'project' in outs:
            self.project = outs['project']
        if 'scheduling' in outs:
            self.scheduling = outs['scheduling']
        if 'scratchDisks' in outs:
            self.scratch_disks = outs['scratchDisks']
        if 'selfLink' in outs:
            self.self_link = outs['selfLink']
        if 'serviceAccount' in outs:
            self.service_account = outs['serviceAccount']
        if 'sourceInstanceTemplate' in outs:
            self.source_instance_template = outs['sourceInstanceTemplate']
        if 'tags' in outs:
            self.tags = outs['tags']
        if 'tagsFingerprint' in outs:
            self.tags_fingerprint = outs['tagsFingerprint']
        if 'zone' in outs:
            self.zone = outs['zone']

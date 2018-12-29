# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class RegionDisk(pulumi.CustomResource):
    """
    Persistent disks are durable storage devices that function similarly to
    the physical disks in a desktop or a server. Compute Engine manages the
    hardware behind these devices to ensure data redundancy and optimize
    performance for you. Persistent disks are available as either standard
    hard disk drives (HDD) or solid-state drives (SSD).
    
    Persistent disks are located independently from your virtual machine
    instances, so you can detach or move persistent disks to keep your data
    even after you delete your instances. Persistent disk performance scales
    automatically with size, so you can resize your existing persistent disks
    or add more persistent disks to an instance to meet your performance and
    storage space requirements.
    
    Add a persistent disk to your instance when you need reliable and
    affordable storage with consistent performance characteristics.
    
    
    To get more information about RegionDisk, see:
    
    * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/regionDisks)
    * How-to Guides
        * [Adding or Resizing Regional Persistent Disks](https://cloud.google.com/compute/docs/disks/regional-persistent-disk)
    
    > **Warning:** All arguments including the disk encryption key will be stored in the raw
    state as plain-text.
    [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
    """
    def __init__(__self__, __name__, __opts__=None, description=None, disk_encryption_key=None, labels=None, name=None, project=None, region=None, replica_zones=None, size=None, snapshot=None, source_snapshot_encryption_key=None, type=None):
        """Create a RegionDisk resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description

        __props__['disk_encryption_key'] = disk_encryption_key

        __props__['labels'] = labels

        __props__['name'] = name

        __props__['project'] = project

        __props__['region'] = region

        if not replica_zones:
            raise TypeError('Missing required property replica_zones')
        __props__['replica_zones'] = replica_zones

        __props__['size'] = size

        __props__['snapshot'] = snapshot

        __props__['source_snapshot_encryption_key'] = source_snapshot_encryption_key

        __props__['type'] = type

        __props__['creation_timestamp'] = None
        __props__['label_fingerprint'] = None
        __props__['last_attach_timestamp'] = None
        __props__['last_detach_timestamp'] = None
        __props__['self_link'] = None
        __props__['source_snapshot_id'] = None
        __props__['users'] = None

        super(RegionDisk, __self__).__init__(
            'gcp:compute/regionDisk:RegionDisk',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


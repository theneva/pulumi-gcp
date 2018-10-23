# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class InterconnectAttachment(pulumi.CustomResource):
    def __init__(__self__, __name__, __opts__=None, description=None, interconnect=None, name=None, project=None, region=None, router=None):
        """Create a InterconnectAttachment resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if description and not isinstance(description, basestring):
            raise TypeError('Expected property description to be a basestring')
        __self__.description = description
        __props__['description'] = description

        if not interconnect:
            raise TypeError('Missing required property interconnect')
        elif not isinstance(interconnect, basestring):
            raise TypeError('Expected property interconnect to be a basestring')
        __self__.interconnect = interconnect
        __props__['interconnect'] = interconnect

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        __props__['name'] = name

        if project and not isinstance(project, basestring):
            raise TypeError('Expected property project to be a basestring')
        __self__.project = project
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        __props__['project'] = project

        if region and not isinstance(region, basestring):
            raise TypeError('Expected property region to be a basestring')
        __self__.region = region
        __props__['region'] = region

        if not router:
            raise TypeError('Missing required property router')
        elif not isinstance(router, basestring):
            raise TypeError('Expected property router to be a basestring')
        __self__.router = router
        __props__['router'] = router

        __self__.cloud_router_ip_address = pulumi.runtime.UNKNOWN
        __self__.creation_timestamp = pulumi.runtime.UNKNOWN
        __self__.customer_router_ip_address = pulumi.runtime.UNKNOWN
        __self__.google_reference_id = pulumi.runtime.UNKNOWN
        __self__.private_interconnect_info = pulumi.runtime.UNKNOWN
        __self__.self_link = pulumi.runtime.UNKNOWN
        """
        The URI of the created resource.
        """

        super(InterconnectAttachment, __self__).__init__(
            'gcp:compute/interconnectAttachment:InterconnectAttachment',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'cloudRouterIpAddress' in outs:
            self.cloud_router_ip_address = outs['cloudRouterIpAddress']
        if 'creationTimestamp' in outs:
            self.creation_timestamp = outs['creationTimestamp']
        if 'customerRouterIpAddress' in outs:
            self.customer_router_ip_address = outs['customerRouterIpAddress']
        if 'description' in outs:
            self.description = outs['description']
        if 'googleReferenceId' in outs:
            self.google_reference_id = outs['googleReferenceId']
        if 'interconnect' in outs:
            self.interconnect = outs['interconnect']
        if 'name' in outs:
            self.name = outs['name']
        if 'privateInterconnectInfo' in outs:
            self.private_interconnect_info = outs['privateInterconnectInfo']
        if 'project' in outs:
            self.project = outs['project']
        if 'region' in outs:
            self.region = outs['region']
        if 'router' in outs:
            self.router = outs['router']
        if 'selfLink' in outs:
            self.self_link = outs['selfLink']

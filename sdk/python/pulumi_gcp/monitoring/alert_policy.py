# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class AlertPolicy(pulumi.CustomResource):
    def __init__(__self__, __name__, __opts__=None, combiner=None, conditions=None, display_name=None, enabled=None, labels=None, notification_channels=None, project=None):
        """Create a AlertPolicy resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not combiner:
            raise TypeError('Missing required property combiner')
        elif not isinstance(combiner, basestring):
            raise TypeError('Expected property combiner to be a basestring')
        __self__.combiner = combiner
        __props__['combiner'] = combiner

        if not conditions:
            raise TypeError('Missing required property conditions')
        elif not isinstance(conditions, list):
            raise TypeError('Expected property conditions to be a list')
        __self__.conditions = conditions
        __props__['conditions'] = conditions

        if not display_name:
            raise TypeError('Missing required property display_name')
        elif not isinstance(display_name, basestring):
            raise TypeError('Expected property display_name to be a basestring')
        __self__.display_name = display_name
        __props__['displayName'] = display_name

        if not enabled:
            raise TypeError('Missing required property enabled')
        elif not isinstance(enabled, bool):
            raise TypeError('Expected property enabled to be a bool')
        __self__.enabled = enabled
        __props__['enabled'] = enabled

        if labels and not isinstance(labels, list):
            raise TypeError('Expected property labels to be a list')
        __self__.labels = labels
        __props__['labels'] = labels

        if notification_channels and not isinstance(notification_channels, list):
            raise TypeError('Expected property notification_channels to be a list')
        __self__.notification_channels = notification_channels
        __props__['notificationChannels'] = notification_channels

        if project and not isinstance(project, basestring):
            raise TypeError('Expected property project to be a basestring')
        __self__.project = project
        __props__['project'] = project

        __self__.creation_record = pulumi.runtime.UNKNOWN
        __self__.name = pulumi.runtime.UNKNOWN

        super(AlertPolicy, __self__).__init__(
            'gcp:monitoring/alertPolicy:AlertPolicy',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'combiner' in outs:
            self.combiner = outs['combiner']
        if 'conditions' in outs:
            self.conditions = outs['conditions']
        if 'creationRecord' in outs:
            self.creation_record = outs['creationRecord']
        if 'displayName' in outs:
            self.display_name = outs['displayName']
        if 'enabled' in outs:
            self.enabled = outs['enabled']
        if 'labels' in outs:
            self.labels = outs['labels']
        if 'name' in outs:
            self.name = outs['name']
        if 'notificationChannels' in outs:
            self.notification_channels = outs['notificationChannels']
        if 'project' in outs:
            self.project = outs['project']

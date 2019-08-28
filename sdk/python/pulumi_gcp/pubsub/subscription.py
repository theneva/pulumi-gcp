# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Subscription(pulumi.CustomResource):
    ack_deadline_seconds: pulumi.Output[float]
    expiration_policy: pulumi.Output[dict]
    labels: pulumi.Output[dict]
    message_retention_duration: pulumi.Output[str]
    name: pulumi.Output[str]
    path: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    push_config: pulumi.Output[dict]
    retain_acked_messages: pulumi.Output[bool]
    topic: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, ack_deadline_seconds=None, expiration_policy=None, labels=None, message_retention_duration=None, name=None, project=None, push_config=None, retain_acked_messages=None, topic=None, __props__=None, __name__=None, __opts__=None):
        """
        A named resource representing the stream of messages from a single,
        specific topic, to be delivered to the subscribing application.
        
        
        To get more information about Subscription, see:
        
        * [API documentation](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.subscriptions)
        * How-to Guides
            * [Managing Subscriptions](https://cloud.google.com/pubsub/docs/admin#managing_subscriptions)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        
        The **expiration_policy** object supports the following:
        
          * `ttl` (`pulumi.Input[str]`)
        
        The **push_config** object supports the following:
        
          * `attributes` (`pulumi.Input[dict]`)
          * `pushEndpoint` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/pubsub_subscription.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['ack_deadline_seconds'] = ack_deadline_seconds
            __props__['expiration_policy'] = expiration_policy
            __props__['labels'] = labels
            __props__['message_retention_duration'] = message_retention_duration
            __props__['name'] = name
            __props__['project'] = project
            __props__['push_config'] = push_config
            __props__['retain_acked_messages'] = retain_acked_messages
            if topic is None:
                raise TypeError("Missing required property 'topic'")
            __props__['topic'] = topic
            __props__['path'] = None
        super(Subscription, __self__).__init__(
            'gcp:pubsub/subscription:Subscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, ack_deadline_seconds=None, expiration_policy=None, labels=None, message_retention_duration=None, name=None, path=None, project=None, push_config=None, retain_acked_messages=None, topic=None):
        """
        Get an existing Subscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        
        The **expiration_policy** object supports the following:
        
          * `ttl` (`pulumi.Input[str]`)
        
        The **push_config** object supports the following:
        
          * `attributes` (`pulumi.Input[dict]`)
          * `pushEndpoint` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/pubsub_subscription.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["ack_deadline_seconds"] = ack_deadline_seconds
        __props__["expiration_policy"] = expiration_policy
        __props__["labels"] = labels
        __props__["message_retention_duration"] = message_retention_duration
        __props__["name"] = name
        __props__["path"] = path
        __props__["project"] = project
        __props__["push_config"] = push_config
        __props__["retain_acked_messages"] = retain_acked_messages
        __props__["topic"] = topic
        return Subscription(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


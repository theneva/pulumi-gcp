# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class DataTransferConfig(pulumi.CustomResource):
    data_refresh_window_days: pulumi.Output[float]
    data_source_id: pulumi.Output[str]
    destination_dataset_id: pulumi.Output[str]
    disabled: pulumi.Output[bool]
    display_name: pulumi.Output[str]
    location: pulumi.Output[str]
    name: pulumi.Output[str]
    params: pulumi.Output[dict]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    schedule: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, data_refresh_window_days=None, data_source_id=None, destination_dataset_id=None, disabled=None, display_name=None, location=None, params=None, project=None, schedule=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a DataTransferConfig resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/bigquery_data_transfer_config.html.markdown.
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

            __props__['data_refresh_window_days'] = data_refresh_window_days
            if data_source_id is None:
                raise TypeError("Missing required property 'data_source_id'")
            __props__['data_source_id'] = data_source_id
            if destination_dataset_id is None:
                raise TypeError("Missing required property 'destination_dataset_id'")
            __props__['destination_dataset_id'] = destination_dataset_id
            __props__['disabled'] = disabled
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['location'] = location
            if params is None:
                raise TypeError("Missing required property 'params'")
            __props__['params'] = params
            __props__['project'] = project
            __props__['schedule'] = schedule
            __props__['name'] = None
        super(DataTransferConfig, __self__).__init__(
            'gcp:bigquery/dataTransferConfig:DataTransferConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, data_refresh_window_days=None, data_source_id=None, destination_dataset_id=None, disabled=None, display_name=None, location=None, name=None, params=None, project=None, schedule=None):
        """
        Get an existing DataTransferConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/bigquery_data_transfer_config.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["data_refresh_window_days"] = data_refresh_window_days
        __props__["data_source_id"] = data_source_id
        __props__["destination_dataset_id"] = destination_dataset_id
        __props__["disabled"] = disabled
        __props__["display_name"] = display_name
        __props__["location"] = location
        __props__["name"] = name
        __props__["params"] = params
        __props__["project"] = project
        __props__["schedule"] = schedule
        return DataTransferConfig(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Trigger(pulumi.CustomResource):
    build: pulumi.Output[dict]
    create_time: pulumi.Output[str]
    description: pulumi.Output[str]
    disabled: pulumi.Output[bool]
    filename: pulumi.Output[str]
    github: pulumi.Output[dict]
    ignored_files: pulumi.Output[list]
    included_files: pulumi.Output[list]
    name: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    substitutions: pulumi.Output[dict]
    trigger_id: pulumi.Output[str]
    trigger_template: pulumi.Output[dict]
    def __init__(__self__, resource_name, opts=None, build=None, description=None, disabled=None, filename=None, github=None, ignored_files=None, included_files=None, name=None, project=None, substitutions=None, trigger_template=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Trigger resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        
        The **build** object supports the following:
        
          * `images` (`pulumi.Input[list]`)
          * `steps` (`pulumi.Input[list]`)
        
            * `args` (`pulumi.Input[list]`)
            * `dir` (`pulumi.Input[str]`)
            * `entrypoint` (`pulumi.Input[str]`)
            * `envs` (`pulumi.Input[list]`)
            * `id` (`pulumi.Input[str]`)
            * `name` (`pulumi.Input[str]`)
            * `secretEnvs` (`pulumi.Input[list]`)
            * `timeout` (`pulumi.Input[str]`)
            * `timing` (`pulumi.Input[str]`)
            * `volumes` (`pulumi.Input[list]`)
        
              * `name` (`pulumi.Input[str]`)
              * `path` (`pulumi.Input[str]`)
        
            * `waitFors` (`pulumi.Input[list]`)
        
          * `tags` (`pulumi.Input[list]`)
        
        The **github** object supports the following:
        
          * `name` (`pulumi.Input[str]`)
          * `owner` (`pulumi.Input[str]`)
          * `pullRequest` (`pulumi.Input[dict]`)
        
            * `branch` (`pulumi.Input[str]`)
            * `commentControl` (`pulumi.Input[str]`)
        
          * `push` (`pulumi.Input[dict]`)
        
            * `branch` (`pulumi.Input[str]`)
            * `tag` (`pulumi.Input[str]`)
        
        The **trigger_template** object supports the following:
        
          * `branchName` (`pulumi.Input[str]`)
          * `commitSha` (`pulumi.Input[str]`)
          * `dir` (`pulumi.Input[str]`)
          * `projectId` (`pulumi.Input[str]`)
          * `repoName` (`pulumi.Input[str]`)
          * `tagName` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloudbuild_trigger.html.markdown.
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

            __props__['build'] = build
            __props__['description'] = description
            __props__['disabled'] = disabled
            __props__['filename'] = filename
            __props__['github'] = github
            __props__['ignored_files'] = ignored_files
            __props__['included_files'] = included_files
            __props__['name'] = name
            __props__['project'] = project
            __props__['substitutions'] = substitutions
            __props__['trigger_template'] = trigger_template
            __props__['create_time'] = None
            __props__['trigger_id'] = None
        super(Trigger, __self__).__init__(
            'gcp:cloudbuild/trigger:Trigger',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, build=None, create_time=None, description=None, disabled=None, filename=None, github=None, ignored_files=None, included_files=None, name=None, project=None, substitutions=None, trigger_id=None, trigger_template=None):
        """
        Get an existing Trigger resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        
        The **build** object supports the following:
        
          * `images` (`pulumi.Input[list]`)
          * `steps` (`pulumi.Input[list]`)
        
            * `args` (`pulumi.Input[list]`)
            * `dir` (`pulumi.Input[str]`)
            * `entrypoint` (`pulumi.Input[str]`)
            * `envs` (`pulumi.Input[list]`)
            * `id` (`pulumi.Input[str]`)
            * `name` (`pulumi.Input[str]`)
            * `secretEnvs` (`pulumi.Input[list]`)
            * `timeout` (`pulumi.Input[str]`)
            * `timing` (`pulumi.Input[str]`)
            * `volumes` (`pulumi.Input[list]`)
        
              * `name` (`pulumi.Input[str]`)
              * `path` (`pulumi.Input[str]`)
        
            * `waitFors` (`pulumi.Input[list]`)
        
          * `tags` (`pulumi.Input[list]`)
        
        The **github** object supports the following:
        
          * `name` (`pulumi.Input[str]`)
          * `owner` (`pulumi.Input[str]`)
          * `pullRequest` (`pulumi.Input[dict]`)
        
            * `branch` (`pulumi.Input[str]`)
            * `commentControl` (`pulumi.Input[str]`)
        
          * `push` (`pulumi.Input[dict]`)
        
            * `branch` (`pulumi.Input[str]`)
            * `tag` (`pulumi.Input[str]`)
        
        The **trigger_template** object supports the following:
        
          * `branchName` (`pulumi.Input[str]`)
          * `commitSha` (`pulumi.Input[str]`)
          * `dir` (`pulumi.Input[str]`)
          * `projectId` (`pulumi.Input[str]`)
          * `repoName` (`pulumi.Input[str]`)
          * `tagName` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloudbuild_trigger.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["build"] = build
        __props__["create_time"] = create_time
        __props__["description"] = description
        __props__["disabled"] = disabled
        __props__["filename"] = filename
        __props__["github"] = github
        __props__["ignored_files"] = ignored_files
        __props__["included_files"] = included_files
        __props__["name"] = name
        __props__["project"] = project
        __props__["substitutions"] = substitutions
        __props__["trigger_id"] = trigger_id
        __props__["trigger_template"] = trigger_template
        return Trigger(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


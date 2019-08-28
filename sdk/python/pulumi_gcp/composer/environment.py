# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Environment(pulumi.CustomResource):
    config: pulumi.Output[dict]
    labels: pulumi.Output[dict]
    name: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, config=None, labels=None, name=None, project=None, region=None, __props__=None, __name__=None, __opts__=None):
        """
        An environment for running orchestration tasks.
        
        Environments run Apache Airflow software on Google infrastructure.
        
        To get more information about Environments, see:
        
        * [API documentation](https://cloud.google.com/composer/docs/reference/rest/)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/composer/docs)
            * [Configuring Shared VPC for Composer Environments](https://cloud.google.com/composer/docs/how-to/managing/configuring-shared-vpc)
        * [Apache Airflow Documentation](http://airflow.apache.org/)
        
        > **Warning:** We **STRONGLY** recommend  you read the [GCP guides](https://cloud.google.com/composer/docs/how-to)
          as the Environment resource requires a long deployment process and involves several layers of GCP infrastructure, 
          including a Kubernetes Engine cluster, Cloud Storage, and Compute networking resources. Due to limitations of the API,
          this provider will not be able to automatically find or manage many of these underlying resources. In particular:
          * It can take up to one hour to create or update an environment resource. In addition, GCP may only detect some 
            errors in configuration when they are used (e.g. ~40-50 minutes into the creation process), and is prone to limited
            error reporting. If you encounter confusing or uninformative errors, please verify your configuration is valid 
            against GCP Cloud Composer before filing bugs against this provider. 
          * **Environments create Google Cloud Storage buckets that do not get cleaned up automatically** on environment 
            deletion. [More about Composer's use of Cloud Storage](https://cloud.google.com/composer/docs/concepts/cloud-storage).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        
        The **config** object supports the following:
        
          * `airflowUri` (`pulumi.Input[str]`)
          * `dagGcsPrefix` (`pulumi.Input[str]`)
          * `gkeCluster` (`pulumi.Input[str]`)
          * `nodeConfig` (`pulumi.Input[dict]`)
        
            * `diskSizeGb` (`pulumi.Input[float]`)
            * `ipAllocationPolicy` (`pulumi.Input[dict]`)
        
              * `clusterIpv4CidrBlock` (`pulumi.Input[str]`)
              * `clusterSecondaryRangeName` (`pulumi.Input[str]`)
              * `servicesIpv4CidrBlock` (`pulumi.Input[str]`)
              * `servicesSecondaryRangeName` (`pulumi.Input[str]`)
              * `useIpAliases` (`pulumi.Input[bool]`)
        
            * `machineType` (`pulumi.Input[str]`)
            * `network` (`pulumi.Input[str]`)
            * `oauthScopes` (`pulumi.Input[list]`)
            * `serviceAccount` (`pulumi.Input[str]`)
            * `subnetwork` (`pulumi.Input[str]`)
            * `tags` (`pulumi.Input[list]`)
            * `zone` (`pulumi.Input[str]`)
        
          * `nodeCount` (`pulumi.Input[float]`)
          * `privateEnvironmentConfig` (`pulumi.Input[dict]`)
        
            * `enablePrivateEndpoint` (`pulumi.Input[bool]`)
            * `masterIpv4CidrBlock` (`pulumi.Input[str]`)
        
          * `softwareConfig` (`pulumi.Input[dict]`)
        
            * `airflowConfigOverrides` (`pulumi.Input[dict]`)
            * `envVariables` (`pulumi.Input[dict]`)
            * `imageVersion` (`pulumi.Input[str]`)
            * `pypiPackages` (`pulumi.Input[dict]`)
            * `pythonVersion` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/composer_environment.html.markdown.
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

            __props__['config'] = config
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['project'] = project
            __props__['region'] = region
        super(Environment, __self__).__init__(
            'gcp:composer/environment:Environment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, config=None, labels=None, name=None, project=None, region=None):
        """
        Get an existing Environment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        
        The **config** object supports the following:
        
          * `airflowUri` (`pulumi.Input[str]`)
          * `dagGcsPrefix` (`pulumi.Input[str]`)
          * `gkeCluster` (`pulumi.Input[str]`)
          * `nodeConfig` (`pulumi.Input[dict]`)
        
            * `diskSizeGb` (`pulumi.Input[float]`)
            * `ipAllocationPolicy` (`pulumi.Input[dict]`)
        
              * `clusterIpv4CidrBlock` (`pulumi.Input[str]`)
              * `clusterSecondaryRangeName` (`pulumi.Input[str]`)
              * `servicesIpv4CidrBlock` (`pulumi.Input[str]`)
              * `servicesSecondaryRangeName` (`pulumi.Input[str]`)
              * `useIpAliases` (`pulumi.Input[bool]`)
        
            * `machineType` (`pulumi.Input[str]`)
            * `network` (`pulumi.Input[str]`)
            * `oauthScopes` (`pulumi.Input[list]`)
            * `serviceAccount` (`pulumi.Input[str]`)
            * `subnetwork` (`pulumi.Input[str]`)
            * `tags` (`pulumi.Input[list]`)
            * `zone` (`pulumi.Input[str]`)
        
          * `nodeCount` (`pulumi.Input[float]`)
          * `privateEnvironmentConfig` (`pulumi.Input[dict]`)
        
            * `enablePrivateEndpoint` (`pulumi.Input[bool]`)
            * `masterIpv4CidrBlock` (`pulumi.Input[str]`)
        
          * `softwareConfig` (`pulumi.Input[dict]`)
        
            * `airflowConfigOverrides` (`pulumi.Input[dict]`)
            * `envVariables` (`pulumi.Input[dict]`)
            * `imageVersion` (`pulumi.Input[str]`)
            * `pypiPackages` (`pulumi.Input[dict]`)
            * `pythonVersion` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/composer_environment.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["config"] = config
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["project"] = project
        __props__["region"] = region
        return Environment(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


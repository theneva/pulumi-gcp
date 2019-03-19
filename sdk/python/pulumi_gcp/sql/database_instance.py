# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class DatabaseInstance(pulumi.CustomResource):
    connection_name: pulumi.Output[str]
    """
    The connection name of the instance to be used in
    connection strings. For example, when connecting with [Cloud SQL Proxy](https://cloud.google.com/sql/docs/mysql/connect-admin-proxy).
    """
    database_version: pulumi.Output[str]
    """
    The MySQL version to
    use. Can be `MYSQL_5_6`, `MYSQL_5_7` or `POSTGRES_9_6` for second-generation
    instances, or `MYSQL_5_5` or `MYSQL_5_6` for first-generation instances.
    See [Second Generation Capabilities](https://cloud.google.com/sql/docs/1st-2nd-gen-differences)
    for more information.
    """
    first_ip_address: pulumi.Output[str]
    """
    The first IPv4 address of any type assigned. This is to
    support accessing the [first address in the list in a terraform output](https://github.com/terraform-providers/terraform-provider-google/issues/912)
    when the resource is configured with a `count`.
    """
    ip_addresses: pulumi.Output[list]
    master_instance_name: pulumi.Output[str]
    """
    The name of the instance that will act as
    the master in the replication setup. Note, this requires the master to have
    `binary_log_enabled` set, as well as existing backups.
    """
    name: pulumi.Output[str]
    """
    The name of the instance. If the name is left
    blank, Terraform will randomly generate one when the instance is first
    created. This is done because after a name is used, it cannot be reused for
    up to [one week](https://cloud.google.com/sql/docs/delete-instance).
    """
    private_ip_address: pulumi.Output[str]
    """
    The first private (`PRIVATE`) IPv4 address assigned. This is
    a workaround for an [issue fixed in Terraform 0.12](https://github.com/hashicorp/terraform/issues/17048)
    but also provides a convenient way to access an IP of a specific type without
    performing filtering in a Terraform config.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    public_ip_address: pulumi.Output[str]
    """
    The first public (`PRIMARY`) IPv4 address assigned. This is
    a workaround for an [issue fixed in Terraform 0.12](https://github.com/hashicorp/terraform/issues/17048)
    but also provides a convenient way to access an IP of a specific type without
    performing filtering in a Terraform config.
    """
    region: pulumi.Output[str]
    """
    The region the instance will sit in. Note, first-generation Cloud SQL instance
    regions do not line up with the Google Compute Engine (GCE) regions, and Cloud SQL is not
    available in all regions - choose from one of the options listed [here](https://cloud.google.com/sql/docs/mysql/instance-locations).
    A valid region must be provided to use this resource. If a region is not provided in the resource definition,
    the provider region will be used instead, but this will be an apply-time error for all first-generation
    instances *and* for second-generation instances if the provider region is not supported with Cloud SQL.
    If you choose not to provide the `region` argument for this resource, make sure you understand this.
    """
    replica_configuration: pulumi.Output[dict]
    """
    The configuration for replication. The
    configuration is detailed below.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    server_ca_cert: pulumi.Output[dict]
    service_account_email_address: pulumi.Output[str]
    """
    The service account email address assigned to the
    instance. This property is applicable only to Second Generation instances.
    """
    settings: pulumi.Output[dict]
    """
    The settings to use for the database. The
    configuration is detailed below.
    """
    def __init__(__self__, resource_name, opts=None, database_version=None, master_instance_name=None, name=None, project=None, region=None, replica_configuration=None, settings=None, __name__=None, __opts__=None):
        """
        Creates a new Google SQL Database Instance. For more information, see the [official documentation](https://cloud.google.com/sql/),
        or the [JSON API](https://cloud.google.com/sql/docs/admin-api/v1beta4/instances).
        
        > **NOTE on `google_sql_database_instance`:** - Second-generation instances include a
        default 'root'@'%' user with no password. This user will be deleted by Terraform on
        instance creation. You should use `google_sql_user` to define a custom user with
        a restricted host and strong password.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_version: The MySQL version to
               use. Can be `MYSQL_5_6`, `MYSQL_5_7` or `POSTGRES_9_6` for second-generation
               instances, or `MYSQL_5_5` or `MYSQL_5_6` for first-generation instances.
               See [Second Generation Capabilities](https://cloud.google.com/sql/docs/1st-2nd-gen-differences)
               for more information.
        :param pulumi.Input[str] master_instance_name: The name of the instance that will act as
               the master in the replication setup. Note, this requires the master to have
               `binary_log_enabled` set, as well as existing backups.
        :param pulumi.Input[str] name: The name of the instance. If the name is left
               blank, Terraform will randomly generate one when the instance is first
               created. This is done because after a name is used, it cannot be reused for
               up to [one week](https://cloud.google.com/sql/docs/delete-instance).
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region the instance will sit in. Note, first-generation Cloud SQL instance
               regions do not line up with the Google Compute Engine (GCE) regions, and Cloud SQL is not
               available in all regions - choose from one of the options listed [here](https://cloud.google.com/sql/docs/mysql/instance-locations).
               A valid region must be provided to use this resource. If a region is not provided in the resource definition,
               the provider region will be used instead, but this will be an apply-time error for all first-generation
               instances *and* for second-generation instances if the provider region is not supported with Cloud SQL.
               If you choose not to provide the `region` argument for this resource, make sure you understand this.
        :param pulumi.Input[dict] replica_configuration: The configuration for replication. The
               configuration is detailed below.
        :param pulumi.Input[dict] settings: The settings to use for the database. The
               configuration is detailed below.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['database_version'] = database_version

        __props__['master_instance_name'] = master_instance_name

        __props__['name'] = name

        __props__['project'] = project

        __props__['region'] = region

        __props__['replica_configuration'] = replica_configuration

        if settings is None:
            raise TypeError('Missing required property settings')
        __props__['settings'] = settings

        __props__['connection_name'] = None
        __props__['first_ip_address'] = None
        __props__['ip_addresses'] = None
        __props__['private_ip_address'] = None
        __props__['public_ip_address'] = None
        __props__['self_link'] = None
        __props__['server_ca_cert'] = None
        __props__['service_account_email_address'] = None

        super(DatabaseInstance, __self__).__init__(
            'gcp:sql/databaseInstance:DatabaseInstance',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


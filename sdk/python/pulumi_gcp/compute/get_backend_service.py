# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetBackendServiceResult:
    """
    A collection of values returned by getBackendService.
    """
    def __init__(__self__, affinity_cookie_ttl_sec=None, backends=None, cdn_policies=None, circuit_breakers=None, connection_draining_timeout_sec=None, consistent_hash=None, creation_timestamp=None, custom_request_headers=None, description=None, enable_cdn=None, fingerprint=None, health_checks=None, iaps=None, load_balancing_scheme=None, locality_lb_policy=None, log_configs=None, name=None, outlier_detections=None, port_name=None, project=None, protocol=None, security_policy=None, self_link=None, session_affinity=None, timeout_sec=None, id=None):
        if affinity_cookie_ttl_sec and not isinstance(affinity_cookie_ttl_sec, float):
            raise TypeError("Expected argument 'affinity_cookie_ttl_sec' to be a float")
        __self__.affinity_cookie_ttl_sec = affinity_cookie_ttl_sec
        if backends and not isinstance(backends, list):
            raise TypeError("Expected argument 'backends' to be a list")
        __self__.backends = backends
        """
        The set of backends that serve this Backend Service.
        """
        if cdn_policies and not isinstance(cdn_policies, list):
            raise TypeError("Expected argument 'cdn_policies' to be a list")
        __self__.cdn_policies = cdn_policies
        if circuit_breakers and not isinstance(circuit_breakers, list):
            raise TypeError("Expected argument 'circuit_breakers' to be a list")
        __self__.circuit_breakers = circuit_breakers
        if connection_draining_timeout_sec and not isinstance(connection_draining_timeout_sec, float):
            raise TypeError("Expected argument 'connection_draining_timeout_sec' to be a float")
        __self__.connection_draining_timeout_sec = connection_draining_timeout_sec
        """
        Time for which instance will be drained (not accept new connections, but still work to finish started ones).
        """
        if consistent_hash and not isinstance(consistent_hash, list):
            raise TypeError("Expected argument 'consistent_hash' to be a list")
        __self__.consistent_hash = consistent_hash
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        __self__.creation_timestamp = creation_timestamp
        if custom_request_headers and not isinstance(custom_request_headers, list):
            raise TypeError("Expected argument 'custom_request_headers' to be a list")
        __self__.custom_request_headers = custom_request_headers
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        Textual description for the Backend Service.
        """
        if enable_cdn and not isinstance(enable_cdn, bool):
            raise TypeError("Expected argument 'enable_cdn' to be a bool")
        __self__.enable_cdn = enable_cdn
        """
        Whether or not Cloud CDN is enabled on the Backend Service.
        """
        if fingerprint and not isinstance(fingerprint, str):
            raise TypeError("Expected argument 'fingerprint' to be a str")
        __self__.fingerprint = fingerprint
        """
        The fingerprint of the Backend Service.
        """
        if health_checks and not isinstance(health_checks, list):
            raise TypeError("Expected argument 'health_checks' to be a list")
        __self__.health_checks = health_checks
        """
        The set of HTTP/HTTPS health checks used by the Backend Service.
        """
        if iaps and not isinstance(iaps, list):
            raise TypeError("Expected argument 'iaps' to be a list")
        __self__.iaps = iaps
        if load_balancing_scheme and not isinstance(load_balancing_scheme, str):
            raise TypeError("Expected argument 'load_balancing_scheme' to be a str")
        __self__.load_balancing_scheme = load_balancing_scheme
        if locality_lb_policy and not isinstance(locality_lb_policy, str):
            raise TypeError("Expected argument 'locality_lb_policy' to be a str")
        __self__.locality_lb_policy = locality_lb_policy
        if log_configs and not isinstance(log_configs, list):
            raise TypeError("Expected argument 'log_configs' to be a list")
        __self__.log_configs = log_configs
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if outlier_detections and not isinstance(outlier_detections, list):
            raise TypeError("Expected argument 'outlier_detections' to be a list")
        __self__.outlier_detections = outlier_detections
        if port_name and not isinstance(port_name, str):
            raise TypeError("Expected argument 'port_name' to be a str")
        __self__.port_name = port_name
        """
        The name of a service that has been added to an instance group in this backend.
        """
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        __self__.protocol = protocol
        """
        The protocol for incoming requests.
        """
        if security_policy and not isinstance(security_policy, str):
            raise TypeError("Expected argument 'security_policy' to be a str")
        __self__.security_policy = security_policy
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        __self__.self_link = self_link
        """
        The URI of the Backend Service.
        """
        if session_affinity and not isinstance(session_affinity, str):
            raise TypeError("Expected argument 'session_affinity' to be a str")
        __self__.session_affinity = session_affinity
        """
        The Backend Service session stickiness configuration.
        """
        if timeout_sec and not isinstance(timeout_sec, float):
            raise TypeError("Expected argument 'timeout_sec' to be a float")
        __self__.timeout_sec = timeout_sec
        """
        The number of seconds to wait for a backend to respond to a request before considering the request failed.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetBackendServiceResult(GetBackendServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackendServiceResult(
            affinity_cookie_ttl_sec=self.affinity_cookie_ttl_sec,
            backends=self.backends,
            cdn_policies=self.cdn_policies,
            circuit_breakers=self.circuit_breakers,
            connection_draining_timeout_sec=self.connection_draining_timeout_sec,
            consistent_hash=self.consistent_hash,
            creation_timestamp=self.creation_timestamp,
            custom_request_headers=self.custom_request_headers,
            description=self.description,
            enable_cdn=self.enable_cdn,
            fingerprint=self.fingerprint,
            health_checks=self.health_checks,
            iaps=self.iaps,
            load_balancing_scheme=self.load_balancing_scheme,
            locality_lb_policy=self.locality_lb_policy,
            log_configs=self.log_configs,
            name=self.name,
            outlier_detections=self.outlier_detections,
            port_name=self.port_name,
            project=self.project,
            protocol=self.protocol,
            security_policy=self.security_policy,
            self_link=self.self_link,
            session_affinity=self.session_affinity,
            timeout_sec=self.timeout_sec,
            id=self.id)

def get_backend_service(name=None,project=None,opts=None):
    """
    Provide access to a Backend Service's attribute. For more information
    see [the official documentation](https://cloud.google.com/compute/docs/load-balancing/http/backend-service)
    and the [API](https://cloud.google.com/compute/docs/reference/latest/backendServices).
    
    :param str name: The name of the Backend Service.
    :param str project: The project in which the resource belongs. If it is not provided, the provider project is used.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/compute_backend_service.html.markdown.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:compute/getBackendService:getBackendService', __args__, opts=opts).value

    return AwaitableGetBackendServiceResult(
        affinity_cookie_ttl_sec=__ret__.get('affinityCookieTtlSec'),
        backends=__ret__.get('backends'),
        cdn_policies=__ret__.get('cdnPolicies'),
        circuit_breakers=__ret__.get('circuitBreakers'),
        connection_draining_timeout_sec=__ret__.get('connectionDrainingTimeoutSec'),
        consistent_hash=__ret__.get('consistentHash'),
        creation_timestamp=__ret__.get('creationTimestamp'),
        custom_request_headers=__ret__.get('customRequestHeaders'),
        description=__ret__.get('description'),
        enable_cdn=__ret__.get('enableCdn'),
        fingerprint=__ret__.get('fingerprint'),
        health_checks=__ret__.get('healthChecks'),
        iaps=__ret__.get('iaps'),
        load_balancing_scheme=__ret__.get('loadBalancingScheme'),
        locality_lb_policy=__ret__.get('localityLbPolicy'),
        log_configs=__ret__.get('logConfigs'),
        name=__ret__.get('name'),
        outlier_detections=__ret__.get('outlierDetections'),
        port_name=__ret__.get('portName'),
        project=__ret__.get('project'),
        protocol=__ret__.get('protocol'),
        security_policy=__ret__.get('securityPolicy'),
        self_link=__ret__.get('selfLink'),
        session_affinity=__ret__.get('sessionAffinity'),
        timeout_sec=__ret__.get('timeoutSec'),
        id=__ret__.get('id'))

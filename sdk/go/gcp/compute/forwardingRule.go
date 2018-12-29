// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// A ForwardingRule resource. A ForwardingRule resource specifies which pool
// of target virtual machines to forward a packet to if it matches the given
// [IPAddress, IPProtocol, portRange] tuple.
// 
// 
// To get more information about ForwardingRule, see:
// 
// * [API documentation](https://cloud.google.com/compute/docs/reference/latest/forwardingRule)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/network/forwarding-rules)
type ForwardingRule struct {
	s *pulumi.ResourceState
}

// NewForwardingRule registers a new resource with the given unique name, arguments, and options.
func NewForwardingRule(ctx *pulumi.Context,
	name string, args *ForwardingRuleArgs, opts ...pulumi.ResourceOpt) (*ForwardingRule, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["backendService"] = nil
		inputs["description"] = nil
		inputs["ipAddress"] = nil
		inputs["ipProtocol"] = nil
		inputs["ipVersion"] = nil
		inputs["labels"] = nil
		inputs["loadBalancingScheme"] = nil
		inputs["name"] = nil
		inputs["network"] = nil
		inputs["networkTier"] = nil
		inputs["portRange"] = nil
		inputs["ports"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["serviceLabel"] = nil
		inputs["subnetwork"] = nil
		inputs["target"] = nil
	} else {
		inputs["backendService"] = args.BackendService
		inputs["description"] = args.Description
		inputs["ipAddress"] = args.IpAddress
		inputs["ipProtocol"] = args.IpProtocol
		inputs["ipVersion"] = args.IpVersion
		inputs["labels"] = args.Labels
		inputs["loadBalancingScheme"] = args.LoadBalancingScheme
		inputs["name"] = args.Name
		inputs["network"] = args.Network
		inputs["networkTier"] = args.NetworkTier
		inputs["portRange"] = args.PortRange
		inputs["ports"] = args.Ports
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["serviceLabel"] = args.ServiceLabel
		inputs["subnetwork"] = args.Subnetwork
		inputs["target"] = args.Target
	}
	inputs["creationTimestamp"] = nil
	inputs["labelFingerprint"] = nil
	inputs["selfLink"] = nil
	inputs["serviceName"] = nil
	s, err := ctx.RegisterResource("gcp:compute/forwardingRule:ForwardingRule", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ForwardingRule{s: s}, nil
}

// GetForwardingRule gets an existing ForwardingRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetForwardingRule(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ForwardingRuleState, opts ...pulumi.ResourceOpt) (*ForwardingRule, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["backendService"] = state.BackendService
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["ipAddress"] = state.IpAddress
		inputs["ipProtocol"] = state.IpProtocol
		inputs["ipVersion"] = state.IpVersion
		inputs["labelFingerprint"] = state.LabelFingerprint
		inputs["labels"] = state.Labels
		inputs["loadBalancingScheme"] = state.LoadBalancingScheme
		inputs["name"] = state.Name
		inputs["network"] = state.Network
		inputs["networkTier"] = state.NetworkTier
		inputs["portRange"] = state.PortRange
		inputs["ports"] = state.Ports
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["selfLink"] = state.SelfLink
		inputs["serviceLabel"] = state.ServiceLabel
		inputs["serviceName"] = state.ServiceName
		inputs["subnetwork"] = state.Subnetwork
		inputs["target"] = state.Target
	}
	s, err := ctx.ReadResource("gcp:compute/forwardingRule:ForwardingRule", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ForwardingRule{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ForwardingRule) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ForwardingRule) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *ForwardingRule) BackendService() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["backendService"])
}

func (r *ForwardingRule) CreationTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

func (r *ForwardingRule) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *ForwardingRule) IpAddress() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ipAddress"])
}

func (r *ForwardingRule) IpProtocol() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ipProtocol"])
}

func (r *ForwardingRule) IpVersion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ipVersion"])
}

func (r *ForwardingRule) LabelFingerprint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["labelFingerprint"])
}

func (r *ForwardingRule) Labels() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["labels"])
}

func (r *ForwardingRule) LoadBalancingScheme() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["loadBalancingScheme"])
}

func (r *ForwardingRule) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *ForwardingRule) Network() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["network"])
}

func (r *ForwardingRule) NetworkTier() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["networkTier"])
}

func (r *ForwardingRule) PortRange() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["portRange"])
}

func (r *ForwardingRule) Ports() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["ports"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *ForwardingRule) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *ForwardingRule) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The URI of the created resource.
func (r *ForwardingRule) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

func (r *ForwardingRule) ServiceLabel() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serviceLabel"])
}

func (r *ForwardingRule) ServiceName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serviceName"])
}

func (r *ForwardingRule) Subnetwork() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subnetwork"])
}

func (r *ForwardingRule) Target() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["target"])
}

// Input properties used for looking up and filtering ForwardingRule resources.
type ForwardingRuleState struct {
	BackendService interface{}
	CreationTimestamp interface{}
	Description interface{}
	IpAddress interface{}
	IpProtocol interface{}
	IpVersion interface{}
	LabelFingerprint interface{}
	Labels interface{}
	LoadBalancingScheme interface{}
	Name interface{}
	Network interface{}
	NetworkTier interface{}
	PortRange interface{}
	Ports interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Region interface{}
	// The URI of the created resource.
	SelfLink interface{}
	ServiceLabel interface{}
	ServiceName interface{}
	Subnetwork interface{}
	Target interface{}
}

// The set of arguments for constructing a ForwardingRule resource.
type ForwardingRuleArgs struct {
	BackendService interface{}
	Description interface{}
	IpAddress interface{}
	IpProtocol interface{}
	IpVersion interface{}
	Labels interface{}
	LoadBalancingScheme interface{}
	Name interface{}
	Network interface{}
	NetworkTier interface{}
	PortRange interface{}
	Ports interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Region interface{}
	ServiceLabel interface{}
	Subnetwork interface{}
	Target interface{}
}

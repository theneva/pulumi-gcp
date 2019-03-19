// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Represents an Autoscaler resource.
// 
// Autoscalers allow you to automatically scale virtual machine instances in
// managed instance groups according to an autoscaling policy that you
// define.
// 
// 
// To get more information about Autoscaler, see:
// 
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/autoscalers)
// * How-to Guides
//     * [Autoscaling Groups of Instances](https://cloud.google.com/compute/docs/autoscaler/)
// 
// <div class = "oics-button" style="float: right; margin: 0 0 -15px">
//   <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=autoscaler_beta&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
//     <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
//   </a>
// </div>
type Autoscalar struct {
	s *pulumi.ResourceState
}

// NewAutoscalar registers a new resource with the given unique name, arguments, and options.
func NewAutoscalar(ctx *pulumi.Context,
	name string, args *AutoscalarArgs, opts ...pulumi.ResourceOpt) (*Autoscalar, error) {
	if args == nil || args.AutoscalingPolicy == nil {
		return nil, errors.New("missing required argument 'AutoscalingPolicy'")
	}
	if args == nil || args.Target == nil {
		return nil, errors.New("missing required argument 'Target'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["autoscalingPolicy"] = nil
		inputs["description"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["target"] = nil
		inputs["zone"] = nil
	} else {
		inputs["autoscalingPolicy"] = args.AutoscalingPolicy
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["target"] = args.Target
		inputs["zone"] = args.Zone
	}
	inputs["creationTimestamp"] = nil
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:compute/autoscalar:Autoscalar", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Autoscalar{s: s}, nil
}

// GetAutoscalar gets an existing Autoscalar resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoscalar(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AutoscalarState, opts ...pulumi.ResourceOpt) (*Autoscalar, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["autoscalingPolicy"] = state.AutoscalingPolicy
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["selfLink"] = state.SelfLink
		inputs["target"] = state.Target
		inputs["zone"] = state.Zone
	}
	s, err := ctx.ReadResource("gcp:compute/autoscalar:Autoscalar", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Autoscalar{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Autoscalar) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Autoscalar) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *Autoscalar) AutoscalingPolicy() *pulumi.Output {
	return r.s.State["autoscalingPolicy"]
}

func (r *Autoscalar) CreationTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

func (r *Autoscalar) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *Autoscalar) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *Autoscalar) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The URI of the created resource.
func (r *Autoscalar) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

func (r *Autoscalar) Target() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["target"])
}

func (r *Autoscalar) Zone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zone"])
}

// Input properties used for looking up and filtering Autoscalar resources.
type AutoscalarState struct {
	AutoscalingPolicy interface{}
	CreationTimestamp interface{}
	Description interface{}
	Name interface{}
	Project interface{}
	// The URI of the created resource.
	SelfLink interface{}
	Target interface{}
	Zone interface{}
}

// The set of arguments for constructing a Autoscalar resource.
type AutoscalarArgs struct {
	AutoscalingPolicy interface{}
	Description interface{}
	Name interface{}
	Project interface{}
	Target interface{}
	Zone interface{}
}

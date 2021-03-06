// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package runtimeconfig

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a RuntimeConfig resource in Google Cloud. For more information, see the
// [official documentation](https://cloud.google.com/deployment-manager/runtime-configurator/),
// or the
// [JSON API](https://cloud.google.com/deployment-manager/runtime-configurator/reference/rest/).
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/runtimeconfig_config.html.markdown.
type Config struct {
	pulumi.CustomResourceState

	// The description to associate with the runtime
	// config.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the runtime config.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewConfig registers a new resource with the given unique name, arguments, and options.
func NewConfig(ctx *pulumi.Context,
	name string, args *ConfigArgs, opts ...pulumi.ResourceOption) (*Config, error) {
	if args == nil {
		args = &ConfigArgs{}
	}
	var resource Config
	err := ctx.RegisterResource("gcp:runtimeconfig/config:Config", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfig gets an existing Config resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigState, opts ...pulumi.ResourceOption) (*Config, error) {
	var resource Config
	err := ctx.ReadResource("gcp:runtimeconfig/config:Config", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Config resources.
type configState struct {
	// The description to associate with the runtime
	// config.
	Description *string `pulumi:"description"`
	// The name of the runtime config.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type ConfigState struct {
	// The description to associate with the runtime
	// config.
	Description pulumi.StringPtrInput
	// The name of the runtime config.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*configState)(nil)).Elem()
}

type configArgs struct {
	// The description to associate with the runtime
	// config.
	Description *string `pulumi:"description"`
	// The name of the runtime config.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Config resource.
type ConfigArgs struct {
	// The description to associate with the runtime
	// config.
	Description pulumi.StringPtrInput
	// The name of the runtime config.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configArgs)(nil)).Elem()
}


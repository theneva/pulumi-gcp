// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package endpoints

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ServiceApi struct {
	Methods []ServiceApiMethod `pulumi:"methods"`
	Name *string `pulumi:"name"`
	Syntax *string `pulumi:"syntax"`
	Version *string `pulumi:"version"`
}

type ServiceApiInput interface {
	pulumi.Input

	ToServiceApiOutput() ServiceApiOutput
	ToServiceApiOutputWithContext(context.Context) ServiceApiOutput
}

type ServiceApiArgs struct {
	Methods ServiceApiMethodArrayInput `pulumi:"methods"`
	Name pulumi.StringPtrInput `pulumi:"name"`
	Syntax pulumi.StringPtrInput `pulumi:"syntax"`
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (ServiceApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceApi)(nil)).Elem()
}

func (i ServiceApiArgs) ToServiceApiOutput() ServiceApiOutput {
	return i.ToServiceApiOutputWithContext(context.Background())
}

func (i ServiceApiArgs) ToServiceApiOutputWithContext(ctx context.Context) ServiceApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceApiOutput)
}

type ServiceApiArrayInput interface {
	pulumi.Input

	ToServiceApiArrayOutput() ServiceApiArrayOutput
	ToServiceApiArrayOutputWithContext(context.Context) ServiceApiArrayOutput
}

type ServiceApiArray []ServiceApiInput

func (ServiceApiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceApi)(nil)).Elem()
}

func (i ServiceApiArray) ToServiceApiArrayOutput() ServiceApiArrayOutput {
	return i.ToServiceApiArrayOutputWithContext(context.Background())
}

func (i ServiceApiArray) ToServiceApiArrayOutputWithContext(ctx context.Context) ServiceApiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceApiArrayOutput)
}

type ServiceApiOutput struct { *pulumi.OutputState }

func (ServiceApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceApi)(nil)).Elem()
}

func (o ServiceApiOutput) ToServiceApiOutput() ServiceApiOutput {
	return o
}

func (o ServiceApiOutput) ToServiceApiOutputWithContext(ctx context.Context) ServiceApiOutput {
	return o
}

func (o ServiceApiOutput) Methods() ServiceApiMethodArrayOutput {
	return o.ApplyT(func (v ServiceApi) []ServiceApiMethod { return v.Methods }).(ServiceApiMethodArrayOutput)
}

func (o ServiceApiOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceApi) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ServiceApiOutput) Syntax() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceApi) *string { return v.Syntax }).(pulumi.StringPtrOutput)
}

func (o ServiceApiOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceApi) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ServiceApiArrayOutput struct { *pulumi.OutputState}

func (ServiceApiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceApi)(nil)).Elem()
}

func (o ServiceApiArrayOutput) ToServiceApiArrayOutput() ServiceApiArrayOutput {
	return o
}

func (o ServiceApiArrayOutput) ToServiceApiArrayOutputWithContext(ctx context.Context) ServiceApiArrayOutput {
	return o
}

func (o ServiceApiArrayOutput) Index(i pulumi.IntInput) ServiceApiOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ServiceApi {
		return vs[0].([]ServiceApi)[vs[1].(int)]
	}).(ServiceApiOutput)
}

type ServiceApiMethod struct {
	Name *string `pulumi:"name"`
	RequestType *string `pulumi:"requestType"`
	ResponseType *string `pulumi:"responseType"`
	Syntax *string `pulumi:"syntax"`
}

type ServiceApiMethodInput interface {
	pulumi.Input

	ToServiceApiMethodOutput() ServiceApiMethodOutput
	ToServiceApiMethodOutputWithContext(context.Context) ServiceApiMethodOutput
}

type ServiceApiMethodArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	RequestType pulumi.StringPtrInput `pulumi:"requestType"`
	ResponseType pulumi.StringPtrInput `pulumi:"responseType"`
	Syntax pulumi.StringPtrInput `pulumi:"syntax"`
}

func (ServiceApiMethodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceApiMethod)(nil)).Elem()
}

func (i ServiceApiMethodArgs) ToServiceApiMethodOutput() ServiceApiMethodOutput {
	return i.ToServiceApiMethodOutputWithContext(context.Background())
}

func (i ServiceApiMethodArgs) ToServiceApiMethodOutputWithContext(ctx context.Context) ServiceApiMethodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceApiMethodOutput)
}

type ServiceApiMethodArrayInput interface {
	pulumi.Input

	ToServiceApiMethodArrayOutput() ServiceApiMethodArrayOutput
	ToServiceApiMethodArrayOutputWithContext(context.Context) ServiceApiMethodArrayOutput
}

type ServiceApiMethodArray []ServiceApiMethodInput

func (ServiceApiMethodArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceApiMethod)(nil)).Elem()
}

func (i ServiceApiMethodArray) ToServiceApiMethodArrayOutput() ServiceApiMethodArrayOutput {
	return i.ToServiceApiMethodArrayOutputWithContext(context.Background())
}

func (i ServiceApiMethodArray) ToServiceApiMethodArrayOutputWithContext(ctx context.Context) ServiceApiMethodArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceApiMethodArrayOutput)
}

type ServiceApiMethodOutput struct { *pulumi.OutputState }

func (ServiceApiMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceApiMethod)(nil)).Elem()
}

func (o ServiceApiMethodOutput) ToServiceApiMethodOutput() ServiceApiMethodOutput {
	return o
}

func (o ServiceApiMethodOutput) ToServiceApiMethodOutputWithContext(ctx context.Context) ServiceApiMethodOutput {
	return o
}

func (o ServiceApiMethodOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceApiMethod) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ServiceApiMethodOutput) RequestType() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceApiMethod) *string { return v.RequestType }).(pulumi.StringPtrOutput)
}

func (o ServiceApiMethodOutput) ResponseType() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceApiMethod) *string { return v.ResponseType }).(pulumi.StringPtrOutput)
}

func (o ServiceApiMethodOutput) Syntax() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceApiMethod) *string { return v.Syntax }).(pulumi.StringPtrOutput)
}

type ServiceApiMethodArrayOutput struct { *pulumi.OutputState}

func (ServiceApiMethodArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceApiMethod)(nil)).Elem()
}

func (o ServiceApiMethodArrayOutput) ToServiceApiMethodArrayOutput() ServiceApiMethodArrayOutput {
	return o
}

func (o ServiceApiMethodArrayOutput) ToServiceApiMethodArrayOutputWithContext(ctx context.Context) ServiceApiMethodArrayOutput {
	return o
}

func (o ServiceApiMethodArrayOutput) Index(i pulumi.IntInput) ServiceApiMethodOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ServiceApiMethod {
		return vs[0].([]ServiceApiMethod)[vs[1].(int)]
	}).(ServiceApiMethodOutput)
}

type ServiceEndpoint struct {
	Address *string `pulumi:"address"`
	Name *string `pulumi:"name"`
}

type ServiceEndpointInput interface {
	pulumi.Input

	ToServiceEndpointOutput() ServiceEndpointOutput
	ToServiceEndpointOutputWithContext(context.Context) ServiceEndpointOutput
}

type ServiceEndpointArgs struct {
	Address pulumi.StringPtrInput `pulumi:"address"`
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ServiceEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpoint)(nil)).Elem()
}

func (i ServiceEndpointArgs) ToServiceEndpointOutput() ServiceEndpointOutput {
	return i.ToServiceEndpointOutputWithContext(context.Background())
}

func (i ServiceEndpointArgs) ToServiceEndpointOutputWithContext(ctx context.Context) ServiceEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointOutput)
}

type ServiceEndpointArrayInput interface {
	pulumi.Input

	ToServiceEndpointArrayOutput() ServiceEndpointArrayOutput
	ToServiceEndpointArrayOutputWithContext(context.Context) ServiceEndpointArrayOutput
}

type ServiceEndpointArray []ServiceEndpointInput

func (ServiceEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpoint)(nil)).Elem()
}

func (i ServiceEndpointArray) ToServiceEndpointArrayOutput() ServiceEndpointArrayOutput {
	return i.ToServiceEndpointArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointArray) ToServiceEndpointArrayOutputWithContext(ctx context.Context) ServiceEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointArrayOutput)
}

type ServiceEndpointOutput struct { *pulumi.OutputState }

func (ServiceEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpoint)(nil)).Elem()
}

func (o ServiceEndpointOutput) ToServiceEndpointOutput() ServiceEndpointOutput {
	return o
}

func (o ServiceEndpointOutput) ToServiceEndpointOutputWithContext(ctx context.Context) ServiceEndpointOutput {
	return o
}

func (o ServiceEndpointOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceEndpoint) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o ServiceEndpointOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceEndpoint) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ServiceEndpointArrayOutput struct { *pulumi.OutputState}

func (ServiceEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpoint)(nil)).Elem()
}

func (o ServiceEndpointArrayOutput) ToServiceEndpointArrayOutput() ServiceEndpointArrayOutput {
	return o
}

func (o ServiceEndpointArrayOutput) ToServiceEndpointArrayOutputWithContext(ctx context.Context) ServiceEndpointArrayOutput {
	return o
}

func (o ServiceEndpointArrayOutput) Index(i pulumi.IntInput) ServiceEndpointOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ServiceEndpoint {
		return vs[0].([]ServiceEndpoint)[vs[1].(int)]
	}).(ServiceEndpointOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceApiOutput{})
	pulumi.RegisterOutputType(ServiceApiArrayOutput{})
	pulumi.RegisterOutputType(ServiceApiMethodOutput{})
	pulumi.RegisterOutputType(ServiceApiMethodArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointOutput{})
	pulumi.RegisterOutputType(ServiceEndpointArrayOutput{})
}
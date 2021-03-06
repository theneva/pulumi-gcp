// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package containeranalysis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type NoteAttestationAuthority struct {
	Hint NoteAttestationAuthorityHint `pulumi:"hint"`
}

type NoteAttestationAuthorityInput interface {
	pulumi.Input

	ToNoteAttestationAuthorityOutput() NoteAttestationAuthorityOutput
	ToNoteAttestationAuthorityOutputWithContext(context.Context) NoteAttestationAuthorityOutput
}

type NoteAttestationAuthorityArgs struct {
	Hint NoteAttestationAuthorityHintInput `pulumi:"hint"`
}

func (NoteAttestationAuthorityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NoteAttestationAuthority)(nil)).Elem()
}

func (i NoteAttestationAuthorityArgs) ToNoteAttestationAuthorityOutput() NoteAttestationAuthorityOutput {
	return i.ToNoteAttestationAuthorityOutputWithContext(context.Background())
}

func (i NoteAttestationAuthorityArgs) ToNoteAttestationAuthorityOutputWithContext(ctx context.Context) NoteAttestationAuthorityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoteAttestationAuthorityOutput)
}

func (i NoteAttestationAuthorityArgs) ToNoteAttestationAuthorityPtrOutput() NoteAttestationAuthorityPtrOutput {
	return i.ToNoteAttestationAuthorityPtrOutputWithContext(context.Background())
}

func (i NoteAttestationAuthorityArgs) ToNoteAttestationAuthorityPtrOutputWithContext(ctx context.Context) NoteAttestationAuthorityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoteAttestationAuthorityOutput).ToNoteAttestationAuthorityPtrOutputWithContext(ctx)
}

type NoteAttestationAuthorityPtrInput interface {
	pulumi.Input

	ToNoteAttestationAuthorityPtrOutput() NoteAttestationAuthorityPtrOutput
	ToNoteAttestationAuthorityPtrOutputWithContext(context.Context) NoteAttestationAuthorityPtrOutput
}

type noteAttestationAuthorityPtrType NoteAttestationAuthorityArgs

func NoteAttestationAuthorityPtr(v *NoteAttestationAuthorityArgs) NoteAttestationAuthorityPtrInput {	return (*noteAttestationAuthorityPtrType)(v)
}

func (*noteAttestationAuthorityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NoteAttestationAuthority)(nil)).Elem()
}

func (i *noteAttestationAuthorityPtrType) ToNoteAttestationAuthorityPtrOutput() NoteAttestationAuthorityPtrOutput {
	return i.ToNoteAttestationAuthorityPtrOutputWithContext(context.Background())
}

func (i *noteAttestationAuthorityPtrType) ToNoteAttestationAuthorityPtrOutputWithContext(ctx context.Context) NoteAttestationAuthorityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoteAttestationAuthorityPtrOutput)
}

type NoteAttestationAuthorityOutput struct { *pulumi.OutputState }

func (NoteAttestationAuthorityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NoteAttestationAuthority)(nil)).Elem()
}

func (o NoteAttestationAuthorityOutput) ToNoteAttestationAuthorityOutput() NoteAttestationAuthorityOutput {
	return o
}

func (o NoteAttestationAuthorityOutput) ToNoteAttestationAuthorityOutputWithContext(ctx context.Context) NoteAttestationAuthorityOutput {
	return o
}

func (o NoteAttestationAuthorityOutput) ToNoteAttestationAuthorityPtrOutput() NoteAttestationAuthorityPtrOutput {
	return o.ToNoteAttestationAuthorityPtrOutputWithContext(context.Background())
}

func (o NoteAttestationAuthorityOutput) ToNoteAttestationAuthorityPtrOutputWithContext(ctx context.Context) NoteAttestationAuthorityPtrOutput {
	return o.ApplyT(func(v NoteAttestationAuthority) *NoteAttestationAuthority {
		return &v
	}).(NoteAttestationAuthorityPtrOutput)
}
func (o NoteAttestationAuthorityOutput) Hint() NoteAttestationAuthorityHintOutput {
	return o.ApplyT(func (v NoteAttestationAuthority) NoteAttestationAuthorityHint { return v.Hint }).(NoteAttestationAuthorityHintOutput)
}

type NoteAttestationAuthorityPtrOutput struct { *pulumi.OutputState}

func (NoteAttestationAuthorityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NoteAttestationAuthority)(nil)).Elem()
}

func (o NoteAttestationAuthorityPtrOutput) ToNoteAttestationAuthorityPtrOutput() NoteAttestationAuthorityPtrOutput {
	return o
}

func (o NoteAttestationAuthorityPtrOutput) ToNoteAttestationAuthorityPtrOutputWithContext(ctx context.Context) NoteAttestationAuthorityPtrOutput {
	return o
}

func (o NoteAttestationAuthorityPtrOutput) Elem() NoteAttestationAuthorityOutput {
	return o.ApplyT(func (v *NoteAttestationAuthority) NoteAttestationAuthority { return *v }).(NoteAttestationAuthorityOutput)
}

func (o NoteAttestationAuthorityPtrOutput) Hint() NoteAttestationAuthorityHintOutput {
	return o.ApplyT(func (v NoteAttestationAuthority) NoteAttestationAuthorityHint { return v.Hint }).(NoteAttestationAuthorityHintOutput)
}

type NoteAttestationAuthorityHint struct {
	HumanReadableName string `pulumi:"humanReadableName"`
}

type NoteAttestationAuthorityHintInput interface {
	pulumi.Input

	ToNoteAttestationAuthorityHintOutput() NoteAttestationAuthorityHintOutput
	ToNoteAttestationAuthorityHintOutputWithContext(context.Context) NoteAttestationAuthorityHintOutput
}

type NoteAttestationAuthorityHintArgs struct {
	HumanReadableName pulumi.StringInput `pulumi:"humanReadableName"`
}

func (NoteAttestationAuthorityHintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NoteAttestationAuthorityHint)(nil)).Elem()
}

func (i NoteAttestationAuthorityHintArgs) ToNoteAttestationAuthorityHintOutput() NoteAttestationAuthorityHintOutput {
	return i.ToNoteAttestationAuthorityHintOutputWithContext(context.Background())
}

func (i NoteAttestationAuthorityHintArgs) ToNoteAttestationAuthorityHintOutputWithContext(ctx context.Context) NoteAttestationAuthorityHintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoteAttestationAuthorityHintOutput)
}

type NoteAttestationAuthorityHintOutput struct { *pulumi.OutputState }

func (NoteAttestationAuthorityHintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NoteAttestationAuthorityHint)(nil)).Elem()
}

func (o NoteAttestationAuthorityHintOutput) ToNoteAttestationAuthorityHintOutput() NoteAttestationAuthorityHintOutput {
	return o
}

func (o NoteAttestationAuthorityHintOutput) ToNoteAttestationAuthorityHintOutputWithContext(ctx context.Context) NoteAttestationAuthorityHintOutput {
	return o
}

func (o NoteAttestationAuthorityHintOutput) HumanReadableName() pulumi.StringOutput {
	return o.ApplyT(func (v NoteAttestationAuthorityHint) string { return v.HumanReadableName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NoteAttestationAuthorityOutput{})
	pulumi.RegisterOutputType(NoteAttestationAuthorityPtrOutput{})
	pulumi.RegisterOutputType(NoteAttestationAuthorityHintOutput{})
}

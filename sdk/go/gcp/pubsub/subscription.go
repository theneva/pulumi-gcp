// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/pubsub_subscription.html.markdown.
type Subscription struct {
	s *pulumi.ResourceState
}

// NewSubscription registers a new resource with the given unique name, arguments, and options.
func NewSubscription(ctx *pulumi.Context,
	name string, args *SubscriptionArgs, opts ...pulumi.ResourceOpt) (*Subscription, error) {
	if args == nil || args.Topic == nil {
		return nil, errors.New("missing required argument 'Topic'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["ackDeadlineSeconds"] = nil
		inputs["expirationPolicy"] = nil
		inputs["labels"] = nil
		inputs["messageRetentionDuration"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["pushConfig"] = nil
		inputs["retainAckedMessages"] = nil
		inputs["topic"] = nil
	} else {
		inputs["ackDeadlineSeconds"] = args.AckDeadlineSeconds
		inputs["expirationPolicy"] = args.ExpirationPolicy
		inputs["labels"] = args.Labels
		inputs["messageRetentionDuration"] = args.MessageRetentionDuration
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["pushConfig"] = args.PushConfig
		inputs["retainAckedMessages"] = args.RetainAckedMessages
		inputs["topic"] = args.Topic
	}
	inputs["path"] = nil
	s, err := ctx.RegisterResource("gcp:pubsub/subscription:Subscription", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Subscription{s: s}, nil
}

// GetSubscription gets an existing Subscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscription(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SubscriptionState, opts ...pulumi.ResourceOpt) (*Subscription, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["ackDeadlineSeconds"] = state.AckDeadlineSeconds
		inputs["expirationPolicy"] = state.ExpirationPolicy
		inputs["labels"] = state.Labels
		inputs["messageRetentionDuration"] = state.MessageRetentionDuration
		inputs["name"] = state.Name
		inputs["path"] = state.Path
		inputs["project"] = state.Project
		inputs["pushConfig"] = state.PushConfig
		inputs["retainAckedMessages"] = state.RetainAckedMessages
		inputs["topic"] = state.Topic
	}
	s, err := ctx.ReadResource("gcp:pubsub/subscription:Subscription", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Subscription{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Subscription) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Subscription) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// This value is the maximum time after a subscriber receives a message before the subscriber should acknowledge the
// message. After message delivery but before the ack deadline expires and before the message is acknowledged, it is an
// outstanding message and will not be delivered again during that time (on a best-effort basis). For pull subscriptions,
// this value is used as the initial value for the ack deadline. To override this value for a given message, call
// subscriptions.modifyAckDeadline with the corresponding ackId if using pull. The minimum custom deadline you can specify
// is 10 seconds. The maximum custom deadline you can specify is 600 seconds (10 minutes). If this parameter is 0, a
// default value of 10 seconds is used. For push delivery, this value is also used to set the request timeout for the call
// to the push endpoint. If the subscriber never acknowledges the message, the Pub/Sub system will eventually redeliver the
// message.
func (r *Subscription) AckDeadlineSeconds() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["ackDeadlineSeconds"])
}

// A policy that specifies the conditions for this subscription's expiration. A subscription is considered active as long
// as any connected subscriber is successfully consuming messages from the subscription or is issuing operations on the
// subscription. If expirationPolicy is not set, a default policy with ttl of 31 days will be used. If it is set but left
// empty, the resource never expires. The minimum allowed value for expirationPolicy.ttl is 1 day.
func (r *Subscription) ExpirationPolicy() *pulumi.Output {
	return r.s.State["expirationPolicy"]
}

// A set of key/value label pairs to assign to this Subscription.
func (r *Subscription) Labels() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["labels"])
}

// How long to retain unacknowledged messages in the subscription's backlog, from the moment a message is published. If
// retainAckedMessages is true, then this also configures the retention of acknowledged messages, and thus configures how
// far back in time a subscriptions.seek can be done. Defaults to 7 days. Cannot be more than 7 days ('"604800s"') or less
// than 10 minutes ('"600s"'). A duration in seconds with up to nine fractional digits, terminated by 's'. Example:
// '"600.5s"'.
func (r *Subscription) MessageRetentionDuration() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["messageRetentionDuration"])
}

// Name of the subscription.
func (r *Subscription) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *Subscription) Path() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["path"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *Subscription) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// If push delivery is used with this subscription, this field is used to configure it. An empty pushConfig signifies that
// the subscriber will pull and ack messages using API methods.
func (r *Subscription) PushConfig() *pulumi.Output {
	return r.s.State["pushConfig"]
}

// Indicates whether to retain acknowledged messages. If 'true', then messages are not expunged from the subscription's
// backlog, even if they are acknowledged, until they fall out of the messageRetentionDuration window.
func (r *Subscription) RetainAckedMessages() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["retainAckedMessages"])
}

// A reference to a Topic resource.
func (r *Subscription) Topic() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["topic"])
}

// Input properties used for looking up and filtering Subscription resources.
type SubscriptionState struct {
	// This value is the maximum time after a subscriber receives a message before the subscriber should acknowledge the
	// message. After message delivery but before the ack deadline expires and before the message is acknowledged, it is an
	// outstanding message and will not be delivered again during that time (on a best-effort basis). For pull subscriptions,
	// this value is used as the initial value for the ack deadline. To override this value for a given message, call
	// subscriptions.modifyAckDeadline with the corresponding ackId if using pull. The minimum custom deadline you can specify
	// is 10 seconds. The maximum custom deadline you can specify is 600 seconds (10 minutes). If this parameter is 0, a
	// default value of 10 seconds is used. For push delivery, this value is also used to set the request timeout for the call
	// to the push endpoint. If the subscriber never acknowledges the message, the Pub/Sub system will eventually redeliver
	// the message.
	AckDeadlineSeconds interface{}
	// A policy that specifies the conditions for this subscription's expiration. A subscription is considered active as long
	// as any connected subscriber is successfully consuming messages from the subscription or is issuing operations on the
	// subscription. If expirationPolicy is not set, a default policy with ttl of 31 days will be used. If it is set but left
	// empty, the resource never expires. The minimum allowed value for expirationPolicy.ttl is 1 day.
	ExpirationPolicy interface{}
	// A set of key/value label pairs to assign to this Subscription.
	Labels interface{}
	// How long to retain unacknowledged messages in the subscription's backlog, from the moment a message is published. If
	// retainAckedMessages is true, then this also configures the retention of acknowledged messages, and thus configures how
	// far back in time a subscriptions.seek can be done. Defaults to 7 days. Cannot be more than 7 days ('"604800s"') or less
	// than 10 minutes ('"600s"'). A duration in seconds with up to nine fractional digits, terminated by 's'. Example:
	// '"600.5s"'.
	MessageRetentionDuration interface{}
	// Name of the subscription.
	Name interface{}
	Path interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// If push delivery is used with this subscription, this field is used to configure it. An empty pushConfig signifies that
	// the subscriber will pull and ack messages using API methods.
	PushConfig interface{}
	// Indicates whether to retain acknowledged messages. If 'true', then messages are not expunged from the subscription's
	// backlog, even if they are acknowledged, until they fall out of the messageRetentionDuration window.
	RetainAckedMessages interface{}
	// A reference to a Topic resource.
	Topic interface{}
}

// The set of arguments for constructing a Subscription resource.
type SubscriptionArgs struct {
	// This value is the maximum time after a subscriber receives a message before the subscriber should acknowledge the
	// message. After message delivery but before the ack deadline expires and before the message is acknowledged, it is an
	// outstanding message and will not be delivered again during that time (on a best-effort basis). For pull subscriptions,
	// this value is used as the initial value for the ack deadline. To override this value for a given message, call
	// subscriptions.modifyAckDeadline with the corresponding ackId if using pull. The minimum custom deadline you can specify
	// is 10 seconds. The maximum custom deadline you can specify is 600 seconds (10 minutes). If this parameter is 0, a
	// default value of 10 seconds is used. For push delivery, this value is also used to set the request timeout for the call
	// to the push endpoint. If the subscriber never acknowledges the message, the Pub/Sub system will eventually redeliver
	// the message.
	AckDeadlineSeconds interface{}
	// A policy that specifies the conditions for this subscription's expiration. A subscription is considered active as long
	// as any connected subscriber is successfully consuming messages from the subscription or is issuing operations on the
	// subscription. If expirationPolicy is not set, a default policy with ttl of 31 days will be used. If it is set but left
	// empty, the resource never expires. The minimum allowed value for expirationPolicy.ttl is 1 day.
	ExpirationPolicy interface{}
	// A set of key/value label pairs to assign to this Subscription.
	Labels interface{}
	// How long to retain unacknowledged messages in the subscription's backlog, from the moment a message is published. If
	// retainAckedMessages is true, then this also configures the retention of acknowledged messages, and thus configures how
	// far back in time a subscriptions.seek can be done. Defaults to 7 days. Cannot be more than 7 days ('"604800s"') or less
	// than 10 minutes ('"600s"'). A duration in seconds with up to nine fractional digits, terminated by 's'. Example:
	// '"600.5s"'.
	MessageRetentionDuration interface{}
	// Name of the subscription.
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// If push delivery is used with this subscription, this field is used to configure it. An empty pushConfig signifies that
	// the subscriber will pull and ack messages using API methods.
	PushConfig interface{}
	// Indicates whether to retain acknowledged messages. If 'true', then messages are not expunged from the subscription's
	// backlog, even if they are acknowledged, until they fall out of the messageRetentionDuration window.
	RetainAckedMessages interface{}
	// A reference to a Topic resource.
	Topic interface{}
}

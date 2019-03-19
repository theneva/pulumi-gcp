// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfunctions

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get information about a Google Cloud Function. For more information see
// the [official documentation](https://cloud.google.com/functions/docs/)
// and [API](https://cloud.google.com/functions/docs/apis).
func LookupFunction(ctx *pulumi.Context, args *GetFunctionArgs) (*GetFunctionResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["region"] = args.Region
	}
	outputs, err := ctx.Invoke("gcp:cloudfunctions/getFunction:getFunction", inputs)
	if err != nil {
		return nil, err
	}
	return &GetFunctionResult{
		AvailableMemoryMb: outputs["availableMemoryMb"],
		Description: outputs["description"],
		EntryPoint: outputs["entryPoint"],
		EnvironmentVariables: outputs["environmentVariables"],
		EventTriggers: outputs["eventTriggers"],
		HttpsTriggerUrl: outputs["httpsTriggerUrl"],
		Labels: outputs["labels"],
		RetryOnFailure: outputs["retryOnFailure"],
		Runtime: outputs["runtime"],
		ServiceAccountEmail: outputs["serviceAccountEmail"],
		SourceArchiveBucket: outputs["sourceArchiveBucket"],
		SourceArchiveObject: outputs["sourceArchiveObject"],
		SourceRepositories: outputs["sourceRepositories"],
		Timeout: outputs["timeout"],
		TriggerBucket: outputs["triggerBucket"],
		TriggerHttp: outputs["triggerHttp"],
		TriggerTopic: outputs["triggerTopic"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getFunction.
type GetFunctionArgs struct {
	// The name of a Cloud Function.
	Name interface{}
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The region in which the resource belongs. If it
	// is not provided, the provider region is used.
	Region interface{}
}

// A collection of values returned by getFunction.
type GetFunctionResult struct {
	// Available memory (in MB) to the function.
	AvailableMemoryMb interface{}
	// Description of the function.
	Description interface{}
	// Name of a JavaScript function that will be executed when the Google Cloud Function is triggered.
	EntryPoint interface{}
	EnvironmentVariables interface{}
	// A source that fires events in response to a condition in another service. Structure is documented below.
	EventTriggers interface{}
	// If function is triggered by HTTP, trigger URL is set here.
	HttpsTriggerUrl interface{}
	// A map of labels applied to this function.
	Labels interface{}
	RetryOnFailure interface{}
	// The runtime in which the function is running.
	Runtime interface{}
	ServiceAccountEmail interface{}
	// The GCS bucket containing the zip archive which contains the function.
	SourceArchiveBucket interface{}
	// The source archive object (file) in archive bucket.
	SourceArchiveObject interface{}
	SourceRepositories interface{}
	// Function execution timeout (in seconds).
	Timeout interface{}
	TriggerBucket interface{}
	// If function is triggered by HTTP, this boolean is set.
	TriggerHttp interface{}
	TriggerTopic interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}

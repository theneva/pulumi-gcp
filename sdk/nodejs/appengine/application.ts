// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Allows creation and management of an App Engine application.
 * 
 * > App Engine applications cannot be deleted once they're created; you have to delete the
 *    entire project to delete the application. This provider will report the application has been
 *    successfully deleted; this is a limitation of this provider, and will go away in the future.
 *    This provider is not able to delete App Engine applications.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/app_engine_application.html.markdown.
 */
export class Application extends pulumi.CustomResource {
    /**
     * Get an existing Application resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationState, opts?: pulumi.CustomResourceOptions): Application {
        return new Application(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:appengine/application:Application';

    /**
     * Returns true if the given object is an instance of Application.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Application {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Application.__pulumiType;
    }

    /**
     * Identifier of the app, usually `{PROJECT_ID}`
     */
    public /*out*/ readonly appId!: pulumi.Output<string>;
    /**
     * The domain to authenticate users with when using App Engine's User API.
     */
    public readonly authDomain!: pulumi.Output<string>;
    /**
     * The GCS bucket code is being stored in for this app.
     */
    public /*out*/ readonly codeBucket!: pulumi.Output<string>;
    /**
     * The GCS bucket content is being stored in for this app.
     */
    public /*out*/ readonly defaultBucket!: pulumi.Output<string>;
    /**
     * The default hostname for this app.
     */
    public /*out*/ readonly defaultHostname!: pulumi.Output<string>;
    /**
     * A block of optional settings to configure specific App Engine features:
     */
    public readonly featureSettings!: pulumi.Output<outputs.appengine.ApplicationFeatureSettings>;
    /**
     * The GCR domain used for storing managed Docker images for this app.
     */
    public /*out*/ readonly gcrDomain!: pulumi.Output<string>;
    /**
     * The [location](https://cloud.google.com/appengine/docs/locations)
     * to serve the app from.
     */
    public readonly locationId!: pulumi.Output<string>;
    /**
     * Unique name of the app, usually `apps/{PROJECT_ID}`
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The project ID to create the application under.
     * ~>**NOTE**: GCP only accepts project ID, not project number. If you are using number,
     * you may get a "Permission denied" error.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The serving status of the app.
     */
    public readonly servingStatus!: pulumi.Output<string>;
    /**
     * A list of dispatch rule blocks. Each block has a `domain`, `path`, and `service` field.
     */
    public /*out*/ readonly urlDispatchRules!: pulumi.Output<outputs.appengine.ApplicationUrlDispatchRule[]>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationArgs | ApplicationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ApplicationState | undefined;
            inputs["appId"] = state ? state.appId : undefined;
            inputs["authDomain"] = state ? state.authDomain : undefined;
            inputs["codeBucket"] = state ? state.codeBucket : undefined;
            inputs["defaultBucket"] = state ? state.defaultBucket : undefined;
            inputs["defaultHostname"] = state ? state.defaultHostname : undefined;
            inputs["featureSettings"] = state ? state.featureSettings : undefined;
            inputs["gcrDomain"] = state ? state.gcrDomain : undefined;
            inputs["locationId"] = state ? state.locationId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["servingStatus"] = state ? state.servingStatus : undefined;
            inputs["urlDispatchRules"] = state ? state.urlDispatchRules : undefined;
        } else {
            const args = argsOrState as ApplicationArgs | undefined;
            if (!args || args.locationId === undefined) {
                throw new Error("Missing required property 'locationId'");
            }
            inputs["authDomain"] = args ? args.authDomain : undefined;
            inputs["featureSettings"] = args ? args.featureSettings : undefined;
            inputs["locationId"] = args ? args.locationId : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["servingStatus"] = args ? args.servingStatus : undefined;
            inputs["appId"] = undefined /*out*/;
            inputs["codeBucket"] = undefined /*out*/;
            inputs["defaultBucket"] = undefined /*out*/;
            inputs["defaultHostname"] = undefined /*out*/;
            inputs["gcrDomain"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["urlDispatchRules"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Application.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Application resources.
 */
export interface ApplicationState {
    /**
     * Identifier of the app, usually `{PROJECT_ID}`
     */
    readonly appId?: pulumi.Input<string>;
    /**
     * The domain to authenticate users with when using App Engine's User API.
     */
    readonly authDomain?: pulumi.Input<string>;
    /**
     * The GCS bucket code is being stored in for this app.
     */
    readonly codeBucket?: pulumi.Input<string>;
    /**
     * The GCS bucket content is being stored in for this app.
     */
    readonly defaultBucket?: pulumi.Input<string>;
    /**
     * The default hostname for this app.
     */
    readonly defaultHostname?: pulumi.Input<string>;
    /**
     * A block of optional settings to configure specific App Engine features:
     */
    readonly featureSettings?: pulumi.Input<inputs.appengine.ApplicationFeatureSettings>;
    /**
     * The GCR domain used for storing managed Docker images for this app.
     */
    readonly gcrDomain?: pulumi.Input<string>;
    /**
     * The [location](https://cloud.google.com/appengine/docs/locations)
     * to serve the app from.
     */
    readonly locationId?: pulumi.Input<string>;
    /**
     * Unique name of the app, usually `apps/{PROJECT_ID}`
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The project ID to create the application under.
     * ~>**NOTE**: GCP only accepts project ID, not project number. If you are using number,
     * you may get a "Permission denied" error.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The serving status of the app.
     */
    readonly servingStatus?: pulumi.Input<string>;
    /**
     * A list of dispatch rule blocks. Each block has a `domain`, `path`, and `service` field.
     */
    readonly urlDispatchRules?: pulumi.Input<pulumi.Input<inputs.appengine.ApplicationUrlDispatchRule>[]>;
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    /**
     * The domain to authenticate users with when using App Engine's User API.
     */
    readonly authDomain?: pulumi.Input<string>;
    /**
     * A block of optional settings to configure specific App Engine features:
     */
    readonly featureSettings?: pulumi.Input<inputs.appengine.ApplicationFeatureSettings>;
    /**
     * The [location](https://cloud.google.com/appengine/docs/locations)
     * to serve the app from.
     */
    readonly locationId: pulumi.Input<string>;
    /**
     * The project ID to create the application under.
     * ~>**NOTE**: GCP only accepts project ID, not project number. If you are using number,
     * you may get a "Permission denied" error.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The serving status of the app.
     */
    readonly servingStatus?: pulumi.Input<string>;
}

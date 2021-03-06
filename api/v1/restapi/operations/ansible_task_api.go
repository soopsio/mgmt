// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/soopsio/mgmt/api/v1/restapi/operations/access_keys"
	"github.com/soopsio/mgmt/api/v1/restapi/operations/policies"
	"github.com/soopsio/mgmt/api/v1/restapi/operations/tasks"
)

// NewAnsibleTaskAPI creates a new AnsibleTask instance
func NewAnsibleTaskAPI(spec *loads.Document) *AnsibleTaskAPI {
	return &AnsibleTaskAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		AccessKeysCreateAccessKeyHandler: access_keys.CreateAccessKeyHandlerFunc(func(params access_keys.CreateAccessKeyParams) middleware.Responder {
			return middleware.NotImplemented("operation AccessKeysCreateAccessKey has not yet been implemented")
		}),
		PoliciesCreatePolicyHandler: policies.CreatePolicyHandlerFunc(func(params policies.CreatePolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation PoliciesCreatePolicy has not yet been implemented")
		}),
		TasksCreateTaskHandler: tasks.CreateTaskHandlerFunc(func(params tasks.CreateTaskParams) middleware.Responder {
			return middleware.NotImplemented("operation TasksCreateTask has not yet been implemented")
		}),
		PoliciesDelPolicyHandler: policies.DelPolicyHandlerFunc(func(params policies.DelPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation PoliciesDelPolicy has not yet been implemented")
		}),
		AccessKeysDeleteAccessKeyHandler: access_keys.DeleteAccessKeyHandlerFunc(func(params access_keys.DeleteAccessKeyParams) middleware.Responder {
			return middleware.NotImplemented("operation AccessKeysDeleteAccessKey has not yet been implemented")
		}),
		TasksDeleteTaskHandler: tasks.DeleteTaskHandlerFunc(func(params tasks.DeleteTaskParams) middleware.Responder {
			return middleware.NotImplemented("operation TasksDeleteTask has not yet been implemented")
		}),
		AccessKeysGetAccessKeyHandler: access_keys.GetAccessKeyHandlerFunc(func(params access_keys.GetAccessKeyParams) middleware.Responder {
			return middleware.NotImplemented("operation AccessKeysGetAccessKey has not yet been implemented")
		}),
		AccessKeysGetAccessKeyPoliciesHandler: access_keys.GetAccessKeyPoliciesHandlerFunc(func(params access_keys.GetAccessKeyPoliciesParams) middleware.Responder {
			return middleware.NotImplemented("operation AccessKeysGetAccessKeyPolicies has not yet been implemented")
		}),
		AccessKeysGetAccessKeysHandler: access_keys.GetAccessKeysHandlerFunc(func(params access_keys.GetAccessKeysParams) middleware.Responder {
			return middleware.NotImplemented("operation AccessKeysGetAccessKeys has not yet been implemented")
		}),
		PoliciesGetPolicyHandler: policies.GetPolicyHandlerFunc(func(params policies.GetPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation PoliciesGetPolicy has not yet been implemented")
		}),
		TasksGetTaskHandler: tasks.GetTaskHandlerFunc(func(params tasks.GetTaskParams) middleware.Responder {
			return middleware.NotImplemented("operation TasksGetTask has not yet been implemented")
		}),
		TasksGetTasksHandler: tasks.GetTasksHandlerFunc(func(params tasks.GetTasksParams) middleware.Responder {
			return middleware.NotImplemented("operation TasksGetTasks has not yet been implemented")
		}),
	}
}

/*AnsibleTaskAPI Ansible 管理接口 */
type AnsibleTaskAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// AccessKeysCreateAccessKeyHandler sets the operation handler for the create access key operation
	AccessKeysCreateAccessKeyHandler access_keys.CreateAccessKeyHandler
	// PoliciesCreatePolicyHandler sets the operation handler for the create policy operation
	PoliciesCreatePolicyHandler policies.CreatePolicyHandler
	// TasksCreateTaskHandler sets the operation handler for the create task operation
	TasksCreateTaskHandler tasks.CreateTaskHandler
	// PoliciesDelPolicyHandler sets the operation handler for the del policy operation
	PoliciesDelPolicyHandler policies.DelPolicyHandler
	// AccessKeysDeleteAccessKeyHandler sets the operation handler for the delete access key operation
	AccessKeysDeleteAccessKeyHandler access_keys.DeleteAccessKeyHandler
	// TasksDeleteTaskHandler sets the operation handler for the delete task operation
	TasksDeleteTaskHandler tasks.DeleteTaskHandler
	// AccessKeysGetAccessKeyHandler sets the operation handler for the get access key operation
	AccessKeysGetAccessKeyHandler access_keys.GetAccessKeyHandler
	// AccessKeysGetAccessKeyPoliciesHandler sets the operation handler for the get access key policies operation
	AccessKeysGetAccessKeyPoliciesHandler access_keys.GetAccessKeyPoliciesHandler
	// AccessKeysGetAccessKeysHandler sets the operation handler for the get access keys operation
	AccessKeysGetAccessKeysHandler access_keys.GetAccessKeysHandler
	// PoliciesGetPolicyHandler sets the operation handler for the get policy operation
	PoliciesGetPolicyHandler policies.GetPolicyHandler
	// TasksGetTaskHandler sets the operation handler for the get task operation
	TasksGetTaskHandler tasks.GetTaskHandler
	// TasksGetTasksHandler sets the operation handler for the get tasks operation
	TasksGetTasksHandler tasks.GetTasksHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *AnsibleTaskAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *AnsibleTaskAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *AnsibleTaskAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *AnsibleTaskAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *AnsibleTaskAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *AnsibleTaskAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *AnsibleTaskAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the AnsibleTaskAPI
func (o *AnsibleTaskAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.AccessKeysCreateAccessKeyHandler == nil {
		unregistered = append(unregistered, "access_keys.CreateAccessKeyHandler")
	}

	if o.PoliciesCreatePolicyHandler == nil {
		unregistered = append(unregistered, "policies.CreatePolicyHandler")
	}

	if o.TasksCreateTaskHandler == nil {
		unregistered = append(unregistered, "tasks.CreateTaskHandler")
	}

	if o.PoliciesDelPolicyHandler == nil {
		unregistered = append(unregistered, "policies.DelPolicyHandler")
	}

	if o.AccessKeysDeleteAccessKeyHandler == nil {
		unregistered = append(unregistered, "access_keys.DeleteAccessKeyHandler")
	}

	if o.TasksDeleteTaskHandler == nil {
		unregistered = append(unregistered, "tasks.DeleteTaskHandler")
	}

	if o.AccessKeysGetAccessKeyHandler == nil {
		unregistered = append(unregistered, "access_keys.GetAccessKeyHandler")
	}

	if o.AccessKeysGetAccessKeyPoliciesHandler == nil {
		unregistered = append(unregistered, "access_keys.GetAccessKeyPoliciesHandler")
	}

	if o.AccessKeysGetAccessKeysHandler == nil {
		unregistered = append(unregistered, "access_keys.GetAccessKeysHandler")
	}

	if o.PoliciesGetPolicyHandler == nil {
		unregistered = append(unregistered, "policies.GetPolicyHandler")
	}

	if o.TasksGetTaskHandler == nil {
		unregistered = append(unregistered, "tasks.GetTaskHandler")
	}

	if o.TasksGetTasksHandler == nil {
		unregistered = append(unregistered, "tasks.GetTasksHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *AnsibleTaskAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *AnsibleTaskAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *AnsibleTaskAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *AnsibleTaskAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *AnsibleTaskAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *AnsibleTaskAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the ansible task API
func (o *AnsibleTaskAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *AnsibleTaskAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/accesskeys"] = access_keys.NewCreateAccessKey(o.context, o.AccessKeysCreateAccessKeyHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/policies"] = policies.NewCreatePolicy(o.context, o.PoliciesCreatePolicyHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/tasks"] = tasks.NewCreateTask(o.context, o.TasksCreateTaskHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/policies/{id}"] = policies.NewDelPolicy(o.context, o.PoliciesDelPolicyHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/accesskeys/{key_id}"] = access_keys.NewDeleteAccessKey(o.context, o.AccessKeysDeleteAccessKeyHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/tasks/{task_id}"] = tasks.NewDeleteTask(o.context, o.TasksDeleteTaskHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/accesskeys/{key_id}"] = access_keys.NewGetAccessKey(o.context, o.AccessKeysGetAccessKeyHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/accesskeys/{key_id}/policies"] = access_keys.NewGetAccessKeyPolicies(o.context, o.AccessKeysGetAccessKeyPoliciesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/accesskeys"] = access_keys.NewGetAccessKeys(o.context, o.AccessKeysGetAccessKeysHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/policies/{id}"] = policies.NewGetPolicy(o.context, o.PoliciesGetPolicyHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tasks/{task_id}"] = tasks.NewGetTask(o.context, o.TasksGetTaskHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tasks"] = tasks.NewGetTasks(o.context, o.TasksGetTasksHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *AnsibleTaskAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *AnsibleTaskAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

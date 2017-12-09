// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"crypto/tls"
	"net/http"
	"time"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/soopsio/mgmt/api/v1/custom/processer"
	"github.com/soopsio/mgmt/api/v1/restapi/operations"
	"github.com/soopsio/mgmt/api/v1/restapi/operations/access_keys"
	"github.com/soopsio/mgmt/api/v1/restapi/operations/policies"
	"github.com/soopsio/mgmt/api/v1/restapi/operations/tasks"
	"go.uber.org/zap"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name AnsibleTask --spec ../../../conf/swagger.yaml

// 自定义代码部分 -----
var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
}

type ApiConfig struct {
	l *zap.Logger
}

func (ac *ApiConfig) SetLogger(l *zap.Logger) error {
	if l != nil {
		ac.l = l
	}
	return nil
}

// 初始化 API 配置
func NewApiConfig() *ApiConfig {
	log, _ := zap.NewProduction()
	return &ApiConfig{
		l: log,
	}
}

// 自定义代码部分 -----

func configureFlags(api *operations.AnsibleTaskAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.AnsibleTaskAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.AccessKeysCreateAccessKeyHandler = access_keys.CreateAccessKeyHandlerFunc(func(params access_keys.CreateAccessKeyParams) middleware.Responder {
		return middleware.NotImplemented("operation access_keys.CreateAccessKey has not yet been implemented")
	})
	api.PoliciesCreatePolicyHandler = policies.CreatePolicyHandlerFunc(func(params policies.CreatePolicyParams) middleware.Responder {
		return middleware.NotImplemented("operation policies.CreatePolicy has not yet been implemented")
	})
	api.TasksCreateTaskHandler = tasks.CreateTaskHandlerFunc(func(params tasks.CreateTaskParams) middleware.Responder {
		logger.Info("#########", zap.Any("params", params))
		startTime := time.Now()
		resp := processer.CreateTask(params)
		logger.Info("######### request Time", zap.Any("response", resp), zap.Any("params", params), zap.Time("startTime", startTime))

		return resp
		// return middleware.NotImplemented("operation tasks.CreateTask has not yet been implemented")
	})
	api.PoliciesDelPolicyHandler = policies.DelPolicyHandlerFunc(func(params policies.DelPolicyParams) middleware.Responder {
		return middleware.NotImplemented("operation policies.DelPolicy has not yet been implemented")
	})
	api.AccessKeysDeleteAccessKeyHandler = access_keys.DeleteAccessKeyHandlerFunc(func(params access_keys.DeleteAccessKeyParams) middleware.Responder {
		return middleware.NotImplemented("operation access_keys.DeleteAccessKey has not yet been implemented")
	})
	api.TasksDeleteTaskHandler = tasks.DeleteTaskHandlerFunc(func(params tasks.DeleteTaskParams) middleware.Responder {
		return middleware.NotImplemented("operation tasks.DeleteTask has not yet been implemented")
	})
	api.AccessKeysGetAccessKeyHandler = access_keys.GetAccessKeyHandlerFunc(func(params access_keys.GetAccessKeyParams) middleware.Responder {
		return middleware.NotImplemented("operation access_keys.GetAccessKey has not yet been implemented")
	})
	api.AccessKeysGetAccessKeyPoliciesHandler = access_keys.GetAccessKeyPoliciesHandlerFunc(func(params access_keys.GetAccessKeyPoliciesParams) middleware.Responder {
		return middleware.NotImplemented("operation access_keys.GetAccessKeyPolicies has not yet been implemented")
	})
	api.AccessKeysGetAccessKeysHandler = access_keys.GetAccessKeysHandlerFunc(func(params access_keys.GetAccessKeysParams) middleware.Responder {
		return middleware.NotImplemented("operation access_keys.GetAccessKeys has not yet been implemented")
	})
	api.PoliciesGetPolicyHandler = policies.GetPolicyHandlerFunc(func(params policies.GetPolicyParams) middleware.Responder {
		return middleware.NotImplemented("operation policies.GetPolicy has not yet been implemented")
	})
	api.TasksGetTaskHandler = tasks.GetTaskHandlerFunc(func(params tasks.GetTaskParams) middleware.Responder {
		return processer.GetTaskStatus(params.TaskID)
		// return middleware.NotImplemented("operation tasks.GetTask has not yet been implemented")

	})
	api.TasksGetTasksHandler = tasks.GetTasksHandlerFunc(func(params tasks.GetTasksParams) middleware.Responder {
		return middleware.NotImplemented("operation tasks.GetTasks has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
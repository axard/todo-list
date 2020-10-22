// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/axard/todo-list/internal/restapi/operations"
	"github.com/axard/todo-list/internal/restapi/operations/todos"
)

// nolint: lll
//go:generate swagger generate server --target ../../internal --name TodoList --spec ../../api/swagger.yaml --model-package restmodels --principal interface{} --exclude-main

func configureFlags(api *operations.TodoListAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TodoListAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.TodosCreateOneHandler == nil {
		api.TodosCreateOneHandler = todos.CreateOneHandlerFunc(func(params todos.CreateOneParams) middleware.Responder {
			return middleware.NotImplemented("operation todos.CreateOne has not yet been implemented")
		})
	}

	if api.TodosDeleteOneHandler == nil {
		api.TodosDeleteOneHandler = todos.DeleteOneHandlerFunc(func(params todos.DeleteOneParams) middleware.Responder {
			return middleware.NotImplemented("operation todos.DeleteOne has not yet been implemented")
		})
	}

	if api.TodosReadTodosHandler == nil {
		api.TodosReadTodosHandler = todos.ReadTodosHandlerFunc(func(params todos.ReadTodosParams) middleware.Responder {
			return middleware.NotImplemented("operation todos.ReadTodos has not yet been implemented")
		})
	}

	if api.TodosUpdateOneHandler == nil {
		api.TodosUpdateOneHandler = todos.UpdateOneHandlerFunc(func(params todos.UpdateOneParams) middleware.Responder {
			return middleware.NotImplemented("operation todos.UpdateOne has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

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
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// nolint: lll
// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

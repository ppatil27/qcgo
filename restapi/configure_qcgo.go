// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"qcgo/restapi/operations"
	"qcgo/restapi/operations/health_checks"
	"qcgo/restapi/operations/service"
)

//go:generate swagger generate server --target ..\..\qcgo --name Qcgo --spec ..\api\swagger.yaml --principal interface{} --default-scheme https

func configureFlags(api *operations.QcgoAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.QcgoAPI) http.Handler {
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

	if api.ServiceDeleteResourcesByIDHandler == nil {
		api.ServiceDeleteResourcesByIDHandler = service.DeleteResourcesByIDHandlerFunc(func(params service.DeleteResourcesByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation service.DeleteResourcesByID has not yet been implemented")
		})
	}
	if api.HealthChecksHealthCheckHandler == nil {
		api.HealthChecksHealthCheckHandler = health_checks.HealthCheckHandlerFunc(func(params health_checks.HealthCheckParams) middleware.Responder {
			return middleware.NotImplemented("operation health_checks.HealthCheck has not yet been implemented")
		})
	}
	if api.ServiceUpdateResourcesByIDHandler == nil {
		api.ServiceUpdateResourcesByIDHandler = service.UpdateResourcesByIDHandlerFunc(func(params service.UpdateResourcesByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation service.UpdateResourcesByID has not yet been implemented")
		})
	}
	if api.ServiceGetResourcesByIDHandler == nil {
		api.ServiceGetResourcesByIDHandler = service.GetResourcesByIDHandlerFunc(func(params service.GetResourcesByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation service.GetResourcesByID has not yet been implemented")
		})
	}
	if api.ServicePostResourceByAppCodeHandler == nil {
		api.ServicePostResourceByAppCodeHandler = service.PostResourceByAppCodeHandlerFunc(func(params service.PostResourceByAppCodeParams) middleware.Responder {
			return middleware.NotImplemented("operation service.PostResourceByAppCode has not yet been implemented")
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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

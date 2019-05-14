// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"TopNService/restapi/operations"
	"TopNService/restapi/operations/topn_microservice"
)

//go:generate swagger generate server --target ..\..\TopNService --name TopNMicorservice --spec ..\swaggerDefination.yaml

func configureFlags(api *operations.TopNMicorserviceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TopNMicorserviceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.TopnMicroserviceGetV1PullmetricsHandler == nil {
		api.TopnMicroserviceGetV1PullmetricsHandler = topn_microservice.GetV1PullmetricsHandlerFunc(func(params topn_microservice.GetV1PullmetricsParams) middleware.Responder {
			return middleware.NotImplemented("operation topn_microservice.GetV1Pullmetrics has not yet been implemented")
		})
	}
	if api.TopnMicroservicePostV1AnalyticsHandler == nil {
		api.TopnMicroservicePostV1AnalyticsHandler = topn_microservice.PostV1AnalyticsHandlerFunc(func(params topn_microservice.PostV1AnalyticsParams) middleware.Responder {
			return middleware.NotImplemented("operation topn_microservice.PostV1Analytics has not yet been implemented")
		})
	}
	if api.TopnMicroservicePostV1GettopnHandler == nil {
		api.TopnMicroservicePostV1GettopnHandler = topn_microservice.PostV1GettopnHandlerFunc(func(params topn_microservice.PostV1GettopnParams) middleware.Responder {
			return middleware.NotImplemented("operation topn_microservice.PostV1Gettopn has not yet been implemented")
		})
	}

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

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

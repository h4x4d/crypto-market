// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/h4x4d/crypto-market/stack_connector/internal/restapi/operations"
)

//go:generate swagger generate server --target ../../internal --name MarketStackConnector --spec ../../api/swagger/stack_connector.yaml --principal interface{}

func configureFlags(api *operations.MarketStackConnectorAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MarketStackConnectorAPI) http.Handler {
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

	// Applies when the "api_key" header is set
	if api.APIKeyAuth == nil {
		api.APIKeyAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (api_key) api_key from header param [api_key] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.CancelBidHandler == nil {
		api.CancelBidHandler = operations.CancelBidHandlerFunc(func(params operations.CancelBidParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.CancelBid has not yet been implemented")
		})
	}
	if api.CreateBidHandler == nil {
		api.CreateBidHandler = operations.CreateBidHandlerFunc(func(params operations.CreateBidParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateBid has not yet been implemented")
		})
	}
	if api.GetBidByIDHandler == nil {
		api.GetBidByIDHandler = operations.GetBidByIDHandlerFunc(func(params operations.GetBidByIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetBidByID has not yet been implemented")
		})
	}
	if api.GetBidsHandler == nil {
		api.GetBidsHandler = operations.GetBidsHandlerFunc(func(params operations.GetBidsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetBids has not yet been implemented")
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

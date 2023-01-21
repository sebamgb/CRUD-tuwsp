package routes

import (
	"net/http"

	"tuwsp/handlers"
	"tuwsp/middlewares"
	"tuwsp/server"

	"github.com/gorilla/mux"
)

// BindRoute implement Router mux of gorilla for call
func BindRoute(s server.Server, r *mux.Router) {
	r.Path("/login").Handler(handlers.Login(s))
	r.Path("/signup").Handler(handlers.Signup(s))
	// public := r.NewRoute().Subrouter()
	private := r.NewRoute().Subrouter()
	// Insert paths
	// user:
	private.
		Use(middlewares.AuthMiddleware(s))
	private.
		Handle("/auths", handlers.InsertAuthHandler(s)).
		Methods(http.MethodPost)
	private.
		Handle("/users", handlers.InsertUserHandler(s)).
		Methods(http.MethodPost)
	private.
		Handle("/info-users", handlers.InsertInfoUserHandler(s)).
		Methods(http.MethodPost)
	// url:
	private.
		Handle("/protocols", handlers.InsertProtocolHandler(s)).
		Methods(http.MethodPost)
	private.
		Handle("/urls", handlers.InsertUrlHandler(s)).
		Methods(http.MethodPost)
	private.
		Handle("/endpoints", handlers.InsertEndpointHandler(s)).
		Methods(http.MethodPost)
	private.
		Handle("/query-keys", handlers.InsertQueryKeyHandler(s)).
		Methods(http.MethodPost)
	private.
		Handle("/query-values", handlers.InsertQueryValueHandler(s)).
		Methods(http.MethodPost)
	// Get paths
	// user:
	// url:
}

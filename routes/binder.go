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
	public := r.NewRoute().Subrouter()
	private := r.PathPrefix("/api-tuwsp/v1").Subrouter()
	private.
		Use(middlewares.AuthMiddleware(s))
	// entrypoint
	public.
		Path("/login").
		Handler(handlers.LoginHandler(s)).
		Methods(http.MethodPost)
	public.
		Path("/validate").
		Handler(handlers.ValidateHandler(s)).
		Methods(http.MethodPost)
	public.
		Path("/signup").
		Handler(handlers.SignupHandler(s)).
		Methods(http.MethodPost)
	// private gets
	private.
		Handle("/me", handlers.MeHandler(s)).
		Methods(http.MethodGet)
	// Insert paths
	// user:
	public.
		Handle("/users", handlers.InsertUserHandler(s)).
		Methods(http.MethodPost)
	public.
		Handle("/info-users", handlers.InsertInfoUserHandler(s)).
		Methods(http.MethodPost)
	// url:
	public.
		Handle("/protocols", handlers.InsertProtocolHandler(s)).
		Methods(http.MethodPost)
	public.
		Handle("/urls", handlers.InsertUrlHandler(s)).
		Methods(http.MethodPost)
	public.
		Handle("/endpoints", handlers.InsertEndpointHandler(s)).
		Methods(http.MethodPost)
	public.
		Handle("/query-keys", handlers.InsertQueryKeyHandler(s)).
		Methods(http.MethodPost)
	public.
		Handle("/query-values", handlers.InsertQueryValueHandler(s)).
		Methods(http.MethodPost)
	// Get paths
	// user:
	public.
		Handle("/auths-by-id", handlers.GetAuthByIdHandler(s)).
		Methods(http.MethodGet)
	public.
		Handle("/auths-by-email", handlers.GetAuthByEmailHandler(s)).
		Methods(http.MethodGet)
	public.
		Handle("/users-by-nick-name", handlers.GetUserByNickNameHandler(s)).
		Methods(http.MethodGet)
	public.
		Handle("/users-by-id", handlers.GetUserByIdHandler(s)).
		Methods(http.MethodGet)
	public.
		Handle("/info-users-by-id", handlers.GetInfoUserByUserIdHandler(s)).
		Methods(http.MethodGet)
	public.
		Handle("/info-users-phone", handlers.GetInfoUserByPhoneHandler(s)).
		Methods(http.MethodGet)
	// url:
	public.
		Handle("/protocols", handlers.GetProtocolHandler(s)).
		Methods(http.MethodGet)
	public.
		Handle("/urls", handlers.GetUrlHandler(s)).
		Methods(http.MethodGet)
	public.
		Handle("/endpoints", handlers.GetEndpointHandler(s)).
		Methods(http.MethodGet)
	public.
		Handle("/query-keys", handlers.GetQueryKeyHandler(s)).
		Methods(http.MethodGet)
	public.
		Handle("/query-values", handlers.GetQueryValueHandler(s))

	// Update paths
	// user:
	// url:
}

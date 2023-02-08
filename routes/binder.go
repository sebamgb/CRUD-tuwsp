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
	public.
		Path("/form").
		Handler(handlers.GetFormByTitleHandler(s)).
		Methods(http.MethodGet)
	// Insert paths
	// info:
	private.
		Handle("/forms", handlers.InsertFormHandler(s)).
		Methods(http.MethodPost)
	private.
		Handle("/logins", handlers.InsertLoginHandler(s)).
		Methods(http.MethodPost)
	public.
		Handle("/signups", handlers.InsertSignupHandler(s)).
		Methods(http.MethodPost)
	private.
		Handle("/dashboards", handlers.InsertDashboardHandler(s)).
		Methods(http.MethodPost)
	private.
		Handle("/auths", handlers.InsertAuthHandler(s)).
		Methods(http.MethodPost)
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
	// info:
	private.
		Handle("/me", handlers.MeHandler(s)).
		Methods(http.MethodGet)
	private.
		Handle("/login", handlers.GetLoginByAuthIdHandler(s)).
		Methods(http.MethodGet)
	private.
		Handle("/signup", handlers.GetSignupByIdHandler(s)).
		Methods(http.MethodGet)
	private.
		Handle("/dashboard", handlers.GetDashboardByAuthIdHandler(s)).
		Methods(http.MethodGet)
	// user:
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
	// info:
	private.
		Handle("/forms/{id}", handlers.UpdateFormHandler(s)).
		Methods(http.MethodPut)
	private.
		Handle("/logins/{id}", handlers.UpdateLoginHandler(s)).
		Methods(http.MethodPut)
	private.
		Handle("/signups/{id}", handlers.UpdateSignupHandler(s)).
		Methods(http.MethodPut)
	private.
		Handle("/dashboards/{id}", handlers.UpdateDashboardHandler(s)).
		Methods(http.MethodPut)
	private.
		Handle("/auths/{id}", handlers.UpdateAuthHandler(s)).
		Methods(http.MethodPut)
	// user:
	private.
		Handle("/users/{id}", handlers.UpdateUserHandler(s)).
		Methods(http.MethodPut)
	private.
		Handle("/info-users/{id}", handlers.UpdateInfoUserHandler(s)).
		Methods(http.MethodPut)
	// url:
	private.
		Handle("/protocols/{id}", handlers.UpdateProtocolHandler(s)).
		Methods(http.MethodPut)
	private.
		Handle("/urls/{id}", handlers.UpdateUrlHandler(s)).
		Methods(http.MethodPut)
	private.
		Handle("/endpoints/{id}", handlers.UpdateEndpointHandler(s)).
		Methods(http.MethodPut)
	private.
		Handle("/query-keys/{id}", handlers.UpdateQueryKeyHandler(s)).
		Methods(http.MethodPut)
	private.
		Handle("/query-values/{id}", handlers.UpdateQueryValueHandler(s)).
		Methods(http.MethodPut)
	// Delete paths
	// info:
	private.
		Handle("/forms/{id}", handlers.DeleteFormHandler(s)).
		Methods(http.MethodDelete)
	private.
		Handle("/logins/{id}", handlers.DeleteLoginHandler(s)).
		Methods(http.MethodDelete)
	private.
		Handle("/signups/{id}", handlers.DeleteSignupHandler(s)).
		Methods(http.MethodDelete)
	private.
		Handle("/dashboards/{id}", handlers.DeleteDashboardHandler(s)).
		Methods(http.MethodDelete)
	private.
		Handle("/auths/{id}", handlers.DeleteAuthHandler(s)).
		Methods(http.MethodDelete)
	// user:
	private.
		Handle("/users/{id}", handlers.DeleteUserHandler(s)).
		Methods(http.MethodDelete)
	private.
		Handle("/info-users/{id}", handlers.DeleteInfoUserHandler(s)).
		Methods(http.MethodDelete)
	// url:
	private.
		Handle("/protocols/{id}", handlers.DeleteProtocolHandler(s)).
		Methods(http.MethodDelete)
	private.
		Handle("/urls/{id}", handlers.DeleteUrlHandler(s)).
		Methods(http.MethodDelete)
	private.
		Handle("/endpoints/{id}", handlers.DeleteEndpointHandler(s)).
		Methods(http.MethodDelete)
	private.
		Handle("/query-keys/{id}", handlers.DeleteQueryKeyHandler(s)).
		Methods(http.MethodDelete)
	private.
		Handle("/query-values/{id}", handlers.DeleteQueryValueHandler(s)).
		Methods(http.MethodDelete)
	// List paths
	// info:
	private.
		Handle("/forms", handlers.ListFormsHandler(s)).
		Methods(http.MethodGet)
	private.
		Handle("/logins", handlers.ListLoginsHandler(s)).
		Methods(http.MethodGet)
	private.
		Handle("/signups", handlers.ListSignupsHandler(s)).
		Methods(http.MethodGet)
	private.
		Handle("/dashboards", handlers.ListDashboardsHandler(s)).
		Methods(http.MethodGet)
	private.
		Handle("/auths", handlers.ListAuthsHandler(s)).
		Methods(http.MethodGet)
	// user:
	private.
		Handle("/users", handlers.ListUsersHandler(s)).
		Methods(http.MethodGet)
	private.
		Handle("/info-users", handlers.ListInfoUsersHandler(s)).
		Methods(http.MethodGet)
	// url:
	private.
		Handle("/protocols", handlers.ListProtocolsHandler(s)).
		Methods(http.MethodGet)
	private.
		Handle("/urls", handlers.ListURLsHandler(s)).
		Methods(http.MethodGet)
	private.
		Handle("/endpoints", handlers.ListEndpointsHandler(s)).
		Methods(http.MethodGet)
	private.
		Handle("/query-keys", handlers.ListQueryKeysHandler(s)).
		Methods(http.MethodGet)
	private.
		Handle("/query-values", handlers.ListQueryValuesHandler(s)).
		Methods(http.MethodGet)
}

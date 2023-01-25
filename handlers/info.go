package handlers

import (
	"net/http"

	"tuwsp/models"
	"tuwsp/repository"
	"tuwsp/server"
)

/* -----Inserts-----*/

// InsertLoginHandler handle an insert for logins in db
func InsertLoginHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loginRequest := models.Login{}
		decode(r, w, &loginRequest)

	}
}

/* -----Gets-----*/

// GetFormByTitleHandler handle a select for forms in db
func GetFormByTitle(S server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		form, err := repository.
			GetFormByTitle(r.Context(), params.Get("q"))
		internalErr(w, err)
		encode(w, &form)
	}
}

// GetLoginByAuthIdHandler handle a select for logins in db
func GetLoginByAuthIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		login, err := repository.
			GetLoginByAuthId(r.Context(), params.Get("q"))
		internalErr(w, err)
		encode(w, &login)
	}
}

package handlers

import (
	"net/http"
	"time"

	"tuwsp/models"
	"tuwsp/repository"
	"tuwsp/server"

	"github.com/golang-sql/civil"
	"github.com/segmentio/ksuid"
)

/* -----Inserts-----*/

type InsertLoginResponse struct {
	Id string `json:"id"`
}

// InsertLoginHandler handle an insert for logins in db
func InsertLoginHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loginRequest := models.Login{}
		// decode request into body
		decode(r, w, &loginRequest)
		// login created at now
		loginRequest.CreatedAt = civil.DateTimeOf(time.Now())
		// new random ksuid
		id, err := ksuid.NewRandom()
		if err != nil {
			panic(err)
		}
		// login id random
		loginRequest.Id = id.String()
		// insert login
		err = repository.InsertIntoLogins(r.Context(), &loginRequest)
		internalErr(w, err)
		// encoding response
		encode(w, &InsertLoginResponse{Id: loginRequest.Id})
	}
}

type InsertFormResponse struct {
	Id string `json:"id"`
}

// InsertFormHandler handle an insert for forms in db
func InsertFormHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		formRequest := models.Form{}
		// decode request into body
		decode(r, w, &formRequest)
		// new random ksuid
		id, err := ksuid.NewRandom()
		if err != nil {
			panic(err)
		}
		// form id random
		formRequest.Id = id.String()
		// insert form
		err = repository.InsertIntoForms(r.Context(), &formRequest)
		internalErr(w, err)
		// encoding response
		encode(w, &InsertFormResponse{Id: formRequest.Id})
	}
}

type InsertSignupResponse struct {
	Id string `json:"id"`
}

// InsertSignupHandler handle an insert for signups in db
func InsertSignupHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		signupRequest := models.Signup{}
		// decode request into body
		decode(r, w, &signupRequest)
		// new random ksuid
		id, err := ksuid.NewRandom()
		if err != nil {
			panic(err)
		}
		// signup id random
		signupRequest.Id = id.String()
		// insert signup
		err = repository.InsertIntoSignups(r.Context(), &signupRequest)
		internalErr(w, err)
		// encoding response
		encode(w, &InsertSignupResponse{Id: signupRequest.Id})
	}
}

type InsertDashboardResponse struct {
	Id string `json:"id"`
}

// InsertDashboardHandler handle an insert for dashboards in db
func InsertDashboardHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dashboardRequest := models.Dashboard{}
		// decode request into body
		decode(r, w, &dashboardRequest)
		// new random ksuid
		id, err := ksuid.NewRandom()
		if err != nil {
			panic(err)
		}
		// dashboard id random
		dashboardRequest.Id = id.String()
		// insert dashboard
		err = repository.InsertIntoDashboards(r.Context(), &dashboardRequest)
		internalErr(w, err)
		// encoding response
		encode(w, &InsertDashboardResponse{Id: dashboardRequest.Id})
	}
}

type InsertAuthResponse struct {
	Id string `json:"id"`
}

// InsertAuthHandler handle the insert of auths
func InsertAuthHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var authRequest = models.Auth{}
		// decode request into body
		decode(r, w, &authRequest)
		// new random ksuid
		id, err := ksuid.NewRandom()
		if err != nil {
			panic(err)
		}
		// auth id random
		authRequest.Id = id.String()
		// auth created at now
		authRequest.CreatedAt = civil.DateTimeOf(time.Now())
		// auth signup id random
		authRequest.SignupId = id.String()
		// insert auth
		err = repository.InsertIntoAuths(r.Context(), &authRequest)
		internalErr(w, err)
		// encoding response
		encode(w, &InsertAuthResponse{Id: authRequest.Id})
	}
}

/* -----Gets-----*/

// GetFormByTitleHandler handle a select for forms in db
func GetFormByTitleHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting queryparams
		params := r.URL.Query()
		// getting form by title
		form, err := repository.
			GetFormByTitle(r.Context(), params.Get("q"))
		internalErr(w, err)
		// encoding response
		encode(w, &form)
	}
}

// GetLoginByAuthIdHandler handle a select for logins in db
func GetLoginByAuthIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting queryparams
		params := r.URL.Query()
		// getting login by auth id
		login, err := repository.
			GetLoginByAuthId(r.Context(), params.Get("q"))
		internalErr(w, err)
		// encoding response
		encode(w, &login)
	}
}

// GetSignupByIdHandler handle a select for signups in db
func GetSignupByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting queryparams
		params := r.URL.Query()
		// getting signup by id
		signup, err := repository.
			GetSignupById(r.Context(), params.Get("q"))
		internalErr(w, err)
		// encoding response
		encode(w, &signup)
	}
}

// GetDashboardByAuthIdHandler handle a select for dashboards in db
func GetDashboardByAuthIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting queryparams
		params := r.URL.Query()
		// getting dashboard by auth id
		dashboard, err := repository.
			GetDashboardByAuthId(r.Context(), params.Get("q"))
		internalErr(w, err)
		// encoding response
		encode(w, &dashboard)
	}
}

// GetAuthByIdHandler handle the select for auths
func GetAuthByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting queryparams
		params := r.URL.Query()
		// getting auth by id
		auth, err := repository.
			GetAuthById(r.Context(), params.Get("q"))
		internalErr(w, err)
		// encoding response
		encode(w, &auth)
	}
}

// GetAuthByEmailHandler handle the select for auths
func GetAuthByEmailHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting queryparams
		params := r.URL.Query()
		// getting auth by email
		auth, err := repository.
			GetAuthByEmail(r.Context(), params.Get("q"))
		internalErr(w, err)
		// encoding response
		encode(w, &auth)
	}
}

/* -----Updates-----*/

type UpdateLoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// UpdateLoginHandler handle an update for logins in db
func UpdateLoginHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loginRequest := models.Login{}
		// decode request into body
		decode(r, w, &loginRequest)
		// updating login
		err := repository.UpdateLogins(r.Context(), &loginRequest)
		internalErr(w, err)
		// encoding response
		encode(w, &UpdateLoginResponse{Success: true, Message: "Login updated successfully"})
	}
}

type UpdateFormResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// UpdateFormHandler handle an update for forms in db
func UpdateFormHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		formRequest := models.Form{}
		// decode request into body
		decode(r, w, &formRequest)
		// updating form
		err := repository.UpdateForms(r.Context(), &formRequest)
		internalErr(w, err)
		// encoding response
		encode(w, &UpdateFormResponse{Success: true, Message: "Form updated successfully"})
	}
}

type UpdateSignupResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// UpdateSignupHandler handle an update for signups in db
func UpdateSignupHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		signupRequest := models.Signup{}
		// decode request into body
		decode(r, w, &signupRequest)
		// updating signup
		err := repository.UpdateSignups(r.Context(), &signupRequest)
		internalErr(w, err)
		// encoding response
		encode(w, &UpdateSignupResponse{Success: true, Message: "Signup updated successfully"})
	}
}

type UpdateDashboardResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// UpdateDashboardHandler handle an update for dashboards in db
func UpdateDashboardHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dashboardRequest := models.Dashboard{}
		// decode request into body
		decode(r, w, &dashboardRequest)
		// updating dashboard
		err := repository.UpdateDashboards(r.Context(), &dashboardRequest)
		internalErr(w, err)
		// encoding response
		encode(w, &UpdateDashboardResponse{Success: true, Message: "Dashboard updated successfully"})
	}
}

type UpdateAuthResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// UpdateAuthHandler handle the update of auths
func UpdateAuthHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authRequest := models.Auth{}
		// decode request into body
		decode(r, w, &authRequest)
		// updating auth
		err := repository.UpdateAuths(r.Context(), &authRequest)
		internalErr(w, err)
		// encoding response
		encode(w, &UpdateAuthResponse{Success: true, Message: "Auth updated successfully"})
	}
}

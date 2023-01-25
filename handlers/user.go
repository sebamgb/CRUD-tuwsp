package handlers

import (
	"net/http"
	"time"

	"tuwsp/models"
	"tuwsp/repository"
	"tuwsp/server"

	"github.com/golang-sql/civil"
	"github.com/gorilla/mux"
)

func Login(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func Signup(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

/*----- Inserts -----*/

// InsertAuthHandler handle the insert of auths
func InsertAuthHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var authRequest = models.Auth{}
		decode(r, w, &authRequest)
		// auth created now
		authRequest.CratedAt = civil.DateTimeOf(time.Now())
		internalErr(w, repository.
			InsertIntoAuths(r.Context(), &authRequest))
		encode(w, &authRequest)
	}
}

// InsertUserHandler handle the insert of users
func InsertUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRequest := models.User{}
		decode(r, w, &userRequest)
		internalErr(w, repository.
			InsertIntoUsers(r.Context(), &userRequest))
		encode(w, &userRequest)
	}
}

// InsertInfoUserHandler handle the insert of info_users
func InsertInfoUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		infoUserRequest := models.InfoUser{}
		decode(r, w, &infoUserRequest)
		internalErr(w, repository.
			InsertIntoInfoUsers(r.Context(), &infoUserRequest))
		encode(w, &infoUserRequest)
	}
}

/*----- Gets -----*/

// GetAuthHandler handle the select for auths
func GetAuthHandler(S server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		auth, err := repository.
			GetAuthById(r.Context(), params["q"])
		if err != nil {
			return
		}
		encode(w, &auth)
	}
}

// GetUserHandler handle the select for users
func GetUserHandler(S server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		user, err := repository.
			GetUserByNickName(r.Context(), params["q"])
		if err != nil {
			return
		}
		encode(w, &user)
	}
}

// GetInfoUserHandler handle the select for auths
func GetInfoUserHandler(S server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		infouser, err := repository.
			GetInfoUserByUserId(r.Context(), params["q"])
		if err != nil {
			return
		}
		encode(w, &infouser)
	}
}

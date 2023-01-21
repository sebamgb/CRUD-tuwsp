package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"tuwsp/models"
	"tuwsp/repository"
	"tuwsp/server"

	"github.com/golang-sql/civil"
)

// decode translate PostForm to json and unmarshal to interface
func decode(r *http.Request, w http.ResponseWriter, a any) {
	r.ParseForm()
	fmt.Printf("r.PostForm: %v\n", r.PostForm)
	if bytes, err := json.Marshal(r.PostForm); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err = json.Unmarshal(bytes, a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// internalErr handle error with internalServerError of http
func internalErr(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Login(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func Signup(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// InsertAuthHandler handle the insert of auths
func InsertAuthHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var authRequest = models.Auth{}
		decode(r, w, &authRequest)
		authRequest.CratedAt = civil.DateTimeOf(time.Now())
		internalErr(w, repository.
			InsertIntoAuths(r.Context(), &authRequest))
	}
}

// InsertUserHandler handle the insert of users
func InsertUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRequest := models.User{}
		decode(r, w, &userRequest)
		internalErr(w, repository.
			InsertIntoUsers(r.Context(), &userRequest))
	}
}

// InsertInfoUserHandler handle the insert of info_users
func InsertInfoUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		infoUserRequest := models.InfoUser{}
		decode(r, w, &infoUserRequest)
		infoUserRequest.Birthday = civil.DateOf(time.Now())
		internalErr(w, repository.
			InsertIntoInfoUsers(r.Context(), &infoUserRequest))
	}
}

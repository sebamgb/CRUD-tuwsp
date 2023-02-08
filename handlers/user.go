package handlers

import (
	"net/http"
	"strconv"

	"tuwsp/models"
	"tuwsp/repository"
	"tuwsp/server"

	"github.com/gorilla/mux"
)

type ValidateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// ValidateHandler Compare request with db data
func ValidateHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data = models.Login{}
		// decoding request into body
		decode(r, w, &data)
		// Valida los datos contra la base de datos
		success, message := validateAgainstDB(data.Email, data.Password, r.Context())
		// encoding response
		encode(w, &ValidateResponse{Success: success, Message: message})
	}
}

/*----- Inserts -----*/

type InsertUserResponse struct {
	Id string `json:"id"`
}

// InsertUserHandler handle the insert of users
func InsertUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRequest := models.User{}
		// decoding request into body
		decode(r, w, &userRequest)
		// inserting user
		err := repository.InsertIntoUsers(r.Context(), &userRequest)
		internalErr(w, err)
		// encoding response
		encode(w, &InsertUserResponse{Id: userRequest.Id})
	}
}

type InsertInfoUserResponse struct {
	Id string `json:"id"`
}

// InsertInfoUserHandler handle the insert of info_users
func InsertInfoUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		infoUserRequest := models.InfoUser{}
		// decoding request into body
		decode(r, w, &infoUserRequest)
		// inserting info_user
		err := repository.InsertIntoInfoUsers(r.Context(), &infoUserRequest)
		internalErr(w, err)
		// encoding response
		encode(w, &InsertInfoUserResponse{Id: infoUserRequest.Id})
	}
}

/*----- Gets -----*/

// GetUserByNickNameHandler handle the select for users
func GetUserByNickNameHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting queryparams
		params := r.URL.Query()
		// getting user
		user, err := repository.
			GetUserByNickName(r.Context(), params.Get("q"))
		internalErr(w, err)
		// encoding response
		encode(w, &user)
	}
}

// GetUserByIdHandler handle the select for users
func GetUserByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting queryparams
		params := r.URL.Query()
		// getting user
		user, err := repository.
			GetUserById(r.Context(), params.Get("q"))
		internalErr(w, err)
		// encoding response
		encode(w, &user)
	}
}

// GetInfoUserByUserIdHandler handle the select for info_users
func GetInfoUserByUserIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting queryparams
		params := r.URL.Query()
		// getting info_user
		infouser, err := repository.
			GetInfoUserByUserId(r.Context(), params.Get("q"))
		internalErr(w, err)
		// encoding response
		encode(w, &infouser)
	}
}

// GetInfoUserByPhoneHandler handle the select for info_users
func GetInfoUserByPhoneHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// trying to convert to int by queryparams
		phone, err := strconv.Atoi(r.URL.Query().Get("q"))
		internalErr(w, err)
		// getting info_user
		infouser, err := repository.
			GetInfoUserByPhone(r.Context(), phone)
		internalErr(w, err)
		// encoding response
		encode(w, &infouser)
	}
}

/*----- Updates -----*/

type UpdateUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// UpdateUserHandler handle the update of users
func UpdateUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting mux vars
		vars := mux.Vars(r)
		// getting id from mux vars
		id := vars["id"]
		// getting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			userRequest := models.User{
				Id: id,
			}
			// decoding request into body
			decode(r, w, &userRequest)
			// updating user
			err := repository.UpdateUsers(r.Context(), &userRequest)
			internalErr(w, err)
			// encoding response
			encode(w, &UpdateUserResponse{Success: true, Message: "User updated successfully", Author: claims.AuthId})
		})
	}
}

type UpdateInfoUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// UpdateInfoUserHandler handle the update of info_users
func UpdateInfoUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting mux vars
		vars := mux.Vars(r)
		// getting id from mux vars
		id := vars["id"]
		// getting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			infoUserRequest := models.InfoUser{
				Id: id,
			}
			// decoding request into body
			decode(r, w, &infoUserRequest)
			// updating info_user
			err := repository.UpdateInfoUsers(r.Context(), &infoUserRequest)
			internalErr(w, err)
			// encoding response
			encode(w, &UpdateInfoUserResponse{Success: true, Message: "InfoUser updated successfully", Author: claims.AuthId})
		})
	}
}

/*----- Deletes -----*/

type DeleteUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// DeleteUserHandler handle the delete of users
func DeleteUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting mux vars
		vars := mux.Vars(r)
		// getting id from mux vars
		id := vars["id"]
		// getting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			userRequest := models.User{
				Id: id,
			}
			// deleting user
			err := repository.DeleteUsers(r.Context(), &userRequest)
			internalErr(w, err)
			// encoding response
			encode(w, &DeleteUserResponse{Success: true, Message: "User deleted successfully", Author: claims.AuthId})
		})
	}
}

type DeleteInfoUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// DeleteInfoUserHandler handle the delete of info_users
func DeleteInfoUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting mux vars
		vars := mux.Vars(r)
		// getting id from mux vars
		id := vars["id"]
		// getting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			infoUserRequest := models.InfoUser{
				Id: id,
			}
			// deleting info_user
			err := repository.DeleteInfoUsers(r.Context(), &infoUserRequest)
			internalErr(w, err)
			// encoding response
			encode(w, &DeleteInfoUserResponse{Success: true, Message: "InfoUser deleted successfully", Author: claims.AuthId})
		})
	}
}

/* ----- Lists ----- */

// ListUsersHandler handle the sellect all of users
func ListUsersHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			// list users
			users, err := repository.ListUsers(r.Context())
			internalErr(w, err)
			// encoding response
			encode(w, &users)
		})
	}
}

// ListInfoUsersHandler handle the sellect all of info_users
func ListInfoUsersHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			// list info_users
			infoUsers, err := repository.ListInfoUsers(r.Context())
			internalErr(w, err)
			// encoding response
			encode(w, &infoUsers)
		})
	}
}

package handlers

import (
	"net/http"
	"strconv"
	"time"

	"tuwsp/models"
	"tuwsp/repository"
	"tuwsp/server"

	"github.com/golang-jwt/jwt"
	"github.com/golang-sql/civil"
	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

const HASHCOST = 13

type SignupLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Success bool   `json:"success"`
}

type SignupResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

// LoginHandler handle logins from auth
func LoginHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SignupLoginRequest{}
		// decode request into body
		decode(r, w, &request)
		// get auth
		user, err := repository.GetAuthByEmail(r.Context(), request.Email)
		internalErr(w, err)
		// unathorized requests
		if user == nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		// compare password hashed
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		// making claim for token
		claim := models.AppClaims{
			AuthId: user.Id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(2 * time.Hour * 24).Unix(),
			},
		}
		// confection token with claims previosly do
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		tokenString, err := token.SignedString([]byte(s.Config().JWTSecret))
		internalErr(w, err)
		// encoding response
		encode(w, &LoginResponse{Token: tokenString, Success: true})
	}
}

type ValidateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// ValidateHandler Compare request with db data
func ValidateHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data = SignupLoginRequest{}
		// decoding request into body
		decode(r, w, &data)
		// Valida los datos contra la base de datos
		success, message := validateAgainstDB(data.Email, data.Password, r.Context())
		if !success {
			w.WriteHeader(http.StatusBadRequest)
		}
		// encoding response
		encode(w, &ValidateResponse{Success: success, Message: message})
	}
}

// SignupHandler handle register of users
func SignupHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SignupLoginRequest{}
		// decoding request into body
		decode(r, w, &request)
		// hasshing password request
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), HASHCOST)
		internalErr(w, err)
		// random id
		id, err := ksuid.NewRandom()
		internalErr(w, err)
		// making user to db
		var user = models.Auth{
			Email:     request.Email,
			CreatedAt: civil.DateTimeOf(time.Now()),
			Password:  string(hashedPassword),
			Id:        id.String(),
		}
		// insert of user
		err = repository.InsertIntoAuths(r.Context(), &user)
		internalErr(w, err)
		// encoding response
		encode(w, &SignupResponse{Id: user.Id, Email: user.Email})
	}
}

// MeHandler for get the user auth from db locking for the token
func MeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			// getting auth from db
			auth, err := repository.GetAuthById(r.Context(), claims.AuthId)
			internalErr(w, err)
			// encoding response
			encode(w, auth)
		})
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

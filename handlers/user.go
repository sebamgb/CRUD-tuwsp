package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"tuwsp/models"
	"tuwsp/repository"
	"tuwsp/server"

	"github.com/golang-jwt/jwt"
	"github.com/golang-sql/civil"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

const HASHCOST = 13

type SignupLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type SignupResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

// LoginHandler handle logins from auth
func LoginHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SignupLoginRequest{}
		// info body decode
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
		// making body response
		encode(w, &LoginResponse{
			Token: tokenString,
		})
	}
}

// ValidateHandler Compare request with db data
func ValidateHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data = SignupLoginRequest{}

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		// Valida los datos contra la base de datos
		success, message := validateAgainstDB(data.Email, data.Password, r.Context())
		if !success {
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(&struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}{Success: success, Message: message})
	}
}

// SignupHandler handle register of users
func SignupHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SignupLoginRequest{}
		// info body decode
		decode(r, w, &request)
		// hasshing password request
		hashedPassword, err := bcrypt.
			GenerateFromPassword([]byte(request.Password), HASHCOST)
		internalErr(w, err)
		// random id
		id, err := ksuid.NewRandom()
		internalErr(w, err)
		// making user to db
		var user = models.Auth{
			Email:    request.Email,
			CratedAt: civil.DateTimeOf(time.Now()),
			Password: string(hashedPassword),
			Id:       id.String(),
		}
		// insert of user
		err = repository.InsertIntoAuths(r.Context(), &user)
		internalErr(w, err)
		// making body response
		encode(w, &SignupResponse{
			Id:    user.Id,
			Email: user.Email,
		})
	}
}

// MeHandler for get the user auth from db locking for the token
func MeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// receiving token in format string
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		//  Parsing token string
		token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})
		unathorizedError(w, err)
		// validating token
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			user, err := repository.GetAuthById(r.Context(), claims.AuthId)
			internalErr(w, err)
			encode(w, &user)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

/*----- Inserts -----*/

// InsertAuthHandler handle the insert of auths
// func InsertAuthHandler(s server.Server) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var authRequest = models.Auth{}
// 		decode(r, w, &authRequest)
// 		// auth created now
// 		authRequest.CratedAt = civil.DateTimeOf(time.Now())
// 		internalErr(w, repository.
// 			InsertIntoAuths(r.Context(), &authRequest))
// 		encode(w, &authRequest)
// 	}
// }

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

// GetAuthByIdHandler handle the select for auths
func GetAuthByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		auth, err := repository.
			GetAuthById(r.Context(), params.Get("q"))
		internalErr(w, err)
		encode(w, &auth)
	}
}

// GetAuthByEmailHandler handle the select for auths
func GetAuthByEmailHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		auth, err := repository.
			GetAuthByEmail(r.Context(), params.Get("q"))
		internalErr(w, err)
		encode(w, &auth)
	}
}

// GetUserByNickNameHandler handle the select for users
func GetUserByNickNameHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		user, err := repository.
			GetUserByNickName(r.Context(), params.Get("q"))
		internalErr(w, err)
		encode(w, &user)
	}
}

// GetUserByIdHandler handle the select for users
func GetUserByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		user, err := repository.
			GetUserById(r.Context(), params.Get("q"))
		internalErr(w, err)
		encode(w, &user)
	}
}

// GetInfoUserByUserIdHandler handle the select for info_users
func GetInfoUserByUserIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		infouser, err := repository.
			GetInfoUserByUserId(r.Context(), params.Get("q"))
		internalErr(w, err)
		encode(w, &infouser)
	}
}

// GetInfoUserByPhoneHandler handle the select for info_users
func GetInfoUserByPhoneHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var infouser *models.InfoUser
		phone, err := strconv.Atoi(r.URL.Query().Get("q"))
		if err != nil {
			return
		} else {
			infouser, err = repository.
				GetInfoUserByPhone(r.Context(), phone)
			internalErr(w, err)
		}
		encode(w, &infouser)
	}
}

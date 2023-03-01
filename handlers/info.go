package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
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

/* ----- Inserts -----*/

type InsertLoginResponse struct {
	Id      string `json:"id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// InsertLoginHandler handle an insert for logins in db
func InsertLoginHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "tuwsper", func(claims *models.AppClaims) {
			loginRequest := models.Login{
				AuthId: claims.AuthId,
			}
			// decode request into body
			if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
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
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			// encoding response
			encode(w, &InsertLoginResponse{Id: loginRequest.Id, Success: true, Message: "Login created successfully", Author: claims.AuthId})
		})
	}
}

type KeyValueResponse struct {
	Id      string `json:"id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// InsertKeyValueLabelshandler handle an insert for key_value_labels in db
func InsertKeyValueLabelsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keyValueRequest := models.KeyValue{}
		// decode request into body
		if err := json.NewDecoder(r.Body).Decode(&keyValueRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// new random ksuid
		id, err := ksuid.NewRandom()
		if err != nil {
			panic(err)
		}
		// key_value id random
		keyValueRequest.Id = id.String()
		// insert key_value
		err = repository.InsertIntoKeyValueLabels(r.Context(), &keyValueRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		// encoding response
		encode(w, &KeyValueResponse{Id: keyValueRequest.Id, Success: true, Message: "KeyValue created successfully", Author: keyValueRequest.Author})
	}
}

// InsertKeyValueplaceholdershandler handle an insert for key_value_placeholders in db
func InsertKeyValuePlaceholdersHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keyValueRequest := models.KeyValue{}
		// decode request into body
		if err := json.NewDecoder(r.Body).Decode(&keyValueRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// new random ksuid
		id, err := ksuid.NewRandom()
		if err != nil {
			panic(err)
		}
		// key_value id random
		keyValueRequest.Id = id.String()
		// insert key_value
		err = repository.InsertIntoKeyValuePlaceholders(r.Context(), &keyValueRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		// encoding response
		encode(w, &KeyValueResponse{Id: keyValueRequest.Id, Success: true, Message: "KeyValue created successfully", Author: keyValueRequest.Author})
	}
}

type InsertFormResponse struct {
	Id      string `json:"id"`
	Sucess  bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// InsertFormHandler handle an insert for forms in db
func InsertFormHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		formRequest := models.Form{}
		// decode request into body
		if err := json.NewDecoder(r.Body).Decode(&formRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// new random ksuid
		id, err := ksuid.NewRandom()
		if err != nil {
			panic(err)
		}
		// form id random
		formRequest.Id = id.String()
		// insert form
		err = repository.InsertIntoForms(r.Context(), &formRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		// encoding response
		encode(w, &InsertFormResponse{Id: formRequest.Id, Sucess: true, Message: "Form created successfully", Author: formRequest.Author})
	}
}

type InsertSignupResponse struct {
	Id      string `json:"id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// InsertSignupHandler handle an insert for signups in db
func InsertSignupHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "admin", func(claims *models.AppClaims) {
			signupRequest := models.Signup{}
			// decode request into body
			if err := json.NewDecoder(r.Body).Decode(&signupRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// new random ksuid
			id, err := ksuid.NewRandom()
			if err != nil {
				panic(err)
			}
			// signup id random
			signupRequest.Id = id.String()
			// insert signup
			err = repository.InsertIntoSignups(r.Context(), &signupRequest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			// encoding response
			encode(w, &InsertSignupResponse{Id: signupRequest.Id, Success: true, Message: "Signup created successfully", Author: claims.AuthId})
		})
	}
}

const HASHCOST = 13

type SignupResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

// SignupHandler handle regists an user in db
func SignupHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = models.Signup{}
		// decoding request into body
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// hasshing password request
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), HASHCOST)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// random id
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// translate body response to InsertSignupResponse
		var signupResponse = InsertSignupResponse{}
		// making user to db
		var user = models.Auth{
			SignupId:  signupResponse.Id,
			Email:     request.Email,
			CreatedAt: civil.DateTimeOf(time.Now()),
			Password:  string(hashedPassword),
			Id:        id.String(),
		}
		// insert of user
		err = repository.InsertIntoAuths(r.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		// encoding response
		encode(w, &SignupResponse{Id: user.Id, Email: user.Email})
	}
}

type InsertDashboardResponse struct {
	Id      string `json:"id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// InsertDashboardHandler handle an insert for dashboards in db
func InsertDashboardHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "tuwsper", func(claims *models.AppClaims) {
			dashboardRequest := models.Dashboard{
				Owner: claims.AuthId,
			}
			// decode request into body
			if err := json.NewDecoder(r.Body).Decode(&dashboardRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// new random ksuid
			id, err := ksuid.NewRandom()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// dashboard id random
			dashboardRequest.Id = id.String()
			// insert dashboard
			err = repository.InsertIntoDashboards(r.Context(), &dashboardRequest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			// encoding response
			encode(w, &InsertDashboardResponse{Id: dashboardRequest.Id, Success: true, Message: "Dashboard created successfully", Author: claims.AuthId})
		})
	}
}

type InsertAuthResponse struct {
	Id      string `json:"id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// InsertAuthHandler handle the insert of auths
func InsertAuthHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "admin", func(claims *models.AppClaims) {
			authRequest := models.Auth{}
			// decode request into body
			if err := json.NewDecoder(r.Body).Decode(&authRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// new random ksuid
			id, err := ksuid.NewRandom()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// hashing password
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authRequest.Password), HASHCOST)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// auth password hashed
			authRequest.Password = string(hashedPassword)
			// auth id random
			authRequest.Id = id.String()
			// auth created at now
			authRequest.CreatedAt = civil.DateTimeOf(time.Now())
			// signup id as claim id auth
			authRequest.SignupId = claims.AuthId
			// insert auth
			err = repository.InsertIntoAuths(r.Context(), &authRequest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			// encoding response
			encode(w, &InsertAuthResponse{Id: authRequest.Id, Success: true, Message: "Auth created successfully", Author: claims.AuthId})
		})
	}
}

/* ----- Gets -----*/

type ValidateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ValidateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// ValidateHandler Compare request with db data
func ValidateHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data = ValidateRequest{}
		// decoding request into body
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Validating form
		success, message := validateAgainstDB(data.Email, data.Password, r.Context())
		// encoding response
		encode(w, &ValidateResponse{Success: success, Message: message})
	}
}

// GetKeyValueLabelsByIdHandler handle a select for key_value_labels in db
func GetKeyValueLabelsByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting queryparams
		params := r.URL.Query()
		// getting keyvalue by id
		keyvalue, err := repository.
			GetKeyValueLabelsById(r.Context(), params.Get("q"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		// encoding response
		encode(w, &keyvalue)
	}
}

// GetKeyValuePlaceholdersByLabelIdHandler handle a select for key_value_placeholders in db
func GetKeyValuePlaceholdersByLabelIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting queryparams
		params := r.URL.Query()
		// getting keyvalue by label id
		keyvalue, err := repository.
			GetKeyValuePlaceholdersByLabelId(r.Context(), params.Get("q"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		// encoding response
		encode(w, &keyvalue)
	}
}

// GetLoginByAuthIdHandler handle a select for logins in db
func GetLoginByAuthIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "tuwsper", func(claims *models.AppClaims) {
			// getting login by auth id
			login, err := repository.
				GetLoginByAuthId(r.Context(), claims.AuthId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)
			// encoding response
			encode(w, &login)
		})
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Success bool   `json:"success"`
}

// LoginHandler handle logins from auth
func LoginHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = LoginRequest{}
		// decode request into body
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// get auth
		user, err := repository.GetAuthByEmail(r.Context(), request.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
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
			AuthId:   user.Id,
			SignupId: user.SignupId,
			Role:     user.Role,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(2 * time.Hour * 24).Unix(),
			},
		}
		// confection token with claims previosly do
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		tokenString, err := token.SignedString([]byte(s.Config().JWTSecret))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		// encoding response
		encode(w, &LoginResponse{Token: tokenString, Success: true})
	}
}

// GetSignupByIdHandler handle a select for signups in db
func GetSignupByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "tuwsper", func(claims *models.AppClaims) {
			// getting signup by id
			signup, err := repository.
				GetSignupById(r.Context(), claims.SignupId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			// encoding response
			encode(w, &signup)
		})
	}
}

// GetDashboardByAuthIdHandler handle a select for dashboards in db
func GetDashboardByAuthIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "tuwsper", func(claims *models.AppClaims) {
			// getting dashboard by auth id
			dashboard, err := repository.
				GetDashboardByAuthId(r.Context(), claims.AuthId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			// encoding response
			encode(w, &dashboard)
		})
	}
}

// MeHandler for get the user auth from db locking for the token
func MeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validating token
		validateTokenAndRole(w, token, "tuwsper", func(claims *models.AppClaims) {
			// getting auth from db
			auth, err := repository.GetAuthById(r.Context(), claims.AuthId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			// encoding response
			encode(w, auth)
		})
	}
}

// GetFormByIdHandler handle a select for forms in db
func GetFormByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting queryparams
		params := r.URL.Query()
		// getting form by id
		form, err := repository.GetFormById(r.Context(), params.Get("q"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		// encoding response
		encode(w, &form)
	}
}

/* ----- Updates ----- */

type UpdateLoginInLogInResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// UpdateLoginInLogInHandler handle an update for logins in db
func UpdateLoginInLogInHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			// getting mux vars
			vars := mux.Vars(r)
			// getting id from vars
			id, ok := vars["id"]
			if !ok {
				http.Error(w, "id is required", http.StatusBadRequest)
				return
			}
			// making request
			loginRequest := models.Login{
				Id: id,
			}
			// decode request into body
			if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// updating login
			err = repository.UpdateLoginsInLogIn(r.Context(), &loginRequest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			// encoding response
			encode(w, &UpdateLoginInLogInResponse{Success: true, Message: fmt.Sprintf("%s successfully logged in", claims.AuthId)})
		} else {
			http.Error(w, "invalid token", http.StatusUnauthorized)
		}
	}
}

type UpdateLoginInLogOutResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// UpdateLoginInLogOutHandler handle an update for logins in db
func UpdateLoginInLogOutHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			// getting mux vars
			vars := mux.Vars(r)
			// getting id from vars
			id, ok := vars["id"]
			if !ok {
				http.Error(w, "id is required", http.StatusBadRequest)
				return
			}
			// making request
			loginRequest := models.Login{
				Id: id,
			}
			// decode request into body
			if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// updating login
			err = repository.UpdateLoginsInLogOut(r.Context(), &loginRequest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			// encoding response
			encode(w, &UpdateLoginInLogOutResponse{Success: true, Message: fmt.Sprintf("%s successfully logged out", claims.AuthId)})
		} else {
			http.Error(w, "invalid token", http.StatusUnauthorized)
		}
	}
}

type UpdateFormResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// UpdateFormHandler handle an update for forms in db
func UpdateFormHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "admin", func(claims *models.AppClaims) {
			// getting mux vars
			vars := mux.Vars(r)
			// getting id from vars
			id, ok := vars["id"]
			if !ok {
				http.Error(w, "id is required", http.StatusBadRequest)
				return
			}
			// making request
			formRequest := models.Form{
				Id: id,
			}
			// decode request into body
			if err := json.NewDecoder(r.Body).Decode(&formRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// updating form
			err := repository.UpdateForms(r.Context(), &formRequest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			// encoding response
			encode(w, &UpdateFormResponse{Success: true, Message: "Form updated successfully"})
		})
	}
}

type UpdateSignupResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// UpdateSignupHandler handle an update for signups in db
func UpdateSignupHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "tuwsper", func(claims *models.AppClaims) {
			// getting mux vars
			vars := mux.Vars(r)
			// getting id from vars
			id, ok := vars["id"]
			if !ok {
				http.Error(w, "id is required", http.StatusBadRequest)
				return
			}
			// making request
			signupRequest := models.Signup{
				Id: id,
			}
			// decode request into body
			if err := json.NewDecoder(r.Body).Decode(&signupRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// updating signup
			err := repository.UpdateSignups(r.Context(), &signupRequest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			// encoding response
			encode(w, &UpdateSignupResponse{Success: true, Message: "Signup updated successfully"})
		})
	}
}

type UpdateDashboardResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// UpdateDashboardHandler handle an update for dashboards in db
func UpdateDashboardHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "tuwsper", func(claims *models.AppClaims) {
			// getting mux vars
			vars := mux.Vars(r)
			// getting id from vars
			id, ok := vars["id"]
			if !ok {
				http.Error(w, "id is required", http.StatusBadRequest)
				return
			}
			// making request
			dashboardRequest := models.Dashboard{
				Id: id,
			}
			// decode request into body
			if err := json.NewDecoder(r.Body).Decode(&dashboardRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// updating dashboard
			err := repository.UpdateDashboards(r.Context(), &dashboardRequest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			// encoding response
			encode(w, &UpdateDashboardResponse{Success: true, Message: "Dashboard updated successfully"})
		})
	}
}

type UpdateAuthResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// UpdateAuthHandler handle the update of auths
func UpdateAuthHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "tuwsper", func(claims *models.AppClaims) {
			// getting mux vars
			vars := mux.Vars(r)
			// getting id from vars
			id, ok := vars["id"]
			if !ok {
				http.Error(w, "id is required", http.StatusBadRequest)
				return
			}
			// making request
			authRequest := models.Auth{
				Id: id,
			}
			// decode request into body
			if err := json.NewDecoder(r.Body).Decode(&authRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// updating auth
			err := repository.UpdateAuths(r.Context(), &authRequest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			// encoding response
			encode(w, &UpdateAuthResponse{Success: true, Message: "Auth updated successfully"})
		})
	}
}

/* ----- Deletes ----- */

type DeleteLoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// DeleteLoginHandler handle the delete of logins
func DeleteLoginHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "admin", func(claims *models.AppClaims) {
			// getting mux vars
			vars := mux.Vars(r)
			// getting id from vars
			id, ok := vars["id"]
			if !ok {
				http.Error(w, "id is required", http.StatusBadRequest)
				return
			}
			// deleting login
			err := repository.DeleteLogins(r.Context(), id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			// encoding response
			encode(w, &DeleteLoginResponse{Success: true, Message: "Login deleted successfully", Author: claims.AuthId})
		})
	}
}

type DeleteFormResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// DeleteFormHandler handle the delete of forms
func DeleteFormHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "admin", func(claims *models.AppClaims) {
			// getting mux vars
			vars := mux.Vars(r)
			// getting id from vars
			id, ok := vars["id"]
			if !ok {
				http.Error(w, "id is required", http.StatusBadRequest)
				return
			}
			// deleting form
			err := repository.DeleteForms(r.Context(), id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			// encoding response
			encode(w, &DeleteFormResponse{Success: true, Message: "Form deleted successfully", Author: claims.AuthId})
		})
	}
}

type DeleteSignupResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// DeleteSignupHandler handle the delete of signups
func DeleteSignupHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "admin", func(claims *models.AppClaims) {
			// getting mux vars
			vars := mux.Vars(r)
			// getting id from vars
			id, ok := vars["id"]
			if !ok {
				http.Error(w, "id is required", http.StatusBadRequest)
				return
			}
			// deleting signup
			err := repository.DeleteSignups(r.Context(), id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			// encoding response
			encode(w, &DeleteSignupResponse{Success: true, Message: "Signup deleted successfully", Author: claims.AuthId})
		})
	}
}

type DeleteDashboardResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// DeleteDashboardHandler handle the delete of dashboards
func DeleteDashboardHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "tuwsper", func(claims *models.AppClaims) {
			// getting mux vars
			vars := mux.Vars(r)
			// getting id from vars
			id, ok := vars["id"]
			if !ok {
				http.Error(w, "id is required", http.StatusBadRequest)
				return
			}
			// deleting dashboard
			err := repository.DeleteDashboards(r.Context(), id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			// encoding response
			encode(w, &DeleteDashboardResponse{Success: true, Message: "Dashboard deleted successfully", Author: claims.AuthId})
		})
	}
}

type DeleteAuthResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// DeleteAuthHandler handle the delete of auths
func DeleteAuthHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "tuwsper", func(claims *models.AppClaims) {
			// getting mux vars
			vars := mux.Vars(r)
			// getting id from vars
			id, ok := vars["id"]
			if !ok {
				http.Error(w, "id is required", http.StatusBadRequest)
				return
			}
			// deleting auth
			err := repository.DeleteAuths(r.Context(), id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			// encoding response
			encode(w, &DeleteAuthResponse{Success: true, Message: "Auth deleted successfully", Author: claims.AuthId})
		})
	}
}

/* ----- Lists ----- */

// ListFormsHandler handle the sellect all of forms
func ListFormsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "admin", func(claims *models.AppClaims) {
			// listing forms
			forms, err := repository.ListForms(r.Context())
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)
			// encoding response
			encode(w, &forms)
		})
	}
}

// ListLoginsHandler handle the sellect all of logins
func ListLoginsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "admin", func(claims *models.AppClaims) {
			// listing logins
			logins, err := repository.ListLogins(r.Context())
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)
			// encoding response
			encode(w, &logins)
		})
	}
}

// ListSignupsHandler handle the sellect all of signups
func ListSignupsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "admin", func(claims *models.AppClaims) {
			// listing signups
			signups, err := repository.ListSignups(r.Context())
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)
			// encoding response
			encode(w, &signups)
		})
	}
}

// ListDashboardsHandler handle the sellect all of dashboards
func ListDashboardsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "admin", func(claims *models.AppClaims) {
			// listing dashboards
			dashboards, err := repository.ListDashboards(r.Context())
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)
			// encoding response
			encode(w, &dashboards)
		})
	}
}

// ListAuthsHandler handle the sellect all of auths
func ListAuthsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token
		token, err := getToken(s, r, "Authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate token
		validateTokenAndRole(w, token, "tuwsper", func(claims *models.AppClaims) {
			// listing auths
			auths, err := repository.ListAuths(r.Context())
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)
			// encoding response
			encode(w, &auths)
		})
	}
}

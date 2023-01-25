package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"tuwsp/models"
	"tuwsp/repository"
	"tuwsp/server"

	"github.com/golang-jwt/jwt"
)

// decode do a decode with package json from new decoder of interface
func decode(r *http.Request, w http.ResponseWriter, a any) {
	if err := json.
		NewDecoder(r.Body).Decode(a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// encode do a encode with package json from new encoder of interface
func encode(w http.ResponseWriter, a any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(a); err != nil {
		internalErr(w, err)
	}
}

// validateAgainsyDB do consusult to db about the user
func validateAgainstDB(email string, password string, ctx context.Context) (bool, string) {
	user, err := repository.GetAuthByEmail(ctx, email)
	if err != nil || user == nil || user.Password != password {
		return false, "credenciales invalidas"
	}
	return true, ""
}

// internalErr handle error with internalServerError of http
func internalErr(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// unauthorizedError hndle an unathorized error
func unathorizedError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
}

// getToken get a token from header
func getToken(s server.Server, r *http.Request, authorization string) (*jwt.Token, error) {
	// getting token from header
	tokenString := strings.TrimSpace(r.Header.Get(authorization))
	//  Parsing token string
	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config().JWTSecret), nil
	})
	return token, err
}

// validateToken validate token
func validateToken(w http.ResponseWriter, token *jwt.Token, f func(claims *models.AppClaims)) {
	if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
		f(claims)
	} else {
		internalErr(w, fmt.Errorf("invalid token"))
	}
}

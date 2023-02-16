package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"tuwsp/models"
	"tuwsp/repository"
	"tuwsp/server"

	"github.com/golang-jwt/jwt"
)

// encode do a encode with package json from new encoder of interface
func encode(w http.ResponseWriter, a any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(a); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// validateAgainsyDB do consusult to db about the user
func validateAgainstDB(email string, password string, ctx context.Context) (bool, string) {
	user, err := repository.GetAuthByEmail(ctx, email)
	if err != nil || user.Password != password {
		return false, "email o contraseña invalidos"
	}
	return true, ""
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

// validateTokenAndRole validate token and role
func validateTokenAndRole(w http.ResponseWriter, token *jwt.Token, role string, f func(claims *models.AppClaims)) {
	if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
		if claims.Role != role {
			http.Error(w, "no tienes permisos", http.StatusUnauthorized)
			return
		}
		f(claims)
	} else {
		http.Error(w, "invalid token", http.StatusUnauthorized)
	}
}

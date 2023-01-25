package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/golang-sql/civil"
)

type AppClaims struct {
	AuthId   string `json:"auth_id"`
	SignupId string `json:"signup_id"`
	jwt.StandardClaims
}

type Auth struct {
	Id        string         `json:"id"`
	Email     string         `json:"email"`
	CreatedAt civil.DateTime `json:"created_at"`
	Password  string         `json:"password"`
	SignupId  string         `json:"signup_id"`
}

package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/golang-sql/civil"
)

type AppClaims struct {
	AuthId string `json:"auth_id"`
	jwt.StandardClaims
}

type Auth struct {
	Id       string         `url:"id"`
	Email    string         `url:"email"`
	CratedAt civil.DateTime `url:"created_at"`
	Password string         `url:"password"`
}

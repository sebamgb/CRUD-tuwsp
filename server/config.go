package server

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

var DatabaseUrl []string

type Config struct {
	port        string
	dataBaseUrl string
	driverDb    string
	JWTSecret   string
}

// NewConfig builder of Config
func NewConfig(DRIVER_DB, SERVER_DB, USER_DB, PASSWORD_DB, PORT_DB, DATABASE string) (*Config, error) {
	query := url.Values{}
	query.Add("database", DATABASE)

	u := &url.URL{
		Scheme:   DRIVER_DB,
		User:     url.UserPassword(USER_DB, PASSWORD_DB),
		Host:     fmt.Sprintf("%s:%s", SERVER_DB, PORT_DB),
		RawQuery: query.Encode(),
	}
	Port := os.Getenv("PORT_APP")
	Jwt := os.Getenv("JWT_SECRET")
	prefix := strings.Index(Port, ":")
	sb := strings.Builder{}
	sb.WriteString(":")
	sb.WriteString(Port)
	if prefix != 0 {
		Port = sb.String()
	}
	return &Config{
		dataBaseUrl: u.String(),
		port:        Port,
		driverDb:    DRIVER_DB,
		JWTSecret:   Jwt,
	}, nil
}

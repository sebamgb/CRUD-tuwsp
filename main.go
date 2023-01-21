package main

import (
	"context"
	"log"
	"os"

	"tuwsp/routes"
	"tuwsp/server"

	"github.com/joho/godotenv"
)

var DRIVER_DB, SERVER_DB, USER_DB, PASSWORD_DB, PORT_DB, DATABASE string

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
		return
	}
	DRIVER_DB = os.Getenv("DRIVER_DB")
	SERVER_DB = os.Getenv("SERVER_DB")
	USER_DB = os.Getenv("USER_DB")
	PASSWORD_DB = os.Getenv("PASSWORD_DB")
	PORT_DB = os.Getenv("PORT_DB")
	DATABASE = os.Getenv("DATABASE")
}

func main() {
	if config, err := server.
		NewConfig(DRIVER_DB, SERVER_DB, USER_DB, PASSWORD_DB, PORT_DB, DATABASE); err != nil {
		log.Fatal(err)
	} else if s, err := server.NewServer(context.Background(), config); err != nil {
		log.Fatal(err)
	} else if err := s.
		Up(routes.BindRoute); err != nil {
		log.Fatal(err)
		return
	}
}

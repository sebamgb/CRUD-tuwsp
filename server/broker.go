package server

import (
	"log"
	"net/http"
	"time"

	"tuwsp/db"
	"tuwsp/repository"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Broker struct {
	config *Config
	router *mux.Router
	servr  *http.Server
}

type binder func(s Server, r *mux.Router)

// Config() method of Broker that return a Config
func (b *Broker) Config() *Config {
	return b.config
}

// Up() method of Broker to up the server
func (b *Broker) Up(binder binder) (err error) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	handler := cors.Default().Handler(b.router)
	b.servr = &http.Server{
		Handler: handler,
		Addr:    b.config.port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	repo := db.GetSQLServer(b.config.driverDb, b.config.dataBaseUrl)
	repository.SetRepository(repo)
	defer repository.Close()
	log.Println("Starting to listen in port", b.config.port)
	if err = b.servr.ListenAndServe(); err != nil {
		return
	}
	return
}

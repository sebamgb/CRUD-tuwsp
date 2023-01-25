package handlers

import (
	"encoding/json"
	"net/http"

	"tuwsp/models"
	"tuwsp/repository"
	"tuwsp/server"

	"github.com/gorilla/mux"
)

/*----- Inserts -----*/

// InsertProtocolHandler handle the insert of protocols
func InsertProtocolHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		protocolRequest := models.Protocol{}
		decode(r, w, &protocolRequest)
		internalErr(w, repository.
			InsertIntoProtocols(r.Context(), &protocolRequest))
		encode(w, &protocolRequest)
	}
}

// InsertUrlHandler handle the insert of urls
func InsertUrlHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlRequest := models.Url{}
		decode(r, w, &urlRequest)
		internalErr(w, repository.
			InsertIntoURLs(r.Context(), &urlRequest))
		encode(w, &urlRequest)
	}
}

// InsertEndpointHandler handle the insert of endpoints
func InsertEndpointHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		endpointRequest := models.Endpoint{}
		decode(r, w, &endpointRequest)
		internalErr(w, repository.
			InsertIntoEndpoints(r.Context(), &endpointRequest))
		encode(w, &endpointRequest)
	}
}

// InsertQueryKeyHandler handle the insert of query_keys
func InsertQueryKeyHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		querykeyRequest := models.QueryKey{}
		decode(r, w, &querykeyRequest)
		internalErr(w, repository.
			InsertIntoQueryKeys(r.Context(), &querykeyRequest))
		encode(w, &querykeyRequest)
	}
}

// InsertQueryValueHandler handle the insert of query_values
func InsertQueryValueHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryvalueRequest := models.QueryValue{}
		decode(r, w, &queryvalueRequest)
		internalErr(w, repository.
			InsertIntoQueryValues(r.Context(), &queryvalueRequest))
		encode(w, &queryvalueRequest)
	}
}

/*----- Gets -----*/

// GetProtocolHandler handle the select of protocols
func GetProtocolHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		protocol, err := repository.
			GetProtocolById(r.Context(), params["q"])
		if err != nil {
			return
		}
		encode(w, &protocol)
		json.NewEncoder(w).Encode(&protocol)
	}
}

// GetUrlHandler handle the select of urls
func GetUrlHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		url, err := repository.
			GetUrlById(r.Context(), params["q"])
		if err != nil {
			return
		}
		encode(w, &url)
	}
}

// GetEndpointHandler handle the select of endpoints
func GetEndpointHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		endpoint, err := repository.
			GetEndPointByUrlId(r.Context(), params["q"])
		if err != nil {
			return
		}
		encode(w, &endpoint)
	}
}

// GetQueryKeyHandler handle the select of query_keys
func GetQueryKeyHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		querykey, err := repository.
			GetQueryKeyByUrlId(r.Context(), params["q"])
		if err != nil {
			return
		}
		encode(w, &querykey)
	}
}

// GetQueryValueHandler handle the select of query_values
func GetQueryValueHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		queryvalue, err := repository.
			GetQueryValueByUserId(r.Context(), params["q"])
		if err != nil {
			return
		}
		encode(w, &queryvalue)
	}
}

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"tuwsp/models"
	"tuwsp/repository"
	"tuwsp/server"
)

// InsertProtocolHandler handle the insert of protocols
func InsertProtocolHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		protocolRequest := models.Protocol{}
		decode(r, w, &protocolRequest)
		internalErr(w, repository.
			InsertIntoProtocols(r.Context(), &protocolRequest))
	}
}

// InsertUrlHandler handle the insert of urls
func InsertUrlHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlRequest := models.Url{}
		decode(r, w, &urlRequest)
		internalErr(w, repository.
			InsertIntoURLs(r.Context(), &urlRequest))
	}
}

// InsertEndpointHandler handle the insert of endpoints
func InsertEndpointHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		endpointRequest := models.Endpoint{}
		decode(r, w, &endpointRequest)
		internalErr(w, repository.
			InsertIntoEndpoints(r.Context(), &endpointRequest))
	}
}

// InsertQueryKeyHandler handle the insert of query_keys
func InsertQueryKeyHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		querykeyRequest := models.QueryKey{}
		decode(r, w, &querykeyRequest)
		internalErr(w, repository.
			InsertIntoQueryKeys(r.Context(), &querykeyRequest))
	}
}

// InsertQueryValueHandler handle the insert of query_values
func InsertQueryValueHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryvalueRequest := models.QueryValue{}
		decode(r, w, &queryvalueRequest)
		internalErr(w, repository.
			InsertIntoQueryValues(r.Context(), &queryvalueRequest))
	}
}

// GetProtocolHandler handle the select of protocols
func GetProtocolHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		protocol, err := repository.
			GetProtocolById(r.Context(), id)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		json.NewEncoder(w).Encode(protocol)
	}
}

// GetUrlHandler handle the select of urls
func GetUrlHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		url, err := repository.
			GetUrlById(r.Context(), id)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		json.NewEncoder(w).Encode(url)
	}
}

// GetEndpointHandler handle the select of endpoints
func GetEndpointHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		endpoint, err := repository.
			GetEndPointByUrlId(r.Context(), id)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		json.NewEncoder(w).Encode(endpoint)
	}
}

// GetQueryKeyHandler handle the select of query_keys
func GetQueryKeyHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		querykey, err := repository.
			GetQueryKeyByUrlId(r.Context(), id)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		json.NewEncoder(w).Encode(querykey)
	}
}

// GetQueryValueHandler handle the select of query_values
func GetQueryValueHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		queryvalue, err := repository.
			GetQueryValueByUrlId(r.Context(), id)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		json.NewEncoder(w).Encode(queryvalue)
	}
}

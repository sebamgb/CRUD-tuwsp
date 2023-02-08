package handlers

import (
	"net/http"

	"tuwsp/models"
	"tuwsp/repository"
	"tuwsp/server"

	"github.com/gorilla/mux"
)

/*----- Inserts -----*/

type InsertProtocolResponse struct {
	Id string `json:"id"`
}

// InsertProtocolHandler handle the insert of protocols
func InsertProtocolHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		protocolRequest := models.Protocol{}
		// decode request into body
		decode(r, w, &protocolRequest)
		// insert protocol
		err := repository.InsertIntoProtocols(r.Context(), &protocolRequest)
		internalErr(w, err)
		// encoding response
		encode(w, &InsertProtocolResponse{Id: protocolRequest.Id})
	}
}

type InsertUrlResponse struct {
	Id string `json:"id"`
}

// InsertUrlHandler handle the insert of urls
func InsertUrlHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlRequest := models.Url{}
		// decode request into body
		decode(r, w, &urlRequest)
		// insert url
		err := repository.InsertIntoURLs(r.Context(), &urlRequest)
		internalErr(w, err)
		// encoding response
		encode(w, &InsertUrlResponse{Id: urlRequest.Id})
	}
}

type InsertEndpointResponse struct {
	EndPoint string `json:"endpoint"`
	UrlId    string `json:"url_id"`
}

// InsertEndpointHandler handle the insert of endpoints
func InsertEndpointHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		endpointRequest := models.Endpoint{}
		// decode request into body
		decode(r, w, &endpointRequest)
		// insert endpoint
		err := repository.InsertIntoEndpoints(r.Context(), &endpointRequest)
		internalErr(w, err)
		// encoding response
		encode(w, &InsertEndpointResponse{EndPoint: endpointRequest.Endpoint, UrlId: endpointRequest.UrlId})
	}
}

type InsertQueryKeyResponse struct {
	QueryKey string `json:"query_key"`
	UrlId    string `json:"url_id"`
}

// InsertQueryKeyHandler handle the insert of query_keys
func InsertQueryKeyHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		querykeyRequest := models.QueryKey{}
		// decode request into body
		decode(r, w, &querykeyRequest)
		// insert query_key
		err := repository.InsertIntoQueryKeys(r.Context(), &querykeyRequest)
		internalErr(w, err)
		// encoding response
		encode(w, &InsertQueryKeyResponse{QueryKey: querykeyRequest.KeyParam, UrlId: querykeyRequest.UrlId})
	}
}

type InsertQueryValueResponse struct {
	QueryValue string `json:"query_value"`
	UserId     string `json:"user_id"`
}

// InsertQueryValueHandler handle the insert of query_values
func InsertQueryValueHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryvalueRequest := models.QueryValue{}
		// decode request into body
		decode(r, w, &queryvalueRequest)
		// insert query_value
		err := repository.InsertIntoQueryValues(r.Context(), &queryvalueRequest)
		internalErr(w, err)
		// encoding response
		encode(w, &InsertQueryValueResponse{QueryValue: queryvalueRequest.ValueParam, UserId: queryvalueRequest.UserId})
	}
}

/*----- Gets -----*/

// GetProtocolHandler handle the select of protocols
func GetProtocolHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get queryparams
		params := r.URL.Query()
		// getting protocol
		protocol, err := repository.
			GetProtocolById(r.Context(), params.Get("q"))
		internalErr(w, err)
		// encoding response
		encode(w, &protocol)
	}
}

// GetUrlHandler handle the select of urls
func GetUrlHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get queryparams
		params := r.URL.Query()
		// getting url
		url, err := repository.
			GetUrlById(r.Context(), params.Get("q"))
		internalErr(w, err)
		// encoding response
		encode(w, &url)
	}
}

// GetEndpointHandler handle the select of endpoints
func GetEndpointHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get queryparams
		params := r.URL.Query()
		// getting endpoint
		endpoint, err := repository.
			GetEndPointByUrlId(r.Context(), params.Get("q"))
		internalErr(w, err)
		// encoding response
		encode(w, &endpoint)
	}
}

// GetQueryKeyHandler handle the select of query_keys
func GetQueryKeyHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get queryparams
		params := r.URL.Query()
		// getting query_key
		querykey, err := repository.
			GetQueryKeyByUrlId(r.Context(), params.Get("q"))
		internalErr(w, err)
		// encoding response
		encode(w, &querykey)
	}
}

// GetQueryValueHandler handle the select of query_values
func GetQueryValueHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get queryparams
		params := r.URL.Query()
		// getting query_value
		queryvalue, err := repository.
			GetQueryValueByUserId(r.Context(), params.Get("q"))
		internalErr(w, err)
		// encoding response
		encode(w, &queryvalue)
	}
}

/*----- Updates -----*/

type UpdateProtocolResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// UpdateProtocolHandler handle the update of protocols
func UpdateProtocolHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting mux vars
		vars := mux.Vars(r)
		// getting id from mux vars
		id := vars["id"]
		// geting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			protcolRequest := models.Protocol{
				Id: id,
			}
			// decode request into body
			decode(r, w, &protcolRequest)
			// updatting protocol
			err := repository.UpdateProtocols(r.Context(), &protcolRequest)
			internalErr(w, err)
			// encoding response
			encode(w, &UpdateProtocolResponse{Success: true, Message: "Protocol updated successfully", Author: claims.AuthId})
		})
	}
}

type UpdateUrlResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// UpdateUrlHandler handle the update of urls
func UpdateUrlHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting mux vars
		vars := mux.Vars(r)
		// getting id from mux vars
		id := vars["id"]
		// geting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			urlRequest := models.Url{
				Id: id,
			}
			// decode request into body
			decode(r, w, &urlRequest)
			// updatting url
			err := repository.UpdateURLs(r.Context(), &urlRequest)
			internalErr(w, err)
			// encoding response
			encode(w, &UpdateUrlResponse{Success: true, Message: "Url updated successfully", Author: claims.AuthId})
		})
	}
}

type UpdateEndpointResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// UpdateEndpointHandler handle the update of endpoints
func UpdateEndpointHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting mux vars
		vars := mux.Vars(r)
		// getting id from mux vars
		id := vars["id"]
		// geting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			endpointRequest := models.Endpoint{
				UrlId: id,
			}
			// decode request into body
			decode(r, w, &endpointRequest)
			// updatting endpoint
			err := repository.UpdateEndpoints(r.Context(), &endpointRequest)
			internalErr(w, err)
			// encoding response
			encode(w, &UpdateEndpointResponse{Success: true, Message: "Endpoint updated successfully", Author: claims.AuthId})
		})
	}
}

type UpdateQueryKeyResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// UpdateQueryKeyHandler handle the update of query_keys
func UpdateQueryKeyHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting mux vars
		vars := mux.Vars(r)
		// getting id from mux vars
		id := vars["id"]
		// geting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			queryKeyRequest := models.QueryKey{
				UrlId: id,
			}
			// decode request into body
			decode(r, w, &queryKeyRequest)
			// updatting query_key
			err := repository.UpdateQueryKeys(r.Context(), &queryKeyRequest)
			internalErr(w, err)
			// encoding response
			encode(w, &UpdateQueryKeyResponse{Success: true, Message: "QueryKey updated successfully", Author: claims.AuthId})
		})
	}
}

type UpdateQueryValueResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// UpdateQueryValueHandler handle the update of query_values
func UpdateQueryValueHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting mux vars
		vars := mux.Vars(r)
		// getting id from mux vars
		id := vars["id"]
		// geting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			queryValueRequest := models.QueryValue{
				UserId: id,
			}
			// decode request into body
			decode(r, w, &queryValueRequest)
			// updatting query_value
			err := repository.UpdateQueryValues(r.Context(), &queryValueRequest)
			internalErr(w, err)
			// encoding response
			encode(w, &UpdateQueryValueResponse{Success: true, Message: "QueryValue updated successfully", Author: claims.AuthId})
		})
	}
}

/*----- Deletes -----*/

type DeleteProtocolResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// DeleteProtocolHandler handle the delete of protocols
func DeleteProtocolHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting mux vars
		vars := mux.Vars(r)
		// getting id from mux vars
		id := vars["id"]
		// geting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			protocolRequest := models.Protocol{
				Id: id,
			}
			// deleting protocol
			err := repository.DeleteProtocols(r.Context(), &protocolRequest)
			internalErr(w, err)
			// encoding response
			encode(w, &DeleteProtocolResponse{Success: true, Message: "Protocol deleted successfully", Author: claims.AuthId})
		})
	}
}

type DeleteUrlResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// DeleteUrlHandler handle the delete of urls
func DeleteUrlHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting mux vars
		vars := mux.Vars(r)
		// getting id from mux vars
		id := vars["id"]
		// geting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			urlRequest := models.Url{
				Id: id,
			}
			// deleting url
			err := repository.DeleteURLs(r.Context(), &urlRequest)
			internalErr(w, err)
			// encoding response
			encode(w, &DeleteUrlResponse{Success: true, Message: "Url deleted successfully", Author: claims.AuthId})
		})
	}
}

type DeleteEndpointResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// DeleteEndpointHandler handle the delete of endpoints
func DeleteEndpointHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting mux vars
		vars := mux.Vars(r)
		// getting id from mux vars
		id := vars["id"]
		// geting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			endpointRequest := models.Endpoint{
				UrlId: id,
			}
			// deleting endpoint
			err := repository.DeleteEndpoints(r.Context(), &endpointRequest)
			internalErr(w, err)
			// encoding response
			encode(w, &DeleteEndpointResponse{Success: true, Message: "Endpoint deleted successfully", Author: claims.AuthId})
		})
	}
}

type DeleteQueryKeyResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// DeleteQueryKeyHandler handle the delete of query_keys
func DeleteQueryKeyHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting mux vars
		vars := mux.Vars(r)
		// getting id from mux vars
		id := vars["id"]
		// geting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			queryKeyRequest := models.QueryKey{
				UrlId: id,
			}
			// deleting query_key
			err := repository.DeleteQueryKeys(r.Context(), &queryKeyRequest)
			internalErr(w, err)
			// encoding response
			encode(w, &DeleteQueryKeyResponse{Success: true, Message: "QueryKey deleted successfully", Author: claims.AuthId})
		})
	}
}

type DeleteQueryValueResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

// DeleteQueryValueHandler handle the delete of query_values
func DeleteQueryValueHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting mux vars
		vars := mux.Vars(r)
		// getting id from mux vars
		id := vars["id"]
		// geting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			queryValueRequest := models.QueryValue{
				UserId: id,
			}
			// deleting query_value
			err := repository.DeleteQueryValues(r.Context(), &queryValueRequest)
			internalErr(w, err)
			// encoding response
			encode(w, &DeleteQueryValueResponse{Success: true, Message: "QueryValue deleted successfully", Author: claims.AuthId})
		})
	}
}

/* ----- Lists ----- */

// ListProtocolsHandler handle the sellect all of protocols
func ListProtocolsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// geting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			// list protocols
			protocols, err := repository.ListProtocols(r.Context())
			internalErr(w, err)
			// encoding response
			encode(w, &protocols)
		})
	}
}

// ListURLsHandler handle the sellect all of urls
func ListURLsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// geting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			// list urls
			urls, err := repository.ListURLs(r.Context())
			internalErr(w, err)
			// encoding response
			encode(w, &urls)
		})
	}
}

// ListEndpointsHandler handle the sellect all of endpoints
func ListEndpointsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// geting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			// list endpoints
			endpoints, err := repository.ListEndpoints(r.Context())
			internalErr(w, err)
			// encoding response
			encode(w, &endpoints)
		})
	}
}

// ListQueryKeysHandler handle the sellect all of query_keys
func ListQueryKeysHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// geting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			// list query_keys
			queryKeys, err := repository.ListQueryKeys(r.Context())
			internalErr(w, err)
			// encoding response
			encode(w, &queryKeys)
		})
	}
}

// ListQueryValuesHandler handle the sellect all of query_values
func ListQueryValuesHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// geting token
		token, err := getToken(s, r, "Authorization")
		unathorizedError(w, err)
		// validating token
		validateToken(w, token, func(claims *models.AppClaims) {
			// list query_values
			queryValues, err := repository.ListQueryValues(r.Context())
			internalErr(w, err)
			// encoding response
			encode(w, &queryValues)
		})
	}
}

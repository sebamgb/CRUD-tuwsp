package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// decode do a decode with package json from new decoder of interface
func decode(r *http.Request, w http.ResponseWriter, a any) {
	if err := json.
		NewDecoder(r.Body).Decode(a); err != nil {
		http.Error(w, "decode:"+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("a: %v\n", a)
}

// encode do a encode with package json from new encoder of interface
func encode(w http.ResponseWriter, a any) {
	internalErr(w, json.NewEncoder(w).Encode(a))
	fmt.Printf("a: %v\n", a)
}

// internalErr handle error with internalServerError of http
func internalErr(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, "internal db:"+err.Error(), http.StatusInternalServerError)
		return
	}
}

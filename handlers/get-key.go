// Package handlers contains the functions that will handle http requests
package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/netrebel/kube-api/storage"
)

// GetKey returns an http.Handler that can get a key registered by Gorilla mux
// as "key" in the path. It gets the value of the key from db
func GetKey(db storage.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]
		log.Printf("Getting %s", key)
		if key == "" {
			http.Error(w, "missing key name in query string", http.StatusBadRequest)
			return
		}
		val, err := db.Get(key)
		if err == storage.ErrNotFound {
			http.Error(w, "404 not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, fmt.Sprintf("error getting value from database: %s", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(val)
	}
}

package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/netrebel/kube-api/storage"
)

// PutKey returns an http.Handler that can set a value for the key registered by Gorilla
// mux as "key" in the path. It expects the value to be in the body of the PUT request
func PutKey(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]
		if key == "" {
			http.Error(w, "missing key name in path", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		val, err := ioutil.ReadAll(r.Body)
		log.Printf("Saving %s:%s", key, val)
		if err != nil {
			http.Error(w, "error reading PUT body", http.StatusBadRequest)
			return
		}
		if err := db.Set(key, val); err != nil {
			http.Error(w, fmt.Sprintf("error setting value in DB - %s", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

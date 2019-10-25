package main

import (
	"github.com/netrebel/kube-api/src/handlers"
	"github.com/netrebel/kube-api/src/storage"
	"log"
	"net/http"
)

func main() {
	db := storage.NewInMemoryDB()

	// this creates a new http.ServeMux, which is used to register handlers to execute in response to routes
	mux := http.NewServeMux()
	mux.HandleFunc("/get", handlers.GetKey(db))
	mux.Handle("/put", handlers.PutKey(db))

	log.Printf("serving on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

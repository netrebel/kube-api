package main

import (
	"log"
	"net/http"
	"os"

	"github.com/netrebel/kube-api/handlers"
	"github.com/netrebel/kube-api/storage"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	db := storage.NewInMemoryDB()

	// this creates a new http.ServeMux, which is used to register handlers to execute in response to routes
	mux := http.NewServeMux()
	mux.HandleFunc("/get", handlers.GetKey(db))
	mux.Handle("/put", handlers.PutKey(db))

	// Configure Logging
	logFileLocation := os.Getenv("LOG_FILE_LOCATION")
	if logFileLocation != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   logFileLocation,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
	}

	// Start Server
	log.Println("Serving on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

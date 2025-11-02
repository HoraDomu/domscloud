package main

import (
	"log"
	"net/http"

	"domscloud/modules/file"
	"domscloud/modules/health"
	"domscloud/modules/logging"
)

func main() {
	// File routes
	http.HandleFunc("/upload", file.UploadHandler)
	http.HandleFunc("/download", file.DownloadHandler)
	http.HandleFunc("/list", file.ListHandler)

	// Logging routes
	http.HandleFunc("/log", logging.LogHandler)

	// Health route
	http.HandleFunc("/health", health.HealthHandler)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}

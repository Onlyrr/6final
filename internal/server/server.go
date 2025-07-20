package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type AppServer struct {
	Log        *log.Logger
	HTTPServer *http.Server
}

func NewServer(Log *log.Logger) *AppServer {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler(Log))

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     Log,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &AppServer{
		Log:        Log,
		HTTPServer: srv,
	}
}

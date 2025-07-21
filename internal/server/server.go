package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/handlers"
)

type MyServer struct {
	Log    *log.Logger
	Server *http.Server
}

func NewServer(log *log.Logger) *MyServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     log,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &MyServer{
		Log:    log,
		Server: srv,
	}
}

func (s *MyServer) Run() error {
	s.Log.Println("Сервер запущен")
	return s.Server.ListenAndServe()
}

package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/server"
)

func main() {
	logger := log.New(os.Stdout, "[http] ", log.LstdFlags)
	srv := server.NewServer(logger)

	if err := srv.Run(); err != nil {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

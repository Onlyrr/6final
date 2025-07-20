package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "app: ", log.LstdFlags)

	appServer := server.NewServer(logger)

	logger.Println("Запуск сервера")
	if err := appServer.HTTPServer.ListenAndServe(); err != nil {
		logger.Fatal("Ошибка запуска сервера:", err)
	}
}

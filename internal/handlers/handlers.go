package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("сервер не поддерживает: %v", r.Method), http.StatusMethodNotAllowed)
		return
	}

	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("ошибка чтения файла: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	_, err = w.Write([]byte(data))
	if err != nil {
		http.Error(w, fmt.Sprintf("ошибка записи: %v", err), http.StatusInternalServerError)
		return
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("неправильный метод: %v", r.Method), http.StatusMethodNotAllowed)
		return
	}
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, fmt.Sprintf("ошибка при получении файла: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, fmt.Sprintf("ошибка чтения файла: %v", err), http.StatusInternalServerError)
		return
	}

	result, err := service.Analysis(string(data))
	if err != nil {
		http.Error(w, fmt.Sprintf("ошибка конвертации: %v", err), http.StatusInternalServerError)
		return
	}

	fileName := time.Now().UTC().String() + ".txt"

	newFile, err := os.Create(fileName)
	if err != nil {
		http.Error(w, fmt.Sprintf("ошибка при создании файла: %v", err), http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = newFile.WriteString(result)
	if err != nil {
		http.Error(w, fmt.Sprintf("ошибка записи в файл: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	_, err = w.Write([]byte(result))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error writing: %v", err), http.StatusInternalServerError)
		return
	}
}

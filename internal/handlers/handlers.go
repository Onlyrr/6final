package handlers

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func UploadHandler(log *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			log.Println("Ошибка парсинга:", err)
			http.Error(w, "Ошибка обработки:", http.StatusInternalServerError)
			return
		}

		file, _, err := r.FormFile("file")
		if err != nil {
			log.Println("Ошибка извлечения:", err)
			http.Error(w, "Ошибка загрузки:", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			log.Println("Ошибка чтения файла:", err)
			http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
			return
		}

		inputStr := string(content)
		result, err := service.Analysis(inputStr)
		if err != nil {
			log.Println("Ошибка определения содержимого:", err)
			http.Error(w, "Ошибка обработки данных", http.StatusInternalServerError)
			return
		}

		filename := "result:" + time.Now().UTC().Format("20060102_150405") + filepath.Ext("file")

		err = os.WriteFile(filename, []byte(result), 0644)
		if err != nil {
			log.Println("Ошибка записи файла:", err)
			http.Error(w, "Ошибка сохранения", http.StatusInternalServerError)
			return
		}

		w.Write([]byte(filename))
	}
}

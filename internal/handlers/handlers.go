package handlers

import (
	"net/http"
	"strings"
	"os"
	"io"
	"path/filepath"
	"time"
	service "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("index.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ошибка чтения файла"))
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()

    data, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "ошибка чтения", http.StatusInternalServerError)
        return
    }

    message := string(data)

    result, err := service.TextDetector(message)
    if err != nil {
        http.Error(w, "ошибка обработки", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte(result))
}
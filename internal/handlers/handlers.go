package handlers

import (
	"net/http"
	"os"
	"io"
	"path/filepath"
	"time"
	"strings"
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

    var data []byte

    contentType := r.Header.Get("Content-Type")

    if strings.Contains(contentType, "multipart/form-data") {
        err := r.ParseMultipartForm(10 << 20)
        if err != nil {
            http.Error(w, "ошибка парсинга формы", http.StatusInternalServerError)
            return
        }

        for _, files := range r.MultipartForm.File {
            for _, header := range files {
                file, err := header.Open()
                if err != nil {
                    continue
                }
                defer file.Close()

                data, err = io.ReadAll(file)
                if err != nil {
                    http.Error(w, "ошибка чтения файла", http.StatusInternalServerError)
                    return
                }
                break
            }
            if data != nil {
                break
            }
        }

    } else {
        var err error
        data, err = io.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "ошибка чтения тела", http.StatusInternalServerError)
            return
        }
    }

    message := string(data)

    result, err := service.TextDetector(message)
    if err != nil {
        http.Error(w, "ошибка обработки", http.StatusInternalServerError)
        return
	}

    fileName := time.Now().UTC().Format("20060102_150405") +  filepath.Ext(header.Filename)

    file, err := os.Create(fileName)
    if err != nil {
        http.Error(w, "ошибка создания файла", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    _, err = file.Write([]byte(result))
    if err != nil {
        http.Error(w, "ошибка записи файла", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte(result))
}
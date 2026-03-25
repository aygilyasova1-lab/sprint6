package handlers

import (
	"net/http"
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
		w.Write([]byte("ошибка при чтении файла"))
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("upload")
	if err == nil {
		defer file.Close()
		data, err = io.ReadAll(file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		    w.Write([]byte("ошибка  при чтении файла"))
		}
		} else {
		data, err = io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("ошибка при чтении запроса"))
			} 
		}
	message := string(data)
	result, err := service.TextDetector(message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ошибка конвертации данных"))
		return
	}

	fileName := time.Now().UTC().Format("20060102_150405") + filepath.Ext(header.Filename)

	localFile, err := os.Create(fileName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ошибка создания файла"))
		return
	}

	defer localFile.Close()

	_, err = localFile.Write([]byte(result))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ошибка при записи данных"))
		return
	} 
	
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(result))

}
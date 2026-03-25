package handlers

import (
	"net/http"
	"os"
	"io"
	"path/filepath"
	"time"
	service "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
	"strings"
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
	var data []byte
	var fileName string
	contentType := r.Header.Get("Content-Type")

	if strings.Contains(contentType, "multipart/form-data") {

		file, header, err := r.FormFile("upload")
 		if err != nil {
			data, err = io.ReadAll(r.Body)
   			if err != nil {
   				http.Error(w, "ошибка при чтении данных", http.StatusInternalServerError)
   				return
  				}
			fileName = "body_" + time.Now().UTC().Format("20060102_150405") + ".txt"
  		} else {
   		defer file.Close()
  		data, err = io.ReadAll(file)
		if err != nil {
			http.Error(w, "ошибка при чтении файла", http.StatusInternalServerError)
			return
 			}
   		fileName = time.Now().UTC().Format("20060102_150405") + filepath.Ext(header.Filename)
		}
	} else {
		var err error
		data, err = io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "ошибка при чтении тела запроса", http.StatusInternalServerError)
			return
 			}
 		 fileName = "body_" + time.Now().UTC().Format("20060102_150405") + ".txt"
 		}

 	message := string(data)

	result, err := service.TextDetector(message)
 	if err != nil {
		result = message 
 	}

	localFile, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "ошибка создания файла", http.StatusInternalServerError)
  		return
		}

 	defer localFile.Close()
	_, err = localFile.Write([]byte(result))
	if err != nil {
		http.Error(w, "ошибка при записи данных", http.StatusInternalServerError)
		return
		}

	w.Header().Set("Content-Type", "text/plain")
 	w.Write([]byte(result))
}
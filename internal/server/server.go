package server

import (
	"time"
	"net/http"
	"log"
	handler "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)
type Server struct {
	Logger *log.Logger
	HTTP *http.Server
}

func NewServer(logger *log.Logger) *Server {
	router := http.NewServeMux()

	router.HandleFunc("/", handler.RootHandler)
	router.HandleFunc("/upload", handler.UploadHandler)

	srv := &http.Server {
		Addr:		   ":8080",
		Handler:	   router,
		ErrorLog:	   logger,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  10 * time.Second,
		IdleTimeout:   15 * time.Second,
	}

	return &Server{
		Logger:  logger,
		HTTP: 	 srv,
	}
}
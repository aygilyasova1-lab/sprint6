package main

import (
	"os"
	"log"
	"net/http"
	server "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	morseLogger := log.New(os.Stdout, "server: ", log.LstdFlags)

	morseServer := server.NewServer(morseLogger)

	err := morseServer.HTTP.ListenAndServe()
	if err != nil {
		morseLogger.Fatal(err)
	}

}

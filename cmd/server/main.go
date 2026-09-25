package main

import (
	"log"
	"net/http"

	"github.com/dhxli/daysNewYear/internal/httpapi"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/days-left", httpapi.DaysLeftHandler)

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

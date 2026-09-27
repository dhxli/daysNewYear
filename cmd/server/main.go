package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/dhxli/daysNewYear/internal/httpapi"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/days-left", httpapi.DaysLeftHandler)
	mux.HandleFunc("/healthz", httpapi.HealthzHandler)

	handler := httpapi.LoggingMiddleware(mux)

	slog.Info("server starting", "addr", ":8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

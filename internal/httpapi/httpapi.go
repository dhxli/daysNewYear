package httpapi

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/dhxli/daysNewYear/internal/daysleft"
)

type response struct {
	Date     string `json:"date"`
	DaysLeft int    `json:"days_left"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type healthResponse struct {
	Status string `json:"status"`
}

func DaysLeftHandler(w http.ResponseWriter, r *http.Request) {
	dateParam := r.URL.Query().Get("date")
	slog.Info("получен запрос", "path", r.URL.Path, "date", dateParam)

	var target time.Time
	if dateParam == "" {
		target = time.Now()
	} else {
		parsed, err := time.Parse("2006-01-02", dateParam)
		if err != nil {
			slog.Warn("некорректная дата", "input", dateParam, "error", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorResponse{Error: "неверный формат даты, ожидается YYYY-MM-DD"})
			return
		}
		target = parsed
	}

	resp := response{
		Date:     target.Format("2006-01-02"),
		DaysLeft: daysleft.Until(target),
	}

	slog.Info("запрос обработан", "date", resp.Date, "days_left", resp.DaysLeft)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(healthResponse{Status: "ok"})
}

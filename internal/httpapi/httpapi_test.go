package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDaysLeftHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/days-left?date=2024-12-31", nil)
	rec := httptest.NewRecorder()

	DaysLeftHandler(rec, req)

	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("ожидался статус 200, получен %d", res.StatusCode)
	}

	var body response
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("не удалось разобрать JSON: %v", err)
	}

	if body.DaysLeft != 1 {
		t.Errorf("ожидалось DaysLeft = 1, получено %d", body.DaysLeft)
	}
}

func TestDaysLeftHandler_BadInput(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/days-left?date=неверно", nil)
	rec := httptest.NewRecorder()

	DaysLeftHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("ожидался статус 400, получен %d", rec.Code)
	}
}

func TestHealthzHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	HealthzHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("ожидался статус 200, получен %d", rec.Code)
	}

	var body healthResponse
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("не удалось разобрать JSON: %v", err)
	}

	if body.Status != "ok" {
		t.Errorf("ожидался status=ok, получено %q", body.Status)
	}
}

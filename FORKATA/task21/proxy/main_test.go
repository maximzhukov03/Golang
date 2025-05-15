package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProxy(t *testing.T){
	// r := chi.NewRouter()⁡
	reverseProx := NewReverseProxy("localhost", "1212")
	handler := reverseProx.ReverseProxy(http.NotFoundHandler())

	req := httptest.NewRequest("GET", "/api", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Body.String() != "Hello from API"{
		t.Errorf("ОЖИДАЛОСЬ %s ПОЛУЧИЛИ %s", "Hello from API", rec.Body.String())
	}

	if rec.Code != 200{
		t.Errorf("ОЖИДАЛОСЬ %d ПОЛУЧИЛИ %d", 200, rec.Code)
	}

}

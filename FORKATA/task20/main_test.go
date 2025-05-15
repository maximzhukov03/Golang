package main

import (
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func TestHttp(t *testing.T){
	r := chi.NewRouter()
	RouterConfigure(r)
	Logger()
	defer logger.Sync()
	tester := []struct{
		method string
		ph string
		code int
		body string
	}{
		{"GET", "/hello", 200, "Hello"},
		{"GET", "/hell", 404, "404 page not found\n"},

	}

	for _, test := range tester{
		req := httptest.NewRequest(test.method, test.ph, nil)
		rec := httptest.NewRecorder()
		
		r.ServeHTTP(rec, req)

		if rec.Code != test.code{
			t.Errorf("ОЖИДАЛОСЬ %d ПОЛУЧИЛИ: %d", test.code, rec.Code)
		}

		if rec.Body.String() != test.body{
			t.Errorf("ОЖИДАЛОСЬ %s ПОЛУЧИЛИ: %s", test.body, rec.Body.String())
		}
	}
}

func TestLogger(t *testing.T){
	Logger()
	if logger == nil{
		t.Error("ОШИБКА В ЛОГГЕРЕ")
	}
}


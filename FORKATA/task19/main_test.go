package main

import (
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func TestHttp(t *testing.T){
	r := chi.NewRouter()
	RouterConfigure(r)
	tester := []struct{
		method string
		ph string
		code int
		body string
	}{
		{"GET", "/1", 200, "Обработка 1го маршрута"},
		{"GET", "/2", 200, "Обработка 2го маршрута"},
		{"GET", "/3", 200, "Обработка 3го маршрута"},
		{"POST", "/3", 200, "Обработка 3го маршрута"},
		{"GET", "/", 404, "404 page not found\n"},

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
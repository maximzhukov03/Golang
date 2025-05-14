package main

import (
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T){
	r := RouterConfigure()
	tester := []struct{
		method string
		ph string
		code int
		body string
	}{
		{"GET", "/group1/1", 200, "Group 1 Привет, мир 1"},
		{"GET", "/group1/2", 200, "Group 1 Привет, мир 2"},
		{"GET", "/group1/3", 200, "Group 1 Привет, мир 3"},
		{"GET", "/group2/1", 200, "Group 2 Привет, мир 1"},
		{"GET", "/group2/2", 200, "Group 2 Привет, мир 2"},
		{"GET", "/group2/3", 200, "Group 2 Привет, мир 3"},
		{"GET", "/group3/1", 200, "Group 3 Привет, мир 1"},
		{"GET", "/group3/2", 200, "Group 3 Привет, мир 2"},
		{"GET", "/group3/3", 200, "Group 3 Привет, мир 3"},
		{"GET", "/group2/", 200, "Привет, мир 2"},
		{"GET", "/group3/", 200, "Привет, мир 3"},

	}

	for _, test := range tester{
		req := httptest.NewRequest(test.method, test.ph, nil)
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		if rec.Code != test.code{
			t.Errorf("ОЖИДАЛСЯ КОД %d ПОЛУЧИЛИ: %d", test.code, rec.Code)
		}

		if rec.Body.String() != test.body{
			t.Errorf("ОЖИДАЛОСЬ ТЕЛО ЗАПРОСА %s ПОЛУЧИЛИ: %s", test.body, rec.Body.String())
		}
	}
}
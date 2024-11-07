package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStudentsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/students", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StudentsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := Student{"aaa", 16, 87}
	var actual Student
	if err := json.NewDecoder(rr.Body).Decode(&actual); err != nil {
		t.Errorf("could not decode response: %v", err)
	}

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}

	contentType := rr.Header().Get("content-type")
	if contentType != "application/json" {
		t.Errorf("content-type header does not match: got %v want %v", contentType, "application/json")
	}
}

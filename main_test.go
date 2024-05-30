package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// TestAddHandler tests the AddHandler function
func TestAddHandler(t *testing.T) {
	tests := []struct {
		num1      string
		num2      string
		expected  string
		wantError bool
	}{
		{"1", "2", "Result: 4", false},
		{"10", "20", "Result: 30", false},
	}

	for _, tt := range tests {
		req := httptest.NewRequest("GET", "/add/"+tt.num1+"/"+tt.num2, nil)
		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/add/{num1}/{num2}", AddHandler).Methods("GET")
		router.ServeHTTP(rr, req)

		if status := rr.Code; tt.wantError {
			if status != http.StatusBadRequest {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
			}
		} else {
			if status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}
			if rr.Body.String() != tt.expected {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tt.expected)
			}
		}
	}
}

// TestStartHandler tests the StartHandler function
func TestStartHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/", StartHandler).Methods("GET")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := "Perfect"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// AddHandler handles the addition of two numbers from the URL
func AddHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	num1Str, num2Str := vars["num1"], vars["num2"]

	num1, err1 := strconv.Atoi(num1Str)
	num2, err2 := strconv.Atoi(num2Str)

	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	result := num1 + num2
	fmt.Fprintf(w, "Result: %d", result)
}

// Handler handles the addition of two numbers from the URL
func StartHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Perfect")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/add/{num1}/{num2}", AddHandler).Methods("GET")
	r.HandleFunc("/", StartHandler).Methods("GET")

	http.Handle("/", r)
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

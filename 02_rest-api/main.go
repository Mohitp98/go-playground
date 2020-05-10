package main

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	Country   string   `json:"country"`
	Languages []string `json:"languages,omitempty" binding:"required"`
}

func person(w http.ResponseWriter, req *http.Request) {

	// gathering data
	data := Person{
		Name:    "Mohit",
		Age:     23,
		Country: "India",
	}

	// Marshaling data into []byte format ie. For application layer
	body, err := json.Marshal(data)
	if err != nil {

	}

	// setting header
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Writing response body
	w.Write(body)
}

func main() {

	// Get person data
	http.HandleFunc("/person", person)

	http.ListenAndServe(":8080", nil)
}

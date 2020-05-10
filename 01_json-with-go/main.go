package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Person :
type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	Country   string   `json:"country"`
	Languages []string `json:"languages"`
}

func main() {

	var person Person
	// sample json
	body := []byte(`{"name":"Mohit","age": 23,"country": "India", "languages": ["Golang", "Python", "C++"]}`)
	fmt.Printf("Before Unmarshaling ([]byte) : %v \n", body)

	// Decoding json
	if err := json.Unmarshal(body, &person); err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Printf("After Unmarshaling (actual data) : %+v \n", person)
}

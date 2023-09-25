package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	type Person struct {
		Name       *string `json:"name"`
		Age        int     `json:"age"`
		City       string
		Occupation string
	}

	var jsonStr = `{
	"name": "Jane",
	"age": 24,
	"city": "New York",
	"occupation": "Doctor"
}`

	var p Person
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		log.Fatalf("json.Unmarshal failed with '%s'\n", err)
	}
	fmt.Printf("Person struct parsed from JSON: %#v\n", p)
	fmt.Printf("Name: %#v\n", *p.Name)

}

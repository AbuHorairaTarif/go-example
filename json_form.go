package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Bird struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does"`
	Dimension   Dimension
	CreatedAt   time.Time
}

type Dimension struct {
	Height int
	Width  int
}

func main() {
	// Unstructured JSON to MAP

	birdJson := `{
		"birds":
		{
			"pigeon":"likes to perch on rocks",
			"eagle":"bird of prey"
		},
		"animals":"none"
		}`

	var result map[string]any
	json.Unmarshal([]byte(birdJson), &result)

	birds := result["birds"].(map[string]any)
	for key, value := range birds {
		fmt.Println(key, value.(string))
	}

	// =============================  Decoding JSON to map

	// {
	// 	"birdType": "pigeon",
	// 	"what it does": "likes to perch on rocks"
	//   }

	// birdJson := `{
	// 	"birdType": "pigeon",
	// 	"what it does": "likes to perch on rocks"}`

	// var bird Bird

	// json.Unmarshal([]byte(birdJson), &bird)
	// fmt.Println(bird.Species)

	// ================================= time.Date()

	// dateJson := `"2021-10-18T11:08:47.577Z"`

	// var date time.Time
	// json.Unmarshal([]byte(dateJson), &date)
	// fmt.Println(date)

	// ========================= primitive data type manipulation =========================
	// numberJson := []byte("3")
	// floatJson := "3.1412"
	// stringJson := `"bird"`

	// var n int
	// var pi float64
	// var str string

	// json.Unmarshal((numberJson), &n)
	// fmt.Println(n)
	// // 3

	// json.Unmarshal([]byte(floatJson), &pi)
	// fmt.Println(pi)
	// // 3.1412

	// json.Unmarshal([]byte(stringJson), &str)
	// fmt.Println(str)

	// jsonBird := `
	// 	{
	// 	"species":"pigeon",
	// 	"decription":"likes to perch on rocks",
	// 	"dimension":{"height":24,"width":10}}
	// `
	// var birds Bird

	// json.Unmarshal([]byte(jsonBird), &birds)
	// fmt.Println(birds)

	// 	type Person struct {
	// 		Name       *string `json:"name"`
	// 		Age        int     `json:"age"`
	// 		City       string
	// 		Occupation string
	// 	}

	// 	var jsonStr = `{
	// 	"name": "Jane",
	// 	"age": 24,
	// 	"city": "New York",
	// 	"occupation": "Doctor"
	// }`

	// 	var p Person
	// 	err := json.Unmarshal([]byte(jsonStr), &p)
	// 	if err != nil {
	// 		log.Fatalf("json.Unmarshal failed with '%s'\n", err)
	// 	}
	// 	fmt.Printf("Person struct parsed from JSON: %#v\n", p)
	// 	fmt.Printf("Name: %#v\n", *p.Name)

}

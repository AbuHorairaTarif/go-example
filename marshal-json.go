package main

import (
	"encoding/json"
	"fmt"
)

type BirdData struct {
	Species string `json:"birdType"`
	// Description string `json:"what it does,omitempty"`
	Dimensions Dimensions
}

type Dimensions struct {
	Height int
	Width  int
}

func (d Dimensions) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%dx%d"`, d.Height, d.Width)), nil
}

func main() {

	bird := BirdData{
		Species: "Crow",
		Dimensions: Dimensions{
			Height: 23,
			Width:  10,
		},
	}

	birdData, _ := json.Marshal(bird)
	fmt.Println(string(birdData))

	// birdData := map[string]any{
	// 	"species": map[string]string{
	// 		"bird1": "pigeon",
	// 		"bird2": "crow",
	// 	},
	// 	"total bird": 2,
	// }

	// data, _ := json.Marshal(birdData)
	// fmt.Println(string(data))

	// pigeon := &BirdData{
	// 	Species:     "Pigeon",
	// 	Description: "No",
	// }

	// data, _ := json.Marshal(pigeon)
	// fmt.Println(string(data))

	// The keys need to be strings, the values can be
	// any serializable value
	// birdData := map[string]any{
	// 	"birdSounds": map[string]string{
	// 		"pigeon": "coo",
	// 		"eagle":  "squawk",
	// 	},
	// 	"total birds": 2,
	// }

	// // JSON encoding is done the same way as before
	// data, _ := json.Marshal(birdData)
	// fmt.Println(string(data))

	// json.Unmarshal([]byte(data), &n)
}

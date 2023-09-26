package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Response struct {
	Status  string
	Code    string
	Total   int
	Persons []Person `json:"data"`
}

type Person struct {
	Firstname string
	Lastname  string
	Email     string
	Phone     string
	Birthday  string
	Gender    string
	Address   Address
	Website   string
	Image     string
}

type Address struct {
	Street         string
	StreetName     string
	BuildingNumber string
	City           string
	Zipcode        string
	Country        string
	CountyCode     string `json:"county_code"`
	Latitude       float64
	Longitude      float64
}

func main() {
	myFile, err := os.Open("myjson.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("myjson.json file opened successfully !!!")

	defer myFile.Close()

	byteResult, _ := io.ReadAll(myFile)
	var response Response

	json.Unmarshal([]byte(byteResult), &response)

	for i, v := range response.Persons {
		fmt.Println("Person #", i+1, "---- Full Name :", v.Firstname, v.Lastname)
		fmt.Println(strings.Repeat("#", 40))
		fmt.Println("Gender:", v.Gender)
		fmt.Println("Email Address:", v.Email)
		fmt.Println("Phone Number:", v.Phone)
		fmt.Print("Address:", v.Address.BuildingNumber, ", ",
			v.Address.Street, ", ", v.Address.StreetName, ", ",
			v.Address.City, ", ", v.Address.Zipcode, ", ",
			v.Address.Country, "\n")
	}

}

package main

import (
	"encoding/json"
	"fmt"
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

	jsonData := `{
		"status": "OK",
		"code": 200,
		"total": 2,
		"data": [
		  {
			"id": 1,
			"firstname": "Nelle",
			"lastname": "Hilpert",
			"email": "roger51@yahoo.com",
			"phone": "+9505791682530",
			"birthday": "2014-07-17",
			"gender": "female",
			"address": {
			  "id": 0,
			  "street": "166 Pacocha Lane",
			  "streetName": "Ryan Way",
			  "buildingNumber": "47676",
			  "city": "Lake Nasir",
			  "zipcode": "75970-8502",
			  "country": "Mauritania",
			  "county_code": "LS",
			  "latitude": 86.461538,
			  "longitude": 12.052732
			},
			"website": "http://rogahn.org",
			"image": "http://placeimg.com/640/480/people"
		  },
		  {
			"id": 2,
			"firstname": "Octavia",
			"lastname": "Cummings",
			"email": "percival90@yahoo.com",
			"phone": "+3905590191171",
			"birthday": "2015-08-05",
			"gender": "female",
			"address": {
			  "id": 0,
			  "street": "5490 McCullough Valley Apt. 583",
			  "streetName": "Rose Pass",
			  "buildingNumber": "798",
			  "city": "East Keegan",
			  "zipcode": "30047-5225",
			  "country": "Burundi",
			  "county_code": "SD",
			  "latitude": -65.171837,
			  "longitude": -47.821106
			},
			"website": "http://sporer.com",
			"image": "http://placeimg.com/640/480/people"
		  }
		]
	  }`
	// https://fakerapi.it/api/v1/persons?_quantity=2&_gender=female&_birthday_start=2005-01-01
	var response Response
	json.Unmarshal([]byte(jsonData), &response)

	// fmt.Println(response)
	for i, v := range response.Persons {
		fmt.Println("Person", i+1, "Full Name :", v.Firstname, v.Lastname)
		fmt.Println(strings.Repeat("#", 40))
		fmt.Println("Gender:", v.Gender)
		fmt.Println("Email Address:", v.Email)
		fmt.Println("Phone Number:", v.Phone)
		fmt.Print("Address:", v.Address.BuildingNumber, ", ",
			v.Address.Street, ", ", v.Address.StreetName, ", ",
			v.Address.City, ", ", v.Address.Zipcode, ", ",
			v.Address.Country)
	}

}

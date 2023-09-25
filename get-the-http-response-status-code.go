package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("https://golangcode.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Println("Http Status code is in 2xx range")
	} else {
		fmt.Println("Arg Broken")
	}
}

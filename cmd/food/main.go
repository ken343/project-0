package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ken343/project-0/cmd/food/internal/place"
)

func main() {
	var name string = os.Getenv("USERNAME")
	fmt.Printf("Howdy %s, here are your options for tonight.\n\n", name)

	resp, err := http.Get(url) //url is globally located in config.go
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var myPlace place.Places = loadPlaces(resp.Body)
	place.PrintSuggestions(myPlace)
}

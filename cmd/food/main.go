package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/ken343/project-0/cmd/food/localpkg/place"
)

func main() {
	var name string

	if runtime.GOOS == "windows" {
		name = os.Getenv("USERNAME")
	} else {
		name = os.Getenv("USER")
	}
	fmt.Printf("Howdy %s, here are your options for tonight.\n\n", name)

	resp, err := http.Get(url) //url is globally located in config.go
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var myPlace place.Places = LoadPlaces(resp.Body)
	place.PrintSuggestions(myPlace)
}

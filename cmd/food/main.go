package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/ken343/project-0/localpkg/myerror"

	"github.com/ken343/project-0/localpkg/place"
)

func main() {
	var name string

	if runtime.GOOS == "windows" {
		name = os.Getenv("USERNAME")
	} else {
		name = os.Getenv("USER")
	}
	fmt.Printf("Howdy %s, here are your options for tonight.\n\n", name)

	//Query my server to get response with Place being returned as the payload.

	//debug print
	fmt.Println(url)

	resp, err := http.Get(url) //url is globally located in config.go

	myerror.Check(err)
	defer resp.Body.Close()

	var clientPlace place.Places
	json.NewDecoder(resp.Body).Decode(&clientPlace)

	place.PrintSuggestions(clientPlace)
}

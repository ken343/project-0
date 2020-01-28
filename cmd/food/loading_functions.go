package main

import (
	"encoding/json"
	"io"
	"log"

	"github.com/ken343/project-0/localpkg/place"
)

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}

}

// LoadPlaces takes in data from a reader containing json and outputs
// a place.
func LoadPlaces(myReader io.Reader) place.Places {

	myDecoder := json.NewDecoder(myReader)
	var myPlaces place.Places
	err := myDecoder.Decode(&myPlaces)
	if err != nil {
		log.Fatal(err)
	}
	return myPlaces
}

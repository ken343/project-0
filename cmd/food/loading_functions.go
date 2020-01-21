package main

import (
	"encoding/json"
	"io"
	"log"

	"github.com/ken343/project-0/cmd/food/localpkg/place"
)

func handleError(e error) {
	if e != nil {
		panic(e)
	}

}

func loadPlaces(myReader io.Reader) place.Places {

	myDecoder := json.NewDecoder(myReader)
	var myPlaces place.Places
	err := myDecoder.Decode(&myPlaces)
	if err != nil {
		log.Fatal(err)
	}
	return myPlaces
}

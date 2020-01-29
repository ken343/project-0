package main

import (
	"encoding/json"
	"io"

	"github.com/ken343/project-0/localpkg/myerror"
	"github.com/ken343/project-0/localpkg/place"
)

// LoadPlaces takes in data from a reader containing json and outputs
// a place.
func LoadPlaces(myReader io.Reader) place.Places {

	myDecoder := json.NewDecoder(myReader)
	var myPlaces place.Places
	err := myDecoder.Decode(&myPlaces)

	myerror.Check(err)

	return myPlaces
}

package main

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/ken343/project-0/cmd/food/internal/place"

	"github.com/ken343/project-0/cmd/food/internal/restaurant"
)

func handleError(e error) {
	if e != nil {
		panic(e)
	}

}
func loadRestaurants() []restaurant.Restaurant {
	var restaurants []restaurant.Restaurant
	f, err := os.Open(LISTFILE)
	handleError(err)

	fileInfo, err := f.Stat()
	handleError(err)
	fileSize := fileInfo.Size()

	var buffer []byte = make([]byte, fileSize)
	f.Read(buffer)
	handleError(err)
	json.Unmarshal(buffer, &restaurants)
	return restaurants
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

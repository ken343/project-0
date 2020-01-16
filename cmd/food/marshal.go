package main

import (
	"encoding/json"
	"os"

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

package main

import (
	"encoding/json"
	"os"
)

func handleError(e error) {
	if e != nil {
		panic(e)
	}

}
func loadRestaurants() []Restaurant {
	var restaurants []Restaurant
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

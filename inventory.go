package main

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func handleError(e error) {
	if e != nil {
		panic(e)
	}

}
func loadRestaurants(r []Restaurant) {
	f, err := os.Open(LISTFILE)
	handleError(err)

	fileInfo, err := f.Stat()
	handleError(err)
	fileSize := fileInfo.Size()

	var buffer []byte = make([]byte, fileSize)

}

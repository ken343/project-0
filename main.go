package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	RestaurantFilter(&restSlice, *pFoodType, *pMaxPrice, *pMaxDistance)
	PrintSuggestions(restSlice)

	const LISTFILE = "./restaurant_list.txt"
	f, err := os.Open(LISTFILE) //File type implements Reader!
	if err != nil {
		f.Close()
		log.Fatal(err)
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	scan.Split(bufio.ScanLines)

	var lines []string
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
	fmt.Println(lines)
}

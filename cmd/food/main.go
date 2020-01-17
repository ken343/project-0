package main

import (
	"fmt"

	"github.com/ken343/project-0/cmd/food/internal/restaurant"
)

func main() {

	var msg string = "Howdy User, here are your options for tonight."
	fmt.Println(msg + "\n")
	restSlice := loadRestaurants()

	newSlice := restaurant.Filter(restSlice, *pFoodType, *pMaxPrice, *pMaxDistance)
	restaurant.PrintSuggestions(newSlice)

	fmt.Println(newSlice)
}

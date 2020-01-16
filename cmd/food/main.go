package main

import (
	"fmt"
)

func main() {

	var msg string = "Howdy User, here are your options for tonight."
	fmt.Println(msg + "\n")
	restSlice := loadRestaurants()
	RestaurantFilter(&restSlice, *pFoodType, *pMaxPrice, *pMaxDistance)
	PrintSuggestions(restSlice)

	fmt.Println(restSlice)
}

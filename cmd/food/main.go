package main

import (
	"fmt"

	"github.com/ken343/project-0/cmd/food/internal/restaurant"
)

func main() {
	var msg string = "Howdy User, here are your options for tonight."
	fmt.Println(msg + "\n")

	restSlice := loadRestaurants()

	restSlice = RestaurantFilter(restSlice, *pFoodType, *pMaxPrice, *pMaxDistance)

	restaurant.PrintSuggestions(restSlice)
}

// RestaurantFilter will take in the command line flags and use them to pare down the array of restaurants
// down to reasonable list.
func RestaurantFilter(r []restaurant.Restaurant, food string, price float64, distance float64) []restaurant.Restaurant {
	if food != "" {
		r = restaurant.FilterCuisine(r, restaurant.EQL, food)
	}
	if price > 0 {
		r = restaurant.FilterPrice(r, restaurant.LTE, price)
	}
	if distance > 0 {
		r = restaurant.FilterDistance(r, restaurant.LTE, distance)
	}
	return r
}

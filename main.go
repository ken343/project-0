package main

import (
	"fmt"
)

func main() {

	restSlice := loadRestaurants()
	RestaurantFilter(&restSlice, *pFoodType, *pMaxPrice, *pMaxDistance)
	PrintSuggestions(restSlice)

	fmt.Println(restSlice)
}

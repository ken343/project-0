package main

func main() {
	RestaurantFilter(&restSlice, *pFoodType, *pMaxPrice, *pMaxDistance)
	PrintSuggestions(restSlice)
}

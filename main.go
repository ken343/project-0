package main

import "fmt"

func main() {
	GetRecommendation(restSlice[0])
	GetRecommendation(restSlice[1])
	fmt.Printf("Cuiseine: %s\n", *pFoodType)
	fmt.Printf("Distance: %.2f\n", *pMaxDistance)
	fmt.Printf("Price: %.2f\n", *pMaxPrice)
}

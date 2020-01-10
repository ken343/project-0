package main

import "fmt"

func main() {
	GetRecommendation(restSlice[0])
	GetRecommendation(restSlice[1])
	fmt.Printf("Cuiseine: %s\n", *pFoodType)
	fmt.Printf("Distance: %.2f\n", *pMaxDistance)
	fmt.Printf("Price: %.2f\n", *pMaxPrice)

	fmt.Println()

	for _, v := range restSlice {
		fmt.Println(v)
	}

	fmt.Println("Filter by Price")

	FilterPrice(&restSlice, 15.0)

	for i := 0; i < len(restSlice); i++ {
		fmt.Println(restSlice[i])
	}
}

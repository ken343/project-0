// Package dining is a package dedicated for representing restaurants and the
// auxilory functions that operate on them.
package main

import "fmt"

// Restaurant is a data structure that represents restaurants with a
// name, type of cuisine, maximum price, and a distance from Liv+.
type Restaurant struct {
	Name     string
	Cuisine  string
	Price    float64
	Distance float64
}

// GetRecommendation prints out restaurant information in formatted string.
func GetRecommendation(r Restaurant) {
	fmt.Printf("The %s restaurant, %s, is availabe %.1f miles away and has prices below $%.2f.\n", r.Cuisine, r.Name, r.Distance, r.Price)
}

// FilterPrice takes in a slice of restaurants and returns a new slice of restaurants that all have prices
// below the filter value.
func FilterPrice(r *[]Restaurant, filter float64) {

	surrogateSlice := make([]Restaurant, 0)

	for i := 0; i < len(*r); i++ {
		if (*r)[i].Price <= filter {
			surrogateSlice = append(surrogateSlice, (*r)[i])
		}
	}

	*r = surrogateSlice
}

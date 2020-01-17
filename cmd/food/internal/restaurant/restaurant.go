// Package restaurant is a package dedicated for representing restaurants and the
// auxilory functions that operate on them.
package restaurant

import (
	"fmt"
)

// These flags will determine what kind of filter logic will be applied when
// comparing the "compare" value to the filterable slice. These flags also apply when
// comparing types aphabetically.
const (
	LT  = "less than"                // Filter out values less than the compare value.
	LTE = "less than or equal to"    // Filter out values less than or equal to the compare value.
	GT  = "greater than"             // Filter out values greater than the compare value.
	GTE = "greater than or equal to" // Filter out values greater than or equal to the compare value.
	EQL = "equal to"                 // Filter out values equal to the compare value.
)

// Restaurant is a data structure that represents restaurants with a
// name, type of cuisine, maximum price, and a distance from Liv+.
type Restaurant struct {
	Name     string
	Cuisine  string
	Price    float64
	Distance float64
}

// FilterCuisine takes a slice of restaurants, one of the restaurant package flag constants,
// and a value for each element to be compared to.
func FilterCuisine(r []Restaurant, flag string, compare string) []Restaurant {
	newSlice := make([]Restaurant, 0)
	for i := 0; i < len(r); i++ {
		if compareString(r[i].Cuisine, compare, flag) {
			newSlice = append(newSlice, r[i])
		}
	}
	return newSlice
}

// FilterPrice takes a slice of restaurants, one of the restaurant package flag constants,
// and a value for each element to be compared to.
func FilterPrice(r []Restaurant, flag string, compare float64) []Restaurant {
	newSlice := make([]Restaurant, 0)
	for i := 0; i < len(r); i++ {
		if compareFloat64(r[i].Price, compare, flag) {
			newSlice = append(newSlice, r[i])
		}
	}
	return newSlice
}

// FilterDistance takes a slice of restaurants, one of the restaurant package flag constants,
// and a value for each element to be compared to.
func FilterDistance(r []Restaurant, flag string, compare float64) []Restaurant {
	newSlice := make([]Restaurant, 0)
	for i := 0; i < len(r); i++ {
		if compareFloat64(r[i].Distance, compare, flag) {
			newSlice = append(newSlice, r[i])
		}
	}
	return newSlice
}

// GetRecommendation prints out restaurant information in formatted string.
func GetRecommendation(r Restaurant) {
	fmt.Printf("The %s restaurant, %s, is availabe %.1f miles away and has prices below $%.2f.\n", r.Cuisine, r.Name, r.Distance, r.Price)
}

// PrintSuggestions will iterate through a slice of restaurants and output the formatted
// strings produced by the GetRecommendation() function.
func PrintSuggestions(r []Restaurant) {
	for i, v := range r {
		fmt.Printf("%d) ", i+1)
		GetRecommendation(v)
	}
}

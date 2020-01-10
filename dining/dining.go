// Package dining is a package dedicated for representing restaurants and the
//auxilory functions that operate on them.
package dining

import "fmt"

// Restaurant is a data structure that represents restaurants with a
//name, type of cuisine, maximum price, and a distance from Liv+.
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

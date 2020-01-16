// Package restaurant is a package dedicated for representing restaurants and the
// auxilory functions that operate on them.
package restaurant

import (
	"fmt"

	"github.com/ken343/project-0/pkg/filter"
)

// These flags are intended to be used with the filter package methods,
// SetOrdinalOption and SetLexicalOption.
const (
	NAME     = "name"     // Filter will use Restaurant Name for lexical filtering.
	CUISINE  = "cuisine"  // Filter will use Restaurant Cuisine for lexical filtering.
	PRICE    = "price"    // Filter will use Restaurant Price for ordinal filtering.
	DISTANCE = "distance" // Filter will use Restaurant Distance for ordinal filtering.
)

// Restaurant is a data structure that represents restaurants with a
// name, type of cuisine, maximum price, and a distance from Liv+.
type Restaurant struct {
	Name          string
	Cuisine       string
	Price         float64
	Distance      float64
	ordinalOption string
	lexicalOption string
}

// SetOrdinalOption uses one of flags provided (e.g Restaurnat.PRICE ) to set which
// numerical field will be used by the Restaurant struct for determining filter behavior.
func (r Restaurant) SetOrdinalOption(option string) {
	r.ordinalOption = option
}

// SetLexicalOption uses one of flags provided (e.g Restaurnat.CUISINE) to set which
// alphabetical field will be used by the Restaurant struct for determiing filter behavior.
func (r Restaurant) SetLexicalOption(option string) {
	r.lexicalOption = option
}

// OrdinalVal returns the struct value set by the SetOrdinalOption method.
// It is used internally to implent the behavior of filter.Ordinal()
func (r Restaurant) OrdinalVal() float64 {
	switch r.ordinalOption {
	case PRICE:
		return r.Price
	case DISTANCE:
		return r.Distance
	default:
		panic("Need a proper Restaurant ordinal option.")
	}
}

// LexicalVal returns the struct value set by the SetLexicalOption method.
// It is used internally to implent the behavior of filter.Lexical()
func (r Restaurant) LexicalVal() string {
	switch r.lexicalOption {
	case NAME:
		return r.Name
	case CUISINE:
		return r.Cuisine
	default:
		panic("Need a proper Restaurant lexical option")
	}
}

// GetRecommendation prints out restaurant information in formatted string.
func GetRecommendation(r Restaurant) {
	fmt.Printf("The %s restaurant, %s, is availabe %.1f miles away and has prices below $%.2f.\n", r.Cuisine, r.Name, r.Distance, r.Price)
}

// Filter will pare down the input slice of Restaurants structs to only include the
// restaurants with a certain type of food and limits on price/distance.
func Filter(r []Restaurant, foodType string, maxPrice float64, maxDistance float64) []Restaurant {
	//only executes filter functions if flags are set for that value
	//TO-DO turn this []Restaurant to []Filterable logic into its own function!
	var fs []filter.Filterable = make([]filter.Filterable, 0)
	for _, v := range r {
		fs = append(fs, v)
	}

	if foodType != "" {
		for _, v := range fs {
			v.SetLexicalOption(CUISINE)
		}
		fs = filter.Lexical(fs, filter.EQL, foodType)
	}
	if maxPrice > 0 {
		for _, v := range fs {
			v.SetOrdinalOption(PRICE)
		}
		fs = filter.Ordinal(fs, filter.LTE, maxPrice)
	}
	if maxDistance > 0 {
		for _, v := range fs {
			v.SetOrdinalOption(DISTANCE)
		}
		fs = filter.Ordinal(fs, filter.LTE, maxDistance)
	}

	helpMe := make([]Restaurant, 0)
	for _, v := range fs {
		helpMe = append(helpMe, v.(Restaurant))
	}
	return helpMe
}

// PrintSuggestions will iterate through a slice of restaurants and output the formatted
// strings produced by the GetRecommendation() function
func PrintSuggestions(r []Restaurant) {
	for i, v := range r {
		fmt.Printf("%d) ", i+1)
		GetRecommendation(v)
	}
}

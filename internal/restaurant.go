// Package restaurant is a package dedicated for representing restaurants and the
// auxilory functions that operate on them.
package restaurant

import "fmt"

const (
	NAME     = "name"
	CUISINE  = "cuisine"
	PRICE    = "price"
	DISTANCE = "distance"
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

func (r *Restaurant) SetOrdinalOption(option string) {
	r.ordinalOption = option
}

func (r *Restaurant) SetLexicalOption(option string) {
	r.lexicalOption = option
}

func (r *Restaurant) OrdinalVal() float64 {
	switch r.ordinalOption {
	case PRICE:
		return r.Price
	case DISTANCE:
		return r.Distance
	default:
		panic("Need a proper Restaurant ordinal option.")
	}
}

func (r *Restaurant) LexicalVal() string {
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

// FilterPrice pares down the input slice of Restaurant structs to only include restaurants within
// the filtered maximum price.
func FilterPrice(r *[]Restaurant, filter float64) {

	surrogateSlice := make([]Restaurant, 0)

	for i := 0; i < len(*r); i++ {
		if (*r)[i].Price <= filter {
			surrogateSlice = append(surrogateSlice, (*r)[i])
		}
	}

	*r = surrogateSlice
}

// FilterFood pares down the input slice of Restaurant structs to only include restaurants within
// the filtered type of cuisine.
func FilterFood(r *[]Restaurant, filter string) {

	surrogateSlice := make([]Restaurant, 0)

	for i := 0; i < len(*r); i++ {
		if (*r)[i].Cuisine == filter {
			surrogateSlice = append(surrogateSlice, (*r)[i])
		}
	}

	*r = surrogateSlice
}

// FilterDistance pares down the input slice of Restaurant structs to only include restaurants within
// the filtered distance.
func FilterDistance(r *[]Restaurant, filter float64) {

	surrogateSlice := make([]Restaurant, 0)

	for i := 0; i < len(*r); i++ {
		if (*r)[i].Distance <= filter {
			surrogateSlice = append(surrogateSlice, (*r)[i])
		}
	}

	*r = surrogateSlice
}

// RestaurantFilter will pare down the input slice of Restaurants structs to only include the
// restaurants with a certain type of food and limits on price/distance.
func RestaurantFilter(r *[]Restaurant, foodType string, maxPrice float64, maxDistance float64) {
	//only executes filter functions if flags are set for that value
	if foodType != "" {
		FilterFood(r, foodType)
	}
	if maxPrice > 0 {
		FilterPrice(r, maxPrice)
	}
	if maxDistance > 0 {
		FilterDistance(r, maxDistance)
	}
}

// PrintSuggestions will iterate through a slice of restaurants and output the formatted
// strings produced by the GetRecommendation() function
func PrintSuggestions(r []Restaurant) {
	for i, v := range r {
		fmt.Printf("%d) ", i+1)
		GetRecommendation(v)
	}
}

package dining

import "fmt"

// STRUCT SECTION
type Restaurant struct {
	Name     string
	Cuisine  string
	Price    float64
	Distance float64
}

// Prints out restaurant information in formatted string.
func GetRecommendation(r Restaurant) {
	fmt.Printf("The %s restaurant, %s, is availabe %.1f miles away and has prices below $%.2f.", r.Cuisine, r.Name, r.Distance, r.Price)
}

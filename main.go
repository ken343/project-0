package main

import "fmt"

//STRUCT SECTION

type restaurant struct {
	name     string
	cuisine  string
	price    float64
	distance float64
}

//FUNCTION SECTION

//Prints out restaurant information in formatted string.
func printrestaurant(r restaurant) {
	fmt.Printf("The %s restaurant, %s, is availabe %.1f miles away and has prices below $%.2f.", r.cuisine, r.name, r.distance, r.price)
}

//VAR DECLARATION SECTION

var mcd restaurant

//PROGRAM
func init() {
	mcd = restaurant{
		name:     "McDonalds",
		cuisine:  "fast-food",
		price:    10,
		distance: 5,
	}
}

func main() {
	printrestaurant(mcd)
}

package main

import "flag"

// LISTFILE provides the name of the file where all the restaurant json data is located.
const LISTFILE = "./restaurant_list.txt"

var restSlice []Restaurant

var pFoodType *string
var pMaxPrice *float64
var pMaxDistance *float64

func init() {
	// Set up slice of
	restSlice = append(restSlice, Restaurant{
		Name:     "McDonalds",
		Cuisine:  "fast-food",
		Price:    10,
		Distance: 5,
	})

	restSlice = append(restSlice, Restaurant{
		Name:     "Pot Belly's",
		Cuisine:  "sandwich",
		Price:    20,
		Distance: 1.1,
	})

	restSlice = append(restSlice, Restaurant{
		Name:     "Red Lobster",
		Cuisine:  "sea-food",
		Price:    30,
		Distance: 20.5,
	})

	restSlice = append(restSlice, Restaurant{
		Name:     "Eddie V's",
		Cuisine:  "fancy",
		Price:    50,
		Distance: 10.5,
	})

	// Set up flags for use in filter function
	pFoodType = flag.String("type", "", "This flag will determine what kind of food you are interested in.")
	pMaxPrice = flag.Float64("price", 0.0, "This flag will represent the most you are willing to pay for a meal.")
	pMaxDistance = flag.Float64("distance", 0.0, "This flag will set the max distance you are willing to drive")
	flag.Parse()

}

package main

import "flag"

var restSlice []Restaurant

var pFoodType *string
var pMaxPrice *float64
var pMaxDistance *float64

func init() {
	// Set up array of 
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

	// Set up flags for use in filter function
	pFoodType = flag.String("type", "", "This flag will determine what kind of food you are interested in.")
	pMaxPrice = flag.Float64("price", 0.0, "This flag will represent the most you are willing to pay for a meal.")
	pMaxDistance = flag.Float64("distance", 0.0, "This flag will set the max distance you are willing to drive")
	flag.Parse()

}

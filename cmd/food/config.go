package main

import (
	"flag"

	"github.com/ken343/project-0/cmd/food/localpkg/place"
)

var pMaxPrice *float64
var pFoodType *string
var pMaxDistance *float64

var url string

func init() {
	// Set up flags for use in filter function
	pFoodType = flag.String("type", "", "This flag will determine what kind of food you are interested in.")
	pMaxPrice = flag.Float64("price", 0.0, "This flag will represent the most you are willing to pay for a meal.")
	pMaxDistance = flag.Float64("distance", 0.0, "This flag will set the max distance you are willing to drive")
	flag.Parse()

	url = place.ProduceQueryString(*pFoodType, *pMaxDistance, *pMaxPrice)
}

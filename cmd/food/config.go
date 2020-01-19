package main

import (
	"flag"
	"log"

	"github.com/ken343/project-0/cmd/food/internal/place"
)

// LISTFILE provides the name of the file where all the restaurant json data is located.
const LISTFILE = "C:/Users/kmcdo/go/src/github.com/ken343/project-0/cmd/food/internal/restaurant/restaurant_list.json"

// These values are equivalent to what google recognizes price level in there "maxprice" query.
const (
	EXTRA  = 4 // Prices $30 dollars and up
	HIGH   = 3 // Prices below $30 dollars
	MEDIUM = 2 // Prices below $20 dollars
	LOW    = 1 // Prices below $10 dollars
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

	radius := *pMaxDistance * 1609.34 // 1 mile == 1609.34 meters
	var priceLevel int64
	if *pMaxPrice == 0 || *pMaxPrice >= 30 {
		priceLevel = EXTRA
	} else if *pMaxPrice < 10 {
		priceLevel = LOW
	} else if *pMaxPrice < 20 {
		priceLevel = MEDIUM
	} else {
		log.Fatalf("Max Price not set to proper valid number: %d", *pMaxPrice)
	}

	url = place.ProduceQueryString(query, radius, priceLevel)
}

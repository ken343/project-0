package main

import (
	"flag"
	"fmt"
	"log"
)

var pMaxPrice *float64
var pFoodType *string
var pMaxDistance *float64

var url string

// These values are equivalent to what google recognizes as price level in their "maxprice" query.
const (
	EXTRA  = 4 // Prices $30 dollars and up
	HIGH   = 3 // Prices below $30 dollars
	MEDIUM = 2 // Prices below $20 dollars
	LOW    = 1 // Prices below $10 dollars
)

// Keys for creating query string.
const (
	URL           = "https://maps.googleapis.com/maps/api/place/textsearch/json?query=" // Core URL for API
	DEFAULTQUERY  = ""                                                                  // type of 'place' that the API should look for
	LOCATION      = "32.727220,-97.105940"                                              // Latitude and Longitude of Latitude                                                            // About 12 mile search radius
	DEFAULTRADIUS = 45000                                                               // About 28 miles around liv+
)

func init() {
	// Set up flags for use in filter function
	pFoodType = flag.String("type", "", "This flag will determine what kind of food you are interested in.")
	pMaxPrice = flag.Float64("price", 0.0, "This flag will represent the most you are willing to pay for a meal.")
	pMaxDistance = flag.Float64("distance", 0.0, "This flag will set the max distance you are willing to drive")
	flag.Parse()

	var clientQuery string = constructQuery(*pFoodType, *pMaxDistance, *pMaxPrice)
	url = "http//localhost:8080" + clientQuery

}

// Configure call to food server
func constructQuery(query string, distance float64, prices float64) string {

	// Make sure that each vaue is equivalent to some sensible defaults
	// assuming that the user did not set a value for the CLI flags.
	var stringQuery string
	if query != "" {
		stringQuery = query + "+"
	} else {
		stringQuery = DEFAULTQUERY
	}

	var radius float64
	if distance != 0 {
		radius = distance * 1609.34
	} else {
		radius = DEFAULTRADIUS * 1609.34 // 1 mile == 1609.34 meters
	}

	var priceLevel int64
	if prices == 0 || prices >= 30 {
		priceLevel = EXTRA
	} else if prices < 10 {
		priceLevel = LOW
	} else if prices < 20 {
		priceLevel = MEDIUM
	} else {
		log.Fatalf("Max Price not set to proper valid number: %f", prices)
	}

	//Debug Query String
	msg := fmt.Sprintf("?food=%s&distance=%f&pricelevel=%d", stringQuery, radius, priceLevel)
	return msg
}

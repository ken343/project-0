package place

import (
	"fmt"
	"log"
)

// These values are equivalent to what google recognizes price level in there "maxprice" query.
const (
	EXTRA  = 4 // Prices $30 dollars and up
	HIGH   = 3 // Prices below $30 dollars
	MEDIUM = 2 // Prices below $20 dollars
	LOW    = 1 // Prices below $10 dollars
)

// Keys for creating query string.
const (
	URL           = "https://maps.googleapis.com/maps/api/place/textsearch/json?query=" // Core URL for API
	KEY           = "AIzaSyA87nRXQgzR60DLwDe2jvPzDK56_4JWAN8"                           // Googe Place API Key
	DEFAULTQUERY  = "food"                                                              // type of 'place' that the API should look for
	LOCATION      = "32.727220,-97.105940"                                              // Latitude and Longitude of Latitude                                                            // About 12 mile search radius
	DEFAULTRADIUS = "20000"                                                             // About 12 miles around liv+
)

// ProduceQueryString function assembles an url to be used with an HTTP GET method.
// The url will look for chow within a custom radius around liv+.
func ProduceQueryString(query string, distance float64, prices float64) string {

	var stringQuery string
	if stringQuery == "" {
		stringQuery = DEFAULTQUERY
	} else {
		stringQuery = query
	}

	radius := distance * 1609.34 // 1 mile == 1609.34 meters

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

	return fmt.Sprintf("%s%s&location=%s&radius=%f&maxprice=%d&key=%s", URL, stringQuery, LOCATION, radius, priceLevel, KEY)
}

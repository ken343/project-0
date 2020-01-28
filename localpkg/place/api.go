package place

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

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

// ProduceQueryString function assembles a url to be used with an HTTP GET method.
// The url will look for chow within a custom radius around liv+.
func ProduceQueryString(query string, distance float64, prices float64) string {

	KEY := ProduceKey("./api-key.txt")

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
	msg := fmt.Sprintf("%s%sfood&location=%s&radius=%f&maxprice=%d&key=%s", URL, stringQuery, LOCATION, radius, priceLevel, KEY)
	return msg
}

// ProduceKey will spawn the api key necessary for the running the program.
// The key is hidden in a text file so that it remain off of github.
func ProduceKey(filepath string) string {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("My file %s did not open: %w", filepath, err)
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	size := fileInfo.Size()

	var keyBytes []byte = make([]byte, size)

	keyBytes, err = ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return string(keyBytes)
}

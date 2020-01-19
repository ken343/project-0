package place

import (
	"fmt"
	"strconv"
)

// Keys for creating query string.
const (
	URL           = "https://maps.googleapis.com/maps/api/place/textsearch/json?query=" // Core URL for API
	KEY           = "AIzaSyA87nRXQgzR60DLwDe2jvPzDK56_4JWAN8"                           // Googe Place API Key
	QUERY         = "food"                                                              // type of 'place' that the API should look for
	LOCATION      = "32.727220,-97.105940"                                              // Latitude and Longitude of Latitude                                                            // About 12 mile search radius
	DEFAULTRADIUS = "20000"                                                             // About 12 miles around liv+
)

// ProduceQueryString function assembles an url to be used with an HTTP GET method.
// The url will look for chow within a custom radius around liv+.
func ProduceQueryString(radius float64) string {
	var stringRadius string
	if radius == 0 {
		stringRadius = DEFAULTRADIUS
	} else {
		stringRadius = strconv.FormatFloat(radius, 'f', 1, 64)
	}
	return fmt.Sprintf("%s%s&location=%s&radius=%s&key=%s", URL, QUERY, LOCATION, stringRadius, KEY)
}

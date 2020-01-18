package place

import "fmt"

// Keys for creating query string.
const (
	URL      = "https://maps.googleapis.com/maps/api/place/textsearch/json?query=" // Core URL for API
	KEY      = "AIzaSyA87nRXQgzR60DLwDe2jvPzDK56_4JWAN8"                           // Googe Place API Key
	QUERY    = "food"                                                              // type of 'place' that the API should look for
	LOCATION = "32.727220,-97.105940"                                              // Latitude and Longitude of Liv+
	RADIUS   = "20000"                                                             // About 12 mile search radius
)

// PrintKey tests function to test concatanation of api key.
func PrintKey() {
	fmt.Printf("%s%s&location=%s&radius=%s&key=%s", URL, QUERY, LOCATION, RADIUS, KEY)
}

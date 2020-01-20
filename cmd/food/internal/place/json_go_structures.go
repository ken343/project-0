//    places, err := UnmarshalPlaces(bytes)
//    bytes, err = places.Marshal()

package place

import (
	"encoding/json"
	"fmt"
)

const (
	american  = "american"
	barbecue  = "barbecue"
	chinese   = "chinese"
	french    = "french"
	hamburger = "hamburger"
	italian   = "italian"
	japanese  = "japanese"
	mexican   = "mexican"
	pizza     = "pizza"
	seafood   = "seafood"
	steak     = "steak"
	sushi     = "sushi"
	thai      = "thai"
)

var cuisines []string = []string{american, barbecue, chinese, french, hamburger, italian, japanese, mexican, pizza, seafood, steak, sushi, thai}

// UnmarshalPlaces was the original way to take teh api data.
// Will likely delete in the future.
func UnmarshalPlaces(data []byte) (Places, error) {
	var r Places
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal was original way to take send back data in api format.
// Will likely delete in the future.
func (r *Places) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Places wraps around the api's retrurned results.
// only the results array was useful.
type Places struct {
	Results []Result `json:"results"`
}

// Result contains the important restaurant data.
type Result struct {
	Name             string        `json:"name"`
	FormattedAddress string        `json:"formatted_address"`
	OpeningHours     *OpeningHours `json:"opening_hours,omitempty"`
	PriceLevel       *int64        `json:"price_level,omitempty"`
	Types            []string      `json:"types"`
}

// OpeningHours may be used in outpout to indicate if a place is open or not.
type OpeningHours struct {
	OpenNow bool `json:"open_now"`
}

func (r Result) String() string {
	var isOpen string
	if r.OpeningHours.OpenNow {
		isOpen = "Yes"
	} else {
		isOpen = "No"
	}
	return fmt.Sprintf("%s:\n[Address: %s | Price: %d | Open? %s]\n",
		r.Name, r.FormattedAddress, *r.PriceLevel, isOpen)
}

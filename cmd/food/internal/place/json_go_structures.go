package place

import (
	"fmt"
	"log"
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

	var priceText string
	switch *r.PriceLevel {
	case 1:
		priceText = "Low"
	case 2:
		priceText = "Medium"
	case 3:
		priceText = "High"
	case 4:
		priceText = "Extra-High"
	default:
		log.Fatal("Line 64: String() method - Invalid r.PriceLevel")
	}

	return fmt.Sprintf("%s:\nAddress: %s\nPrice: %s | Open? %s\n",
		r.Name, r.FormattedAddress, priceText, isOpen)
}

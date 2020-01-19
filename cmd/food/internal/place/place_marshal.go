//    places, err := UnmarshalPlaces(bytes)
//    bytes, err = places.Marshal()

package place

import "encoding/json"

func UnmarshalPlaces(data []byte) (Places, error) {
	var r Places
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Places) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Places struct {
	HTMLAttributions []interface{} `json:"html_attributions"`
	NextPageToken    string        `json:"next_page_token"`
	Results          []Result      `json:"results"`
	Status           string        `json:"status"`
}

type Result struct {
	FormattedAddress string        `json:"formatted_address"`
	Geometry         Geometry      `json:"geometry"`
	Icon             string        `json:"icon"`
	ID               string        `json:"id"`
	Name             string        `json:"name"`
	OpeningHours     *OpeningHours `json:"opening_hours,omitempty"`
	Photos           []Photo       `json:"photos"`
	PlaceID          string        `json:"place_id"`
	PlusCode         PlusCode      `json:"plus_code"`
	PriceLevel       *int64        `json:"price_level,omitempty"`
	Rating           float64       `json:"rating"`
	Reference        string        `json:"reference"`
	Types            []string      `json:"types"`
	UserRatingsTotal int64         `json:"user_ratings_total"`
}

type Geometry struct {
	Location Location `json:"location"`
	Viewport Viewport `json:"viewport"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Viewport struct {
	Northeast Location `json:"northeast"`
	Southwest Location `json:"southwest"`
}

type OpeningHours struct {
	OpenNow bool `json:"open_now"`
}

type Photo struct {
	Height           int64    `json:"height"`
	HTMLAttributions []string `json:"html_attributions"`
	PhotoReference   string   `json:"photo_reference"`
	Width            int64    `json:"width"`
}

type PlusCode struct {
	CompoundCode string `json:"compound_code"`
	GlobalCode   string `json:"global_code"`
}

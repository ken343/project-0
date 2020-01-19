package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ken343/project-0/cmd/food/internal/place"
	"github.com/ken343/project-0/cmd/food/internal/restaurant"
)

func main() {
	var name string = os.Getenv("USERNAME")
	fmt.Printf("Howdy %s, here are your options for tonight.\n\n", name)

	resp, err := http.Get(url) //url is globally located in config.go
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var myPlace place.Places = loadPlaces(resp.Body)
	//myPlace = RestaurantFilter(myPlace, *pFoodType, *pMaxPrice, *pMaxDistance)
	fmt.Println(myPlace)
}

// RestaurantFilter will take in the command line flags and use them to pare down the array of restaurants
// down to reasonable list.
func RestaurantFilter(r []restaurant.Restaurant, food string, price float64, distance float64) []restaurant.Restaurant {
	if food != "" {
		r = restaurant.FilterCuisine(r, restaurant.EQL, food)
	}
	if price > 0 {
		r = restaurant.FilterPrice(r, restaurant.LTE, price)
	}
	if distance > 0 {
		r = restaurant.FilterDistance(r, restaurant.LTE, distance)
	}
	return r
}

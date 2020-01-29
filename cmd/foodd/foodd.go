package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ken343/project-0/localpkg/myerror"
	"github.com/ken343/project-0/localpkg/place"
)

func main() {
	fmt.Println("Server is running on port 8080...")

	http.HandleFunc("/chow", func(w http.ResponseWriter, r *http.Request) {
		//Retrieve food type, distance and price from client application.

		r.ParseForm()
		myValues := r.URL.Query()

		fmt.Println("Passed in URL: ", r.URL.RawQuery)
		// Form.Get() is giving hard errors for distance and pricelevel.
		// foodType := r.Form.Get("food")
		// foodDistance := r.Form.Get("distance")
		// foodPrice := r.Form.Get("pricelevel")
		foodType := myValues.Get("food")
		foodDistance := myValues.Get("distance")
		foodPrice := myValues.Get("pricelevel")
		fmt.Println("Before string conversions: foodType=", foodType, "foodDistance=", foodDistance, "foodPrice=", foodPrice)
		floatFoodDistance, err := strconv.ParseFloat("11.11", 64)
		myerror.Check(err)
		floatFoodPrice, err := strconv.ParseFloat("30.2", 64)
		myerror.Check(err)

		//debug print
		// Retrieve Place from API
		fmt.Println("After Parse Form", foodType, foodDistance, foodPrice)
		url := place.ProduceQueryString(foodType, floatFoodDistance, floatFoodPrice)
		resp, err := http.Get(url) //url is globally located in config.go
		myerror.Check(err)
		defer resp.Body.Close()

		// Decode response and "Re-code it back to the client"
		var myPlace place.Places = LoadPlaces(resp.Body)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(myPlace)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Listen & Serve: ", err)
	}
}

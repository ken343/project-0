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

		fmt.Println("Passed in URL: ", r.URL.RawQuery)
		//Form.Get() is giving hard errors for distance and pricelevel.
		foodType := r.Form.Get("food")
		foodDistance := r.Form.Get("distance")
		foodPrice := r.Form.Get("pricelevel")
		r.ParseForm()

		// Convert Strings to float64s
		floatDistance, err := strconv.ParseFloat(foodDistance, 64)
		myerror.Check(err)
		floatPrice, err := strconv.ParseFloat(foodPrice, 64)
		myerror.Check(err)

		//debug print
		// Retrieve Place from API
		fmt.Println("After Parse Form", foodType, foodDistance, foodPrice)
		url := place.ProduceQueryString(foodType, floatDistance, floatPrice)
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

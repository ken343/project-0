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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Retrieve food type, distance and price from client application.

		r.ParseForm()
		foodType := r.Form.Get("food")
		foodDistance := r.Form.Get("distance")
		foodPrice := r.Form.Get("pricelevel")
		floatFoodDistance, err := strconv.ParseFloat(foodDistance, 64)
		myerror.Check(err)
		floatFoodPrice, err := strconv.ParseFloat(foodPrice, 64)
		myerror.Check(err)

		//debug print
		fmt.Println(foodType, foodDistance, foodPrice)
		url := place.ProduceQueryString(foodType, floatFoodDistance, floatFoodPrice)
		resp, err := http.Get(url) //url is globally located in config.go
		myerror.Check(err)
		defer resp.Body.Close()

		var myPlace place.Places = LoadPlaces(resp.Body)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(myPlace)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Listen & Serve: ", err)
	}
}

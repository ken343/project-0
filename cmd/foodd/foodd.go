package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ken343/project-0/localpkg/place"
)

func main() {
	fmt.Println("Server is running on port 8080...")

	http.HandleFunc("/chow", func(w http.ResponseWriter, r *http.Request) {
		//Retrieve food type, distance and price from client application.
		r.Form.
			url = place.ProduceQueryString(*pFoodType, *pMaxDistance, *pMaxPrice)

	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Listen & Serve: ", err)
	}
}

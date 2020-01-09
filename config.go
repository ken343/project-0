package main

import "github.com/ken343/project-0/dining"

var mcd dining.Restaurant
var restSlice []dining.Restaurant

func init() {
	restSlice = append(restSlice, dining.Restaurant{
		Name:     "McDonalds",
		Cuisine:  "fast-food",
		Price:    10,
		Distance: 5,
	})

	restSlice = append(restSlice, dining.Restaurant{
		Name:     "Pot Belly's",
		Cuisine:  "Sandwiches",
		Price:    20,
		Distance: 1.1,
	})
}

package main

import "github.com/ken343/project-0/dining"

var mcd dining.Restaurant

func init() {
	mcd = dining.Restaurant{
		Name:     "McDonalds",
		Cuisine:  "fast-food",
		Price:    10,
		Distance: 5,
	}
}

package main

import "fmt"

type restaurant struct {
	name     string
	cuisine  string
	price    float64
	distance float64
}

//Prints in format: [McDonalds: Fast Food | Under $10 | Under 5 miles]
func printR(r restaurant) {
	fmt.Println(r.name, r.cuisine, r.price, r.distance)
}

func main() {

	var mcd restaurant = restaurant{
		name:     "McDonalds",
		cuisine:  "fast-food",
		price:    10,
		distance: 5,
	}

	printR(mcd)

	var instantRestuant string = "[McDona lds: Fast Food | Under $10 | Under 5 miles]"
	fmt.Println(instantRestuant)
}

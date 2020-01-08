package main

import "fmt"

func init() {
	type Resturant struct {
		name     string
		cuisine  string
		price    float64
		distance float64
	}
}

func main() {
	var instantRestuant string = "[McDonalds: Fast Food | Under $10 | Under 5 miles]"
	fmt.Println(instantRestuant)
}

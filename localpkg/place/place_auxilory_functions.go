package place

import "fmt"

// PrintSuggestions will produce all the results in a ordered list
// for the user to view restaurant results.
func PrintSuggestions(p Places) {
	for i, v := range p.Results {
		fmt.Println("\n-------------------------------------------------------------")
		fmt.Printf("%d) %v\n", i+1, v)
	}
}

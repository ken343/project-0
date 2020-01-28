package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server is running on port 8080...")

	http.HandleFunc("/chow", func(w http.ResponseWriter, r *http.Request) {

	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Listen & Serve: ", err)
	}
}

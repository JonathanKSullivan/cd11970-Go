package main

import (
	"fmt"
	"net/http"
)

var cities = []string{"Tokyo", "Delhi", "Shanghai", "Sao Paulo", "Mexico City"}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func cityList(w http.ResponseWriter, r *http.Request) {
	response := "List of most populous cities.\n"

	for i, city := range cities {
		response += fmt.Sprintf("%d. %s\n", i+1, city)
	}

	fmt.Fprintf(w, response)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/cities", cityList)

	http.ListenAndServe(":3000", nil)
}

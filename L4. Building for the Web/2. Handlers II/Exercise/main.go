package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var cityPopulations = map[string]uint32{
	"Tokyo":       37435191,
	"Delhi":       29399141,
	"Shanghai":    26317104,
	"Sao Paulo":   21846507,
	"Mexico City": 21671908,
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	response := "<h1>City Populations</h1>"

	for city, population := range cityPopulations {
		response += fmt.Sprintf("<h2>%s: %s</h2>", city, strconv.Itoa(int(population)))
	}

	fmt.Fprint(w, response)
}

func main() {
	http.HandleFunc("/", index)

	fmt.Println("Server is running on port :3000")
	http.ListenAndServe(":3000", nil)
}

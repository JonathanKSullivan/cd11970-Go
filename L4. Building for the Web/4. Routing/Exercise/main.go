package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var members = map[string]string{
	"1": "Andy",
	"2": "Peter",
	"3": "Gabriella",
	"4": "Jordy",
}

func getMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(members)
}

func deleteMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var id int = 0
	// Read the HTTP request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	// Encode the request body into a Golang value so that we can work with the data
	json.Unmarshal(reqBody, &id)

	if _, ok := members[string(id)]; ok {
		delete(members, string(id))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(members)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Member not found")
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/members", getMembers).Methods("GET")
	router.HandleFunc("/members", deleteMember).Methods("DELETE")

	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", router)
}

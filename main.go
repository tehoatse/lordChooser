package main

import (
	"fmt"
	"log"
	"net/http"
)

var lords []Lord

func handleRequests() {
	http.HandleFunc("/", apiResponse)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	lords = createLords()
	handleRequests()
}

func apiResponse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"hello world!"}`))
	fmt.Println(getRandomLord(lords))
}

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the homepage")
// 	fmt.Println("Endpoint Hit: homePage")
// }

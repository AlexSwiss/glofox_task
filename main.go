package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Our class model
type class struct {
	Name      string `json:"name"`
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
	Capacity  int    `json: "capacty"`
}

type allClasses []class

var classes = allClasses{
	{
		Name:      "Pilate",
		StartDate: "21 Jun 20",
		EndDate:   "02 Dec 20",
		Capacity:  10,
	},
}

func main() {

	// Initialize mux router
	router := mux.NewRouter().StrictSlash(true)

	// API endpoints
	router.HandleFunc("/classes", createClass)
	router.HandleFunc("/allclasses", getClasses)
	router.HandleFunc("/allclasses/{classID}", getClass)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// create new classes
func createClass(w http.ResponseWriter, r *http.Request) {

}

// Get all classes
func getClasses(w http.ResponseWriter, r *http.Request) {

}

// Get single class
func getClass(w http.ResponseWriter, r *http.Request) {

}

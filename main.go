package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Our class model
type class struct {
	ID        string `json:"ID"`
	Name      string `json:"name"`
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
	Capacity  int    `json: "capacty"`
}

type allClasses []class

var classes = allClasses{
	{
		ID:        "1",
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
	router.HandleFunc("/allclasses/{id}", getClass)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// create new classes
func createClass(w http.ResponseWriter, r *http.Request) {
	var newClass class
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "Kindly enter values with ClassID, Name, StartDate, EndDate and Capacity only")
	}
	json.Unmarshal(reqBody, &newClass)
	classes = append(classes, newClass)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newClass)
}

// Get all classes
func getClasses(w http.ResponseWriter, r *http.Request) {

}

// Get single class
func getClass(w http.ResponseWriter, r *http.Request) {
	classID := mux.Vars(r)["id"]

	for _, singleClass := range classes {
		if singleClass.ID == classID {
			json.NewEncoder(w).Encode(singleClass)
		}
	}
}

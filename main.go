package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/classes", Classes)
	router.HandleFunc("/classes/{classId}", Class)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// Get all classes
func Classes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "view all classes")
}

// Get single class
func Class(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "view single class")
}

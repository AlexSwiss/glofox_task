package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.POST("/classes", addClass)         // Create new class endpoint
	router.GET("/classes_list", listClass)    // Get all classes endpoint
	router.POST("/bookings", addNewBooking)   // add a new Booking
	router.GET("/bookings_list", allBookings) // Get all Bookings
	port := ":8080"
	fmt.Println("server started on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

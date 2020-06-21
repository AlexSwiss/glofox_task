package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.POST("/api/classes", addClass)         // Create new class endpoint
	router.GET("/api/classes_list", listClass)    // Get all classes endpoint
	router.POST("/api/bookings", addNewBooking)   // add a new Booking
	router.GET("/api/bookings_list", allBookings) // Get all Bookings
	port := ":8080"
	fmt.Println("server started on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

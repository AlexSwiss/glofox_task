package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Booking Model for our data
type Booking struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
}

var bookings = []Booking{}

// Function to make a new Booking to the studio
func addNewBooking(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		failureResponse(w, err.Error())
		return
	}

	type newBookingRequestBody struct {
		Name *string `json:"name"`
		Date *string `json:"date"`
	}

	var bookingBody newBookingRequestBody

	err = json.Unmarshal(body, &bookingBody)
	if err != nil {
		failureResponse(w, err.Error())
		return
	}

	if bookingBody.Name == nil {
		failureResponse(w, `name parameter is not specified or empty`)
		return
	}

	if bookingBody.Date == nil {
		failureResponse(w, `Date parameter is not specified or empty`)
		return
	}

	addBooking := Booking{
		ID:   len(classes) + 1,
		Name: *bookingBody.Name,
		Date: *bookingBody.Date,
	}

	bookings = append(bookings, addBooking)

	type Output struct {
		Message string `json:"message"`
	}
	var output Output
	output.Message = "New Booking added succesfully"

	j, err := json.Marshal(output)
	if err != nil {
		failureResponse(w, `Internal Server Error! `)
		return
	}

	successResponse(w, j)

}

// Function to get all bookings by clients
func allBookings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookings, err := json.Marshal(bookings)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(bookings)
}

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Model for our class data
type Class struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
	Capacity  string `json:"capacity"`
}

var classes = []Class{}

// Function to add a new class
func addClass(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		failureResponse(w, err.Error())
		return
	}

	type addClassRequestBody struct {
		Name      *string `json:"name"`
		StartDate *string `json:"startdate"`
		EndDate   *string `json:"enddate"`
		Capacity  *string `json:"capacity"`
	}

	var classBody addClassRequestBody

	err = json.Unmarshal(body, &classBody)
	if err != nil {
		failureResponse(w, err.Error())
		return
	}

	if classBody.Name == nil {
		failureResponse(w, `name parameter is not specified or empty`)
		return
	}

	if classBody.StartDate == nil {
		failureResponse(w, `startdate parameter is not specified or empty`)
		return
	}

	if classBody.EndDate == nil {
		failureResponse(w, `end date parameter is not specified or empty`)
		return
	}

	newClass := Class{
		ID:        len(classes) + 1,
		Name:      *classBody.Name,
		StartDate: *classBody.StartDate,
		EndDate:   *classBody.EndDate,
		Capacity:  *classBody.Capacity,
	}

	classes = append(classes, newClass)

	type Output struct {
		Message string `json:"message"`
	}
	var output Output
	output.Message = "Successful Class Creation"

	j, err := json.Marshal(output)
	if err != nil {
		failureResponse(w, `Internal Server Error! `)
		return
	}

	successResponse(w, j)

}

// Function to get all classes
func listClass(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	classes, err := json.Marshal(classes)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(classes)
}

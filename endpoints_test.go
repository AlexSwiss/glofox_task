package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

const (
	BadRequestCode     = 400
	SuccessRequestCode = 200
)

type TestStruct struct {
	requestBody        string
	expectedStatusCode int
	responseBody       string
	observedStatusCode int
}

//Inside user_test.go
func TestAddClass(t *testing.T) {

	url := "http://localhost:8080/api/classes_list"

	tests := []TestStruct{
		//{`{}`, BadRequestCode, "", 0},
		//{`{"name":"workout101"}`, BadRequestCode, "", 0},
		//{`{"name":"workout101","startdate":"10 Jun 20"}`, BadRequestCode, "", 0},
		//{`{"name":"workout101","startdate":"10 Jun 20","enddate" : "20 Jun 20"}`, BadRequestCode, "", 0},
		{`{"name":"workout101","startdate":"10 Jun 20","enddate" : "20 Jun 20", "capacity": "10"}`, SuccessRequestCode, "", 0},
	}

	for i, testCase := range tests {

		var reader io.Reader
		reader = strings.NewReader(testCase.requestBody) //Convert string to reader

		request, err := http.NewRequest("GET", url, reader)

		res, err := http.DefaultClient.Do(request)

		if err != nil {
			t.Error(err) //Something is wrong while sending request
		}
		body, _ := ioutil.ReadAll(res.Body)

		tests[i].responseBody = strings.TrimSpace(string(body))
		tests[i].observedStatusCode = res.StatusCode

	}

	DisplayTestCaseResults("addClass", tests, t)

}

func TestNewBooking(t *testing.T) {

	url := "http://localhost:8080/api/bookings_list"

	tests := []TestStruct{
		//{`{}`, BadRequestCode, "", 0},
		//{`{"name":"Alexander Swiss"}`, BadRequestCode, "", 0},
		{`{"name":"Alexander Swiss","date":"02 Apr 20"}`, SuccessRequestCode, "", 0},
	}

	for i, testCase := range tests {

		var reader io.Reader
		reader = strings.NewReader(testCase.requestBody) //Convert string to reader

		request, err := http.NewRequest("GET", url, reader)

		res, err := http.DefaultClient.Do(request)

		if err != nil {
			t.Error(err) //Something is wrong while sending request
		}
		body, _ := ioutil.ReadAll(res.Body)

		tests[i].responseBody = strings.TrimSpace(string(body))
		tests[i].observedStatusCode = res.StatusCode

	}

	DisplayTestCaseResults("addNewBooking", tests, t)

}

func DisplayTestCaseResults(functionalityName string, tests []TestStruct, t *testing.T) {

	for _, test := range tests {

		if test.observedStatusCode == test.expectedStatusCode {
			t.Logf("Passed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n", test.requestBody, test.expectedStatusCode, test.responseBody, test.observedStatusCode)
		} else {
			t.Errorf("Failed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n", test.requestBody, test.expectedStatusCode, test.responseBody, test.observedStatusCode)
		}
	}
}

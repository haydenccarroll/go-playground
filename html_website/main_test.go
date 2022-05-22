package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// we create a new HTTP request here. This request is going to be passed to
	//   the handler.
	// first arg is method, 2nd arg is the route, 3rd arg is request body (not implemented yet)
	req, err := http.NewRequest("GET", "", nil)

	// there was an error creating the request, stop and fail test.
	if err != nil {
		t.Fatal(err)
	}

	// golang's httptest library to create a http recorder. think of this recorder as a mini browser, it is
	// the target of our http requests.
	recorder := httptest.NewRecorder()

	// create http handler for our handler function. argument is handler function
	// defined in main.go.
	hf := http.HandlerFunc(home)

	hf.ServeHTTP(recorder, req)

	// if we didnt get the right code from the html request
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler function returned wrong status code: got %v expected %v",
			status, http.StatusOK)
	}

	expected := "Hello world!"
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler function returned wrong body: got %v expected %v",
			actual, expected)
	}

}

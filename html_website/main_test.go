package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterForHome(t *testing.T) {
	// Instantiate the router
	r := newRouter()

	mockServer := httptest.NewServer(r)
	res, err := http.Get(mockServer.URL + "")

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got: %v", res.StatusCode)
	}

	// close res.Body at end of test function
	defer res.Body.Close()

	//
	bStr, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	resStr := string(bStr)
	expected := "Hello World"

	if resStr != expected {
		t.Errorf("Response should be %s got %s instead.", expected, resStr)
	}
}

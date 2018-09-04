//main_test.go

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// create request obj
	req, err := http.NewRequest("GET", "", nil) // method, route, and request body

	// check error and fail if not nill
	if err != nil {
		t.Fatal(err)
	}

	// create recorder from httptest --- like test-client()
	recorder := httptest.NewRecorder()

	// create an http handler using handler function in main, this is the function we are testing
	hf := http.HandlerFunc(handler)

	// serve the http request to the recorder, actually executing the handler being tested
	hf.ServeHTTP(recorder, req)

	// check the status code is what we expect
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v, want %v",
			status, http.StatusOK,
		)
	}

	// check the response body is what we expect
	expected := `Hello Wiorld!`
	actual := recorder.Body.String()

	if actual != expected {
		t.Errorf("handler returned unexpected  body: got %v, want %v", actual, expected)
	}
}

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplication_GetAllDogBreedsJSON(t *testing.T) {
	//Create request
	req, _ := http.NewRequest("GET", "/dogs/breeds", nil)
	//Create response recorder
	rr := httptest.NewRecorder()
	//Create handler
	handler := http.HandlerFunc(testApp.GetAllDogBreedsJSON)
	//Server the handler
	handler.ServeHTTP(rr, req)
	//Check the response
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %d want 200", rr.Code)
	}
}

func TestApplication_GetAllCatBreeds(t *testing.T) {
	//Create request
	req, _ := http.NewRequest("GET", "/cat-breeds", nil)
	//Create response recorder
	rr := httptest.NewRecorder()
	//Create handler
	handler := http.HandlerFunc(testApp.GetAllCatBreeds)
	//Server the handler
	handler.ServeHTTP(rr, req)
	//Check the response
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %d want 200", rr.Code)
	}
}

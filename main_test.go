package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleIndexReturnsWithStatusOK(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	CityHandler(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Non-expected status code%v:\n\tbody: %v", "200", response.Code)
	}
}

func TestArrayContainsCorrectCities(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	CityHandler(response, request)

	expected := []byte(`{"cities":"San Francisco, Amsterdam, Berlin, New York, Tokyo, Warsaw, Singapore"}`)

	if response.Body.String() != string(expected) {
		t.Fatalf("The returned array does not match the expected array:\nexpected: %v", string(expected))
	}
}

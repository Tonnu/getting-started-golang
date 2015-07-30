package main

import (
	"fmt"
	"log"
	"net/http"
)

func CityHandler(res http.ResponseWriter, req *http.Request) {
	data := []byte(`{"cities":"San Francisco, Amsterdam, Berlin, New York, Tokyo, Warsaw"}`)

	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

func main() {
	http.HandleFunc("/", CityHandler)
	http.HandleFunc("/cities.json", CityHandler)
	fmt.Println("started listening...")
	err := http.ListenAndServe("0.0.0.0:3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

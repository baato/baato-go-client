package main

import (
	"fmt"

	baato "github.com/baato/baato-go-client/lib"
	geocode "github.com/baato/baato-go-client/lib/geocode"
)

func main() {

	// initialize Baato core module
	accessToken := "YOUR BAATO ACCESS TOKEN HERE" // Get Baato token from environment
	baatoMap := baato.Baato(accessToken)

	//geocoding API

	// intialize geocoding request options
	var geocodingRequest = geocode.GeocodeRequestOptions{
		Q:     "do",
		Limit: 5,
	}

	geocode, _ := baatoMap.Geocode.GetGeocode(geocodingRequest)

	fmt.Println(geocode.Data)

}

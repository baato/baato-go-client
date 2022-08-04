package main

import (
	"fmt"

	baato "github.com/baato/baato-go-client/lib"
	"github.com/baato/baato-go-client/lib/geocode"
)

func main() {

	// initialize Baato core module
	accessToken := "YOUR BAATO ACCESS TOKEN HERE" // Get Baato token from environment
	baatoMap := baato.Baato(accessToken)

	//reverse geocoding API

	// intialize geocoding request options
	var geocodingRequest = geocode.GeocodeRequestOpts{
		Q: "po",
	}

	geocode, _ := baatoMap.Geocode.GetGeocode(&geocodingRequest)

	fmt.Println(geocode.Data)

}

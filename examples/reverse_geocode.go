package main

import (
	"fmt"

	baato "github.com/baato/baato-go-client/lib"
	reversegeocode "github.com/baato/baato-go-client/lib/reverse_geocode"
)

func main() {
	// initialize Baato core module
	accessToken := "YOUR BAATO ACCESS TOKEN HERE" // Get Baato token from environment
	baatoMap := baato.Baato(accessToken)

	//reverse geocoding API

	// intialize reverse geocoding request options
	var reverseGeocodingRequest = reversegeocode.ReverseGeocodeRequestOptions{
		Lat: 27.717728723291803,
		Lon: 85.32784938812257,
	}

	reverseGeocode, _ := baatoMap.ReverseGeocode.GetReverseGeocode(reverseGeocodingRequest)

	fmt.Println(reverseGeocode.Data)
}

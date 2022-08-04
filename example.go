package main

import (
	"fmt"

	baato "github.com/baato/baato-go-client/lib"
	reversegeocode "github.com/baato/baato-go-client/lib/reverse_geocode"
)

func main() {

	// initialize Baato core module
	accessToken := "YOUR BAATO ACCESS TOKEN HERE"
	baatoMap := baato.Baato(accessToken)

	//reverse geocoding API

	var reverseGeocodingRequest = reversegeocode.ReverseGeocodeRequestOpts{
		Lat: 27.7148,
		Lon: 85.3105,
	}

	reverseGeocode, _ := baatoMap.ReverseGeocode.GetReverseGeocode(&reverseGeocodingRequest)
	fmt.Println(reverseGeocode)
	fmt.Println(reverseGeocode.Data[0].Name + reverseGeocode.Data[0].Address)
}

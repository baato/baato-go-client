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

	var reverseGeocodingOptions = reversegeocode.ReverseGeocodeRequestOpts{}
	reverseGeocodingOptions.Lat = 27.7183
	reverseGeocodingOptions.Lon = 85.3500

	reverseGeocode, _ := baatoMap.ReverseGeocode.GetGeocode(&reverseGeocodingOptions)
	fmt.Println(reverseGeocode)
	fmt.Println(reverseGeocode.Data[0].Name + reverseGeocode.Data[0].Address)
}

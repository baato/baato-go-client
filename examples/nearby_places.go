package main

import (
	"fmt"

	baato "github.com/baato/baato-go-client/lib"
	nearby "github.com/baato/baato-go-client/lib/nearby"
)

func main() {

	// initialize Baato core module
	accessToken := "YOUR BAATO ACCESS TOKEN" // Get Baato token from environment
	baatoMap := baato.Baato(accessToken)

	// intialize nearby places request options
	var nearbyPlacesRequest = nearby.NearbyPlacesRequestOptions{
		Type: "eat",
		Lat:  27.717728723291803,
		Lon:  85.32784938812257,
	}

	nearbyplaces, _ := baatoMap.NearbyPlaces.GetNearbyPlaces(nearbyPlacesRequest)

	fmt.Println(nearbyplaces.Data)

}

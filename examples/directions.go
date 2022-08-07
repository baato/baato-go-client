package main

import (
	"fmt"

	baato "github.com/baato/baato-go-client/lib"
	"github.com/baato/baato-go-client/lib/directions"
)

func main() {

	// initialize Baato core module
	accessToken := "YOUR BAATO ACCESS TOKEN" // Get Baato token from environment
	baatoMap := baato.Baato(accessToken)

	// intialize directions request options
	directionsRequest := directions.DirectionsRequestOptions{
		PointsArray: []string{"27.6733433,85.2763307", "27.67444444,85.28047222"},
		Mode:        "bike",
	}

	directions, _ := baatoMap.Directions.GetDirections(directionsRequest)

	fmt.Println(directions.Data)

}

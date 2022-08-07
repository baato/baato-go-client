# Baato Go Client

![Build](https://github.com/drklrd/baato-go-client/actions/workflows/baato_go_client.yml/badge.svg)
[![Documentation](https://img.shields.io/docsrs/badge?style=plastic)](https://docs.baato.io)
![License](https://img.shields.io/badge/License-MIT-green.svg)


![Baato Go Client Logo](https://github.com/drklrd/baato-go-client/blob/main/goclient.png?raw=true)




Go client to consume Baato APIs easily.


# Initialization
```go
// Import core APIs
import (
	baato "github.com/baato/baato-go-client/lib"
	"github.com/baato/baato-go-client/lib/geocode"
)

// initialize Baato core module
accessToken := "YOUR BAATO ACCESS TOKEN HERE" // Get Baato token from environment
baatoMap := baato.Baato(accessToken)

```

# Available APIs
- Geocoding (Search) API
- ReverseGeocoding (Reverse Search) API
- Nearby Features API
- Directions API


# Usage 

For a detailed usage instructions, please visit documentation for the Baato-Go-Client avaiable [here](https://docs.baato.io). Quick usage examples are shown here.

## Geocode (Search) API
```go
// intialize geocoding request options
var geocodingRequest = geocode.GeocodeRequestOpts{
	Q:     "po",
	Limit: 5,
}

geocode, _ := baatoMap.Geocode.GetGeocode(&geocodingRequest)

fmt.Println(geocode.Data)
```

## ReverseGeocode (Reverse Search) API

```go
// intialize reverse geocoding request options
var reverseGeocodingRequest = reversegeocode.ReverseGeocodeRequestOpts{
	Lat: 27.717728723291803,
	Lon: 85.32784938812257,
}

reverseGeocode, _ := baatoMap.ReverseGeocode.GetReverseGeocode(&reverseGeocodingRequest)

fmt.Println(reverseGeocode.Data)
```


## Nearby Places API

You can use this API to get nearby places around a point that is interesting for you. For a complete list of supported type of places, please see the [docs](https://docs.baato.io/#/v1/services/nearby_places).

```go
// intialize nearby places request options
	var nearbyPlacesRequest = nearby.NearbyPlacesRequestOptions{
		Type: "eat",
		Lat:  27.717728723291803,
		Lon:  85.32784938812257,
	}

	nearbyplaces, _ := baatoMap.NearbyPlaces.GetNearbyPlaces(nearbyPlacesRequest)

	fmt.Println(nearbyplaces.Data)
```

## Directions API

```go
// intialize directions request options. PointsArray represents points that we should pass through.
	directionsRequest := directions.DirectionsRequestOptions{
		PointsArray: []string{"27.6733433,85.2763307", "27.67444444,85.28047222"},
		Mode:        "bike",
	}

	directions, _ := baatoMap.Directions.GetDirections(directionsRequest)

	fmt.Println(directions.Data)

```



# Contribution
You can contribute in many ways (development, issue/bug reporting, feature request, etc.)
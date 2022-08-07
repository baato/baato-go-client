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



# Contribution
You can contribute in many ways (development, issue/bug reporting, feature request, etc.)
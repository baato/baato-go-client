package commons

type Point []float64

type Tags []string

type Centroid struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

type Geometry struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"`
}

type ReverseSearchFeature struct {
	PlaceId  string   `json:"placeId"`
	OsmId    string   `json:"osmID"`
	License  string   `json:"license"`
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Centroid Centroid `json:"centroid"`
	Geometry Geometry `json:"geometry"`
	Tags     Tags     `json:"tags"`
	Score    string   `json:"score"`
}

type ReverseSearchData struct {
	Data []ReverseSearchFeature `json:"data"`
}

type SearchFeature struct {
	PlaceId            string  `json:"placeId"`
	OsmId              string  `json:"osmID"`
	Name               string  `json:"name"`
	Address            string  `json:"address"`
	Type               string  `json:"type"`
	Score              float64 `json:"score"`
	RadialDistanceInKm int     `json:"radialDistanceInKm"`
}

type SearchData struct {
	Data []SearchFeature `json:"data"`
}

type NearbyPlaceFeature struct {
	PlaceId            string   `json:"placeId"`
	OsmId              string   `json:"osmID"`
	License            string   `json:"license"`
	Name               string   `json:"name"`
	Address            string   `json:"address"`
	Type               string   `json:"type"`
	Centroid           Centroid `json:"centroid"`
	Tags               Tags     `json:"tags"`
	Geometry           Geometry `json:"geometry"`
	Score              float64  `json:"score"`
	RadialDistanceInKm int      `json:"radialDistanceInKm"`
	open               string   `json:"open"`
}

type DirectionsFeature struct {
	EncodedPolyline  string      `json:"encodedPolyline"`
	DistanceInMeters float64     `json:"distanceInMeters"`
	TimeInMs         int         `json:"timeInMs"`
	InstructionList  interface{} `json:"instructionList"`
}

type NearbyPlacesData struct {
	Data []NearbyPlaceFeature `json:"data"`
}

type DirectionsData struct {
	Data []DirectionsFeature `json:"data"`
}

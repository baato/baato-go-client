package commons

type Point []float64

type Tags []string

type Centroid struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

type Geometry struct {
	Type        string `json:"type"`
	Coordinates Point  `json:"coordinates"`
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

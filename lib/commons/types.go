package commons

type Point []float64

type Tags []string

type Centroid struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Geometry struct {
	Type        string   `json:"type"`
	Coordinates Centroid `json:"coordinates"`
}

type Feature struct {
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

type FeatureCollection struct {
	Data []Feature `json:"data"`
}

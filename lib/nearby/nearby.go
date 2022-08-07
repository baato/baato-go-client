package nearby

import (
	"github.com/baato/baato-go-client/lib/commons"
	"github.com/google/go-querystring/query"
)

const (
	apiLabel = "search/nearby"
)

type NearbyPlaces struct {
	commons *commons.Commons
}

type NearbyPlacesRequestOptions struct {
	Type string  `url:"type"`
	Lat  float64 `url:"lat,omitempty"`
	Lon  float64 `url:"lon,omitempty"`
}

type NearbyPlacesResponse struct {
	Timestamp string
	Status    int
	Message   string
	*commons.NearbyPlacesData
}

func NewNearbyPlaces(commons *commons.Commons) *NearbyPlaces {
	return &NearbyPlaces{commons}
}

func (n *NearbyPlaces) GetNearbyPlaces(request NearbyPlacesRequestOptions) (*NearbyPlacesResponse, error) {
	values, error := query.Values(request)

	if error != nil {
		return nil, error
	}

	resp := NearbyPlacesResponse{}

	n.commons.BaatoAPIRequest(apiLabel, &values, &resp)

	return &resp, error
}

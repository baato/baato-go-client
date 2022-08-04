package reversegeocode

import (
	"github.com/baato/baato-go-client/lib/commons"
	"github.com/google/go-querystring/query"
)

const (
	apiLabel = "reverse"
)

type ReverseGeocode struct {
	commons *commons.Commons
}

type ReverseGeocodeRequestOpts struct {
	Lat float64 `url:"lat,omitempty"`
	Lon float64 `url:"lon,omitempty"`
}

type ReverseGeocodeResponse struct {
	Timestamp string
	Status    int
	Message   string
	*commons.FeatureCollection
}

func NewReverseGeocode(commons *commons.Commons) *ReverseGeocode {
	return &ReverseGeocode{commons}
}

func (r *ReverseGeocode) GetReverseGeocode(request *ReverseGeocodeRequestOpts) (*ReverseGeocodeResponse, error) {

	values, error := query.Values(request)

	if error != nil {
		return nil, error
	}

	resp := ReverseGeocodeResponse{}

	r.commons.BaatoAPIRequest(apiLabel, &values, &resp)

	return &resp, error
}

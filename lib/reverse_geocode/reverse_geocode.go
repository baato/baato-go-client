package reversegeocode

import (
	"github.com/baato/baato-go-client/lib/commons"
)

type ReverseGeocode struct {
	commons *commons.Commons
}

type ReverseGeocodeRequestOpts struct {
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

func (r *ReverseGeocode) GetGeocode() *ReverseGeocodeResponse {

	resp := ReverseGeocodeResponse{}

	r.commons.APIRequest(&resp)

	return &resp
}

package geocode

import (
	"github.com/baato/baato-go-client/lib/commons"
	"github.com/google/go-querystring/query"
)

const (
	apiLabel = "search"
)

type Geocode struct {
	commons *commons.Commons
}

type GeocodeRequestOptions struct {
	Q     string `url:"q"`
	Limit int    `url:"limit"`
}

type GecodeResponse struct {
	Timestamp string
	Status    int
	Message   string
	*commons.SearchData
}

func NewGeocode(commons *commons.Commons) *Geocode {
	return &Geocode{commons}
}

func (g *Geocode) GetGeocode(request GeocodeRequestOptions) (*GecodeResponse, error) {
	values, error := query.Values(request)

	if error != nil {
		return nil, error
	}

	resp := GecodeResponse{}

	g.commons.BaatoAPIRequest(apiLabel, &values, &resp)

	return &resp, error
}

package directions

import (
	"github.com/baato/baato-go-client/lib/commons"
	"github.com/google/go-querystring/query"
)

const (
	apiLabel = "directions"
)

type Directions struct {
	commons *commons.Commons
}

type DirectionsRequestOptions struct {
	PointsArray  []string `url:"pointsArray"`
	Mode         string   `url:"mode"`
	Alternatives bool     `url:"alternatives"`
	Instructions bool     `url:"instructions"`
}

type DirectionsResponse struct {
	Timestamp string
	Status    int
	Message   string
	*commons.DirectionsData
}

func NewDirections(commons *commons.Commons) *Directions {
	return &Directions{commons}
}

func (d *Directions) GetDirections(request DirectionsRequestOptions) (*DirectionsResponse, error) {

	values, error := query.Values(request)

	// deserialize points array into each points[] query param
	pointsArray := request.PointsArray
	for _, s := range pointsArray {
		values.Add("points[]", s)
	}

	if error != nil {
		return nil, error
	}

	resp := DirectionsResponse{}

	d.commons.BaatoAPIRequest(apiLabel, &values, &resp)

	return &resp, error
}

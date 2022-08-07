package baato

import (
	"github.com/baato/baato-go-client/lib/commons"
	"github.com/baato/baato-go-client/lib/directions"
	"github.com/baato/baato-go-client/lib/geocode"
	"github.com/baato/baato-go-client/lib/nearby"
	reversegeocode "github.com/baato/baato-go-client/lib/reverse_geocode"
)

type BaatoAPI struct {
	commons *commons.Commons

	ReverseGeocode *reversegeocode.ReverseGeocode
	Geocode        *geocode.Geocode
	NearbyPlaces   *nearby.NearbyPlaces
	Directions     *directions.Directions
}

func Baato(accessToken string) *BaatoAPI {

	instance := &BaatoAPI{}
	commons, _ := commons.NewCommons(accessToken)

	instance.commons = commons

	instance.ReverseGeocode = reversegeocode.NewReverseGeocode(instance.commons)
	instance.Geocode = geocode.NewGeocode(instance.commons)
	instance.NearbyPlaces = nearby.NewNearbyPlaces(instance.commons)
	instance.Directions = directions.NewDirections(instance.commons)
	return instance

}

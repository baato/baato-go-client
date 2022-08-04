package baato

import (
	"github.com/baato/baato-go-client/lib/commons"
	"github.com/baato/baato-go-client/lib/geocode"
	reversegeocode "github.com/baato/baato-go-client/lib/reverse_geocode"
)

type BaatoAPI struct {
	commons *commons.Commons

	ReverseGeocode *reversegeocode.ReverseGeocode
	Geocode        *geocode.Geocode
}

func Baato(accessToken string) *BaatoAPI {

	instance := &BaatoAPI{}
	commons, _ := commons.NewCommons(accessToken)

	instance.commons = commons

	instance.ReverseGeocode = reversegeocode.NewReverseGeocode(instance.commons)
	instance.Geocode = geocode.NewGeocode(instance.commons)
	return instance

}

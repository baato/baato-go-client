package baato

import (
	"github.com/baato/baato-go-client/lib/commons"
	reversegeocode "github.com/baato/baato-go-client/lib/reverse_geocode"
)

type BaatoAPI struct {
	commons *commons.Commons

	ReverseGeocode *reversegeocode.ReverseGeocode
}

func Baato(accessToken string) *BaatoAPI {

	instance := &BaatoAPI{}
	commons, _ := commons.NewCommons(accessToken)

	instance.commons = commons

	instance.ReverseGeocode = reversegeocode.NewReverseGeocode(instance.commons)
	return instance

}

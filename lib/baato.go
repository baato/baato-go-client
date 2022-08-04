package baato

import (
	"github.com/baato/baato-go-client/lib/commons"
	reversegeocode "github.com/baato/baato-go-client/lib/reverse_geocode"
)

type BaatoAPI struct {
	commons *commons.Commons

	ReverseGeocode *reversegeocode.ReverseGeocode
}

func Baato(token string) (*BaatoAPI, error) {

	instance := &BaatoAPI{}
	commons, _ := commons.NewCommons(token)

	instance.commons = commons

	instance.ReverseGeocode = reversegeocode.NewReverseGeocode(instance.commons)
	return instance, nil

}

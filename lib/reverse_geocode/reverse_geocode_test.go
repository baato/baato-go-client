package reversegeocode

import (
	"os"
	"testing"

	"github.com/baato/baato-go-client/lib/commons"
)

func TestReverseGeocoder(t *testing.T) {

	c, err := commons.NewCommons(os.Getenv("BAATO_ACCESS_TOKEN"))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	reverse_geocode := NewReverseGeocode(c)

	t.Run("Can Reverse Geocode", func(t *testing.T) {
		var reqOpt ReverseGeocodeRequestOptions
		// intialize reverse geocoding request options

		reqOpt.Lat = 27.717728723291803
		reqOpt.Lon = 85.32784938812257

		res, err := reverse_geocode.GetReverseGeocode(reqOpt)

		if err != nil {
			t.Error(err)
		}

		if res.Status != 200 {
			t.Error("Request Failed with status: ", res.Status)
		}

	})
}

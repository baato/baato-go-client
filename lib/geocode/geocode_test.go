package geocode

import (
	"os"
	"testing"

	"github.com/baato/baato-go-client/lib/commons"
)

func TestGeocoder(t *testing.T) {

	c, err := commons.NewCommons(os.Getenv("BAATO_ACCESS_TOKEN"))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	geocode := NewGeocode(c)

	t.Run("Can Geocode", func(t *testing.T) {
		var reqOpt GeocodeRequestOptions
		// intialize geocoding request options

		reqOpt.Q = "pop"
		reqOpt.Limit = 5

		res, err := geocode.GetGeocode(reqOpt)

		if err != nil {
			t.Error(err)
		}

		if res.Status != 200 {
			t.Error("Request Failed with status: ", res.Status)
		}

	})
}

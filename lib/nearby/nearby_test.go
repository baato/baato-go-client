package nearby

import (
	"os"
	"testing"

	"github.com/baato/baato-go-client/lib/commons"
)

func TestNearby(t *testing.T) {

	c, err := commons.NewCommons(os.Getenv("BAATO_ACCESS_TOKEN"))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	nearby := NewNearbyPlaces(c)

	t.Run("Can Get Nearby places", func(t *testing.T) {
		var reqOpt NearbyPlacesRequestOptions
		// intialize nearby places request options

		reqOpt.Type = "eat"
		reqOpt.Lat = 27.717728723291803
		reqOpt.Lon = 85.32784938812257

		res, err := nearby.GetNearbyPlaces(reqOpt)

		if err != nil {
			t.Error(err)
		}

		if res.Status != 200 {
			t.Error("Request Failed with status: ", res.Status)
		}

	})
}

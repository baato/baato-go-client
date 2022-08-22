package directions

import (
	"os"
	"testing"

	"github.com/baato/baato-go-client/lib/commons"
)

func TestDirections(t *testing.T) {

	c, err := commons.NewCommons(os.Getenv("BAATO_ACCESS_TOKEN"))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	directions := NewDirections(c)

	t.Run("Can Get Directions", func(t *testing.T) {
		var reqOpt DirectionsRequestOptions
		// intialize directions request options

		reqOpt.PointsArray = []string{"27.6733433,85.2763307", "27.67444444,85.28047222"}
		reqOpt.Mode = "bike"

		res, err := directions.GetDirections(reqOpt)

		if err != nil {
			t.Error(err)
		}

		if res.Status != 200 {
			t.Error("Request Failed with status: ", res.Status)
		}

	})
}

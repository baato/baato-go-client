/**

*Commons package for baato-go-client
* Authored by @drklrd
 */

package commons

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Commons struct {
	accessToken string
}

func NewCommons(accessToken string) (*Commons, error) {
	if accessToken == "" {
		return nil, errors.New("Baato access token is required!")
	}

	c := &Commons{}
	c.accessToken = accessToken
	return c, nil

}

func (c *Commons) APIRequest(values *url.Values, response interface{}) error {

	values.Set("key", c.accessToken)

	url := "https://api.baato.io/api/v1/reverse"
	// Invoke request
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.URL.RawQuery = values.Encode()
	if err != nil {
		return nil
	}

	resp, err := http.DefaultClient.Do(request)

	// Read body into buffer
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &response)

	// fmt.Println(string(body))

	return nil

}

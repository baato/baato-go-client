/**

*Commons package for baato-go-client
* Authored by @drklrd
 */

package commons

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Commons struct {
	token string
}

func NewCommons(token string) (*Commons, error) {
	// if token = ""{
	// 	return nil, errors.New("Baato token is required!")
	// }

	c := &Commons{}
	c.token = token
	return c, nil

}

func (c *Commons) APIRequest(response interface{}) error {

	// Invoke request
	request, err := http.NewRequest(http.MethodGet, url, nil)
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

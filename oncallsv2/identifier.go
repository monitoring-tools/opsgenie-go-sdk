package oncallsv2

import (
	"net/url"
)

// Identifier contains uri params
type Identifier struct {
	ID   string `json:"-"`
	Name string `json:"-"`
	Flat bool   `json:"-"`
}

// GenerateUrl generates url to on calls api
func (request *Identifier) GenerateUrl() (string, url.Values, error) {
	baseUrl := "/v2/schedules/"

	params := url.Values{}
	if request != nil {
		if len(request.ID) != 0 {
			baseUrl += request.ID
			params.Set("scheduleIdentifierType", "id")
		} else if len(request.Name) != 0 {
			params.Set("scheduleIdentifierType", "name")
			baseUrl += request.Name
		}

		baseUrl += "/on-calls"

		if request.Flat {
			params.Set("flat", "true")
		}
	}

	return baseUrl, params, nil
}

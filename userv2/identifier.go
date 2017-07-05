package userv2

import (
	"net/url"
	"errors"
)

const (
	ContactExpandableField = "contact"
)

type Identifier struct {
	ID string `json:"-"`
	UserName string `json:"-"`
	Expand string `json:"-"`
}

func (request *Identifier) GenerateUrl() (string, url.Values, error) {
	baseUrl := "/v2/users/"

	if len(request.ID) > 0 {
		baseUrl += "/" + request.ID
	} else if len(request.UserName) > 0 {
		baseUrl += "/" + request.UserName
	} else {
		return "", nil, errors.New("username or id of the user should be provided")
	}

	params := url.Values{}
	if len(request.Expand) > 0 {
		if request.Expand == ContactExpandableField {
			params.Set("expand", request.Expand)
		} else {
			return "", nil, errors.New("the only expandable field for user api is 'contact'")
		}
	}

	return baseUrl, nil, nil
}
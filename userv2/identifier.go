package userv2

import (
	"net/url"
	"errors"
	"fmt"
)

const (
	ContactExpandableField = "contact"

	EscalationsEntity = "escalations"
	TeamsEntity = "teams"
	ForwardingRulesEntity = "forwarding-rules"
	SchedulesEntity = "schedules"
)

type Identifier struct {
	ID string `json:"-"`
	UserName string `json:"-"`
	Entity string `json:"-"`
	Expand string `json:"-"`
}

func (request *Identifier) GenerateUrl() (string, url.Values, error) {
	baseUrl := "/v2/users"

	if len(request.ID) > 0 {
		baseUrl += "/" + request.ID
	} else if len(request.UserName) > 0 {
		baseUrl += "/" + request.UserName
	} else {
		return "", nil, errors.New("username or id of the user should be provided")
	}

	if len(request.Entity) > 0 {
		if request.Entity == EscalationsEntity ||
			request.Entity == TeamsEntity ||
			request.Entity == ForwardingRulesEntity ||
			request.Entity == SchedulesEntity {

			baseUrl += "/" + request.Entity
		} else {
			err := fmt.Errorf(
				"not available user entity, in should one of the %s, %s, %s, %s",
				EscalationsEntity,
				TeamsEntity,
				ForwardingRulesEntity,
				SchedulesEntity,
			)

			return "", nil, err
		}
	}

	params := url.Values{}
	if len(request.Expand) > 0 {
		if request.Expand == ContactExpandableField {
			params.Set("expand", request.Expand)
		} else {
			return "", nil, errors.New("the only expandable field for user api is 'contact'")
		}
	}

	return baseUrl, params, nil
}
package notificationv2

import (
	"net/url"
	"errors"
)

const (
	StatusActionEnable = "enable"
	StatusActionDisable = "disable"
)

// Identifier defined the set of attributes for identification notification.
type Identifier struct {
	UserID     string `json:"-"`
	RuleID string `json:"-"`
	StatusAction string `json:"-"`
}

// GenerateUrl generates API url using specified attributes of identifier.
func (request *Identifier) GenerateUrl() (string, url.Values, error) {
	if request.UserID == "" {
		return "", nil, errors.New("UserID should be provided")
	}

	if len(request.StatusAction) != 0 && request.RuleID == "" {
		return "", nil, errors.New("RuleID should be provided along with disabled/enabled action")
	}

	baseUrl := "/v2/users/" + request.UserID + "/notification-rules"

	if request.RuleID != "" {
		baseUrl += "/" + request.RuleID
	}

	if request.StatusAction == StatusActionEnable {
		baseUrl += "/" + StatusActionEnable
	}

	if request.StatusAction == StatusActionDisable {
		baseUrl += "/" + StatusActionDisable
	}

	return baseUrl, nil, nil
}
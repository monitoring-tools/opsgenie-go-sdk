package userv2

import "time"

type ListUserForwardingRulesResponse struct {
	ForwardingRules []ForwardingRule `json:"data,omitempty"`
	*ResponseMeta
}

type ForwardingRule struct {
	ID        string    `json:"id,omitempty"`
	Alias     string    `json:"alias,omitempty"`
	FromUser  FromUser  `json:"fromUser,omitempty"`
	ToUser    ToUser    `json:"toUser,omitempty"`
	StartDate time.Time `json:"startDate,omitempty"`
	EndDate   time.Time `json:"endDate,omitempty"`
}

type FromUser struct {
	ID       string `json:"id,omitempty"`
	UserName string `json:"username,omitempty"`
}

type ToUser struct {
	ID       string `json:"id,omitempty"`
	UserName string `json:"username,omitempty"`
}

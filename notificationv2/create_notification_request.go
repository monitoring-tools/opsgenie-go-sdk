package notificationv2

type CreateNotificationRequest struct {
	*Identifier
	ApiKey           string
	Name             string `json:"name"`
	ActionType       string `json:"actionType"`
	Criteria         Criteria `json:"criteria"`
	NotificationTime []string `json:"notificationTime"`
	TimeRestriction  TimeRestriction `json:"timeRestriction"`
	Schedules        []Schedule `json:"schedules"`
	Order            int `json:"order"`
	Steps            []Step `json:"steps"`
	Repeat           Repeat `json:"repeat"`
	Enabled          bool `json:"enabled"`
}

func (r *CreateNotificationRequest) GetApiKey() string {
	return r.ApiKey
}

type Criteria struct {
	Type       string `json:"type"`
	Conditions []Condition `json:"conditions"`
}

type Condition struct {
	Field         string `json:"field"`
	Key           string `json:"key"`
	Not           bool `json:"not"`
	Operation     string `json:"operation"`
	ExpectedValue string `json:"expectedValue"`
	Order         int `json:"order"`
}

type Repeat struct {
	LoopAfter int `json:"loopAfter"`
	Enabled   bool `json:"enabled"`
}

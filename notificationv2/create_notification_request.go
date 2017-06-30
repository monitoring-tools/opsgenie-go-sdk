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

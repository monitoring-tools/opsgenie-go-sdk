package notificationv2

type UpdateNotificationRequest struct {
	*Identifier
	ApiKey           string
	Name             string `json:"name"`
	Criteria         Criteria `json:"criteria"`
	NotificationTime []string `json:"notificationTime"`
	TimeRestriction  TimeRestriction `json:"timeRestriction"`
	Schedules        []Schedule `json:"schedules"`
	Steps            []Step `json:"steps"`
	Repeat           Repeat `json:"repeat"`
	Order            int `json:"order"`
	Enabled          bool `json:"enabled"`
}

func (r *UpdateNotificationRequest) GetApiKey() string {
	return r.ApiKey
}

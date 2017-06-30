package notificationv2

type Notification struct {
	ID               string `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	ActionType       string `json:"actionType,omitempty"`
	NotificationTime []string `json:"notificationTime,omitempty"`
	Order            int `json:"order,omitempty"`
	Steps            []Step `json:"steps,omitempty"`
	Schedules        []Schedule `json:"schedules,omitempty"`
	Criteria         Criteria `json:"criteria,omitempty"`
	Enabled          bool `json:"enabled,omitempty"`
	Repeat           Repeat `json:"repeat,omitempty"`
	TimeRestriction  TimeRestriction `json:"timeRestriction,omitempty"`
}

type ListNotificationResponse struct {
	ResponseMeta
	Notifications []Notification `json:"data"`
}

type CreateNotificationResponse struct {
	Notification
	ResponseMeta
}

package userv2

type ListUserSchedulesResponse struct {
	Schedules []Schedule `json:"data,omitempty"`
	*ResponseMeta
}

type Schedule struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
}

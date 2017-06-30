package notificationv2

type TimeRestriction struct {
	Type        string `json:"type,omitempty"`
	Restrictions []Restriction `json:"restrictions,omitempty"`
}

type Restriction struct {
	StartDay string `json:"startDay,omitempty"`
	EndDay string `json:"endDay,omitempty"`
	StartHour int `json:"startHour,omitempty"`
	EndHour   int `json:"endHour,omitempty"`
	StartMin  int `json:"startMin,omitempty"`
	EndMin    int `json:"endMin,omitempty"`
}

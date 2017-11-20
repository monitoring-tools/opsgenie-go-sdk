package oncallsv2

// GetOnCallsResponse is a response of loading on calls users
type GetOnCallsResponse struct {
	ResponseMeta
	Result GetOnCallsResult `json:"data"`
}

// GetOnCallsResult contains on calls users information
type GetOnCallsResult struct {
	Parent             Schedule      `json:"_parent"`
	OnCallParticipants []Participant `json:"onCallParticipants"`
}

// Schedule contains current info about schedule
type Schedule struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

// Participant contains info about on calls team or schedule
type Participant struct {
	ID                 string        `json:"id"`
	Name               string        `json:"name"`
	Type               string        `json:"type"`
	EscalationTime     int           `json:"escalationTime,omitempty"`
	NotifyType         string        `json:"notifyType,omitempty"`
	ForwardedFrom      *Participant   `json:"forwardedFrom,omitempty"`
	OnCallParticipants []Participant `json:"onCallParticipants,omitempty"`
}



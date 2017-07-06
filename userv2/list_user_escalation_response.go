package userv2

type ListUserEscalationResponse struct {
	Escalations []Escalation `json:"data,omitempty"`
	*ResponseMeta
}

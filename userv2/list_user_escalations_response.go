package userv2

type ListUserEscalationsResponse struct {
	Escalations []Escalation `json:"data,omitempty"`
	*ResponseMeta
}

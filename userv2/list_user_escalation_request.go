package userv2

type ListUserEscalationRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey returns api key.
func (r *ListUserEscalationRequest) GetApiKey() string {
	return r.ApiKey
}

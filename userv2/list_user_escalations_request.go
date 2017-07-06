package userv2

type ListUserEscalationsRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey returns api key.
func (r *ListUserEscalationsRequest) GetApiKey() string {
	return r.ApiKey
}

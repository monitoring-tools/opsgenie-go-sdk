package userv2

type ListUserTeamsRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey returns api key.
func (r *ListUserTeamsRequest) GetApiKey() string {
	return r.ApiKey
}

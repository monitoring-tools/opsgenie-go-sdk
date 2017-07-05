package userv2

type GetUserRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey returns api key.
func (r *GetUserRequest) GetApiKey() string {
	return r.ApiKey
}

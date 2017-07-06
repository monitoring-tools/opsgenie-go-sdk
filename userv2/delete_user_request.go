package userv2

type DeleteUserRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey returns api key.
func (r *DeleteUserRequest) GetApiKey() string {
	return r.ApiKey
}

package userv2

type ListUserSchedulesRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey returns api key.
func (r *ListUserSchedulesRequest) GetApiKey() string {
	return r.ApiKey
}

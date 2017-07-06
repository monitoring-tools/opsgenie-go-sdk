package userv2

type ListUserForwardingRulesRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey returns api key.
func (r *ListUserForwardingRulesRequest) GetApiKey() string {
	return r.ApiKey
}

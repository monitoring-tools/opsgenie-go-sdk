package policyv1

type ApiKey struct {
	ApiKey string
}

// GetApiKey return api key
func (r *ApiKey) GetApiKey() string {
	return r.ApiKey
}
package oncallsv2

// ApiKey contains api key
type ApiKey struct {
	ApiKey string
}

// GetApiKey return api key
func (r *ApiKey) GetApiKey() string {
	return r.ApiKey
}

package notificationv2

type EnableNotificationRequest struct {
	*Identifier
	ApiKey string
}

func (r *EnableNotificationRequest) GetApiKey() string {
	return r.ApiKey
}

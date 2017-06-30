package notificationv2

type GetNotificationRequest struct {
	*Identifier
	ApiKey string
}

func (r *GetNotificationRequest) GetApiKey() string {
	return r.ApiKey
}

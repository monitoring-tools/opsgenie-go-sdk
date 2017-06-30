package notificationv2

type ListNotificationRequest struct {
	*Identifier
	ApiKey string
}

func (r *ListNotificationRequest) GetApiKey() string {
	return r.ApiKey
}

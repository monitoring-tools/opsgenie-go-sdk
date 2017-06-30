package notificationv2

type DeleteNotificationRequest struct {
	*Identifier
	ApiKey string
}

func (r *DeleteNotificationRequest) GetApiKey() string {
	return r.ApiKey
}

package notificationv2

type DisableNotificationRequest struct {
	*Identifier
	ApiKey string
}

func (r *DisableNotificationRequest) GetApiKey() string {
	return r.ApiKey
}

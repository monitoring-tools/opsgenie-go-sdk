package notificationv2

type DetailedNotification struct {
	*Notification
}

type DetailedNotificationResponse struct {
	ResponseMeta
	Notification DetailedNotification `json:"data"`
}

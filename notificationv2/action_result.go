package notificationv2

type DeleteNotificationResponse struct {
	ResponseMeta
	ActionResult
}

type ActionResult struct {
	Result string `json:"result"`
}

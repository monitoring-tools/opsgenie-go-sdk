package notificationv2

type DeleteNotificationResponse struct {
	ResponseMeta
	ActionResult
}

type EnableNotificationResponse struct {
	ResponseMeta
	ActionResult
}

type DisableNotificationResponse struct {
	ResponseMeta
	ActionResult
}

type ActionResult struct {
	Result string `json:"result"`
}

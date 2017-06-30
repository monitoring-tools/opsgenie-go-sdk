package notificationv2

const (
	CreateAlertActionType         = "create-alert"
	AcknowledgedAlertActionType   = "acknowledged-alert"
	ClosedAlertActionType         = "closed-alert"
	AssignedAlertActionType       = "assigned-alert"
	AddNoteActionType             = "add-note"
	ScheduleStartActionType       = "schedule-start"
	ScheduleEndActionType         = "schedule-end"
	IncomingCallRoutingActionType = "incoming-call-routing"

	JustBeforeNotificationTime        = "just-before"
	FifteenMinutesAgoNotificationTime = "15-minutes-ago"
	OneHourAgoNotificationTime        = "1-hour-ago"
	OneDayAgoNotificationTime         = "1-day-ago"

	TimeOfDayTimeRestriction           = "time-of-day"
	WeekendAndTimeOfDayTimeRestriction = "weekday-and-time-of-day"

	Monday = "monday"
	Tuesday = "tuesday"
	Wednesday = "wednesday"
	Thursday = "thursday"
	Friday= "friday"
	Saturday = "saturday"
	Sunday = "sunday"

	Minutes = "minutes"

	SMSNotifyMethod = "sms"
	EmailNotifyMethod = "email"
	VoiceNotifyMethod = "voice"
)

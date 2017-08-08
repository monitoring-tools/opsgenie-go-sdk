package notificationv2

const (
	// Types of action that notification rule will have after creating. It is used to create an alert.
	CreateAlertActionType         = "create-alert"
	AcknowledgedAlertActionType   = "acknowledged-alert"
	ClosedAlertActionType         = "closed-alert"
	AssignedAlertActionType       = "assigned-alert"
	AddNoteActionType             = "add-note"
	ScheduleStartActionType       = "schedule-start"
	ScheduleEndActionType         = "schedule-end"
	IncomingCallRoutingActionType = "incoming-call-routing"

	// List of time periods that notification for schedule start/end will be sent.
	JustBeforeNotificationTime        = "just-before"
	FifteenMinutesAgoNotificationTime = "15-minutes-ago"
	OneHourAgoNotificationTime        = "1-hour-ago"
	OneDayAgoNotificationTime         = "1-day-ago"

	// List of time restrictions within alerts will be sent.
	TimeOfDayTimeRestriction           = "time-of-day"
	WeekendAndTimeOfDayTimeRestriction = "weekday-and-time-of-day"

	// The list of week days. These strings are used for generate time restrictions.
	Monday    = "monday"
	Tuesday   = "tuesday"
	Wednesday = "wednesday"
	Thursday  = "thursday"
	Friday    = "friday"
	Saturday  = "saturday"
	Sunday    = "sunday"

	// The text representation of word "minutes".
	Minutes = "minutes"

	// The list of notification types.
	SMSNotifyMethod    = "sms"
	EmailNotifyMethod  = "email"
	VoiceNotifyMethod  = "voice"
	MobileNotifyMethod = "mobile"

	// The list of statuses, which are used for enable/disable notification rules.
	EnableStatusAction  = "enable"
	DisableStatusAction = "disable"

	// The list of matches which are used for creating criteria of notification.
	MatchAllType = "match-all"
	MatchAnyConditionsType = "match-any-condition"
	MatchAllConditionsType = "match-all-conditions"

	// The list of condition operation, which are used for build notification criteria.
	MatchesConditionOperation = "matches"
	EqualsConditionOperation  = "equals"
	IsEmptyConditionOperation  = "is-empty"

	// The list of fields, which are used for build filter of alerts.
	ActionsField = "actions"
	AliasField = "alias"
	DescriptionField = "description"
	EntityField = "entity"
	MessageField = "message"
	RecipientsField = "recipients"
	SourceField = "source"
	TeamsField = "teams"
	ExtraPropertiesField = "extra-properties"
)

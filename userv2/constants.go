package userv2

const (
	// OwnerRole is the text value of standard role "owner"
	OwnerRole = "owner"
	// AdminRole is the text value of standard role "admin"
	AdminRole = "admin"
	// UserRole is the text value of standard role "user"
	UserRole = "user"

	// ContactExpandableField is the query parameter, which is needed to load fully contact data of user.
	ContactExpandableField = "contact"

	// EscalationsEntity is the query parameter to load user escalations.
	EscalationsEntity     = "escalations"
	// EscalationsEntity is the query parameter to load user teams.
	TeamsEntity           = "teams"
	// EscalationsEntity is the query parameter to load user forwarding rules.
	ForwardingRulesEntity = "forwarding-rules"
	// EscalationsEntity is the query parameter to load user schedules.
	SchedulesEntity       = "schedules"
)

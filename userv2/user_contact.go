package userv2

type UserContact struct {
	ID string `json:"id"`
	To string `json:"to"`
	ContactMethod string `json:"contactMethod"`
	Enabled bool `json:"enabled"`
}

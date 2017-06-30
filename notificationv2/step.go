package notificationv2

type Step struct {
	Parent    Parent `json:"_parent"`
	ID        string `json:"id"`
	Contact   Contact `json:"contact,omitempty"`
	SendAfter SendAfter `json:"sendAfter,omitempty"`
	Enabled   bool `json:"enabled,omitempty"`
}

type Parent struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Contact struct {
	Method string `json:"method,omitempty"`
	To     string `json:"to,omitempty"`
}

type SendAfter struct {
	TimeAmount int `json:"timeAmount,omitempty"`
	TimeUnit   string `json:"timeUnit,omitempty"`
}

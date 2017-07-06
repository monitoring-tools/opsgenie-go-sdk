package userv2

type Escalation struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	OwnerTeam OwnerTeam `json:"ownerTeam,omitempty"`
	Rules []Rule `json:"rules,omitempty"`
}

type OwnerTeam struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Rule struct {
	Condition string `json:"condition,omitempty"`
	NotifyType string `json:"notifyType,omitempty"`
	Delay Delay `json:"delay,omitempty"`
	Recipient Recipient `json:"recipient,omitempty"`
}

type Delay struct {
	TimeAmount int `json:"timeAmount,omitempty"`
	TimeUnit string `json:"timeUnit,omitempty"`
}

type Recipient struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

package notificationv2

// Criteria defines the conditions that will be checked before applying notification rules and type of the operations
// that will be applied on these conditions.
type Criteria struct {
	Type       string `json:"type"`
	Conditions []Condition `json:"conditions"`
}

// Condition defines the conditions that will be checked before applying notification rules.
type Condition struct {
	Field         string `json:"field"`
	Key           string `json:"key"`
	Not           bool `json:"not"`
	Operation     string `json:"operation"`
	ExpectedValue string `json:"expectedValue"`
	Order         int `json:"order"`
}

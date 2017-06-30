package notificationv2

type Criteria struct {
	Type       string `json:"type"`
	Conditions []Condition `json:"conditions"`
}

type Condition struct {
	Field         string `json:"field"`
	Key           string `json:"key"`
	Not           bool `json:"not"`
	Operation     string `json:"operation"`
	ExpectedValue string `json:"expectedValue"`
	Order         int `json:"order"`
}

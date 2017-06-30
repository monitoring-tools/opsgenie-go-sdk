package notificationv2

type Repeat struct {
	LoopAfter int `json:"loopAfter"`
	Enabled   bool `json:"enabled"`
}

package userv2

type GetUserResponse struct {
	User User `json:"data"`
	Expandable []string `json:"expandable"`
	*ResponseMeta
}

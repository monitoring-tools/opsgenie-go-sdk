package userv2

type ListUserTeamsResponse struct {
	Teams []Team `json:"data"`
	*ResponseMeta
}

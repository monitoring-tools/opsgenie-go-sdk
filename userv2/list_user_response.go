package userv2

type ListUserResponse struct {
	Users []User `json:"data"`
	*ResponseMeta
}

package userv2

type ListUsersResponse struct {
	Users []User `json:"data"`
	*ResponseMeta
}

package userv2

// DeleteUserResponse is a response of deleting user result
type DeleteUserResponse struct {
	Result string `json:"result"`
	*ResponseMeta
}

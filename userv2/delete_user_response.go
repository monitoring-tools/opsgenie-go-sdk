package userv2

type DeleteUserResponse struct {
	Result string `json:"result"`
	*ResponseMeta
}

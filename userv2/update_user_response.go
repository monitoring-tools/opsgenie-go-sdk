package userv2

type UpdateUserResponse struct {
	Result string `json:"result"`
	*ResponseMeta
}

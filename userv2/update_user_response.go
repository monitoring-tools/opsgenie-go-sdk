package userv2

const SuccessUpdateResultStatus = "updated"

type UpdateUserResponse struct {
	Result string `json:"result"`
	*ResponseMeta
}

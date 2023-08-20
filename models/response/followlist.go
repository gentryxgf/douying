package response

type FollowListResponse struct {
	Response
	FollowList []UserData `json:"user_list"`
}

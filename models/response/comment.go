package response

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment"`
}

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list"`
}

type Comment struct {
	ID         int64  `json:"id"`
	User       Author `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

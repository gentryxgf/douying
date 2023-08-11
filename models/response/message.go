package response

type Message struct {
	SendUserID     int64  `json:"send_user_id"`
	SendUsername   string `json:"send_username"`
	SendUserAvatar string `json:"send_user_avatar"`
	RevUserID      int64  `json:"rev_user_id" `
	RevUsername    string `json:"rev_username"`
	RevUserAvatar  string `json:"rev_user_avatar"`
	IsRead         bool   `json:"is_read" `
	Content        string `json:"content"`
	CreateTime     string `json:"create_time"`
}

type SendMessageResponse struct {
	Response
}

type MessageChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}

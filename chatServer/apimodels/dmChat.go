package apimodels

type DirectMessageChat struct {
	UserID   string    `json:"user_id"`
	Username string    `json:"user_name"`
	Messages []Message `json:"messages"`
}

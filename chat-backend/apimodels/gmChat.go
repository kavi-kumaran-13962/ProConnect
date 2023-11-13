package apimodels

type GroupMessageChat struct {
	GroupID   string    `json:"group_id"`
	Groupname string    `json:"groupname"`
	Messages  []Message `json:"messages"`
}

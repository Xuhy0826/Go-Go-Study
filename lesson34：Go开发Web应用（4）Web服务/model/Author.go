package model

// Author 表示帖子或者评论的作者
type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

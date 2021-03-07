package model

import (
	"fmt"
	"time"
)

// Post 表示论坛中的帖子
type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

// Comment 表示帖子的评论
type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// Author 表示帖子或者评论的作者
type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// News 新闻
type News struct {
	Id          int      `json:"id"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Author      Author   `json:"author"`
	PublishTime JSONTime `json:"publishTime"`
}
type JSONTime struct {
	time.Time
}

func (t *JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.Format("2006-01-02 15:04:05"))), nil
}

func (t *JSONTime) UnmarshalJSON(data []byte) error {
	var err error

	t.Time, err = time.Parse(`"2006-01-02 15:04:05"`, string(data))
	if err != nil {
		return err
	}

	return nil
}

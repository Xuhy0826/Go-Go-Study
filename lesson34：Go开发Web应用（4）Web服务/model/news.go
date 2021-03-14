package model

import (
	"fmt"
	"time"
)

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

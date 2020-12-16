package main

import (
	"fmt"
	"time"
)

type testEntity struct {
	id         int
	msg        string
	createTime time.Time
}

func (entity testEntity) String() string {
	return fmt.Sprintf("entity : {id: %v, msg: %v, createTime: %v }", entity.id, entity.msg, entity.createTime.Format("2006-01-02 15:04:05"))
}

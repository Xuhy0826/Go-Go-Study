package main

import (
	"fmt"
	"strings"
)

var t interface {
	talk() string
}

//满足接口t（1）
type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

//满足接口t（2）
type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

//接口往往声明为类型，并以-er结尾
type talker interface {
	talk() string
}

//入参为任何满足talker接口的值
func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type starship struct {
	laser
}

func main() {
	fmt.Println("lesson17 Interface")
	t = martian{}
	fmt.Println(t.talk()) //nack nack

	t = laser(3)
	fmt.Println(t.talk()) //pew pew pew

	shout(martian{}) //NACK NACK
	shout(laser(2))  //PEW PEW

	s := starship{laser(2)}
	fmt.Println(s.talk()) //pew pew
	shout(s)              //PEW PEW
}

package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

// func birthday(p person) {
// 	p.age++
// }

func birthday(p *person) {
	p.age++
}

func (p *person) birthday() {
	p.age++
}

type stats struct {
	level             int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name string
	stats
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martain struct{}

func (m martain) talk() string {
	return "neck neck"
}

type laser struct{}

func (l *laser) talk() string {
	return "pew pew"
}

func main() {
	fmt.Println("lesson19 Pointer(2)")

	// jack := person{
	// 	name: "Jack",
	// 	age:  12,
	// }
	//函数是以传值方式传递形参
	//birthday(&jack)
	//fmt.Println(jack.age) //13

	jack := &person{
		name: "Jack",
		age:  10,
	}
	jack.birthday()
	fmt.Println(jack.age) //13

	tom := person{
		name: "Tom",
		age:  20,
	}
	tom.birthday()
	fmt.Println(jack.age) //21

	yasuo := character{name: "Yasuo"}

	//levelUp(yasuo.stats) //error: cannot use yasuo.stats (type stats) as type *stats in argument to levelUp

	levelUp(&yasuo.stats)
	fmt.Printf("%+v\n", yasuo) //{name:Yasuo stats:{level:1 endurance:56 health:280}}

	shout(martain{})  //NECK NECK
	shout(&martain{}) //NECK NECK

	//shout(laser{}) //error: laser does not implement talker
	shout(&laser{}) //PEW PEW
}

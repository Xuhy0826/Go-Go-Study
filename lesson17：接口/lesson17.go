package main

import (
	"fmt"
	"strings"
	"time"
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

//将“地球时间”转成“x星时间”（坏版本）
// func xstardate(t time.Time) float64 {
// 	doy := float64(t.YearDay())
// 	h := float64(t.Hour()) / 24.0
// 	return 1000 + doy + h
// }

type xstadater interface {
	YearDay() int
	Hour() int
}

//将“地球时间”转成“x星时间”（好版本）
func xstardate(t xstadater) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

type marsTime int

func (s marsTime) YearDay() int {
	return int(s % 668)
}
func (s marsTime) Hour() int {
	return 0
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

	//顺路探究下时间类型
	t := time.Now()
	fmt.Println(t) //2020-11-23 22:51:33.8848173 +0800 CST m=+0.003079501
	//格式输出
	fmt.Println(t.Format("2006-01-02 15:04:05")) //2020-11-23 22:51:33
	fmt.Println(time.Now().Unix())               //1606143093
	fmt.Println(t.Year())                        //2020
	fmt.Println(t.YearDay())                     //328
	fmt.Println(t.Month())                       //November
	fmt.Println(t.Date())                        //2020 November 23
	fmt.Println(t.Day())                         //23

	today := time.Date(2020, 11, 23, 22, 59, 10, 0, time.UTC)
	fmt.Println(today) //2020-11-23 22:59:10 +0000 UTC

	fmt.Printf("%.1f Curiosity has landed\n", xstardate(today)) //1328.9 Curiosity has landed

	m := marsTime(1452)
	fmt.Printf("%.1f Curiosity has landed\n", xstardate(m)) //1116.0 Curiosity has landed

}

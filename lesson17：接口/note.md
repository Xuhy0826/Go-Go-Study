# 接口

## 接口类型
和其他常见的编程语言一样，Go也有接口，并且其含义是类似的。类型通过方法来表达自己的行为，而接口是来规定类型必须满足的方法，就像是一种约束或者契约。  
首先如何声明接口。使用的关键字是毫无意外的`interface`。
```go
var t interface {
	talk() string
}
```
无论什么类型，只要存在满足接口的方法，就能成为变量`t`的值。
```go
//类型martian 满足了接口t 
type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

//类型laser 满足了接口t
type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", 3)
}

func main() {
	fmt.Println("lesson17 Interface")
	t = martian{}
	fmt.Println(t.talk()) //nack nack

	t = laser(3)
	fmt.Println(t.talk()) //pew pew pew
}
```
`martian`和`laser`两个完全不同的类型都关联了一个空入参且返回参数为`string`的`talk`方法，那么它们就都可以被赋值给变量`t`。
* 为了复用，一般会将接口声明为类型。按照惯例，接口类型的名称常常以`-er`作为后缀。举个例子
```go
···

type talker interface{
    talk() string
}

//入参为任何满足talker接口的值
func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	shout(martian{}) //NACK NACK
	shout(laser(2))  //PEW PEW
}
```
上一节学习了struct嵌入的特性，下面将满足接口的类型嵌入另一个struct中
```go
···

type starship struct {
	laser
}

func main() {
	s := starship{laser(2)}
	fmt.Println(s.talk()) //pew pew
	shout(s)              //PEW PEW
}
```
`laser`嵌入`starship`中，那么直接调用`starship`的`talk()`方法会将`laser`的`talk()`自动转发。更牛逼的是，通过这个转发让`starship`间接的满足了`talker`接口，所以就可以当做入参传入`shout`函数中了。

## 探索接口
先顺路看一下Go的时间类型。需要引入time包。
```go
//顺路探究下时间类型
t := time.Now()
fmt.Println(t) //output: 2020-11-23 22:51:33.8848173 +0800 CST m=+0.003079501
//格式输出
fmt.Println(t.Format("2006-01-02 15:04:05")) //output: 2020-11-23 22:51:33
fmt.Println(time.Now().Unix())               //output: 1606143093
fmt.Println(t.Year())                        //output: 2020
fmt.Println(t.YearDay())                     //output: 328
fmt.Println(t.Month())                       //output: November
fmt.Println(t.Date())                        //output: 2020 November 23
fmt.Println(t.Day())                         //output: 23

today := time.Date(2020, 11, 23, 22, 59, 10, 0, time.UTC)
fmt.Println(today) //output: 2020-11-23 22:59:10 +0000 UTC
```
现在假设需要个时间转换的方法，看例子
```go
//将“地球时间”转成“x星时间”
func xstardate(t time.Time) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour())
	return 1000 + doy + h
}

func main() {
	today := time.Date(2020, 11, 23, 22, 59, 10, 0, time.UTC)
	fmt.Printf("%.1f 探险号飞船着陆\n", xstardate(today)) //output: 1328.9 探险号飞船着陆
}
```
但是现在就存在一个问题了，这个函数只能将“地球时间”进行转换，因为入参类型是固定的`time.Time`。为了达到通用性来解决这个问题，就可以使用接口。   
先声明接口。
```go
type xstadater interface {
	YearDay() int
	Hour() int
}

func xstardate(t xstadater) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}
```
这样定义后方法`xstardate`就具有通用性了，比如现在其他星球的时间只要具有`YearDay()`和`Hour()`就可以进行转换。
```go
type marsTime int

func (s marsTime) YearDay() int {
	return int(s % 668)
}
func (s marsTime) Hour() int {
	return 0
}

func main() {
	//既可以转换“地球时间”
	today := time.Date(2020, 11, 23, 22, 59, 10, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", xstardate(today)) //1328.9 Curiosity has landed

	//也可以转换火星时间
	m := marsTime(1452)
	fmt.Printf("%.1f Curiosity has landed\n", xstardate(m)) //1116.0 Curiosity has landed
}
```
从这里也可以看出Go的灵活，像Java、C#这样的语言可不能将时间类型显示声明自己实现了`xstardater`接口。

## 满足接口
Go标准库导出了很多只有单个方法的接口，可以在自己的代码中实现它们。
> Go通过简单的、通常只有单个方法的接口······来鼓励组合而不是继承，这些接口在各个组件之间形成了简明易懂的界限。 —— Rob Pike   

比如fmt包就声明了以下所示的Stringer接口:
```go
type Stringer interface{
	String() string
}
```
这样一来，只要一个类型关联了`String`方法，那么它的返回值就能够为`Println`，`Sprintf`等函数所用。
```go
type location struct{
	lat, long float64
}

func (l location) String() string{
	return fmt.Sprintf("%v %v", l.lat, l.long)
}

func main() {
	curiosity := location(-4.3413, 12.473)
	fmt.Println(curiosity)
}
```

## 嵌入类型和接口的实现
前面`starship`的例子也看到，由于Golang的嵌入类型的方法提升，如果内部类型实现了接口，那么外部类型也可以“间接”实现接口。看下面的示例

```go
package main

import{
	"fmt"
}

//声明接口
type notifier interface{
	notify()
}

type user struct{
	name string
	email string
}

//user类型值的指针实现接口
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", 
	u.name,
	u.email)
}

//admin代表一个拥有权限的管理员用户
//admin类型嵌入了user类型
type admin struct{
	user
	level string
}

func main(){
	//创建admin用户
	ad := admin{
		user : user{
			name : "john smith",
			email : "john@yahoo.com:",
		},
		level : "super",
	}

	//实现了接口notifier的内部类型user的方法被提升到了外部，所以此时admin类型也实现了接口
	sendNotification(&ad)
}

func sendNotification(n notifier){
	n.notify()
}
```
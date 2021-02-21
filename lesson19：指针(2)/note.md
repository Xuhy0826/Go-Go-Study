# 指针(2)：Pointer

## 实现修改
通过指针可以实现跨越函数和方法边界的修改

### 将指针用作形参
通过前面的学习我们知道函数是以传值的方式传递形参的。
```go
type person struct {
	name string
	age  int
}
//入参是 person 类型
func birthday(p person) {
	p.age++
}

func main() {
	jack := person{
		name: "Jack",
		age:  12,
	}

    //这里会传入一个jack的副本
	birthday(jack)
    //原jack的字段值不会改变
	fmt.Println(jack.age) //12    
}
```
但当指针被被传递至函数时，函数将接收到传入内存地址的副本，在此之后函数就可以通过解引用内存地址来修改指针指向的值。
```go
//入参改为 person 指正类型
func birthday(p *person) {
	p.age++
}

jack := person {
    name: "Jack",
    age:  12,
}

birthday(&jack)
fmt.Println(jack.age) //13
```
### 指针接收者
与形参的写法类似，将指针用作方法接收者（Receiver）时，便可以实现对接收者字段的修改，看示例。
```go
func (p *person) birthday() {
	p.age++
}

func main() {
	jack := &person{
		name: "Jack",
		age:  12,
	}

	jack.birthday()
	fmt.Println(jack.age) //13
}
```
其实，就算在声明struct时不写`&`，仍然可以正常运行。因为Go在变量通过点标记调用方法是会自动使用`&`取得变量的内存地址。
```go
func (p *person) birthday() {
	p.age++
}

func main() {
	tom := person{
		name: "Tom",
		age:  20,
	}
	tom.birthday()
	fmt.Println(jack.age) //21
}
```
可以看到就算不写`(&tom).birthday()`也可以正常运行。   
当然不是所有涉及到struct的方法都要以指针作为参数，需要视情况而定。

### 内部指针
Go提供了叫做**内部指针**的特性，来确定struct中的指定的字段的内存地址。
```go
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

func main() {
	yasuo := character{name: "Yasuo"}

	levelUp(&yasuo.stats)
	fmt.Printf("%+v", yasuo) //{name:Yasuo stats:{level:1 endurance:56 health:280}}
}
```
类似于`&yasuo.stats`这样就可以提供指向struct内部的指针。

### 修改数组
虽然前面说Go中更倾向于使用切片而不是数组，但是也难免会遇到使用数组更加合理的情况。同样使用指针也可以实现对数组元素进行修改的方法。
```go
func reset(board *[8][8]rune){
    board[0][0] = 'r'
}

func main(){
    var board [8][8]rune
    reset(&board)
    fmt.Printf("%c", board[0][0])   //r
}
```

## 隐式指针
不是所有修改都需要显式的使用指针，Go有些地方会“暗中”使用指针。
* 映射也是指针   
前面我们知道，映射的传值和赋值时传递的都不是副本。因为映射实际上就是一种隐式的指针。

* 切片指向数组   
切片在指向数组元素的时候也是使用了指针。之前提到切片其实是一个结构体类型。
```go
type slice struct {
    array unsafe.Pointer 
    len   int
    cap   int
}
```
切片内部三个字段
1. 指向数组的指针
2. 切片的长度
3. 切片的容量  

当切片被直接传递至函数或者方法的时候，切片的内部指针就可以对底层数组数据进行修改。   
**【注】**：指向切片本身的指针唯一的用处就是修改切片本身，包括长度、容量及起始位置。

## 指针和接口
先看示例A，其实和接口那节用到的示例一样
【示例A】
```go
type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

//martain类型实现了接口talker
type martain struct{}

func (m martain) talk() string {
	return "neck neck"
}

func main(){
	
	shout(martain{})  //NECK NECK
	shout(&martain{}) //NECK NECK
}
```
在上面，无论是传递`martian`变量还是传递指向`martian`变量的指针，都可以满足`talker`接口。如果方法使用的是指针接收者，那么情况就不同了。
```go
type laser struct{}

func (l *laser) talk() string {
	return "pew pew"
}

func main(){
	//shout(laser{}) //error: laser does not implement talker
	shout(&laser{}) //PEW PEW
}

```
如果方法使用的是指针接收者，那么只能使用指针来调用该方法。

## 明智地使用指针
**切记：不要过度使用指针**
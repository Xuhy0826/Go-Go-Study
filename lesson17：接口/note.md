# 接口

## 接口类型
和其他常见的编程语言一样，Go也有接口，并且其含义是类似的。类型通过方法表达自己的行为，而接口通过规定类型必须满足的方法来声明。首先如何声明接口。
```
var t interface {
	talk() string
}
```
无论什么类型，只要存在满足接口的方法，就能成为变量`t`的值。
```
//满足接口t（1）
type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

//满足接口t（2）
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
martian和laser两个完全不同的类型都关联了一个空入参且返回参数为string的`talk`方法，那么它们就都可以被赋值给变量`t`。
* 为了复用，一般会将接口声明为类型。按照惯例，接口类型的名称常常以`-er`作为后缀。举个例子
```
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
```
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
# 函数

## 函数声明
#### 使用`func`关键字声明函数
> func  SayHello( word   string )  string   
> 关键字 函数名   ( 参数名 参数类型 )  返回值类型

**【注意】** 大写字母开头的函数、变量都会被导出，对其他包可见（可理解为访问修饰符为`public`），如果是小写字母开头的话，则对其他包不可见（可理解为访问修饰符为`private`）。
举个简单的函数例子。
```go
//声明函数
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

//使用函数
func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Println(kelvin, "°K is", celsius, "°C") //294 °K is 20.850000000000023 °C
}
```
#### 多入参形式   
> func Unix(sec int64, nsec int64) Time

如果多个入参的参数类型相同，可以像下面这样简写
> func Unix(sec, nsec int64) Time

#### 多返回值形式
在上一节遇到的函数`countdown, err := strconv.Atoi("10")`就是一种多返回值的函数，其声明的写法是
> func Atoi(s string) (i int, err error) 

多个返回值写在括号内，并且返回值可以将名字去掉只保留类型：
> func Atoi(s string) (int, error)

举个简单的例子
```go
func reverse(a, b, c string) (string, string, string) {
	return c, b, a
}
```
或者
```go
func reverse1(a, b, c string) (x string, y string, z string) {
	x, y, z = c, b, a
	return
}
```
#### 可变参数函数
例如常用的`Println`函数，它可以接收一个参数、两个参数或更多参数。并且它可以接收不同类型的参数
```go
fmt.Println("Hello","World")
fmt.Println(186, "seconds")
```
其函数声明如下：
> func Println(a ...interface{}) (n int, err error)
1. ...表示函数的参数的数量是可变的
2. 参数a的类型为`interface{}`，是一个空接口。空接口可以接收所有类型
> 所以，...和interface{}结合到一起，就可以接收任意数量任意类型的参数
看个简单的例子
```go
//addAll 多数相加
func addAll(a int, numbers ...int) int {
	sum := a
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func main() {
	sum := addAll(3, 4, 5, 6, 7)
	fmt.Println(sum)  //output: 25
}
```
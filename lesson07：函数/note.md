# 函数

## 函数声明
* 使用func关键字声明函数
> func  SayHello(word string)  string

> 关键字 函数名(参数名 参数类型) 返回值类型

**【注意】** 大写字母开头的函数、变量都会被导出，对其他包可见（就好像访问修饰符为public），如果是小写字母开头的话，则不行。
* 多参数形式   
> func Unix(sec int64, nsec int64) Time

如果多个参数的函数参数类型相同，可以简写
> func Unix(sec, nsec int64) Time

* 多返回值形式
在之前遇到的函数“countdown, err := strconv.Atoi("10")”就是一种多返回值的函数，其声明的写法是
> func Atoi(s string) (i int, err error)  

简写：返回值可以将名字去掉只保留类型：
> func Atoi(s string) (int, error)

* 可变参数函数
例如常用的Println函数，它可以接收一个参数、两个参数或更多参数。并且它可以接收不同类型的参数
```go
fmt.Println("Hello","World")
fmt.Println(186, "seconds")
```
其函数声明如下：
> func Println(a ...interface{}) (n int, err error)
1. ...表示函数的参数的数量是可变的
2. 参数a的类型为interface{}，是一个空接口。空接口可以接收所有类型
> 所以，...和interface{}结合到一起，就可以接收任意数量任意类型的参数

* 最后写两个例子
```go
package main

import (
	"fmt"
)

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

//addAll 多数相加
func addAll(a int, numbers ...int) int {
	sum := a
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("lesson7 函数")
	kelvin := 294.0
	
	celsius := kelvinToCelsius(kelvin)
	fmt.Println(kelvin, "°K is", celsius, "°C") //294 °K is 20.850000000000023 °C

	sum := addAll(3, 4, 5, 6, 7)
	fmt.Println(sum)  //25
}
```
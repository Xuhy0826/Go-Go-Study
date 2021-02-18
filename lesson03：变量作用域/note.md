# 变量作用域
与其他语言的概念类似。

## 变量作用域
Go作用域通常会随着大括号{}的出现而开启和结束。也有特殊情况，比如包级别的作用域
```go
package main

import (
	"fmt"
	"math/rand"
)

var era = "AD" //这个变量的作用域就是在整个包（package）是可用的

func main() {
	var times = 0
	for times < 10 {
		//变量num的作用域就在这个循环中
		var num = rand.Intn(10) + 1
		fmt.Println(num)
		times++
	}  //num的作用域结束
}  //times的作用域结束
```
## 短声明
短声明写法在GO语言中更加流行，并且功能也更加强大
1. 写法：
```go
count := 10
```
等同于
```go
var count = 10
```
2. for中使用短声明  

短声明不光是少打几个字母那么简单，它还可以用在一些无法使用var关键字的地方,比如说，在for循环中，我们无法使用var来定义count，只能在for循环之前定义好先
```go
var count = 0
for count = 10; count > 0; count-- {
    fmt.Println(count)
}
```
可以简写成
```go
for count := 10; count > 0; count-- {
    fmt.Println(count)
}
```
3. if中使用短声明
```go
if num := rand.Intn(3); num == 0 {
    fmt.Println("Space Adventure")
} else if num == 1 {
    fmt.Println("SpaceX")
} else if num == 2 {
    fmt.Println("Virgin Galatic")
}
```
4. Switch中 使用短声明
```go
switch num := rand.Intn(3); num {
case 0:
    fmt.Println("Space Adventure")
case 1:
    fmt.Println("SpaceX")
case 2:
    fmt.Println("Virgin Galatic")
}
```
总得来说，学会多使用短声明
# 一等函数

在Go语言中，函数是 **一等值** ，可以用在整数、字符串或其他类型能够应用的所有地方。也就是函数好像一个对象那样去使用。
* 可以将函数赋值给变量
* 可以将函数传递给函数
* 可以编写创建并返回函数的函数

## 将函数赋值给变量
【示例1code——sensor.go】
```go
package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

//返回模拟温度的传感器
func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

//返回真实温度的传感器
func realSensor() kelvin {
	return 0
}

func main() {
	fmt.Println("lesson9 First-class functions")

	//【case1】将函数赋值给变量
	sensor := fakeSensor
	fmt.Println(sensor()) //156(随机的)

	sensor = realSensor
	fmt.Println(sensor()) //0

	//fmt.Println(sensor) //报错：
}
```
以上，sensor变量的值是函数本身，而不是调用函数获取的结果值。无论赋值给sensor的是fakeSensor还是realSensor，sensor都可以通过sensor()来调用。   
另外，之所以能将realSensor重新赋值给sensor,是由于realSensor和fakeSensor具有相同的函数签名。

## 将函数传递给其他函数
因为变量既可以指向函数，又可以作为参数传递给函数，那么在Go中函数也可以作为参数传递给其他函数。有点类似C#中委托的用法。   
【示例2code——function-parameter.go】
```go
import (
	"fmt"
	"math/rand"
	"time"
)

//【case2】将函数传递给其他函数
//测量温度，使用传入的传感器测量samples次温度
func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v° K\n", k)
		time.Sleep(time.Second)
	}
}

func main() {
    measureTemperature(3, sensor) //0° K (连续打印三次)
}
```

## 声明函数类型
之前我们使用过type关键字来声明类型，当时使用的底层类型是float64来声明了kelvin类型。同样函数也可以这样玩。
```go
type sensor func() kelvin
```
通过这样的声明之后，代码的可读性可以得到提升，并且之前定义的函数measureTemperature签名可以简写成
```go
func measureTemperature(samples int, s sensor)
```

## 闭包和匿名函数
* Go语言支持匿名函数
* 匿名函数在Go中也称为“函数字面量”。
* 因为函数字面量需要保留外部作用域的变量引用，所以函数字面量都是闭包的。

```go
//匿名函数
var f = func() {
	fmt.Println("Dress up for the masquerade")
}

func main() {
    //调用匿名函数
	f() //Dress up for the masquerade

	//将匿名函数赋值给函数中的变量
	ff := func(message string) {
		fmt.Println(message)
	}
	ff("Go to the party") //Go to the party

	//将匿名函数的声明和执行放在一起写
	func() {
		fmt.Println("function anonymous")
	}() //function anonymous
}
```
再来看一下闭包的情形。
```go
//声明并返回一个匿名函数
func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	newSensor := calibrate(realSensor, 5)
	fmt.Println(newSensor()) //5
}
```
以上，calibrate返回的匿名函数引用了被calibrate函数用作形参的s和offset变量。尽管calibrate函数已经返回了，但是被闭包捕获的变量继续存在，因此sensor仍然能够访问这两个变量。
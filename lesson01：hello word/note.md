# Hello World

## 简介
* Golang是编译型语言
* 开发速度快，性能高效的语言
* 更加简便，更加安全的实现并发
* 灵活的且无继承的类型系统，不像传统的面向对象语言，Go更多的是利用 **组合(Composition)** 的设计模式
* 拥有垃圾回收机制

## 环境搭建
不管是Windows、Mac还是Linux，[官网](https://golang.org/dl/ "官网")都提供了相应的安装包，直接下载安装即可。而且从Go1.16版本（Released@2021-02-17）开始已原生支持Apple Silicon M1。安装完成之后在终端能正常执行`go version`说明Go已经安装成功了。  

### 配置环境变量
通过`go env`命令可以查看当前GO的所有环境变量，两个比较重要的是`GOROOT`和`GOPATH`。
1. `GOROOT`，安装Go的路径
2. `GOPATH`，自定义的开发者的workspace  

### IDE
我选择的是VS Code进行开发，当然还有JetBrains家的GoLand，且已经支持Apple Silicon M1了，不介意收费的可以选择GoLand。

## Hello world
第一次尝鲜
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}
```
* 和大多数语言一样，main函数作为程序主入口。
* package表示当前的代码包含在哪个包里，现在粗浅的理解为命名空间的功能类似。一个go文件有且仅输入一个包，一个包可以有多个go文件
* 通过`import+包名`来引入外部代码

## 继续尝鲜
小试一下Go中的一些语法，不系统的学习什么特定的概念，随手尝试一下Go中的一些简单命令，对Go的风格有个第一映像
#### 几种输出控制台方法
Golang中的几种控制台输出方法，类似C#的`Console.Write()`。常用的几种是`fmt.Println()`, `fmt.Print()`, `fmt.Printf()`；
```go
//打印后不换行
fmt.Print("Hello world \n") //Hello world

//打印后换行
fmt.Println("123456")              //123456
fmt.Println("hello", "world", "!") //hello world !
```
使用Printf函数输出格式化的输出
```go
var s = "hi"
fmt.Printf("%s guys\n", s) //hi guys
fmt.Printf("%v guys\n", s) //hi guys
var i = 20
fmt.Printf("i am %v years old\n", i) //i am 20 years old
```
调整输出的对齐格式，使用%后跟数字和v的方式，中间的数字就表示这个占位符的长度，正数表示靠右对齐，负数表示靠左对齐
```go
fmt.Printf("%-15v %6v\n", "abcdefghijklmno", "123") //abcdefghijklmno    123
fmt.Printf("%-15v %6v\n", "abcdefg", "123456")      //abcdefg         123456
fmt.Printf("%-15v %6v\n", "一二三四五六七八九", "123456")    //一二三四五六七八九       123456
```
#### 基本类型
Golang中包含类似的常用基本类型：int，float，double，bool等

#### 常量和变量
const声明常量，var声明变量。Golang中特有的短声明才是常用的写法，后面再议。
```go
const width = 10
var height = 5
var distance, speed = 5600000, 10080

fmt.Println("area = ", width*height)   //area =  50
fmt.Println("time = ", distance/speed) //time =  555
```
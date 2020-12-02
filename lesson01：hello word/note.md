# Hello World

## 简介
* Golang是编译型语言
* 开发速度快，性能高效的语言
* 更加简便，更加安全的实现并发
* 灵活的且无继承的类型系统，不像传统的面向对象语言，Go更多的是利用**组合(Composition)**的设计模式
* 拥有垃圾回收机制

## 环境搭建
不管是Windows还是Mac，官网都提供了相应的安装包，直接下载安装即可。

## Hello world
第一段程序
```
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
* 几种输出控制台方法：fmt.Println(), fmt.Print(), fmt.Printf()；
```
//几种控制台的打印方法
//打印后不换行
fmt.Print("Hello world \n") //Hello world
//打印后换行
fmt.Println("123456")              //123456
fmt.Println("hello", "world", "!") //hello world !
//使用Printf函数输出格式化的输出
//如何调整输出的对齐格式
//使用%后跟数字和v的方式，中间的数字就表示这个占位符的长度，正数表示靠右对齐，负数表示靠左对齐
fmt.Printf("%-15v %6v\n", "abcdefghijklmno", "123") //abcdefghijklmno    123
fmt.Printf("%-15v %6v\n", "abcdefg", "123456")      //abcdefg         123456
fmt.Printf("%-15v %6v\n", "一二三四五六七八九", "123456")    //一二三四五六七八九       123456
```
* 基本类型：int，float，double，bool等
* 常量const和变量var
```
//const声明常量，var声明变量，没什么特别之处
const width = 10
var height = 5
var distance, speed = 5600000, 10080

fmt.Println("area = ", width*height)   //area =  50
fmt.Println("time = ", distance/speed) //time =  555
```
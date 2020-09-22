package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//字符串类型
	str := "peace"
	fmt.Println(str)

	//使用``，可以方便的定义跨行字符串
	fmt.Println(`
	peace be upon you
	upon you be peace`)

	//rune和byte类型
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang) //960 940 969 33
	//通过使用格式化变量%c，可以将代码点表示成字符
	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang) //π ά ω !

	//通过索引的方式访问字符串中的字符
	message := "shalom"
	ch := message[5]
	fmt.Printf("%c\n", ch)

	//字符串不可被修改
	//message[5] = 'd'  //error：cannot assign to message[5]

	//将字符串解码为符文
	question := "今天星期几？"
	fmt.Println(len(question), "bytes")                    //18 bytes
	fmt.Println(utf8.RuneCountInString(question), "runes") //6 runes

	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes", c, size) //First rune: 今 3 bytes

	//通过range关键字进行迭代，i为索引，c为值。有点python的味道
	//遍历字符串，挨个打印出来
	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}
}

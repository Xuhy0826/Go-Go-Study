package main

import "fmt"

func main() {
	//Println功能有点类似js中console.log方法
	fmt.Println("123456")              //123456
	fmt.Println("hello", "world", "!") //hello world !

	fmt.Println(add(10, 2))             //12
	fmt.Println(reverse("a", "b", "c")) //c b a
	fmt.Println(split(14))              //6 8
}

//返回值为一个，如果参数类型相同则可以省略挨个声明参数类型
func add(a, b int) int {
	return a + b
}

//牛逼的地方来了，支持多个返回值
func reverse(a, b, c string) (string, string, string) {
	return c, b, a
}

//又一个牛逼的写法，命名返回值，没有参数的 return 语句返回已命名的返回值。
//直接返回语句应当仅用在下短函数中。否则会影响代码的可读性。
//如果函数中不给y赋值，则y便是默认值（int就是0）
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

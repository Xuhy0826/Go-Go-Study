//声明本代码所属的包
package main

//导入fmt包，使其可用
import "fmt"

//声明main函数，注意Go中的大括号方式只支持这一种形式
func main() {
	//Println功能有点类似js中console.log方法
	fmt.Println("123456")              //123456
	fmt.Println("hello", "world", "!") //hello world !

	//使用Printf函数输出格式化的输出
	//如何调整输出的对齐格式
	//使用%后跟数字跟v的方式，中间的数字就表示这个占位符的长度（中文的话按每个字按两个字符算），正数表示靠右对齐，复数表示靠左对齐
	fmt.Printf("%-15v %6v\n", "abcdefghijklmno", "123")
	fmt.Printf("%-15v %6v\n", "abcdefg", "123456")
	fmt.Printf("%-15v %6v\n", "一二三四五六七八九", "123456")

	//const声明常量，var声明变量，没什么特别之处
	const width = 10
	var height = 5
	var distance, speed = 5600000, 10080

	fmt.Println("area = ", width*height)
	fmt.Println("time = ", distance/speed)

	fmt.Println(add(10, 2))             //12
	fmt.Println(reverse("a", "b", "c")) //c b a
	fmt.Println(split(14))              //6 8
}

//函数的定义方式
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

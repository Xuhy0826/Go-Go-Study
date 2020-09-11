package main

import (
	"fmt"
	"math"
)

func main() {
	//days := 365.2425 //短声明，days会被Go编译器推断为浮点类型默认的float64类型

	var pi64 = math.Pi
	var pi32 float32 = math.Pi //这样声明的变量才会是float32

	fmt.Println(pi64) //3.141592653589793
	fmt.Println(pi32) //3.1415927

	//零值（默认值）
	var price float64
	fmt.Println(price) //0

	//打印浮点数
	third := 1.0 / 3
	fmt.Println(third) //0.3333333333333333
	//下面使用Printf来格式化输出
	fmt.Printf("%v\n", third)     //0.3333333333333333
	fmt.Printf("%f\n", third)     //0.333333
	fmt.Printf("%.2f\n", third)   //0.33，.2f就是表示小数点后保留2位
	fmt.Printf("%4.2f\n", third)  //0.33，4.2f表示总宽（长）度为4，小数点后保留2位
	fmt.Printf("%5.2f\n", third)  // 0.33，5.2f表示总宽（长）度为5，小数点后保留2位，长度不够使用空格来补
	fmt.Printf("%05.2f\n", third) //00.33，05.2f表示总宽（长）度为5，小数点后保留2位，长度不够使用“0”来补

	//浮点类型的精确性
	//计算机只能通过0和1来表示浮点数，所以浮点数会经常受到舍入错误的影响

	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank) //0.30000000000000004

	//为了尽可能的减少这种精度错误，还可以将乘法计算放到除法计算前面执行

	//浮点类型的精确度问题导致浮点数的比较
	fmt.Println(piggyBank == 0.3) //false
	//解决方案
	fmt.Println(math.Abs(piggyBank-0.3) < 0.0001) //true

}

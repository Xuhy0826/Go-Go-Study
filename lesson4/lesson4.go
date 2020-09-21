package main

import (
	"fmt"
	"math"
	"math/big"
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

	/*****************************骄傲的分割线*********************************/
	//Go会默认将整数推断成int类型
	days := 365
	fmt.Printf("Type %T for %[1]v\n", days) //Type int for 365

	/* 使用uint8来表示颜色rgb值，是个很好的选择：
	 * （1）能将变量限制在合法的范围之内
	 * （2）对于未压缩的图片这种需要按顺序存储大量颜色的场景，可以极大的节省空间
	 */
	var red, green, blue uint8 = 0, 141, 213
	fmt.Printf("color：#%02x%02x%02x;\n", red, green, blue) //color：#008dd5;

	//注意“整型环绕”
	var numberA uint8 = 255 //到达类型最大值
	numberA++
	fmt.Println(numberA) //0	环绕

	//使用大数：就是特别大的数
	var distance int64 = 41.3e12 //就是41.3 * 10的12次方
	fmt.Println("distance is ", distance)

	//使用“大数”类型的方法1
	lightSpeed := big.NewInt(299792)
	fmt.Println("distance is ", lightSpeed)

	//使用“大数”类型的方法2
	distance1 := new(big.Int)
	distance1.SetString("24000000000000000000000", 10)
	fmt.Println("distance1 is ", distance1)

	//使用常量来承载很大的数不会引发溢出异常，并且可以正常使用
	const distanceConst = 240000000000000000000000
	fmt.Println("Andromeda Galaxy is ", distanceConst/299792/86400)
}

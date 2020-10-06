package main

import (
	"fmt"
)

func main() {
	//type关键字声明新类型
	type celsius float64
	var temperature celsius = 20
	fmt.Println(temperature) //20

	//celsius与float64具有相同的行为
	const degrees = 20
	var temperature2 celsius = degrees
	temperature2 += 10
	fmt.Println(temperature2) //30

	//celsius与float64不能混用
	//var warmUp float64 = 10
	//temperature += warmUp //报错

	//通过自定义类型可提高代码的可读性，比如20摄氏度和20华氏度就不是一回事儿
	type fahrenheit float64

	var c celsius = 20
	var f fahrenheit = 20

	// if c == f {  //报错

	// }
	// c += f  //报错
}

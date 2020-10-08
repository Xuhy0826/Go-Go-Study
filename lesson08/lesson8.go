package main

import (
	"fmt"
)

//type关键字声明新类型
type celsius float64
type kelvin float64

//kelvinToCelsius converts °K to °C
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15) //需要使用类型转换
}

//为kelvin类型声明方法
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func main() {
	//使用声明的自定义类型
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
	//type fahrenheit float64

	//var c celsius = 20
	//var f fahrenheit = 20

	// if c == f {  //报错

	// }
	// c += f  //报错

	//使用引入了新类型的函数
	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	fmt.Println(k, "°K is", c, "°C") //294 °K is 20.850000000000023 °C

	//使用方法
	c1 := k.celsius()
	fmt.Println(k, "°K is", c1, "°C") //294 °K is 20.850000000000023 °C
}

package main

import (
	"fmt"
	"math/rand"
	"time"
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

//【case2】将函数传递给其他函数
//测量温度，使用传入的传感器测量samples次温度
func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v° K\n", k)
		time.Sleep(time.Second)
	}
}

//匿名函数
var f = func() {
	fmt.Println("Dress up for the masquerade")
}

//sensor函数类型
type sensor func() kelvin

//声明并返回一个匿名函数
func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	fmt.Println("lesson9 First-class functions")

	//【case1】将函数赋值给变量
	sensor := fakeSensor
	fmt.Println(sensor()) //156(随机的)

	sensor = realSensor
	fmt.Println(sensor()) //0

	//fmt.Println(sensor) //报错：

	measureTemperature(3, sensor) //0° K (连续打印三次)

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

	newSensor := calibrate(realSensor, 5)
	fmt.Println(newSensor()) //5
}

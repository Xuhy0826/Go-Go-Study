package main

import (
	"fmt"
)

//kelvinToCelsius converts °K to °C
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

//addAll 多数相加
func addAll(a int, numbers ...int) int {
	sum := a
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("lesson7 函数")
	kelvin := 294.0

	celsius := kelvinToCelsius(kelvin)
	fmt.Println(kelvin, "°K is", celsius, "°C") //294 °K is 20.850000000000023 °C

	sum := addAll(3, 4, 5, 6, 7)
	fmt.Println(sum) //25
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("lesson11 切片")

	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	terrestrial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]
	fmt.Println(terrestrial, gasGiants, iceGiants) //[Mercury Venus Earth Mars] [Jupiter Saturn] [Uranus Neptune]

	//通过索引访问切片
	fmt.Println(gasGiants[1]) //Saturn

	//创建切片的切片
	giants := planets[4:8]
	gas := giants[0:2]
	ice := giants[2:4]
	fmt.Println(gas, ice) //[Jupiter Saturn] [Uranus Neptune]

	//修改切片的值会影响原数组和其他切片
	iceGiantsMarkII := iceGiants
	fmt.Println(iceGiantsMarkII) //[Uranus Neptune]
	iceGiants[1] = "Poseidon"
	fmt.Println(iceGiantsMarkII) //[Uranus Poseidon] 发生了变化
	fmt.Println(ice)             //[Uranus Poseidon]

	//切片可以简写，利用默认值
}

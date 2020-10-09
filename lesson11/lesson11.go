package main

import (
	"fmt"
	"sort"
	"strings"
)

//遍历切片，消除空格
func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

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
	var slice1 = planets[:3]
	var slice2 = planets[4:]
	var slice3 = planets[:]

	fmt.Println(slice1) //[Mercury Venus Earth]
	fmt.Println(slice2) //[Jupiter Saturn Uranus Poseidon]
	fmt.Println(slice3) //[Mercury Venus Earth Mars Jupiter Saturn Uranus Poseidon]

	//字符串的切分
	neptune := "Neptune"
	tune := neptune[3:]
	fmt.Println(tune) //tune

	//切分字符串时，索引是按照字节号码而不是符文号码
	question := "你在学习Go吗？"
	fmt.Println(question[:6]) //你在

	//切片的复合字面量
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Printf("%T", dwarfs) //[]string

	countries := []string{" China ", "  Japan", " USA"}
	hyperspace(countries)
	fmt.Println(strings.Join(countries, "")) //[]stringChinaJapanUSA

	//使用sort方法
	sort.StringSlice(dwarfs).Sort()
	fmt.Println(dwarfs) //[Ceres Eris Haumea Makemake Pluto]

	//简化，自动执行类型转换和排序
	sort.Strings(dwarfs)
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("lesson10 Array")

	//Array of planets
	var planets [8]string
	planets[0] = "Mercury" //Assigns a planet at index 0
	planets[1] = "Venus"
	planets[2] = "Earth"

	earth := planets[2] //retrieves the planet at index 2
	fmt.Println(earth)  //Earth

	//获取数组的长度，通过内建函数len()即可
	fmt.Println(len(planets)) //8

	//复合字面值 初始化数组
	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	//试试其他奇怪的写法
	dwarfs2 := [4]string{"a", "b", "c"} //正常初始化
	fmt.Println(dwarfs2[3])             //输出空
	//或者
	planetCollection := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	fmt.Println(len(planetCollection)) //8
	fmt.Println(planetCollection)      //[Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]

	//使用索引来为数组赋值
	array := [5]int{1: 10, 3: 30}
	fmt.Printf("%+v\n", array) //[0 10 0 30 0]

	//for循环遍历数组
	for i := 0; i < len(dwarfs); i++ {
		dwarf := dwarfs[i]
		fmt.Println(i, dwarf)
	}

	//range 遍历数组
	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}

	planetsMarkII := planets //会将planets的完整副本赋值给planetMarkII
	planets[2] = "whoops"
	//改变了原数组planets，数组planetsMarkII是副本，所以不受影响
	fmt.Println(planets[2])       //whoops
	fmt.Println(planetsMarkII[2]) //Earth

	//二维数组
	var board [8][8]string
	board[0][0] = "r"
	board[0][7] = "r"
	for column := range board[1] {
		board[1][column] = "p"
	}
	fmt.Println(board) //[[r       r] [p p p p p p p p] [       ] [       ] [       ] [       ] [       ] [       ]]
}

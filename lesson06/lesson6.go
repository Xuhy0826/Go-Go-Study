package main

import (
	"fmt"
	"strconv"
)

func main() {
	age := 41
	//整数转浮点
	marsAge := float64(age)

	marsDays := 687.0
	earthDays := 365.2425

	marsAge = marsAge * earthDays / marsDays

	fmt.Println("i am", marsAge, "years old on Mars.") //i am 21.797587336244543 years old on Mars.

	//浮点转整型
	fmt.Println(int(earthDays)) //365

	//rune/byte → string
	var pi rune = 960
	var alpha rune = 940
	fmt.Println(string(pi), string(alpha)) //π ά

	//数字转string
	//方法一
	countdown := 10
	str := "Launch in T minus " + strconv.Itoa(countdown) + " seconds"
	fmt.Println(str) //Launch in T minus 10 seconds

	//方法二
	str = fmt.Sprintf("Launch in T minus %v seconds", countdown)
	fmt.Println(str) //Launch in T minus 10 seconds

	//string转数字
	count, err := strconv.Atoi("10")
	if err != nil {
		//出错
	}
	fmt.Println(count) //10

	//bool转string
	launch := false
	launchText := fmt.Sprintf("%v", launch)
	fmt.Println("Ready for launch:", launchText) //Ready for launch: false

}

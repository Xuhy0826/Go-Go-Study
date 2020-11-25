package main

import "fmt"

func main() {
	fmt.Println("lesson20 NIL")

	var i int
	var s string
	var p *string

	fmt.Printf("%v\n", i) //0
	fmt.Printf("%v\n", s) // (空字符串)
	fmt.Printf("%v\n", p) // <nil>
	//如果试图对nil指针进行解引用
	//fmt.Printf("%v\n", *p) //panic: runtime error: invalid memory address or nil pointer dereference

	//解决方案：
	var nowhere *int

	if nowhere != nil {
		fmt.Println(nowhere)
	}

	//（1）nil函数值
	var fn func(a, b int) int
	fmt.Println(fn == nil) //true

	//（2）nil切片
	var soup []string
	fmt.Println(soup == nil) //true

	//range可以处理nil
	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}
	//len、append也可以处理nil
	fmt.Println(len(soup)) //0
	soup = append(soup, "onion", "carrot")
	fmt.Println(soup) //[onion carrot]

	//（3）nil映射
	var souplist map[string]int
	fmt.Println(souplist == nil) //true

	measurement, ok := souplist["onion"]

	if ok {
		fmt.Println(measurement)
	}

	for ingredient, measurement := range souplist {
		fmt.Println(ingredient, measurement)
	}

	//souplist["onion"] = 1 //panic: assignment to entry in nil map

	//（4）nil接口
	var v interface{}
	fmt.Printf("%T %v %v\n", v, v, v == nil) //<nil> <nil> true

	var po *int
	v = po
	fmt.Printf("%T %v %v\n", v, v, v == nil) //*int <nil> false

	//格式化变量%#v可以同时打印出变量的类型和值
	fmt.Printf("%#v", v) //(*int)(nil)
}

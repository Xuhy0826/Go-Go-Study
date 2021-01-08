package main

import "fmt"

func main() {
	//声明包含5个元素的整型指针的数组
	array := [5]*int{0: new(int), 1: new(int)}

	//为索引为0和1的元素赋值
	*array[0] = 10
	*array[1] = 20

	fmt.Printf("%+v", array)

	var array1 [5]*int

	array1 = array

}

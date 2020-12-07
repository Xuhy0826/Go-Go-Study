package main

import (
	"fmt"

	"demo.Pkg/myutils"
)

func main() {
	sum := myutils.UtilAdd(1, 3)
	fmt.Println(sum) //4

	res := myutils.UtilMultiplyAll(3, 4, 5, 6, 7)
	fmt.Println(res) //2520

	fmt.Println("myutils's name is", myutils.Name)
	//myutils's name is xuhyUtil
}

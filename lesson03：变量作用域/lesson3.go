//lesson03：变量作用域
package main

import (
	"fmt"
	"math/rand"
)

var era = "AD" //这个变量的作用域就是在整个包（package）是可用的

func main() {
	fmt.Println("lesson3：变量作用域")

	var times = 0

	for times < 10 {
		//变量num的作用域就在这个循环中
		var num = rand.Intn(10) + 1
		fmt.Println(num)

		times++
	}

	//短声明
	//(1)写法：
	// var count = 10
	// count := 10 //这和上一行等价，这就是短声明
	//短声明不光是少打几个字母那么简单，它还可以用在一些无法使用var关键字的地方
	//比如说，在for循环中，我们无法使用var来定义count，只能在for循环之前定义好先
	var count = 0
	for count = 10; count > 0; count-- {
		fmt.Println(count)
	}
	fmt.Println(count)
	//这样会使得count作用域有点过大了，因为count在for循环之外没有什么存在的意义
	//使用短声明，就可以避免这种尴尬
	for count := 10; count > 0; count-- {
		fmt.Println(count)
	}

	//再比如，在if语句中使用短声明
	if num := rand.Intn(3); num == 0 {
		fmt.Println("Space Adventure")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else if num == 2 {
		fmt.Println("Virgin Galatic")
	}

	//在比如，在Switch中 使用短声明
	switch num := rand.Intn(3); num {
	case 0:
		fmt.Println("Space Adventure")
	case 1:
		fmt.Println("SpaceX")
	case 2:
		fmt.Println("Virgin Galatic")
	}
}

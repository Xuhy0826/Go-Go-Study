package main

import (
	"fmt"

	"demo26/counters"
	"demo26/entities"
)

func main() {
	//************** 例1 ****************
	//创建一个未公开的类型的变量
	//counter := counters.alertCounter(10) //error! : undeclared name: counters

	// 使用 counters 包公开的 New 函数来创建一个未公开的类型的变量
	counter := counters.New(10)

	fmt.Printf("Counter: %d\n", counter)

	//************** 例2 ****************
	// 创建 entities 包中的 User 类型的值
	// u := entities.User{
	// 	Name:  "Bill",
	// 	email: "bill@email.com",
	// }

	// 创建 entities 包中的 Admin 类型的值
	a := entities.Admin{
		Rights: 10,
	}
	a.Name = "Bill"
	a.Email = "bill@email.com"
	fmt.Printf("User: %v\n", a)
}

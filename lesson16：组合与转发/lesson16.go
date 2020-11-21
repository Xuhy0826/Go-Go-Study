package main

import "fmt"

//员工
//(1) 一般的写法
/*
type employee struct {
	id         int64
	name       string
	department department
	account    account
}
*/
//(2) 实现struct转发的写法
type employee struct {
	id   int64
	name string
	department
	account
}

//部门
type department struct {
	name string
	code string
}

//账户
type account struct {
	logID string
	level int64
}

//为account类型绑定方法
func (acc account) salary() int64 {
	return acc.level * 2500
}

//为employee类型绑定方法，转发account的salary方法
// func (e employee) salary() int64 {
// 	return e.account.salary()
// }

//解释“命名冲突”
//继续为department绑定一个同名的salary方法
func (dept department) salary() int64 {
	return 2500 * int64(len(dept.code))
}

//来解决“命名冲突”
func (e employee) salary() int64 {
	return e.account.salary()
}

func main() {
	fmt.Println("lesson16 Composition and Forwarding")

	dept := department{"Production", "DP"}
	acc := account{logID: "12345", level: 2}

	jack := employee{
		id:         1001,
		name:       "jack",
		department: dept,
		account:    acc,
	}

	fmt.Printf("%+v\n", jack)                                     //{id:1001 name:jack department:{name:Production code:DP} account:{logID:12345 level:2}}
	fmt.Printf("jack's department is %v\n", jack.department.name) //jack's department is Production

	fmt.Printf("jack's salary is %v now\n", jack.account.salary()) //jack's salary is 5000 now

	fmt.Printf("jack's salary is %v now\n", jack.salary()) //jack's salary is 5000 now

	fmt.Println("jack's logID is ", jack.account.logID) //jack's logID is  12345
	fmt.Println("jack's logID is ", jack.logID)         //jack's logID is  12345
}

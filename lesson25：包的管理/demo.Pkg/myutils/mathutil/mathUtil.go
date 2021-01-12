package mathutil

import "demo.Pkg/calc"

//Add 两数相加
func Add(a, b int) int {
	return calc.Add(a, b)
}

//AddAll 多数相加
func AddAll(a int, numbers ...int) int {
	return calc.AddAll(a, numbers...)
}

//Multiply 两数相乘
func Multiply(a, b int) int {
	return calc.Multiply(a, b)
}

//MultiplyAll 多数相乘
func MultiplyAll(a int, numbers ...int) int {
	return calc.MultiplyAll(a, numbers...)
}

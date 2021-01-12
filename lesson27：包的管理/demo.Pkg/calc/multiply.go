package calc

//Multiply 两数相乘
func Multiply(a, b int) int {
	return a * b
}

//MultiplyAll 多数相乘
func MultiplyAll(a int, numbers ...int) int {
	sum := a
	for _, v := range numbers {
		sum *= v
	}
	return sum
}

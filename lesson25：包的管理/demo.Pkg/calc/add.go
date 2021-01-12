package calc

//Add 两数相加
func Add(a, b int) int {
	return a + b
}

//AddAll 多数相加
func AddAll(a int, numbers ...int) int {
	sum := a
	for _, v := range numbers {
		sum += v
	}
	return sum
}

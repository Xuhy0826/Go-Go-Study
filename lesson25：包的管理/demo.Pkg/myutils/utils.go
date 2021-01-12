package myutils

import "demo.Pkg/myutils/mathutil"

//Name 工具名称
const Name = "xuhyUtil"

//UtilAdd 计算两数之和
func UtilAdd(a, b int) int {
	return mathutil.Add(a, b)
}

//UtilAddAll 计算多数之和
func UtilAddAll(a int, numbers ...int) int {
	return mathutil.AddAll(a, numbers...)
}

//UtilMultiply 计算两数之积
func UtilMultiply(a, b int) int {
	return mathutil.Multiply(a, b)
}

//UtilMultiplyAll 计算多数之积
func UtilMultiplyAll(a int, numbers ...int) int {
	return mathutil.MultiplyAll(a, numbers...)
}

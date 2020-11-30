package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

//比较直接的版本
func proverbs(name string) error {
	//创建文件
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	//写文本信息
	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		f.Close()
		return err
	}
	//写文本信息
	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully.")
	if err != nil {
		f.Close()
		return err
	}
	//写文本信息
	_, err = fmt.Fprintln(f, "Don't Panic.")
	f.Close()
	return err
}

//稍显优雅的版本
func proverbsWithDefer(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	//使用defer关键字，表示在函数退出之前，执行f.Close()
	defer f.Close()

	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully.")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(f, "Don't Panic.")
	return err
}

type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return
	}
	_, sw.err = fmt.Fprintln(sw.w, s)
}

func proverbsGracefully(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	sw := safeWriter{w: f}
	sw.writeln("Errors are values.")
	sw.writeln("Don't just check errors, handle them gracefully.")
	sw.writeln("Don't Panic.")
	return sw.err
}

const rows, columns = 9, 9

//Grid 模拟一个9*9的数独网格
type Grid [rows][columns]int8

func inBound(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if row < 0 || row >= columns {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit < 9
}

//Set dede
func (g *Grid) Set(row, column int, digit int8) error {
	var errs SudokuError
	if !inBound(row, column) {
		//return errors.New("out of bound")
		//return ErrBounds
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
	}
	g[row][column] = digit
	return nil
}

//错误类型
var (
	//ErrBounds：“越界”错误
	ErrBounds = errors.New("out of bounds")
	//ErrDigit：非法数字
	ErrDigit = errors.New("invalid digit")
)

//SudokuError ：自定义错误类型
type SudokuError []error

//Error返回一个或多个用逗号分隔的错误
func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

func main() {
	fmt.Println("lesson21 err")

	//第二参数为错误类型
	files, err := ioutil.ReadDir(".")

	//如果其不为空，则是发生了异常
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
	//优雅的错误处理
	proverbsGracefully("test.txt")

	//新的错误
	// var g Grid
	// myErr := g.Set(10, 0, 5)
	// if myErr != nil {
	// 	fmt.Printf("An error occurred: %v\n", myErr) //An error occurred: out of bound
	// 	os.Exit(1)
	// }

	//类型断言
	var g Grid
	errs := g.Set(10, 0, 15)
	if errs != nil {
		if sudokuError, ok := errs.(SudokuError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(sudokuError))
			for _, e := range sudokuError {
				fmt.Printf("- %v\n", e)
			}
		}
		//os.Exit(1)
	}

	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e) //OMG, i'm sorry
		}
	}()

	panic("OMG, i'm sorry")
}

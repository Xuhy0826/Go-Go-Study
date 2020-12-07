# 包的管理
Go语言的每个代码文件都属于一个包。一个包定义一组编译过的代码，包的名称类似命名空间，可以用来间接访问包内声明的函数、类型、常量和接口。这样就可以把不同的包中定义的同名标识符区分开。

## GOROOT和GOPATH
Go两个比较重要的环境变量：
1. `GOROOT`，安装Go的路径
2. `GOPATH`，自定义的开发者的workspace   

在控制台使用命令`go env`可以查看到当前Go定义的环境变量。
Go编译器查找包的顺序：`GOROOT` → `GOPATH`；如果无法找到会引发编译异常。比如引入了`fmt`包，编译器会查找到`C:\Go\src\fmt`。

## 使用go mod
使用`go mod`命令可以帮助我们生成一个项目模块而摆脱必须在`%GOPATH%`下来创建自己项目的尴尬。创建一个示例项目为例。   
首先`cd`到准备存放项目的目录，这里路径名称为`demo.Pkg`。接着执行`go mod init %项目名%`，按照惯例来说路径名和项目名称应该一致比较好。
```
> go mod init demo.Pkg

go: creating new go.mod: module demo.Pkg
```
可以看到提示，初始化成功了。命令帮我创建了`go.mod`文件，包含这个文件的目录便是模块的根目录了。   
打开`go.mod`文件可以看到目前只包括了模块的名称（就是执行`go mod init %项目名%`时输入的名称）和Go的版本。

## 引入自定义包
我们知道引用其他包的关键字是`import`，比如之前每个文件基本都有的`import "fmt"`。这引入的系统的包，这些系统包是存放在`%GOROOT%\src`中的，Go编译器会自动加载。   
（1）定义自己的包   
比如要创建一个用于简单计算的工具包，在根目录下创建目录命名为`calc`，在里面创建文件`add.go`与`multiply.go`。这两个文件的package名可以自定义，但是按照惯例最好与路径名称一致。**注意**所有处于同一个文件夹里的go文件必须使用同一个package名。
> add.go
```
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
```
> multiply.go
```
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
```
（2）引用定义好的包
定义好自己的包之后，main函数就可以进行引用了。注意import的路径规则是`import %模块名%/[路径名]/包名`。如下所示：
```
package main

import (
	"fmt"

	"demo.Pkg/calc"
)

func main() {
	sum := calc.Add(1, 2)
	fmt.Println(sum)    //3
	res := calc.MultiplyAll(2, 3, 4, 5)
	fmt.Println(res)    //120
}
```
了解这些基本规则，可以试试稍微复杂点的引用关系。可以看看本节的示例。
* 本例中的文件路径结构与其包的引用关系图：
![文件与package](https://github.com/Xuhy0826/Golang-Study/blob/master/resource/packageAtch.jpg)
* 本例中的代码引用关系：
![import package](https://github.com/Xuhy0826/Golang-Study/blob/master/resource/importMyPackage.jpg)   

## 支持远程包导入
很常见的情况是包在GitHub上，如果现在要导入一个远程的包，比如要引用postgres的驱动`import "github.com/bmizerany/pq"`, 编译在导入它时，会先在`GOPATH`下搜索这个包，如果没有，会在使用`go get`命令来获取远程的包，并且会把获取到的源代码存储在GOPATH目录下对应URL的目录里。
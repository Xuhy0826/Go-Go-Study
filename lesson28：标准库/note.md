# 标准库

Go的标准库就是一组核心包，每个包按照功能类型聚集相应的类型与方法。并且这些包会和语言一起发布。所以我们开发时优先使用标准库，其次再考虑第三方库来扩展功能。这里先记录几个常用且有用的包。要查阅所有的标准库的包，可以访问http://golang.org/pkg/。  
在安装Go时，标准库的源代码都会安装在$GOROOT/src/pkg路径下。因为类似`go doc`，`go code`，`go build`这些工具都需要读取标准库的源代码，所以这些源代码的存在是有必要的，否则会发生编译错误。

## log包
日志对于开发调试甚至运行都是不可或缺的，标准库中提供了log包，可以对日志做一些基本的配置。并且可以根据需要自定义日志配置。  
log包最基本的功能首先是输出日志，并且可以自己配置需要的前缀，日期时间，日志具体由哪个文件记录及其源代码所在行等。

### 基本使用  
```
package main

import "log"

func init() {
	log.SetPrefix("Trace: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Println("hello")

	//Fatalln 在调用Println 之后会接着调用os.Exit(1)
	log.Fatalln("fatal")

	//Panicln 在调用Println 之后会接着调用panic()
	log.Panicln("panic")
}
```
上面程序运行输出的日志格式类似
```
Trace: 2021/01/01 20:00:05.352948 d:/WorkSpace/GitHub/Go-GO-Study/lesson28：标准库/main.go:11: hello
```
在`init`函数中进行的配置中，第二项配置传入了一些常量，这些常量就定义在log.go中
```
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23

	Ltime                         // the time in the local time zone: 01:23:23

	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.

	Llongfile                     // full file name and line number: /a/b/c/d.go:23

	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile

	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone

	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message

	LstdFlags = Ldate | Ltime     // initial values for the standard logger
)
```
通过这段源码可以顺便学习在Go中如何声明标志常量。`iota`关键字在常量声明区里的作用有2.
1. 让编译器为每个常量复制相同的表达式，直到声明区结束，或者遇到新的赋值语句
2. `iota`初始值为0，每次执行一次自增1  
所以上面的代码的背后是
```
Ldate = 1 << iota           // 1 << 0 即1
Ltime = 1 << iota           // 1 << 1 即2
Lmicroseconds = 1 << iota   // 1 << 2 即4
Lshortfile = 1 << iota      // 1 << 3 即8
LUTC = 1 << iota            // 1 << 4 即16
...
```
其中操作符`<<`是按位左移的意思，`1 << 3`的意思就是将1按位左移3位，那也就变成8了。

### 定制的日志记录器
要创建定制的日志记录器，要创建一个`Logger`类型值。看示例先
```
package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开错误日志文件失败", err)
	}

	/*
	* log.New函数创建并初始化一个Logger类型的值
	* 第一个参数指定日志要写到的目的地
	* 第二个参数指定前缀
	* 第三个参数定义日志记录包含的属性
	 */
	Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

var (
	//Trace 记录所有日志
	Trace *log.Logger
	//Info 记录重要日志
	Info *log.Logger
	//Warning 记录需要注意日志
	Warning *log.Logger
	//Error 记录错误日志
	Error *log.Logger
)

func main() {
	Trace.Println("i have something standard to say")
	Info.Println("Special Info")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}
```
【注】`ioutil.Discard`虽然是一个`io.Writer`，但是所有的Write调用都不会有具体实质性的操作，但会返回成功。所以当某个等级的日志不重要时，可以使用Discard变量来起到禁用的目的。

## json包
json包是用来编解码json格式数据的。在前面的学习中也有过简单接触。
使用http包获取Google搜索API返回的JSON。

## io包
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

func logTest() {
	log.Println("hello")

	//Fatalln 在调用Println 之后会接着调用os.Exit(1)
	//log.Fatalln("fatal")

	//Panicln 在调用Println 之后会接着调用panic()
	//log.Panicln("panic")

	Trace.Println("i have something standard to say")
	Info.Println("Special Info")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}

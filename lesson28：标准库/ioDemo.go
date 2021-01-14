package main

import (
	"bytes"
	"fmt"
	"os"
)

func ioWriterTest() {
	//（1）创建一个 Buffer 值，并将一个字符串写入 Buffer
	// 因为 bytes.Buffer的类型指针 实现了 io.Writer 接口
	var b bytes.Buffer
	b.Write([]byte("Hello "))

	//（2）使用 Fprintf 来将一个字符串拼接到 Buffer 里
	// Fprintf 方法的第一个参数接收 io.Writer 接口
	fmt.Fprintf(&b, "World!")

	//（3）将 Buffer 的内容输出到标准输出
	// os.Stdout 为 *File类型，*File也实现了 io.Writer 接口
	b.WriteTo(os.Stdout)
}

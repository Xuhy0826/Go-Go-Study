# 标准库

Go的标准库就是一组核心包，每个包按照功能类型聚集相应的类型与方法。并且这些包会和语言一起发布。所以我们开发时优先使用标准库，其次再考虑第三方库来扩展功能。这里先记录几个常用且有用的包。要查阅所有的标准库的包，可以访问http://golang.org/pkg/。  
在安装Go时，标准库的源代码都会安装在`$GOROOT/src/pkg`路径下。因为类似`go doc`，`go code`，`go build`这些工具都需要读取标准库的源代码，所以这些源代码的存在是有必要的，否则会发生编译错误。
下面简单学习三个常用的标准库：log，json和io。

## log包
日志对于开发调试甚至运行都是不可或缺的，标准库中提供了log包，可以对日志做一些基本的配置。并且可以根据需要自定义日志配置。  
log包最基本的功能首先是输出日志，并且可以自己配置需要的前缀，日期时间，日志具体由哪个文件记录及其源代码所在行等。

### 基本使用  
```go
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
```go
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
```go
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
```go
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
使用http包获取Google搜索API返回的JSON的示例如下。
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	//根据api返回的相应的json的结构定义结构体
	gResponse struct {
		Items []gResult `json:"items"`
		Kind  string    `json:"kind"`
	}

	gResult struct {
		Kind             string `json:"kind"`
		Title            string `json:"title"`
		HTMLTitle        string `json:"htmlTitle"`
		Link             string `json:"link"`
		DisplayLink      string `json:"displayLink"`
		Snippet          string `json:"snippet"`
		HTMLSnippet      string `json:"htmlSnippet"`
		FormattedURL     string `json:"formattedUrl"`
		HTMLFormattedURL string `json:"htmlFormattedUrl"`
		Mime             string `json:"mime"`
		FileFormat       string `json:"fileFormat"`
	}
)

func jsonTest() {
	uri := "https://www.googleapis.com/customsearch/v1/siterestrict?key=AIzaSyCIivhVfq-5L9yT8RQ9J8olrRV67lE_Ta8&cx=017576662512468239146:omuauf_lfve&q=golang"

	//向api发起搜索，得到响应
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()

	var gr gResponse

	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	//本例的核心：将响应的json文档解码到声明的结构体中
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Printf("%+v\n", gr)
}
```
其中`NewDecoder`根据传入的`io.Reader`接口类型值返回一个指向`Decoder`类型的指针值。`Decode`方法接受`interface{}`类型的值为入参。使用反射，`Decode`方法会拿到传入值的类型信息。然后，在读取 JSON 响应的过程中，`Decode`方法会将对应的响应解码为这个类型的值。  
如果JSON数据是字符串的形式存在，则需要用到json包中的`Unmarshal`函数进行反序列化。
```go
type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Address struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"address"`
}

func jsonDeserilizeTest() Contact {
	var JSON = `{
		"name": "Gopher",
		"title": "programmer",
		"address": {
			"home": "415.333.3333",
			"cell": "415.555.5555"
		}
	}
	`
	var c Contact
	//进行json反序列化
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("Error", err)
		return
	}
	fmt.Printf("%+v", c)
}
```
> 输出：{Name:Gopher Title:programmer Contact:{Home:415.333.3333 Cell:415.555.5555}}  

有时不方便为json数据声明一个固定的类型，可以将json文档解码到一个`map`变量中。那么把上一个示例中代码进行如下修改即可。
```go
var c map[string]interface{}
//进行json反序列化
err := json.Unmarshal([]byte(JSON), &c)
if err != nil {
	log.Println("Error", err)
	return
}

//访问
fmt.Println("Name:", c["name"])
fmt.Println("Title:", c["title"])
fmt.Println("Address")
fmt.Println("H:", c["address"].(map[string]interface{})["home"])
fmt.Println("C:", c["address"].(map[string]interface{})["cell"])
```
之前介绍的是将json文档转成对象，即反序列化。那么再看下使用json包进行对象的序列化。在“lesson14：结构体”中也学习过json序列化功能，当时使用的是json包中的`Marshal`方法将数据编码成json格式，`bytes, err := json.Marshal(spirit)`，序列化后得到的是`bytes`类型。这里再看下另一个方法` MarshalIndent `。这个函数可以将 map 类型的值或者结构类型的值转换为易读格式的 JSON 文档，也就是说他和`Marshal`的区别就是为我们增加了换行和缩进，让json看起来更美观。 
```go
contact := Contact{
	Name:  "Gopher",
	Title: "programmer",
}
contact.Address.Home = "415.333.3333"
contact.Address.Cell = "415.555.5555"

data, err := json.MarshalIndent(contact, "", " ")
if err != nil {
	log.Println("ERROR:", err)
	return
}

fmt.Println(string(data))
```
输出如下，都已实现了换行和缩进。
```json
{  
  "name": "Gopher",  
  "title": "programmer",  
  "address": {  
  	"home": "415.333.3333",  
  	"cell": "415.555.5555"  
   }  
}
```

## io包

io包可以以**流**的形式高效的处理数据，而不用考虑具体的数据是什么。io包中包含`io.Writer`和`io.Reader`这两个接口。所有实现了这两个接口的类型的值，都可以使用 io 包提供的所有功能，也可以用于其他包里接受这两个接口的函数以及方法。   
`io.Writer`的接口声明。只有一个`Write`方法，接收`[]byte`，返回写入的字节数和`error`。
```go
//Write 从 p 里向底层的数据流写入 len(p)字节的数据。这个方法返回从 p 里写出的字节数（0 <= n <= len(p)），以及任何可能导致写入提前结束的错误。 Write 在返回 n < len(p)的时候，必须返回某个非 nil 值的 error。 Write 绝不能改写切片里的数据，哪怕是临时修改也不行
type Writer interface {
	Write(p []byte) (n int, err error)
}
```
源码的注释中写明，`Write`方法的实现需要试图写入被传入的`byte`切片里的所有数据。**如果无法全部写入，那么该方法就一定会返回一个错误。**  
接下来再看一下`Reader`接口的定义。只有一个`Read`方法，接收`[]byte`，返回读入的字节数和`error`。
```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```
`Reader`接口的说明：
1. 接口的实现需要试图读取数据来填满被传入的`byte`切片。允许出现读取的字节数小于 byte 切片的长度，并且如果在读取时已经读到数据但是数据不足以填满 byte 切片时，不应该等待新数据，而是要直接返回已读数据。
2. 当读到最后一个字节时，可以有两种选择。一种是`Read`返回最终读到的字节数，并且返回`EOF`作为错误值，另一种是返回最终读到的字节数，并返回`nil`作为错误值。在后一种情况下，下一次读取的时候，由于没有更多的数据可供读取，需要返回 0 作为读到的字节数，以及`EOF`作为错误值。
3. 调用`Read`时，会返回读取的字节数，都应该优先处理这些读取到的字节，再去检查 EOF 错误值或者其他错误值。
4. 建议`Read`方法的实现永远不要返回 0 个读取字节的同时返回 nil 作为错误值。如果没有读到值，`Read`应该总是返回一个错误。  
接下来看一个示例。
```go
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
```
控制台输出：

> Hello World!
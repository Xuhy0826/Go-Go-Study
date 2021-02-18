# 基本类型之字符串类型

## string类型

关于字符串类型（string），和其他语言一样没什么区别，使用双引号包起来，如
```go
peace := "peace"
var peace = "peace"
var peace string = "peace"
```
使用双引号包起来的字符串称为“字符串字面量”。“字符串字面量”中可以包含转义字符，比如说 `\n` 可以表示换行。另外一种表示字符串字面量的方法是使用反引号`，这种称为“原始字符串字面量”。使用``，可以方便的定义跨行字符串，如
```go
fmt.Println(`
peace be upon you
upon you be peace`)
```
“字符串字面量”和“原始字符串字面量”都是string类型。

## 字符、代码点、符文和字节
* 都知道计算机中字符是通过编码存取的，也就是每个字符都使用一个特定的数字表示。比如A就是65，那么书中将这个65称为字符A的**代码点**。
* rune类型（符文类型）：Go中使用rune类型来表示字符的代码点，该类型本质上是int32类型的别名，也就是说rune和int32可以相互转换
* byte类型：Go中的byte类型不仅可以表示二进制数据，而且被拿来表示ASCII码（ASCII共包含128个字符）。本质上byte类型是uint8类型的别名
```go
var pi rune = 960
var alpha rune = 940
var omega rune = 969
var bang byte = 33

fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang) //960 940 969 33
//通过使用格式化变量%c，可以将代码点表示成字符
fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang) //π ά ω !
```
在Go中使用单引号来表示字符字面量，如果用户声明一个字符变量而没有为其制定类型，那么Go会将其推断成rune类型。下面三种写法是一样的功能。
```go
grade := 'A'
var grade = 'A'
var grade rune = 'A'
```

## 字符串无法修改
字符串虽然是有字符“串”起来的，但是和C#、java等语言一样，Go中的字符串类型也是不可修改的。
```go
//通过索引的方式访问字符串中的字符
message := "shalom"
c := message[5]
fmt.Printf("%c\n", c)

//字符串不可被修改
//message[5] = 'd'  //报错：cannot assign to message[5]
```

## 字符串与符文
Golang中字符串使用utf-8编码。utf-8是一种变长的编码方式，也就是说每个字符的长度可能占用不同的字节长度。比如中文字符就是需要占据两个字节长度，而英文字符或者数字则只需要占据1个字节长度。  
为了方便，Go为此提供了utf包，里面提供了两个实用的方法
* RuneCountInString函数，能够按照特定的字符返回字符串中字符的个数，而不是像len方法一样返回字节的长度
* DecodeRuneInString函数，解码字符串的首个字符并返回解码后的字符占用的字节长度
```go
question := "今天星期几？"
fmt.Println(len(question), "bytes")                    //18 bytes
fmt.Println(utf8.RuneCountInString(question), "runes") //6 runes

c, size := utf8.DecodeRuneInString(question)
fmt.Printf("First rune: %c %v bytes", c, size) //First rune: 今 3 bytes

//遍历字符串，挨个打印出来
for i, c := range question {
    fmt.Printf("%v %c\n", i, c)
}
```
上面的示例中使用到了`range`关键字来进行遍历操作，其中`i`为索引，`c`为值。有点python的味道。
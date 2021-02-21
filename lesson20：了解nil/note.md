# NIL

可以将其与C#中的`null`类比。在Go中如果一个指针没有明确的指向，那么它的值就是nil。
```go
var i int
var s string
var p *string

fmt.Printf("%v\n", i) //0
fmt.Printf("%v\n", s) // (空字符串)
fmt.Printf("%v\n", p) // <nil>
```

## nil可能会引发Panic
如果对一个nil指针进行解引用会引发panic（引发Go程序崩溃的错误）。
```go
var p *string
fmt.Printf("%v\n", p) // <nil>
fmt.Printf("%v\n", *p) //panic: runtime error: invalid memory address or nil pointer dereference
```
避免这种情况的方法可以在解引用之前先判断指针是否是nil
```go
var nowhere *int

if nowhere != nil {
    fmt.Println(nowhere)
}
```
以往的编程经验告诉我们，在方法中如果入参或者接收者是指针类型，那么最好都要进行下空判断来确保安全。
```go
func (p *person) birthday{
    if p == nil{
        return
    }
    p.age++
}
```

## 默认值是nil的情况

### 函数值
当变量被声明为函数类型，在没有被赋值的情况下，其就为nil值。
```go
var fn func(a, b int) int
fmt.Println(fn == nil) //true
```

### 切片
同理，切片在声明之后没有使用复合字面量或者make函数赋值，其值便为nil。
```go
var soup []string
fmt.Println(soup == nil) //true
```
但是一些内置函数和关键字都可以很好的解决nil切片的问题，比如`len`,`append`,`cap`和`range`。
```go
//range可以处理nil
for _, ingredient := range soup {
    fmt.Println(ingredient)
}
//len、append也可以处理nil
fmt.Println(len(soup)) //0
soup = append(soup, "onion", "carrot")
fmt.Println(soup) //[onion carrot]
```
### 映射
同理，映射在声明之后没有使用复合字面量或者make函数赋值，其值便为nil。对nil映射的读取操作不会引发panic，但是**写入操作则会引发panic**。
```go
var souplist map[string]int
fmt.Println(souplist == nil) //true

measurement, ok := souplist["onion"]

if ok {
    fmt.Println(measurement)
}

for ingredient, measurement := range souplist {
    fmt.Println(ingredient, measurement)
}

//souplist["onion"] = 1 //panic: assignment to entry in nil map
```
### 接口
接口类型的变量在未被赋值时的零值是nil，并且它的接口类型和值都是nil。
```go
var v interface{}
fmt.Printf("%T %v %v\n", v, v, v == nil) //<nil> <nil> true
```
值得注意的是，当接口类型的变量被赋值之后，接口就会在内部指向该变量的类型和值。先看下面的示例。
```go
var v interface{}

var po *int
v = po
fmt.Printf("%T %v %v\n", v, v, v == nil) //*int <nil> false
```
在将`po`赋值给`v`之后，`v`的类型就变成了`*int`，虽然值仍然是`nil`，但是Go认定接口类型的变量只有在类型和值都为nil时才等于`nil`。所以`v == nil`的结果是`false`。
```go
//格式化变量 %#v 可以同时打印出变量的类型和值
fmt.Printf("%#v", v)    //(*int)(nil)
```
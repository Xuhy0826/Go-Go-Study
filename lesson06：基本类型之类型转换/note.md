# 基本类型之类型转换

Go中与其他强类型语言（比如C#）类似，类型之间进行操作时，需要经过类型转换否则会报“类型不匹配”的错误。

## 数字类型转换
* 整数类型 → 浮点类型：

```go
age := 41
marsAge := float64(age)
```
---
* 浮点类型 → 整数类型：  

【注意】：浮点型的小数部分是被截断，而不是四舍五入
```go
earthDays := 365.2425
fmt.Println(int(earthDays)) //output: 365
```
在数值类型进行转换时，一样要注意超出范围的问题，比如一个较大float64转成int16时。

## 字符串转换
* rune/byte → string

```go
var pi rune = 960
var alpha rune = 940
fmt.Println(string(pi), string(alpha)) //output: π ά
```
---
* 数字类型 → string

情况特殊一点，为了将一串数组转换为string类型，必须将其中的每个数字都转换为相应的代码点（char）。也就是代表字符0的48~代表字符9的57。我们需要使用到strconv（代表“string conversion”）包提供的Itoa函数来完成这一工作。
```go
countdown := 10
str := "Launch in T minus " + strconv.Itoa(countdown) + " seconds".
```
另一种方法，使用fmt.Sprintf函数，该函数会返回格式化后的string而不是打印
```go
countdown := 9
str := fmt.Sprintf("Launch in T minus %v seconds", countdown)
fmt.Println(str) //Launch in T minus 9 seconds
```
> 注：使用 `strconv.Itoa()` 比 `fmt.Sprintf()` 要快一倍左右
---
* string → 数字
一种不太常用的转换，也是使用strconv包的Atoi函数
```go
count, err := strconv.Atoi("10")
if err != nil {
    //出错
}
fmt.Println(count) //10
```
上面这种写法是之后经常看到的，是Go处理异常的一种常用写法。由于Go的函数可以返回多个值，一般会将可能产生的异常一并返回。

## 布尔类型转换
如果使用`fmt`的`Print`系函数直接打印bool类型，会输出true或false的文本
```go
launch := false
launchText := fmt.Sprintf("%v", launch)
fmt.Println("Ready for launch:", launchText) //Ready for launch: false
```
某些语言中会把1和0当做true和false，Go中是不行的。
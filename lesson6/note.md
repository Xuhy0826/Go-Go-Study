# 基本类型之类型转换

Go中和C#类似，类型之间进行操作时，需要经过类型转换否则会报“类型不匹配”的错误

## 数字类型转换
* 整数类型 → 浮点类型：
```
age := 41
marsAge := float64(age)
```
* 浮点类型 → 整数类型：
```
fmt.Println(int(earthDays))
```
在数值类型进行转换时，一样要注意超出范围的问题，比如一个较大float64转成int16时。

## 字符串转换
* rune/byte → string
```
var pi rune = 960
var alpha rune = 940
fmt.Println(string(pi), string(alpha)) //π ά
```
* 数字类型 → string
情况特殊一点，为了将一串数组转换为string，必须将其中的每个数字都转换为相应的代码点（char）。也就是代表字符0的48~代表字符9的57。我们需要使用到strconv（代表“string conversion”）包提供的Itoa函数来完成这一工作。
```
countdown := 10
str := "Launch in T minus " + strconv.Itoa(countdown) + " seconds".
```
另一种方法，使用fmt.Sprintf函数，该函数会返回格式化后的string而不是打印
```
countdown := 9
str := fmt.Sprintf("Launch in T minus %v seconds", countdown)
fmt.Println(str) //Launch in T minus 9 seconds
```
* string → 数字
一种不太常用的转换，也是使用strconv包的Atoi函数
```
count, err := strconv.Atoi("10")
if err != nil {
    //出错
}
fmt.Println(count) //10
```
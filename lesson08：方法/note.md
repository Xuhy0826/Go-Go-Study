# 方法

## 声明新类型
* 通过type关键字来声明新的类型
通过type关键字，指定名称和一个底层类型便可以声明一个新的类型。
```go
type celsius float64
var temperature celsius = 20
fmt.Println(temperature)
```
如上，声明了一个摄氏度的新类型叫celsius，由于数字字面量20是一个无类型常量，所以int，float64类型或者其他数字类型都可以将其做值。有因为celsius类型和float64具有相同的行为，可以把其当做float64来使用
```go
const degrees = 20
var temperature2 celsius = degrees
temperature2 += 10
fmt.Println(temperature2) //30
```
**【注意】** 虽然声明的新类型和声明时指定的底层类型具有相同的行为与表示，但是这和前面提过的类型别名不同，通过type关键字声明的类型就是一个全新的类型。所以尝试把celsius和float64一起使用会报错“类型不匹配”
```go
var warmUp float64 = 10
temperature += warmUp   //报错
```
通过自定义新类型可以提高代码的可读性。如下面的代码，因为摄氏度和华氏度是两个不同的类型，它们是无法一起直接比较或运算的。
```go
type fahrenheit float64

var c celsius = 20
var f fahrenheit = 20

if c == f {  //报错

}
c += f  //报错
```
## 引入自定义类型
在声明了新的类型之后，我们可以在函数或者其他地方去使用。我们将之前的kelvinToCelsius函数使用新类型进行改写，将参数类型和返回值类型都改用自定义类型。
```go
//type关键字声明新类型
type celsius float64
type kelvin float64

//kelvinToCelsius converts °K to °C
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15) //需要使用类型转换
}

func main() {
	//使用引入了新类型的函数
	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	fmt.Println(k, "°K is", c, "°C") //294 °K is 20.850000000000023 °C

}
```

## 通过方法添加行为
在以往面向对象的语言中，方法一般都是属于某个类的。但是在Go中不一样，Go中没有类或者对象，但是存在方法。在Go中可以为声明的类型（type关键字声明）关联方法。先看如何声明。
```go
type celsius float64
type kelvin float64

//为kelvin类型关联一个方法celsius
func (k kelvin) celsius() celsius{
	return celsius(k - 273.15)
}
```
这里的k称为接收者（receiver），每个方法可以有多个参数，但是只能有一个接收者。   
![方法声明](https://github.com/Xuhy0826/Golang-Study/blob/master/resource/methodDeclare.png)

在声明好方法之后，类型kelvin就可以直接调用方法了，用起来就和在类中定义了方法很类似。
```go
var k kelvin = 294.0
c := k.celsius()
fmt.Println(k, "°K is", c, "°C") //294 °K is 20.850000000000023 °C
```
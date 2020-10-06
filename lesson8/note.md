# 方法

## 声明新类型
* 通过type关键字来声明新的类型
通过type关键字，指定名称和一个底层类型便可以声明一个新的类型。
```
type celsius float64
var temperature celsius = 20
fmt.Println(temperature)
```
如上，声明了一个摄氏度的新类型叫celsius，由于数字字面量20是一个无类型常量，所以int，float64类型或者其他数字类型都可以将其做值。有因为celsius类型和float64具有相同的行为，可以把其当做float64来使用
```
const degrees = 20
var temperature2 celsius = degrees
temperature2 += 10
fmt.Println(temperature2) //30
```
**【注意】** 虽然声明的新类型和声明时指定的底层类型具有相同的行为与表示，但是这和前面提过的类型别名不同，通过type关键字声明的类型就是一个全新的类型。所以尝试把celsius和float64一起使用会报错“类型不匹配”
```
var warmUp float64 = 10
temperature += warmUp   //报错
```
通过自定义新类型可以提高代码的可读性。如下面的代码，因为摄氏度和华氏度是两个不同的类型，它们是无法一起直接比较或运算的。
```
type fahrenheit float64

var c celsius = 20
var f fahrenheit = 20

if c == f {  //报错

}
c += f  //报错
```
## 引入自定义类型
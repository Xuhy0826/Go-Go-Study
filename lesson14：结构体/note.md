# 结构体：struct

为什么要有结构体这种数据类型存在就不需多解释了，面向对象这么多年了。

## 结构的声明
Go中声明struct类型的语法，通过下面示例声明的结构有点类似在C#中创建的匿名类（但是Go中是没有类的），访问字段的方式也是相同dot notation
```
var curiosity struct {
    lat  float64
    long float64
}

curiosity.lat = -4.9773
curiosity.long = 137.4283

fmt.Println(curiosity.lat, curiosity.long) //-4.9773 137.4283
fmt.Println(curiosity)                     //{-4.9773 137.4283}
```
将struct类型的变量赋值给新的变量也是会复制一份完全相同的值传递过去。
```
curiosityMarkII := curiosity
curiosity.lat = 0
fmt.Println(curiosity)       //{0 137.4283}
fmt.Println(curiosityMarkII) //{-4.9773 137.4283}
```

## 使用type声明可复用的结构
上一节使用struct的方式有点类似C#中的匿名类，更常用的使用方法应该还是声明好一个struct后可以到处复用。看实例
```
type location struct {
	lat  float64
	long float64
}

func main() {
    var spirit location
	spirit.lat = -14.5637
	spirit.long = 175.3774

	var opprtunity location
	opprtunity.lat = -1.9473
	opprtunity.long = 352.8434

	fmt.Println(spirit)     //{-14.5637 175.3774}
	fmt.Println(opprtunity) //{-1.9473 352.8434}
}
```
除了挨个给字段赋值，当然也可以在声明的同时使用复合字面量来初始化struct的数据，建议使用显式对字段名进行赋值。
```
spirit := location{lat : -14.5637, long : 175.3774}
```
使用fmt.Printf()格式化输出结构体数据时，可以使用 **%+v** 来显示结构中的字段名
```
fmt.Printf("%v\n", curiosity)  //{0 137.4283}
fmt.Printf("%+v\n", curiosity) //{lat:0 long:137.4283}
```

## Json序列化
在Go语言中，通过json包中的Marshal()方法将数据编码成json格式。注意方法返回的是bytes类型Json数据，我们可以将其转为string类型再使用。
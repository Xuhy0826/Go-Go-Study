# 结构体：struct

为什么要有结构体这种数据类型存在就不需多解释了，面向对象这么多年了。

## 结构的声明
Go中声明struct类型的语法，通过下面示例声明的结构有点类似在C#中创建的匿名类（但是Go中是没有类的），访问字段的方式也是相同dot notation
```go
var curiosity struct {
    lat  float64
    long float64
}

curiosity.lat = -4.9773
curiosity.long = 137.4283

fmt.Println(curiosity.lat, curiosity.long) //-4.9773 137.4283
fmt.Println(curiosity)                     //{-4.9773 137.4283}
```
将struct类型的变量赋值给新的变量也是会**复制一份完全相同的值**传递过去。
```go
curiosityMarkII := curiosity //传递副本
curiosity.lat = 0
fmt.Println(curiosity)       //{0 137.4283}
fmt.Println(curiosityMarkII) //{-4.9773 137.4283}	//原值不改变
```

## 使用type声明可复用的结构
上一节使用struct的方式有点类似C#中的<font color=#0000FF size=14>**匿名类**</font>，更常用的使用方法应该还是声明好一个struct后可以到处复用。看实例
```go
type location struct {
	lat  float64
	long float64
}

func main() {
    var spirit location		//声明一个location类型
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
```go
spirit := location{lat : -14.5637, long : 175.3774}
```
使用fmt.Printf()格式化输出结构体数据时，可以使用 **%+v** 来显示结构中的字段名。
```go
fmt.Printf("%v\n", curiosity)  //{0 137.4283}
fmt.Printf("%+v\n", curiosity) //{lat:0 long:137.4283}
```

## Json序列化
在Go语言中，通过json包中的`Marshal()`方法将数据编码成json格式。注意方法返回的是bytes类型Json数据，我们可以将其转为string类型再使用。
```go
import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	lat  float64
	long float64
}

func main(){
	spirit := location{lat : -14.5637, long : 175.3774}
	bytes, err := json.Marshal(spirit)	//转json
	exitOnError(err)
	fmt.Println(string(bytes))	//转string后输出
}

//exitOnError prints any errors and exits.
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```
上面的示例代码很明了，就是声明struct后转json打印出来。但是运行后发现，输出的仅仅是一个“{}”,也就是说输出了一个没有任何字段的结构体数据。这是因为struct内以大写字母开头的字段才会被json化，而示例代码中的字段都是小写字母。在Go中，大写字母开头的字段就相当于C#中的public属性，这就好理解了。
```go
type locationV2 struct {
	Lat  float64
	Long float64
}

spiritV2 := locationV2{Lat: 12.433, Long: 144.843}
bytes, err := json.Marshal(spiritV2)
exitOnError(err)
fmt.Println(string(bytes)) //{"Lat":12.433,"Long":144.843}
```
其他部分不变，将字段名称的手写字母变成大写后，转json格式后可以看到结果了。   
另外我们那也可以自定义字段转成Json后的字段名，在C#中可以通过在属性上加attribute的方式，在Go中的语法如下
```go
type locationV3 struct {
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

spiritV3 := locationV3{Lat: 12.433, Long: 144.843}
bytesV3, errV3 := json.Marshal(spiritV3)
exitOnError(errV3)
fmt.Println(string(bytesV3)) //{"latitude":12.433,"longitude":144.843}
```
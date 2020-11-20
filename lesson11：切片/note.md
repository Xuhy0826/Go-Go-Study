# 切片：指向数组的窗口

## 切分数组
Go语言中对数组的切片写法类似python，比如上节中的数组planets，planets[0:4]即可获取索引0到索引4的4个元素 **（不包括索引4的元素）**。
```
planets := [...]string{
    "Mercury",
    "Venus",
    "Earth",
    "Mars",
    "Jupiter",
    "Saturn",
    "Uranus",
    "Neptune",
}

terrestrial := planets[0:4]
gasGiants := planets[4:6]
iceGiants := planets[6:8]
fmt.Println(terrestrial, gasGiants, iceGiants) //[Mercury Venus Earth Mars] [Jupiter Saturn] [Uranus Neptune]
```
* 切片仍然可以像正常数组一样根据索引来获取指定元素：
```
fmt.Println(gasGiants[1]) //Saturn
```
* 切片仍然可以像正常数组一样继续创建切片
```
giants := planets[4:8]
gas := giants[0:2]
ice := giants[2:4]
fmt.Println(gas, ice) //[Jupiter Saturn] [Uranus Neptune]
```
* **[注意]** 切片是数组的“视图”，对切片中的元素进行重新赋值的操作，便会导致原数组中元素的更改，也会影响原数组的其他切片。
```
iceGiantsMarkII := iceGiants
fmt.Println(iceGiantsMarkII) //[Uranus Neptune]
iceGiants[1] = "Poseidon"
fmt.Println(iceGiantsMarkII) //[Uranus Poseidon] 发生了变化
fmt.Println(ice)             //[Uranus Poseidon]
```
切片也有简写模式，也就是利用切片的默认值，array[:3]表示从开头切到index为3的地方，array[4:]表示index为4的元素一直切到最后，array[:]表示数组的所有元素了。
```
planets := [...]string{
    "Mercury",
    "Venus",
    "Earth",
    "Mars",
    "Jupiter",
    "Saturn",
    "Uranus",
    "Neptune",
}
//切片可以简写，利用默认值
var slice1 = planets[:3]
var slice2 = planets[4:]
var slice3 = planets[:]

fmt.Println(slice1) //[Mercury Venus Earth]
fmt.Println(slice2) //[Jupiter Saturn Uranus Poseidon]
fmt.Println(slice3) //[Mercury Venus Earth Mars Jupiter Saturn Uranus Poseidon]
```
另外值得一提的是字符串也可以这么玩
```
neptune := "Neptune"
tune := neptune[3:]
fmt.Println(tune) //tune
```
切分字符串时，索引是按照字节号码而不是符文号码
```
question := "你在学习Go吗？"
fmt.Println(question[:6]) //你在
```
Go语言中，函数更加倾向于使用切片作为输入。除了切分数组，另外一个创建切片的简便方法是使用 **切片的复合字面量** 。声明方法看下面的例子，区分一下，使用字面量声明数组的写法是[...]string，声明切片是[]string。
```
//使用“切片复合”
dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
fmt.Printf("%T", dwarfs) //[]string
```
将切片作为参数传入函数，Go的函数是按值传递，传入的是一个完整副本，但是这两个切片都是指向同一个底层数组的，所以在函数中所做的改动会影响到原数组和其他切片。
```
import (
	"fmt"
	"strings"
)

//遍历切片，消除空格
func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func main() {
    countries := []string{" China ", "  Japan", " USA"}
	hyperspace(countries)
	fmt.Println(strings.Join(countries, "")) //[]stringChinaJapanUSA
}
```
## 带有方法的切片
可以使用切片或者数组作为底层类型声明类型，并为其绑定方法。比如标准库sort包声明了一种StringSlice的类型：
```
type StringSlice []string
```
并且为其关联了方法：可按照字母进行排序
```
func (p StringSlice) Sort()
```
【示例】
```
import (
	"fmt"
	"sort"
	"strings"
)

func main() {
    dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

    //使用sort方法对切片进行排序：
    //sort包中含有的StringSlice类型，先将dwarfs转换成StringSlice类型，之后再调用StringSlice类型的Sort方法
	sort.StringSlice(dwarfs).Sort()
	fmt.Println(dwarfs) //[Ceres Eris Haumea Makemake Pluto]

	//简化上面的操作，类型转换和排序一次执行
	sort.Strings(dwarfs)
}
```
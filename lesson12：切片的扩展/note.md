# 切片的扩展

## append函数

* append是内置函数，可以将元素添加到切片中
```go
dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
//使用append函数为切片增加元素
dwarfs = append(dwarfs, "Orcus")
fmt.Println(dwarfs) //	[Ceres Pluto Haumea Makemake Eris Orcus]
```
* 因为切片是指向数组的，调用append函数，会修改原数组元素。看例子会比较清楚
```go
dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris", "Orcus"}
dwarfsSlice := dwarfs[2:4]
dwarfsSlice = append(dwarfsSlice, "EEEE")	//这时使用append会将原数组的元素“Eris”替换掉
fmt.Println(dwarfsSlice)  					//output：[Haumea Makemake EEEE]
fmt.Println(dwarfs)       					//output：[Ceres Pluto Haumea Makemake EEEE Orcus]
```
* append函数是一个可变参数的函数，可以一次添加多个元素
```go
dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")
fmt.Println(dwarfs) //output：[Ceres Pluto Haumea Makemake Eris Salacia Quaoar Sedna]
```
* append函数可直接拼接切片
```go
// 创建两个切片
s1 := []int{1, 2}
s2 := []int{3, 4}

// 将两个切片拼接在一起，并显示结果
fmt.Printf("%v\n", append(s1, s2...))   //output：[1 2 3 4]
```

## 长度和容量
* 切片的长度（length）：切片中元素的个数
* 切片的容量（capacity）：不能简单的理解为“切片对应的底层数组的长度”，比如底层数组长度为10，如果切片从索引为4的地方开始切，那么这个切片的容量就是6。因为切片只能在原数组上往后追加。   
对底层数组容量是 k 的切片 slice[i:j]来说：
1. 长度: j - i
2. 容量: k - i   

获取切片的长度或者容量，Go都已有内置的函数 **`len()`** 和 **`cap()`** 。
```go
package main

import (
	"fmt"
)

//显示切片的长度和容量信息
func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	dump("planets", planets)           //output：planets: length 8, capacity 8 [Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]
	dump("planets[1:4]", planets[1:4]) //output：planets[1:4]: length 3, capacity 7 [Venus Earth Mars]
	dump("planets[2:5]", planets[2:5]) //output：planets[2:5]: length 3, capacity 6 [Earth Mars Jupiter]
	dump("planets[5:7]", planets[5:7]) //output：planets[5:7]: length 2, capacity 3 [Saturn Uranus]
}
```
## 再探append函数
先看下面这个示例
```go
dwarfsRaw := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
dwarfs1 := dwarfsRaw[:]                                  //length=5， capacity=5
dwarfs2 := append(dwarfs1, "Orcus")                      //length=6， capacity=10
dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna") //length=9， capacity=10

dump("dwarfs1", dwarfs1) //dwarfs1: length 5, capacity 5 [Ceres Pluto Haumea Makemake Eris]

dump("dwarfs2", dwarfs2) //dwarfs2: length 6, capacity 10 [Ceres Pluto Haumea Makemake Eris Orcus]

dump("dwarfs3", dwarfs3) //dwarfs3: length 9, capacity 10 [Ceres Pluto Haumea Makemake Eris Orcus Salacia Quaoar Sedna]

fmt.Println(dwarfsRaw) //[Ceres Pluto Haumea Makemake Eris]
```
可以看出，当为切片调用append函数追加元素时，若切片的容量够，则直接追加。若切片的容量不够，Go会将当前切片的底层数组复制到一个新的数组中，新数组的长度是原数组的两倍大。
![示意图](https://github.com/Xuhy0826/Golang-Study/blob/master/resource/AppendFunc.png)
上面的例子中还有一个值得注意的地方，如果我们修改`dwarfs3`中的元素，`dwarfs1`不会受影响，而`dwarfs2`会相应的改变，而`dwarfs2`和`dwarfs3`的内存是共享的，也就是说`dwarfs2`和`dwarfs3`指向的是同一个新创建的数组。
```go
//修改切片dwarfs3的内容
dwarfs3[1] = "A"
//观察 dwarfs1 ，没有变化
dump("dwarfs1", dwarfs1) //dwarfs1: length 5, capacity 5 [Ceres Pluto Haumea Makemake Eris]
//观察 dwarfs2 ，发生变化
dump("dwarfs2", dwarfs2) //dwarfs2: length 6, capacity 10 [Ceres A Haumea Makemake Eris Orcus]
```

## 使用3个索引来切分数组
* Go支持使用3个索引来切片，第三个索引是用来“限制切片的容量”，称其“limited capacity”。   
比如
```go
source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
slice := source[2:3:4]
```
这就表示切片“切取”了source数组索引2到索引3之间的元素，其实也就是索引为2的那个元素，这时切片的长度是1。如果没有第三个参数，这时切片slice的容量会是3。但是此时使用到了第三个参数，也是表示数组的索引值。这个索引值表示这个切片的容量最大到原数组的那个索引。这样一来，声明的`slice`切片容量为2。   
* 对于 slice[i:j:k] 或 [2:3:4]
1. 长度: j – i 或 3 - 2 = 1
2. 容量: k – i 或 4 - 2 = 2   

再看一个稍微复杂的例子体会一下。
```go
planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
terrestrial := planets[0:4:4]
terrestrial1 := planets[0:4]
dump("terrestrial", terrestrial)   	//length 4, capacity 4 [Mercury Venus Earth Mars]
dump("terrestrial1", terrestrial1) 	//length 4, capacity 8 [Mercury Venus Earth Mars]

worlds := append(terrestrial, "Ceres")

dump("planets", planets)         //planets: length 8, capacity 8 [Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]
dump("terrestrial", terrestrial) //terrestrial: length 4, capacity 4 [Mercury Venus Earth Mars]
dump("worlds", worlds)           //worlds: length 5, capacity 8 [Mercury Venus Earth Mars Ceres]
```
上面的例子可以看出，在为指定了容量的切片执行`append`操作时，如果容量不够用了，`append`返回的新切片会指向重新创建的一个新的底层数组，与原有的底层数组分离。这样能够更加安全地进行后续修改。

## 使用make函数对切片进行预分配
当切片的容量不足以执行append时，会创建一个新的数组并进行复制。但是如果使用make函数来声明切片则可以自定义切片的长度和容量。
* make(类型，长度，容量)
* make(类型，长度和容量)
```go
dwarfsWithMake := make([]string, 0, 10)
dwarfsWithMake = append(dwarfsWithMake, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
dump("dwarfsWithMake", dwarfsWithMake) //dwarfsWithMake: length 5, capacity 10 [Ceres Pluto Haumea Makemake Eris]
```
使用make函数来预分配切片的容量可以避免一些情况下的数组复制操作来提升性能

## 声明可变参数的函数
* 声明可变参数的语法是在参数类型前面加上“...”，此时参数类型实际上转变成了一个相应类型的切片类型
* 调用可变参数函数时传递的是多个参数，如果想传入切片，则需要在切片后加上“...”，这样是表示将切片展开
```go
//声明一个新的切片，切片的内容是传入的原切片元素加上前缀
func terraform(prefix string, worlds ...string) []string {

	newWorlds := make([]string, len(worlds))
	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

func main() {
	//调用“可变参数函数”【方式1】
	twoWorld := terraform("New", "Venus", "Mars")
	fmt.Println(twoWorld) //[New Venus New Mars]

	//调用“可变参数函数”【方式2】
	oldWorlds := []string{"Venus", "Mars", "Jupiter"}
	//传入切片，并在切片后加上“...”
	newWorld := terraform("New", oldWorlds...)
	fmt.Println(newWorld) //[New Venus New Mars New Jupiter]
}
```

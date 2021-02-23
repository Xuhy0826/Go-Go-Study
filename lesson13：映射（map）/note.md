# 映射：map
一种用来存储**无序键值对**的类型，并且键和值都可以是任意的类型。Go语言中的map类型其实和C#，python中的Dictionary很类似。

## map的声明

举个例子，下面声明一个键（key）为string类型，值（value）为int类型的map。
> **map[string]int**

对map的读写操作和其他语言中的字典操作很类似，简单过一下
```go
//声明map，使用复合字面量的声明方式
temperature := map[string]int{
    "Mars":  -65,
    "Earth": 15,
}
//读
temp := temperature["Earth"]
fmt.Printf("On avarage the Earth is %v C \n", temp) //On avarage the Earth is 15 C

//写
temperature["Mars"] = 16
temperature["Vemus"] = 464

fmt.Println(temperature) //output：map[Earth:15 Mars:16 Vemus:464]

//如果访问不存在的键，则返回值类型的零值
fmt.Println(temperature["Moon"]) //0
```
读取map中的数据，Go语言中可以使用下面写法来简化代码
```go
if moon, ok := temperature["Moon"]; ok {
    fmt.Printf("On avarage the Earth is %v C \n", moon)
} else {
    fmt.Println("Where is the moon?")
}
```
应该很容易理解，ok（可以使用其他参数名）是个bool类型，若map中存在此键，则ok为true，反之为false。这种语法省去了我们在其他语言中还要去手动调用类似`ContainKey()`的方法来判断字典中是否包含此元素的动作。
> 注：在使用`Map`的时候，使用整型的 key 会比字符串的要快，因为整型比较比字符串比较要快。
## map是不会被复制的
之前的基本类型与数组，在被赋值给新的变量或者传递给函数或方法时都会创建新的副本。而map则不同，不管是赋值还是当做函数的参数，都是会共享底层数据的（简单理解成都是传引用）。看例子
```go
planets := map[string]string{
    "Earth": "Sector ZZ9",
    "Mars":  "Sector ZZ9",
}
planetsMarkII := planets
planets["Earth"] = "Whoops"
fmt.Println(planets)       //output：map[Earth:Whoops Mars:Sector ZZ9]
fmt.Println(planetsMarkII) //output：map[Earth:Whoops Mars:Sector ZZ9]
```
`planets`改变，相应的`planetsMarkII`也是改变的，因为它们指向同一块内存。
* 使用delete方法可以删除map中的元素，例子接上
```go
delete(planets, "Earth")
delete(planets, "Moon")    //移除不存在的元素，不会引发panic
fmt.Println(planets)       //output：map[Mars:Sector ZZ9]
fmt.Println(planetsMarkII) //output：map[Mars:Sector ZZ9]
```
但是在Go中没有提供清空map的方法，因为你花时间清空还不如重新创建一个新的map，因为Go的GC效率是很高的。

## 使用make函数对map进行预分配
声明map有两种方法，一种是之前用过的复合字面量，另一种就是使用make函数。make函数可以接收一个或两个参数，第二个参数用于指定键的数量来预分配空间。看实例。
```go
temperature := make(map[float64]int, 8)
```
为map指定初始大小能够在map变大时减少一些后续操作而提升效率。

## 遍历map
用一个实例来演示下，使用映射和切片来实现对数据分组，将一组数按照10的跨度分成几组。对于映射而言，使用`range`时返回的不是索引和值，而是键值对。
```go
temperatures := []float64{
    -28.0, 32.0, -31.0, -29.0, -23.0, -28.0, -33.0,
}
//键：float64，值；[]float64
groups := make(map[float64][]float64)

for _, t := range temperatures {
    g := math.Trunc(t/10) * 10 //将温度按10度的跨度进行分组
    groups[g] = append(groups[g], t)
}

for g, temperatures := range groups {
    fmt.Printf("%v : %v\n", g, temperatures)
}

/*
上例输出结果：
-20 : [-28 -29 -23 -28]
30 : [32]
-30 : [-31 -33]
*/
```

## 最后注意一点
在Go的`map`中的数据是无序的，如果要排序map中的值，一般会先把`map`转成切片，在调用内置方法对切片进行排序来完成。   
【示例】：将切片去重并排序  
在Go中没有类似`set`这种数据结构，但是可以使用map来实现set的功能
```go
numbers := []float64{
    51.02, 10.2, -5.2, -10.4, 14.2, 10.2, 5.12, 51.02, -30.0, 4.3,
}

set := make(map[float64]bool)

for _, t := range numbers {
    set[t] = true
}

unique := make([]float64, 0, len(set))

for t := range set {
    unique = append(unique, t)
}
//排序
sort.Float64s(unique)
fmt.Println(unique) //output: [-30 -10.4 -5.2 4.3 5.12 10.2 14.2 51.02]
```
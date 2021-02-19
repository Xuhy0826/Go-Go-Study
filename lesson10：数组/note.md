# 数组

数组的数据结构和其他语言中的数组没有什么太大区别。数组的声明写法如下，如果声明一个长度为8的字符串数组。
```go
var planets [8]string
```
不同之处是Go中声明数组时，将中括号和数组长度的声明放在了类型之前。   
另外，关于数组元素的读写操作与其他常见语言无异。
```go
//Array of planets
var planets [8]string
planets[0] = "Mercury" //Assigns a planet at index 0
planets[1] = "Venus"
planets[2] = "Earth"

earth := planets[2] //retrieves the planet at index 2
fmt.Println(earth)  //Earth

//获取数组的长度，通过内建函数len()即可
fmt.Println(len(planets)) //8
```
在Go中，当对数组的读写越界之后也是会引发异常的。

## 使用复合字面值来初始化数组
**复合字面值**（composite literals）是一种给复合类型进行初始化的紧凑语法。通过这个语法可以只需一句代码就将对数组的声明与初始化同时完成。类似C#中的数组初始化器。
```go
dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
```
更高级一点的用法，我们无需显式指定数组的长度，使用“...”来代替，Go编译器会自动帮我计算出数组的长度。
```go
planetCollection := [...]string{
    "Mercury",
    "Venus",
    "Earth",
    "Mars",
    "Jupiter",
    "Saturn",
    "Uranus",
    "Neptune",
}
fmt.Println(len(planetCollection)) //8
```
也可以使用索引来为数组赋初始值。
```go
array := [5]int{1: 10, 3: 30}
fmt.Printf("%+v\n", array) //[0 10 0 30 0]
```

## 遍历数组
### for循环遍历
* 使用 **for** 循环遍历数组
```go
dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

for i := 0; i < len(dwarfs); i++ {
    dwarf := dwarfs[i]
    fmt.Println(i, dwarf)
}
```
* 使用 **range** 关键字来遍历数组，之前也介绍过
```go
dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

//返回的两个参数，一个为索引 一个为对应的值
for i, dwarf := range dwarfs {
    fmt.Println(i, dwarf)
}
```

## 数组总是被深拷贝的
当我们声明了一个数组之后，无论是将这个数组赋值给其他变量还是作为参数传递给某个函数，都是将此数组的一个完整副本传递给变量或者函数的。
【示例】：将数组传递给变量
```go
planetCollection := [...]string{
    "Mercury",
    "Venus",
    "Earth",
    "Mars",
    "Jupiter",
    "Saturn",
    "Uranus",
    "Neptune",
}
planetsMarkII := planets //会将planets的完整副本赋值给planetMarkII
planets[2] = "whoops"    //改变了原数组planets，

//数组planetsMarkII是副本，所以不受影响
fmt.Println(planets[2])       //whoops
fmt.Println(planetsMarkII[2]) //Earth
```
【示例】：将数组传递给函数
```go
func terraform(planets [8]string){
    for i := range planets{
        planets[i] = "New " + planets[i]
    }
}

func main(){
    planetCollection := [...]string{
        "Mercury",
        "Venus",
        "Earth",
        "Mars",
        "Jupiter",
        "Saturn",
        "Uranus",
        "Neptune",
    }
    terraform(planets)
    fmt.Println(planets) //原数组无变动
}
```

## 数组的数组（二维数组）
这个没啥好说的，一个例子应该能理解

【示例】
```go
//二维数组
var board [8][8]string
board[0][0] = "r"
board[0][7] = "r"
for column := range board[1] {
    board[1][column] = "p"
}
fmt.Println(board)	//[[r       r] [p p p p p p p p] [       ] [       ] [       ] [       ] [       ] [       ]]
```

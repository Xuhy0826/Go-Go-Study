# 切片：指向数组的窗口

## 切分数组
Go语言中对数组的切片写法类似python，比如上节中的数组planets，planets[0:4]即可获取索引0到索引4的元素 **（不包括索引4的元素）**。
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
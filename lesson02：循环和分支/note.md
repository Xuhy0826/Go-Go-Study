# 循环和分支
大体和其他语言类似，快速过一遍即可

## 1.判断与分支
其实和其他语言的用法基本一致，分为几种常见结构：
1. if...else...，这个结构和其他语言基本一致
```go
room := "lake"

if room == "cave" {
    fmt.Println("you find yourself in a dimly lit cavern.")
} else if room == "lake" {
    fmt.Println("the ice seems solid enough")
} else if room == "underwater" {
    fmt.Println("the water is freezing cold")
} else {
    fmt.Println("nothing to say")
}
```
2. switch，不用写break，如果想继续向下可以使用fallthrough关键字

(1) 用法1
```go
room := "lake"

switch room {
case "cave":
    fmt.Println("you find yourself in a dimly lit cavern.")
case "lake":
    fmt.Println("the ice seems solid enough")
case "underwater":
    fmt.Println("the water is freezing cold")
default:
    fmt.Println("nothing to say")
}
```
(2) 用法2
```go
switch {
    case val == "jack":
        ...
    case val == "tom"
        ...
    case val == "jerry"
        ...
        fallthrough         //下降至下一分支
    default:
        ...
}
```
## 2.循环体

1. 常见的for循环，与其他语言类似
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

2. 类似while的for循环
```go
var count = 10
for count > 0 {
    fmt.Println(count)
    time.Sleep(time.Second)
    count--
}
```
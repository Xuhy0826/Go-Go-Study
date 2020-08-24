# 循环和分支

## 1.判断与分支
其实和其他语言的用法基本一致，分为几种常见结构：
1. if...else...，这个结构和其他语言基本一致

2. switch，不用写break，如果想继续向下可以使用fallthrough关键字

(1) 用法1
```
switch val{
    case v1:
        ...
    case v2:
        ...
    default:
        ...
}
```
(2) 用法2
```
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

1. 常见的for循环
```
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

2. 类似while的for循环
```
var count = 10
for count > 0 {
    fmt.Println(count)
    time.Sleep(time.Second)
    count--
}
```
# Go中没有Class

之前好几处都已经提到过Go中没有class，没有class要如何实现面向对象呢。

## 为struct绑定方法

和之前为基本类型绑定方法一样，使用一样的语法可以为struct类型绑定方法，看示例
```go
import (
	"fmt"
	"math"
)

//声明coordinate类型
//三维点坐标
type coordinate struct {
	x, y, z float64
}

//为声明的coordinate类型绑定方法distance
//计算点与点之间的距离
func (s coordinate) distance(t coordinate) float64 {
	return math.Pow(s.x-t.x, 2) + math.Pow(s.y-t.y, 2) + math.Pow(s.z-t.z, 2)
}

func main() {
	p1 := coordinate{x: 10.5, y: 20.1, z: 5.21}
	p2 := coordinate{x: 10.5, y: 20.1, z: 5.21}

	fmt.Println("Distance between p1 and p2 is ", p1.distance(p2))
}
```

## 构造函数
同样的，Go中也没有提供构造函数。一般都是使用一般函数来实现这个功能即可，为了使语义上更合理，一般对函数的命名使用New或者new（如果不想公开）开头。举个例子。
```go
//三维点坐标
type coordinate struct {
	x, y, z float64
}

func newCoordinate(x, y, z float64) coordinate {
	return coordinate{x, y, z}
}

func main() {
	p3 := newCoordinate(2.2, 10.3, -4.24)
	fmt.Printf("p3: %+v \n", p3)
}
```
另外，有些构造函数的命名就用`New`即可。比如error包中就包含New函数。因为Go语言中调用这个函数的时候会带上包名，即调用时的写法就是 `**error.New()**`，这样已经达到了语义上的含义，并且更加简洁。

## Class的替代方案
一般使用struct并为其绑定上相应的方法来达到和大多数class同样的效果。其实和第一段示例代码的意思差不多。但是在面向对象中，有关class的继承等在Go中如何体现，下节涉及再讲。
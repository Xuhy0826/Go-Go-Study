package main

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

//Go中没有构造函数，就用一般函数就可以实现功能
func newCoordinate(x, y, z float64) coordinate {
	return coordinate{x, y, z}
}

func main() {
	fmt.Println("lesson15 Go‘s Got no class")

	p1 := coordinate{x: 10.5, y: 20.1, z: 5.21}
	p2 := coordinate{x: 10.5, y: 20.1, z: 5.21}
	fmt.Println("Distance between p1 and p2 is ", p1.distance(p2))

	p3 := newCoordinate(2.2, 10.3, -4.24)
	fmt.Printf("p3: %+v \n", p3)
}

package main

import (
	"fmt"
	"image"
	"log"
	"sync"
	"time"
)

//Visited 用于记录网页是否被访问过
type Visited struct {
	mu      sync.Mutex
	visited map[string]int
}

//VisitLink 记录本次针对网址url的访问，更新对url的访问总次数
func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

//长时间运行的工作进程
// func worker() {
// 	n := 0
// 	next := time.After(time.Second) //创建一个计时器通道
// 	for {
// 		select {
// 		case <-next: //等待计时器触发
// 			n++
// 			fmt.Println(n)
// 			next = time.After(time.Second) //为下一次循环创建新的计时器
// 		}
// 	}
// }
func worker() {
	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 0}
	next := time.After(time.Second)
	for {
		select {
		case <-next:
			pos = pos.Add(direction)
			fmt.Println("current position is ", pos)
			next = time.After(time.Second)
		}
	}
}

//命令类型
type command int

const (
	right = command(0) //简单的代表向右转
	left  = command(1) //简单的代表向左转
)

//RoverDriver 用来控制探测器
type RoverDriver struct {
	commandc chan command
}

//NewRoverDriver 创建通道并启动工作进程
func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			}
			log.Printf("new direction %v", direction)

		case <-nextMove:
			pos = pos.Add(direction)
			log.Printf("move to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

//Left 会将探测器转向左方
func (r *RoverDriver) Left() {
	r.commandc <- left
}

//Right 会将探测器转向右方
func (r *RoverDriver) Right() {
	r.commandc <- right
}

var mu sync.Mutex

func main() {
	fmt.Println("lesson23 Concurrent state")

	mu.Lock()
	defer mu.Unlock()

	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(2 * time.Second)
	r.Right()
	time.Sleep(4 * time.Second)
}

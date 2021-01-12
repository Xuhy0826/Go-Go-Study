# 并发状态

在使用goroutine后不得不面对一个问题，就是当多个goroutine同时去操作同一个共享值就会发生并发问题，这种情况被称作竞争状态（race candition）。在其他语言中我们往往会通过上锁来解决这类问题，那么在Golang中该如何解决。

## 互斥锁
Go中提供了互斥锁`Mutex`(mutual exclusive)，存在于包sync中。从名字大概就可以理解其意思。   
互斥锁中有`Lock`和`Unlock`两个方法。`Lock`就是上锁，`Unlock`就是解锁。如果有goroutine尝试在互斥锁已经锁定的情况下再次调用`Lock`方法，那么它将等待直到解除锁定之后才能再次上锁。为了防止锁来锁去引发不可预知的错误，通常情况下只会在包的内部使用。
```
import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func main() {
	mu.Lock()
	defer mu.Unlock()  //通常为了防止有多个return方法而忘记解锁，Unlock通常都用defer来写
}
```
将sync.Mutex用作结构成员的做法是一种常见的模式。下面是个示例，现在有多个goroutine来执行网络爬虫，需要有个结构来存储所有被爬的网页的次数。如果使用映射来存储，在多个goroutine尝试更新映射时，就会产生竟态条件。那么这时就需要一个互斥锁来保护。
```
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
```
* 使用互斥锁时要小心陷入**死锁**。

## 原子函数
原子函数能够以很底层的加锁机制来同步访问整型变量和指针。atmoic包中两个有用的原子函数是 LoadInt64 和 StoreInt64。这两个函数提供了一种安全地读
和写一个整型值的方式。如下示例
```
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// shutdown 是通知正在执行的 goroutine 停止工作的标志
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 创建两个 goroutine
	go doWork("A")
	go doWork("B")

	// 给定 goroutine 执行的时间
	time.Sleep(1 * time.Second)

	// 该停止工作了，安全地设置 shutdown 标志
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()
}

// doWork 用来模拟执行工作的 goroutine，检测之前的 shutdown 标志来决定是否提前终止
func doWork(name string) {
	defer wg.Done()
	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
```
atmoic包的`AddInt64`函数。这个函数会同步整型值的加法，方法是强制同一时刻只能有一个**goroutine**运行并完成这个加法操作，例`atomic.AddInt64(&counter, 1)`，类似于C#中的原子累加器`Interlocked.Increment(ref ActivityCount);`。


## 长时间运行的工作进程
我们将一直存在并且独立运行的goroutine称为“工作进程”（worker）。比如一些定时执行某些功能的工作进程，如网站的轮询器等。在C#中我们可以使用一个定时器来完成这样的需求，那么在Golang中我们该如何搭建一个较为通用的工作进程。
```
//长时间运行的工作进程
func worker() {
	n := 0
	next := time.After(time.Second) //创建一个计时器通道
	for {
		select {
		case <-next: //等待计时器触发
			n++
			fmt.Println(n)
			next = time.After(time.Second) //为下一次循环创建新的计时器
		}
	}
}
```
其实上面的示例完全可以不用select和time.After，直接用一个time.Sleep来实现。这里主要为了方便将这个示例拓展成等待多个通道的工作进程。

### 综合示例
背景描述：现在有一个在火星表面行走的探测器，通过遥控发送命令可控制探测器的行走。探测器有一个工作进程来接受命令进行移动，并且定时刷新探测器的位置。   
首先我们将上面的工作进程进行改写。由于需要记录位置，image包中的Point结构很适合，它可以存储x轴和y轴的坐标，并且有一个Add方法可以将一个坐标点与另一个坐标点相加。
```
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
```
上面这个探测器中的工作进程功能还比较单一，只会让探测器直线前进。为了让它能够听从遥控器发送的指令进行转方向、停止、或者加速等，需要添加一个命令通道来发送命令。当工作进程从命令通道中接收到命令后便会立即执行相应的命令。
```
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

func main() {
	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(2 * time.Second)
	r.Right()
	time.Sleep(4 * time.Second)
}
```
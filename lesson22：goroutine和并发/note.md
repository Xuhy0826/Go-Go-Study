# goroutine和并发

在Go中如何实现其他语言中的多线程，答案就是使用**goroutine**。在Go语言中，独立运行的任务就被成为**goroutine**。可以类比成C#中的Thread，但是他们不是完全一样。goroutine的创建效率非常高，并且Go也能够很简洁地协同多个并发操作。

## 启动goroutine
在执行的操作之前加上一个`go`关键字即可，就是这么简单。看一个简单直接的例子。
```
import (
	"fmt"
	"time"
)

func sleepyGopher() {
	time.Sleep(time.Second * 3)
	fmt.Println("...snore...")
}

func main() {
	fmt.Println("lesson22 goroutine and concurrency")

	//通过关键字go，启动goroutine
	go sleepyGopher()
	fmt.Println("this is main func")
	time.Sleep(time.Second * 4)
}
```
执行结果是，控制台会输出`.this is main func`，接着3秒之后，控制台会输出`...snore...`。但是注意，因为在main函数返回时，该程序运行的所有goroutine都会被回收，这就是为什么例子中的main函数需要一个比`sleepyGopher`函数长的等待时间。

## 启动多个goroutine
每次使用`go`关键字都会创建一个新的`goroutine`。
```
import (
	"fmt"
	"time"
)

func sleepyGopher() {
	time.Sleep(time.Second * 3)
	fmt.Println("...snore...")
}

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher()
	}
	time.Sleep(time.Second * 4)
}
```
带参数的函数，一样可以简单的使用go关键字启动goroutine。为了标记每个goroutine，接下来为函数传入一个id。
```
func sleepyGopher(id int) {
	time.Sleep(time.Second * 3)
	fmt.Println("...snore...", id)
}

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher(i)
	}
	time.Sleep(time.Second * 4)
}
```
输出：
```
...snore... 0
...snore... 3
...snore... 2
...snore... 1
...snore... 4

```
其实每次的输出都可能不一样，可以看出，goroutine的执行顺序不是我们可以控制的。   
看完上面的例子后，抛出两个问题
1. main函数不得不Sleep一定的时间来确保所有的goroutine全部执行完毕。那么如果goroutine中执行的不是上面的这种可知具体耗时的操作（比如数据库操作，网络访问等），那么如何确定goroutine什么时候结束呢。
2. 不同的goroutine之间如何传递数据   
接下来的**通道**即可解决这两个问题。

## 通道
* 通道（channel）可以在多个goroutine之间安全地传递值。可以类比想象成我们平时用的消息队列，可以向通道中写入值，可以从通道中取出值。   
* 跟Go中的其他类型一样，可以将通道作为变量，传递至函数，结构中的字段。
* 创建通道的方法：使用内置的`make`函数。并且还要指定相应的类型。
```
c := make(chan int)
```
上面这个通道只能传递int类型。
* 对通道中读写值使用左箭头操作符`<-`:
- 向通道中写入值：`c <- 7`
- 从通道中读取值：`r := <- c`   
不论是向通道中写入值或者读取值，都会形成阻塞。比如发送操作会等待直到有另一个goroutine尝试对相同的通道执行读取操作为止。相同的，读取操作会等待直到有另一个goroutine尝试对相同的通道执行写入操作。
```
import (
	"fmt"
	"math/rand"
	"time"
)

func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	fmt.Println("...snore...", id)
	c <- id
}

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	for i := 0; i < 5; i++ {
		gopherID := <-c
		fmt.Println("gopher", gopherID, "has finished sleeping")
	}
}
```
上面的5个goroutine都向通道`c`中写入了ID值，main函数会等待到5个goroutine全部执行结束，即是向通道中写入值之后，才会返回，这样我们就可以不需再像之前一样让main函数也Sleep一段时间来确保goroutine的执行了。

## 使用select处理多个通道
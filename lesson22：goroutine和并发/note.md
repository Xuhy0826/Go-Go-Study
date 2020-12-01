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
执行结果是，控制台会输出`.this is main func`，接着3秒之后，控制台会输出`...snore...`。但是注意，因为在main函数返回时，该程序运行的所有goroutine都会被回收，这就是为什么例子中的main函数需要一个比`sleepyGopher`函数长的等待时间。当然处理这个问题有很多方法，比如使用sync包中的`sync.WaitGroup`。

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
然而在很多时候，我们不希望程序一直阻塞在等待通道处。根据以往的经验，我们可以想到的是为这些等待设置超时时间。   
Go提供了`time.After`函数来设置超时时间，`time.After`函数会返回一个通道，这个通道会在特定的时间后接收到一个值（由Go运行时发送）。如果我们不想程序一直等待所有的goroutine完成而设置一个超时时间，一个思路是让程序同时等待由`time.After`函数返回的计时通道和其他通道，如果计时通道的值返回了就不再去等待其他通道了。
为了实现这个功能，Go提供的`select`语句很方便，其语法与`switch`很相似，某个case的准备就绪就会执行相应的操作。这样一来，我们就可以同时监视多个通道了。
```
timeout := time.After(2 * time.Second)
for i := 0; i < 5; i++ {
	select {
	case gopherID := <-c:
		fmt.Println("gopher", gopherID, "has finished sleeping")
	case <-timeout: //等待直到超时
		fmt.Println("my patience ran out")
		return
	}
}
```
即使程序已经停止等待goroutine，但只要main函数没有返回，仍在运行的goroutine就仍在占用内存。所以我们应该主动去关闭无用的goroutine。可以使用通道的`Close()`方法，后面再看。   
**【注意】** 如果select语句不包含任何分支，将会永远等待下去。   
#### nil通道
如果不使用make函数初始化通道变量，那么和映射和切片一样，通道的值是nil。对其进行发送和接收操作不会引起panic，但是会导致永久阻塞。但是如果对nil通道执行`Close()`方法则会引发panic。
* nil通道的应用场景   
nil通道的存在并不是一无是处的，比如，一个包含select语句的循环，如果我们不希望循环每次循环都等待所有case，那么可以先将某些通道设置为nil，等到待发送的值准备好之后再将这些nil通道初始化，再去执行发送操作。

#### 阻塞和死锁
goroutine在等待和发送通道时会引起阻塞，等待时程序会一直监视通道啥时候来值。但是这种阻塞和那些空转死循环不一样，除了goroutine本身所占的少量资源外，goroutine并不消耗任何其他资源。当一个或多个goroutine因为某些永远无法发生的事情而被阻塞时，这就发生了**死锁**。
```
func main(){
	c := make(chan, int)
	<- c
}
```

## 实践一下：流水线
接下来做一个实例综合运用本节学的几个语法。用3个goroutine来形成一个流水线作业。
```
//上游
func sourceGopher(downStream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downStream <- v
	}
	downStream <- ""
}

//中游
func filterGopher(upStream, downStream chan string) {
	for {
		item := <-upStream
		if item == "" {
			downStream <- ""
			return
		}
		if !strings.Contains(item, "bad") {
			downStream <- item
		}
	}
}

//下游
func printGopher(upStream chan string) {
	for {
		v := <-upStream
		if v == "" {
			return
		}
		fmt.Println(v)
	}
}
```
```
func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go sourceGopher(c1)
	go filterGopher(c1, c2)
	printGopher(c2)
}
```
上面的示例中我们是使用空字符串作为一个结束标志，但是这不是很稳定的做法。如果上游的字符串数组中包含一个空字符串，那么流程会被意外关闭。实际上更好的做法是使用`close`函数来关闭通道。
```
close(c1)
```
如果向已关闭的通道执行写入会引发panic，读取已关闭的通道会得到相应的零值。   
**【注意】** 如果在循环中读取一个已关闭的通道，并且没有检查该通道是否已关闭，那么这个循环会一直运转下去并消耗大量的性能。所以务必对可能关闭的通道检查是否关闭。   
检查通道是否关闭的写法：
```
v, ok := <-c
```
如果ok是false，那么说明通道c已关闭。   
那么上游和中游的代码可以做如下优化
```
//上游
func sourceGopher(downStream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downStream <- v
	}
	//downStream <- ""
	close(downStream)
}

//中游
func filterGopher(upStream, downStream chan string) {
	for {
		//item := <-upStream
		item, ok := <-upStream
		if !ok {
			//downStream <- ""
			close(downStream)
			return
		}
		if !strings.Contains(item, "bad") {
			downStream <- item
		}
	}
}
```
另外，由于“从通道中读取值，直到通道被关闭为止”这个操作很常用，所以Go提供了快捷的方法。使用`range`来读取通道，程序会在通道关闭前一直去获取通道的值。   
这样一来，中游的代码可以再次优化。
```
func filterGopher(upStream, downStream chan string) {
	//使用range来读取通道的值
	for item := range upStream {
		if !strings.Contains(item, "bad") {
			downStream <- item
		}
	}
	close(downStream)
}
```
下游的代码同理
```
func printGopher(upStream chan string) {
	//使用range来读取通道的值
	for v := range upStream {
		fmt.Println(v)
	}
}
```

## 总结
goroutine是Go语言中十分重要的知识点，涉及的内容不少，做个总结
1. 使用go关键字启动goroutine
2. 通道用于多个goroutine之间传值
3. 使用make函数创建通道
4. 使用`<-`读写通道
5. 使用select语句同时等待多个通道
6. 使用time.After函数实现超时
7. 使用close函数关闭通道
8. 使用range语句读取通道的值
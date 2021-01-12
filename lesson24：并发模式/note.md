# 并发模式

通过3个在实际工程中可以使用的包来理解不同的并发模式。

## runner
可创建一个执行任务的对象，称为runner。runner可以接收多个任务并在后台依次执行。在执行过程中可以通过向其发送终止信号结束执行。如果执行耗时超过了初始设定的超时时间，也会结束执行。这是一种很有用的模式，接下来就是如何构建这个runner的示例。其中如何判断执行时间是否超时和是否有终止信号都是通过通道来监视的。   
在设计上，可支持以下终止点：
* 程序可以在分配的时间内完成工作，正常终止；
* 程序没有及时完成工作，“自杀”；
* 接收到操作系统发送的中断事件，程序立刻试图清理状态并停止工作。
首先`runner`结构中包含四个字段，其中有3个都是通道类型，用来辅助管理生命周期。
```
type Runner struct {
	// tasks表示要执行的任务集合
	tasks []func(int)

	// interrupt 通道获取从操作系统发来的信号
	interrupt chan os.Signal

	// timeout 报告处理任务已经超时，标记为“<-chan”表示此通道为一个单向只读通道，因为向通道中传入数据是Go运行时
	timeout <-chan time.Time

	// complete 通道报告处理任务已经完成
	complete chan error
}
```
interrupt 通道收发 os.Signal 接口类型的值，用来从主机操作系统接收中
断事件。其中`os.Signal`接口类型是用来从主机操作系统中接收中断事件的。
```
type Signal interface {
    String() string
    Signal() 
}
```
`complete`是一个收发`error`接口类型值的通道。当执行过程中发生错误会将错误传入通道。main函数便可获取错误。如果正常执行完任务，便返回一个nil值。  
`timeout`是用来监视超时的通道。达到初始设定的超时时间，通道中可获取到值，此时runner便会终止任务。  
下面，针对可能发生的错误类型预先定义好两个错误变量
```
// 超时的错误，这个错误值会在收到超时事件时返回
var ErrTimeOut = errors.New("received timeout")

// 中断的错误，会在收到操作系统的中断事件时返回
var ErrInterrupt = errors.New("received interrupt")
```
为了方便创建一个任务执行者runner，为其定义工厂函数。
```
//New 设定一个超时时间，返回一个新的准备使用的 Runner 类型的指针
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}
```
在New函数中初始化了`Runner`的所有通道，其中task字段的零值是nil，已经满足初始化的要求，所以没有被明确初始化。  
但是值得注意的是 interrupt 通道被初始化为缓冲区容量为 1 的通道。这是因为这可以保证通道至少能接收一个来自语言运行时的 os.Signal 值，确保语言运行时发送这个事件的时候不会被阻塞。  
但是 complete 通道被初始化为无缓冲的通道，是因为需要使用它来控制我们整个程序是否终止。
接下来为Runner类型关联几个需要的方法。  
（1）注册任务
```
// Add 为 Runner 添加任务。任务是一个接收一个int类型的ID作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}
```
（2）后台按顺序执行每个任务。但是在执行之前，会先调用`gotInterrupt`方法来检查是否有要从操作系统接收的事件。
```
// run 依次执行已注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		//执行前先判断是否有中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		//执行任务
		task(id)
	}

	//任务执行完成
	return nil
}

// gotInterrupt 验证是否接收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	// 当接收到中断信号时
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	// 防止阻塞，使其继续正常运行
	default:
		return false
	}
}
```
值得注意的是，`gotInterrupt`方法中用到了带default分支的select语句。在没有default的select语句中，如果等待的几个通道都没有值的话就会阻塞。如果有default，通道都没有值的话，就会执行 default 分支。  
（3）公开方法`Start`，开启runner执行
```
// Start 执行所有任务，并且监视通道事件
func (r *Runner) Start() error {
	//设置我们希望获取哪些系统信号
	signal.Notify(r.interrupt, os.Interrupt)

	// goroutine 执行任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 当任务处理完成时
	case err := <-r.complete:
		return err
	// 当任务处理程序运行超时
	case <-r.timeout:
		return ErrTimeout
	}
}
```
`Start`方法中声明了一个匿名函数，并单独启动`goroutine`来执行。在 goroutine 的内部调用了`run`方法，并将这个方法返回的`error`接口值发送到`complete`通道。创建 goroutine 后，`Start`进入一个`select`语句，阻塞等待两个事件中的任意一个。如果从`complete`通道接收到`error`接口值，那么该 goroutine 要么在规定的时间内完成了分配的工作，要么收到了操作系统的中断信号。无论哪种情况，收到的`error`接口值都会被返回，随后方法终止。如果从`timeout`通道接收到`time.Time`值，就表示 goroutine 没有在规定的时间内完成工作。这种情况下，程序会返回`ErrTimeout`变量。  
下面可以看到，在main.go中如何使用Runner来执行任务。
```
package main

import (
	"log"
	"os"
	"time"

	"demo24/runner"
)

// timeout 规定了必须在多少秒内处理完成
const timeout = 3 * time.Second

func main() {
	log.Println("Starting work.")

	r := runner.New(timeout)
	// 加入要执行的任务
	r.Add(createTask(), createTask(), createTask())
	// 执行任务并处理结果
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}

```
## pool

## worker
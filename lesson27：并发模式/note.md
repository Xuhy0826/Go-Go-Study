# 并发模式

通过3个在实际工程中可以使用的包来理解不同的并发模式。

## runner
可创建一个执行任务的对象，称为`runner`。`runner`可以接收多个任务并在后台依次执行。在执行过程中可以通过向其发送终止信号结束执行。如果执行耗时超过了初始设定的超时时间，也会结束执行。这是一种很有用的模式，接下来就是如何构建这个`runner`的示例。其中如何判断执行时间是否超时和是否有终止信号都是通过通道来监视的。   
在设计上，可支持以下终止点：
* 程序可以在分配的时间内完成工作，正常终止；
* 程序没有及时完成工作，“自杀”；
* 接收到操作系统发送的中断事件，程序立刻试图清理状态并停止工作。
首先`runner`结构中包含四个字段，其中有3个都是通道类型，用来辅助管理生命周期。
```go
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
`interrupt` 通道收发 `os.Signal` 接口类型的值，用来从主机操作系统接收中断事件。其中`os.Signal`接口类型是用来从主机操作系统中接收中断事件的。
```go
type Signal interface {
    String() string
    Signal() 
}
```
`complete`是一个收发`error`接口类型值的通道。当执行过程中发生错误会将错误传入通道。`main`函数便可获取错误。如果正常执行完任务，便返回一个nil值。  
`timeout`是用来监视超时的通道。达到初始设定的超时时间，通道中可获取到值，此时runner便会终止任务。  
下面，针对可能发生的错误类型预先定义好两个错误变量
```go
// 超时的错误，这个错误值会在收到超时事件时返回
var ErrTimeOut = errors.New("received timeout")

// 中断的错误，会在收到操作系统的中断事件时返回
var ErrInterrupt = errors.New("received interrupt")
```
为了方便创建一个任务执行者runner，为其定义工厂函数。
```go
//New 设定一个超时时间，返回一个新的准备使用的 Runner 类型的指针
func New(timeout time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(timeout),
	}
}
```
在`New`函数中初始化了`Runner`的所有通道，其中`task`字段的零值是`nil`，已经满足初始化的要求，所以没有被明确初始化。  
但是值得注意的是 `interrupt` 通道被初始化为缓冲区容量为 1 的通道。这是因为这可以保证通道至少能接收一个来自语言运行时的 `os.Signal` 值，确保语言运行时发送这个事件的时候不会被阻塞。  
但是 `complete` 通道被初始化为无缓冲的通道，是因为需要使用它来控制我们整个程序是否终止。
接下来为Runner类型关联几个需要的方法。  
（1）注册任务
```go
// Add 为 Runner 添加任务。任务是一个接收一个int类型的ID作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}
```
（2）后台按顺序执行每个任务。但是在执行之前，会先调用`gotInterrupt`方法来检查是否有要从操作系统接收的事件。
```go
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
值得注意的是，`gotInterrupt`方法中用到了带`default`分支的`select`语句。在没有`default`的`select`语句中，如果等待的几个通道都没有值的话就会阻塞。如果有`default`，通道都没有值的话，就会执行 `default` 分支。  
（3）公开方法`Start`，开启runner执行
```go
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
> signal.Notify()函数
> 声明：`func Notify(c chan<- os.Signal, sig ...os.Signal)`
> Notify函数让signal包将输入信号转发到c。如果没有列出要传递的信号，会将所有输入信号传递到c；否则只传递列出的输入信号。

`Start`方法中声明了一个匿名函数，并单独启动`goroutine`来执行。在 `goroutine` 的内部调用了`run`方法，并将这个方法返回的`error`接口值发送到`complete`通道。创建 goroutine 后，`Start`进入一个`select`语句，阻塞等待两个事件中的任意一个。如果从`complete`通道接收到`error`接口值，那么该 goroutine 要么在规定的时间内完成了分配的工作，要么收到了操作系统的中断信号。无论哪种情况，收到的`error`接口值都会被返回，随后方法终止。如果从`timeout`通道接收到`time.Time`值，就表示 goroutine 没有在规定的时间内完成工作。这种情况下，程序会返回`ErrTimeout`变量。  
下面可以看到，在main.go中如何使用Runner来执行任务。
```go
package main

import (
	"log"
	"os"
	"time"

	"demo27/runner"
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
pool包用于展示如何使用有缓冲的通道实现资源池，来管理可以在任意数量的`goroutine`之间共享及独立使用的资源。想想很多地方都会有资源池的实践，最常见的比如数据库的连接池，网络连接池。当一个`goroutine`需要从池中获取一个资源，可以向池中申请，使用完再归还到池中。  
首先定义资源池的结构体，其中包含四个字段。
```go
// Pool 管理一组可以安全地在多个goroutine间共享的资源。且资源必须实现了io.Closer接口
type Pool struct {
	//互斥锁
	m         sync.Mutex

	//保存共享的资源
	resources chan io.Closer

	//工厂函数：当池需要一个新资源且池中没有可用资源时，用这个函数创建。
	factory   func() (io.Closer, error)

	//标志量，Pool 是否已经被关闭
	closed    bool
}
```
如果资源池已被关闭，再去请求池中的资源时会返回错误。预先定义好这个错误。
```go
// ErrPoolClosed 描述一个请求已关闭池的错误
var ErrPoolClosed = errors.New("Pool has been closed")
```
定义资源池结构体的工厂函数，`New`函数
```go
// New 创建一个用来管理资源的池。
// 需要一个工厂函数并规定池的大小
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}
```
显而易见，在创建资源池时要指定池的大小，并且指定生成资源的方法。   
定义好结构体之后，接下来为其关联3个必要的方法，分别是`获取资源`,`释放资源`,`关闭资源池`。而且必须要做好线程安全。  
（1）获取资源的方法`Acquire`
```go
// Acquire 从池中获取资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	// 先检查池中是否还有空闲的资源
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil

	// 因为没有空闲资源，创建新的资源返回
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}
```
（2）释放资源的方法`Release`
```go
// Release 释放资源，将使用后的资源放回池里
func (p *Pool) Release(r io.Closer) {
	//上锁，保证本操作与Close操作的安全
	p.m.Lock()
	defer p.m.Unlock()

	//若池已关闭，则直接关闭资源即可
	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("Release:", "In Queue")

	//如果池已满，则直接关闭这个资源
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}
```
（3）关闭资源池的方法`Close`
```go
// Close 会让资源池停止工作，并关闭所有现有的资源
func (p *Pool) Close() {
	//上锁，保证本操作与Release的安全
	p.m.Lock()
	defer p.m.Unlock()

	//若池已关闭，直接返回
	if p.closed {
		return
	}

	//关闭池
	p.closed = true

	// 在清空通道里的资源之前，将通道关闭
	// !!!如果不这样做，会发生死锁，下面的for循环会一直等待p.resources的值!!!
	close(p.resources)

	// 关闭所有资源
	for r := range p.resources {
		r.Close()
	}
}
```
需要注意的是`Release`和`Close`方法使用同一个互斥锁上锁的，这样可以阻止这两个方法在不同 goroutine 里同时运行。
看下如何在main.go中使用
```go
package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"demo27/pool"
)

const (
	maxGoroutines   = 25 // 要使用的 goroutine 的数量
	pooledResources = 2  // 池中的资源的数量
)

// dbConnection 模拟要共享的资源
type dbConnection struct {
	ID int32
}

// Close 实现了 io.Closer 接口，以便 dbConnection可以被池管理。
// Close 用来完成任意资源的释放管理
func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

// idCounter 用来给每个连接(dbConnection)分配一个独一无二的 id
var idCounter int32

// createConnection 是一个工厂函数，会调用这个函数创建新资源
func createConnection() (io.Closer, error) {
	//使用原子方法
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)

	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// 创建用来管理连接的池
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	// 使用池里的连接来完成查询
	for query := 0; query < maxGoroutines; query++ {
		// 每个 goroutine 需要自己复制一份要
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	// 等待 goroutine 结束
	wg.Wait()
	log.Println("Shutdown Program.")

	//关闭池
	p.Close()
}

// performQueries 用来测试连接的资源池
func performQueries(query int, p *pool.Pool) {
	//从池中请求连接
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	//释放资源
	defer p.Release(conn)
	// 用等待来模拟查询响应
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
```

## work
work包的场景是多个 `goroutine` 等待一个work工作队列中派发来的工作任务，这些`goroutine`拿到任务后并发执行，执行完工作的`goroutine`回来继续等任务。在这种情况下，使用无缓冲的通道要比随意指定一个缓冲区大小的有缓冲的通道好，因为这个情况下既不需要一组 `goroutine` 配合执行，各干各的，来一个任务，任意一个goroutine领走去执行就完事了。  
lesson24的练习其实就是这种场景。
```go
package work

import "sync"

// Worker 必须满足接口类型，才能使用工作池
type Worker interface {
	Task()
}

// Pool 提供一个 goroutine 池， 这个池可以完成任何已提交的 Worker 任务
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New 创建一个新工作池
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

// Run 提交工作任务到工作池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown 等待所有 goroutine 停止工作
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
```
接下来让看一下 main.go 如何使用。示例的工作任务很简单，只是打印出name字段。
```go
// names 提供了一组用来显示的名字
var names = []string{
	"steve",
	"bob",
	"mary",
	"jason",
	"therese",
}

type namePrinter struct {
	name string
}

// Task 实现 Worker 接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func testWork() {
	// 创建有两个goroutine的工作池
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}

			go func() {
				//将任务提交
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	p.Shutdown()
}
```
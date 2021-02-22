package main

import (
	"io"
	"log"
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"demo27/pool"
	"demo27/runner"
	"demo27/work"
)

func main() {

	//testRunner()

	//testPool()

	testWork()
}

//**************************测试Runner**************************
// timeout 规定了必须在多少秒内处理完成
const timeout = 3 * time.Second

// 测试Runner
func testRunner() {
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

//**************************************************************

//**************************测试Pool****************************
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

// 测试Pool
func testPool() {
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

	defer p.Release(conn)
	// 用等待来模拟查询响应
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}

//**************************************************************

//**************************测试Runner**************************
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

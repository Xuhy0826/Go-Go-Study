package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// counter 是所有 goroutine 都要增加其值的变量
	counter  int64
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	// 分配一个逻辑处理器给调度器使用
	// GOMAXPROCS函数可以更改调度器可以使用的逻辑处理器的数量
	//runtime.GOMAXPROCS(1)
	// 给每个可用的核心分配一个逻辑处理器
	//runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(2)

	fmt.Println("Start Goroutines")

	// go func() {
	// 	defer wg.Done()

	// 	for count := 0; count < 3; count++ {
	// 		for char := 'a'; char < 'a'+26; char++ {
	// 			fmt.Printf("%c ", char)
	// 		}
	// 	}
	// }()

	// go func() {
	// 	defer wg.Done()

	// 	for count := 0; count < 3; count++ {
	// 		for char := 'A'; char < 'A'+26; char++ {
	// 			fmt.Printf("%c ", char)
	// 		}
	// 	}
	// }()

	//go printPrime("A")
	//go printPrime("B")

	//go incCounter(1)
	//go incCounter(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)

	fmt.Println("Waiting To Finish")

	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()

	//fmt.Println("Final Counter:", counter)
}

func printPrime(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		counter++

		//这个函数会同步整型值的加法，方法是强制同一时刻只能有一个 goroutine 运行并完成这个加法操作。
		// 安全地对 counter 加 1
		atomic.AddInt64(&counter, 1)
		// 当前 goroutine 从线程退出，并放回到队列
		runtime.Gosched()
	}
}

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

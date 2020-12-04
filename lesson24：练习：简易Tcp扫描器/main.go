package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
	"time"
)

/*【需求】：
扫描一个给定ip的指定区间的端口，检测tcp是否可以连通。
要求开启一定数量的工作进程并行执行检测，获取所有打开的端口号并按序打印出来。
如开启100个工作进程并发执行检测，检测IP为129.204.242.187的1~1024端口哪些是打开的，并按从小到大输出这些打开的端口。
*/

//工作进程：检测tcp端口是否打开
func checkPortWork(host string, tasks chan int, resultReceiver chan int, timeout time.Duration /*, wg *sync.WaitGroup*/) {
	for port := range tasks {
		url := fmt.Sprintf("%s:%d", host, port)
		//fmt.Println("try to dial port ", port)
		fmt.Print(". ")
		//检测端口
		_, err := net.DialTimeout("tcp", url, timeout)
		if err != nil {
			resultReceiver <- 0
		} else {
			resultReceiver <- port
		}
		//wg.Done()
	}
}

func main() {
	//使用flag包，可以实现在执行时给定参数
	hostname := flag.String("hostname", "129.204.242.187", "hostname to test")
	portBegin := flag.Int("portBegin", 1, "the port on which the scanning starts")
	portEnd := flag.Int("portEnd", 1024, "the port on which the scanning ends")
	timeout := flag.Duration("timeout", time.Millisecond*300, "timeout")
	workerNum := flag.Int("workerNum", 100, "the number of worker")
	flag.Parse()

	//wg := &sync.WaitGroup{}
	tasks := make(chan int, *workerNum) //容量100的通道
	resultReceiver := make(chan int)    //接收测试结果

	fmt.Printf("scanning the %s\n", *hostname)

	if *portBegin > *portEnd || *portEnd <= 0 || *portEnd >= 63528 {
		fmt.Println("args is invalid")
	}

	//开启工作进程
	for i := 1; i <= cap(tasks); i++ {
		go checkPortWork(*hostname, tasks, resultReceiver, *timeout)
	}

	//下发“检测”任务
	go func(begin, end int) {
		for i := begin; i <= end; i++ {
			tasks <- i
		}
	}(*portBegin, *portEnd)

	//接收工作进程返回的检测结果
	results := []int{}
	for i := *portBegin; i <= *portEnd; i++ {
		port := <-resultReceiver
		if port != 0 {
			results = append(results, port)
		}
	}

	//排序并输出结果
	sort.Ints(results)
	for _, port := range results {
		fmt.Printf("\n port %d is open \n", port)
	}

	close(tasks)
	close(resultReceiver)

	fmt.Println("scanning taks finished")
}

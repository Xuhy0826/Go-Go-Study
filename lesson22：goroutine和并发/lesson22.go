package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	fmt.Println("...snore...", id)
	c <- id
}

//**********************************************************************
//实践：
//**********************************************************************
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
	// for {
	// 	//item := <-upStream
	// 	item, ok := <-upStream
	// 	if !ok {
	// 		//downStream <- ""
	// 		close(downStream)
	// 		return
	// 	}
	// 	if !strings.Contains(item, "bad") {
	// 		downStream <- item
	// 	}
	// }
	for item := range upStream {
		if !strings.Contains(item, "bad") {
			downStream <- item
		}
	}
	close(downStream)
}

//下游
func printGopher(upStream chan string) {
	// for {
	// 	v := <-upStream
	// 	if v == "" {
	// 		return
	// 	}
	// 	fmt.Println(v)
	// }
	for v := range upStream {
		fmt.Println(v)
	}
}

func main() {
	fmt.Println("lesson22 goroutine and concurrency")

	//通过关键字go，启动goroutine
	// go sleepyGopher()
	// fmt.Println("this is main func")
	// time.Sleep(time.Second * 4)

	// for i := 0; i < 5; i++ {
	// 	go sleepyGopher(i)
	// }
	// time.Sleep(time.Second * 4)

	// c := make(chan int)
	// for i := 0; i < 5; i++ {
	// 	go sleepyGopher(i, c)
	// }

	// for i := 0; i < 5; i++ {
	// 	gopherID := <-c
	// 	fmt.Println("gopher", gopherID, "has finished sleeping")
	// }

	// timeout := time.After(2 * time.Second)
	// for i := 0; i < 5; i++ {
	// 	select {
	// 	case gopherID := <-c:
	// 		fmt.Println("gopher", gopherID, "has finished sleeping")
	// 	case <-timeout: //等待直到超时
	// 		fmt.Println("my patience ran out")
	// 		return
	// 	}
	// }

	//实践：
	c1 := make(chan string)
	c2 := make(chan string)
	go sourceGopher(c1)
	go filterGopher(c1, c2)
	printGopher(c2)
}

package main

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
	fmt.Println("lesson22 goroutine and concurrency")

	//通过关键字go，启动goroutine
	// go sleepyGopher()
	// fmt.Println("this is main func")
	// time.Sleep(time.Second * 4)

	// for i := 0; i < 5; i++ {
	// 	go sleepyGopher(i)
	// }
	// time.Sleep(time.Second * 4)

	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	for i := 0; i < 5; i++ {
		gopherID := <-c
		fmt.Println("gopher", gopherID, "has finished sleeping")
	}
}

package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func main() {
	fmt.Println("lesson23 Concurrent state")

	mu.Lock()
	defer mu.Unlock()
}

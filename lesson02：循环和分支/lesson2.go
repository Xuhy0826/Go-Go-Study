//lesson02：循环和分支
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("lesson2：循环和分支")

	room := "lake"
	//if
	if room == "cave" {
		fmt.Println("you find yourself in a dimly lit cavern.")
	} else if room == "lake" {
		fmt.Println("the ice seems solid enough")
	} else if room == "underwater" {
		fmt.Println("the water is freezing cold")
	} else {
		fmt.Println("nothing to say")
	}

	//switch
	switch room {
	case "cave":
		fmt.Println("you find yourself in a dimly lit cavern.")
	case "lake":
		fmt.Println("the ice seems solid enough")
	case "underwater":
		fmt.Println("the water is freezing cold")
	default:
		fmt.Println("nothing to say")
	}

	//关于循环
	//（1）for循环
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//(2)功能类似while循环
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
}

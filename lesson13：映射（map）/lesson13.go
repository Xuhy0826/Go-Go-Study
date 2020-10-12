package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("lesson13 map")

	//声明map，使用复合字面量的声明方式
	temperature := map[string]int{
		"Mars":  -65,
		"Earth": 15,
	}
	//读
	temp := temperature["Earth"]
	fmt.Printf("On avarage the Earth is %v C \n", temp) //On avarage the Earth is 15 C

	//写
	temperature["Mars"] = 16
	temperature["Vemus"] = 464

	fmt.Println(temperature) //map[Earth:15 Mars:16 Vemus:464]

	//如果访问不存在的键，则反映值类型的零值
	fmt.Println(temperature["Moon"]) //0

	//逗号与ok
	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On avarage the Earth is %v C \n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}

	//map是不会被复制的
	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}
	planetsMarkII := planets
	planets["Earth"] = "Whoops"
	fmt.Println(planets)       //map[Earth:Whoops Mars:Sector ZZ9]
	fmt.Println(planetsMarkII) //map[Earth:Whoops Mars:Sector ZZ9]

	//移除map中的元素（key-value对）
	delete(planets, "Earth")
	delete(planets, "Moon")    //移除不存在的元素，不会引发panic
	fmt.Println(planets)       //map[Mars:Sector ZZ9]
	fmt.Println(planetsMarkII) //map[Mars:Sector ZZ9]

	//综合实例
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -28.0, -33.0,
	}
	//键：float64，值；[]float64
	groups := make(map[float64][]float64)

	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10 //将温度按10度的跨度进行分组
		groups[g] = append(groups[g], t)
	}

	for g, temperatures := range groups {
		fmt.Printf("%v : %v\n", g, temperatures)
	}

	/*
		上例输出结果：
		-20 : [-28 -29 -23 -28]
		30 : [32]
		-30 : [-31 -33]
	*/
}

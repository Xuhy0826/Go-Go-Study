package main

import "fmt"

type location struct {
	lat  float64
	long float64
}

func main() {
	fmt.Println("lesson14 struct")
	//声明结构
	var curiosity struct {
		lat  float64
		long float64
	}

	curiosity.lat = -4.9773
	curiosity.long = 137.4283

	fmt.Println(curiosity.lat, curiosity.long) //-4.9773 137.4283
	fmt.Println(curiosity)                     //{-4.9773 137.4283}

	//struct是值传递的
	curiosityMarkII := curiosity
	curiosity.lat = 0
	fmt.Println(curiosity)       //{0 137.4283}
	fmt.Println(curiosityMarkII) //{-4.9773 137.4283}

	var spirit location
	spirit.lat = -14.5637
	spirit.long = 175.3774

	var opprtunity location
	opprtunity.lat = -1.9473
	opprtunity.long = 352.8434

	fmt.Println(spirit)     //{-14.5637 175.3774}
	fmt.Println(opprtunity) //{-1.9473 352.8434}

	//两种输出方式
	fmt.Printf("%v\n", curiosity)  //{0 137.4283}
	fmt.Printf("%+v\n", curiosity) //{lat:0 long:137.4283}

	// lats := []float64{-4.5422, 8.152, -2.5152, 4.215}
	// longs := []float64{215.21, 125.14, 23.145, 135.512}

	//struct和切片的结合使用
	locations := []location{
		{lat: -4.5422, long: 215.21},
		{lat: 8.152, long: 125.14},
		{lat: -2.5152, long: 23.145},
		{lat: 4.215, long: 135.512},
	}

	for _, loc := range locations {
		fmt.Printf("%+v\n", loc)
	}
	/*
	*	lat:-4.5422 long:215.21}
	*	{lat:8.152 long:125.14}
	*	{lat:-2.5152 long:23.145}
	*	{lat:4.215 long:135.512}
	 */
}

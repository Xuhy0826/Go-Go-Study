package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func main() {
	fmt.Println("lesson18 Pointer(1)")

	answer := 42
	fmt.Println(&answer) //0xc0000a00b0

	address := &answer
	fmt.Println(*address) //42

	fmt.Printf("address is a %T\n", address) //address is a *int

	//fmt.Println(*answer) //*用在int类型前会报错

	china := "China"
	var home *string
	fmt.Printf("home is a %T\n", home) //home is a *string

	home = &china
	fmt.Println(*home) //China

	//home = &answer //error:cannot use &answer (type *int) as type *string in assignment

	var administrator *string

	//指针指向第一个人
	scolese := "Christopher J. Scolese"
	administrator = &scolese
	fmt.Println(*administrator) //Christopher J. Scolese

	//指针指向第二个人
	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(*administrator) //Charles F. Bolden

	//修改bolden变量，使用指针访问可以看到变量的更改
	bolden = "Charles Frank Bolden Jr."
	fmt.Println(*administrator) //Charles Frank Bolden Jr.

	//也可以通过“解引用”间接改变变量
	*administrator = "Maj. Gen. Charles Frank Bolden Jr."
	fmt.Println(bolden) //Maj. Gen. Charles Frank Bolden Jr.

	//把指针赋值给变量，将会产生一个指向相同变量的指针。
	major := administrator
	*major = "Maj. General Charles Frank Bolden Jr."
	fmt.Println(bolden) //Maj. General Charles Frank Bolden Jr.

	fmt.Println(administrator == major) //true

	//但是解引用将指针指向的变量赋值给另一个变量将产生一个副本
	charles := *major
	*major = "Charles Bolden"
	fmt.Println(charles) //Maj. General Charles Frank Bolden Jr.
	fmt.Println(bolden)  //Charles Bolden

	//两个string变量指向不同的地址，但是只要他们的字符串值相同，那么判等时就是ture
	charles = "Charles Bolden"
	fmt.Println(bolden == charles)   //true
	fmt.Println(&bolden == &charles) //false

	//指向结构体的指针
	timmy := &person{
		name: "Timothy",
		age:  10,
	}

	timmy.superpower = "fly"
	fmt.Printf("%+v\n", timmy) //&{name:Timothy superpower:fly age:10}

	fmt.Println(timmy) //&{Timothy fly 10}

	//指向数组的指针
	superpowers := &[3]string{"flight", "invisibility", "super-strength"}
	fmt.Println(superpowers[1])   //invisibility
	fmt.Println(superpowers[1:3]) //[invisibility super-strength]

}

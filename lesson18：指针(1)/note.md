# 指针(1)：Pointer

## &和*
* &：地址操作符，用以得到变量的内存地址
* *：解引用，提供内存地址指向的值
```
answer := 42
fmt.Println(&answer) //0xc0000a00b0

address := &answer
fmt.Println(*address) //42

fmt.Printf("address is a %T\n", address) //address is a *int
```
上面的address变量实际上是一个`*int`类型的指针，它可以指向类型为int的其他变量。  
* 指针类型也是一种类型，也可以用在变量声明，函数形参，返回值类型，结构字段类型等。
```
china := "China"
var home *string
fmt.Printf("home is a %T\n", home) //home is a *string

home = &china
fmt.Println(*home) //China
```
但是如果上面的home变量不可以指向除了string类型之外的其他类型，这使得Go相对于C来说更加安全。
```
home = &answer //error:cannot use &answer (type *int) as type *string in assignment
```

## 指针的作用
看示例代码，一目了然
```
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
```

### 指向结构的指针
对于指向结构的指针，Go的设计者为其做了优化。比如在复合字面量的前面可以放`&`，但是在访问字段时，前面可以不用加`*`，因为Go会自动实施指针解引用。
```
type person struct {
	name, superpower string
	age              int
}

timmy := &person{
    name: "Timothy",
    age:  10,
}

timmy.superpower = "fly"
fmt.Printf("%+v\n", timmy) //&{name:Timothy superpower:fly age:10}
```
**注意：** 字符串字面量和整数浮点数字面量之前不能放置`&`。

### 指向数组的指针
和指向结构的指针类似，对于数组的复合字面量，Go会自动实施指针解引用。
```
superpowers := &[3]string{"flight", "invisibility", "super-strength"}
fmt.Println(superpowers[1])   //invisibility
fmt.Println(superpowers[1:3]) //[invisibility super-strength]
```
# 公开和未公开标识符

之前都接触过，一个包中命名是大写字母开头的标识符是公开的(public)，被包外的代码可见。以小写字母开头的是未公开的(private)，则包外的代码不可见。如下示例
> demo/counters/counters.go
```go
package counters

//声明了未公开的类型
type alertCounter int
```
> demo/main.go
```go
package main

import (
	"fmt"

	"demo26/counters"
)

func main() {
	//创建一个未公开的类型的变量
	counter := counters.alertCounter(10)  //undeclared name: counters

	fmt.Printf("Counter: %d\n", counter)
}
```
由于counters包里的`alertCounter`是小写字母开头，所以该变量是未公开的，包外无法访问。如果改为`AlertCounter`则不会产生编译错误了。另一个办法则是为其定义构造函数的方式，如下示例将counters包里的实现改为工厂模式。
> demo/counters/counters.go
```go
package counters

//声明了未公开的类型
type alertCounter int

// New 创建并返回一个未公开的alertCounter 类型的值
func New(value int) alertCounter {
    return alertCounter(value)
}
```
> demo/main.go
```go
package main

import (
	"fmt"

	"demo26/counters"
)

func main() {
	//创建一个未公开的类型的变量
	//counter := counters.alertCounter(10) //error! : undeclared name: counters

	// 使用 counters 包公开的 New 函数来创建一个未公开的类型的变量
	counter := counters.New(10)

	fmt.Printf("Counter: %d\n", counter)
}
```
将工厂函数命名为`New`是Go语言的一个习惯。但是值得注意的是这个`New`函数创建了一个未公开的类型并赋值给了调用者，这个程序可以编译并且运行。
解释：
1. 公开或者未公开的标识符，不是一个值。就是说未公开的是`alertCounter`这个类型而不是`alertCounter`类型的变量值。
2. **短变声明操作符**有能力捕获引用的类型，并创建一个未公开的类型的变量。永远不能显式创建一个未公开的类型的变量，不过**短变声明操作符**可以这么做。

#### 可见标识符
再看一个例子
> demo/entities/entities.go
```go
package entities

// User 在程序里定义一个用户类型
type User struct {
	Name  string
	email string
}
```
> demo/main.go
```go
package main

import (
	"fmt"

	"demo26/counters"
)

func main() {
	// 创建 entities 包中的 User 类型的值
	u := entities.User{
		Name:  "Bill",
		email: "bill@email.com",  //unknown field email in struct literal
	}
}
```
很好理解，因为结构entities的字段`email`是小写字母开头，所以该字段是未公开的。相反`Name`则是公开的。如果将`email`改为`Email`，则代码便可正常编译。如果继续将entities中的代码做如下修改。
> demo/entities/entities.go
```go
package entities

// user 在程序里定义一个用户类型
type user struct {
	Name  string
	Email string
}

// Admin 在程序里定义了管理员
type Admin struct {
	user   // 嵌入的类型是未公开的
	Rights int
}
```
将`User`改为`user`，即改为未公开类型。再声明多一个公开类型`Admin`，嵌入一个未公开的`user`类型。此时main函数再做如下调整。
> demo/main.go
```go
package main

import (
	"fmt"

	"demo26/entities"
)

func main() {
	// 创建 entities 包中的 Admin 类型的值
	a := entities.Admin{
		Rights: 10,
	}


	a.Name = "Bill"
	a.Email = "bill@email.com"
	fmt.Printf("User: %v\n", a)
}
```
创建了 entities 包中的`Admin`类型的值。由于内部类型`user`是未公开的，所以无法直接通过结构字面量的方式初始化该内部类型`user`。但是由于嵌入类型的字段是公开的，又由于内部类型的标识符**提升**到了外部类型，所以这些内嵌类型的公开字段也可以通过外部类型的字段来访问。
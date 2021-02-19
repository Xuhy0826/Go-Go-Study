# 组合（Composition）与转发（Forwarding）

## 组合：合并结构
简单理解，就是一个复杂点的结构体，可以使用一些的简单的结构体来组成，这样在维护和语义上都更加直观和清晰。举个栗子。
```go
//员工
type employee struct {
	id         int64
	name       string
	department department
	account    account
}

//部门
type department struct {
	name string
	code string
}

//账户
type account struct {
	logID string
	level int64
}

func main() {
	dept := department{"Production", "DP"}
	acc := account{logID: "12345", level: 2}

	jack := employee{
		id:         1001,
		name:       "jack",
		department: dept,
		account:    acc,
	}

	fmt.Printf("%+v\n", jack)     
	//output: {id:1001 name:jack department:{name:Production code:DP} account:{logID:12345 level:2}}
	fmt.Printf("jack's department is %v\n", jack.department.name) 
	//output: jack's department is Production
}
```
示例很简单，用法和其他语言很类似，这里就不赘述了。

## 方法的转发
如果为“内结构”绑定方法，“外结构”一样可以调用。
```go
...

//为account类型绑定方法
func (acc account) salary() int64 {
	return acc.level * 2500
}

func main() {
    ...

	fmt.Printf("jack's salary is %v now\n", jack.account.salary()) //jack's salary is 5000 now
}
```
如果想直接在`employee`这个结构上直接使用`salary()`方法，可以为`employee`也绑定一个方法来转发`account`的`salary()`方法。
```
···

//为employee类型绑定方法，转发account的salary方法
func (e employee) salary() int64 {
	return e.account.salary()
}

func main() {
    ...

	fmt.Printf("jack's salary is %v now\n", jack.salary()) //jack's salary is 5000 now
```
其实像上面这样来实现方法的转发还是较为麻烦，Go语言中可以通过**struct嵌入**来实现方法的转发。
* 实现struct嵌入的写法：在struct中只给定字段类型，不写字段名即可。不写字段名的话，就是默认使用这个类型名称作为其字段的名称了。
改写之前的`employee`来实现struct转发，这次我们就不需要再为`employee`手动绑定`salary()`方法了。
```go
type employee struct {
	id   int64
	name string
	department
	account
}
//部门
type department struct {
	name string
	code string
}

//账户
type account struct {
	logID string
	level int64
}

//为account类型绑定方法
func (acc account) salary() int64 {
	return acc.level * 2500
}

func main() {
	dept := department{"Production", "DP"}
	acc := account{logID: "12345", level: 2}
	jack := employee{
		id:         1001,
		name:       "jack",
		department: dept,
		account:    acc,
	}
	//调用方法
	fmt.Printf("jack's salary is %v now\n", jack.salary()) //jack's salary is 5000 now

    //访问字段
	fmt.Println("jack's logID is ", jack.account.logID) //jack's logID is 12345

    //jack.logID等价于jack.account.logID
    fmt.Println("jack's logID is ", jack.logID)         //jack's logID is  12345
}
```
* 值得注意，不仅仅是struct，可以转发任意的类型，用法都是一样的。

## 命名冲突
如果使用**struct嵌入**，如果嵌入的两个struct拥有相同的方法名或者字段名，就会遇到**命名冲突**的问题，Go无法确定你到底想调用哪个方法了，发生歧义。
```go
···

//为account类型绑定方法
func (acc account) salary() int64 {
	return acc.level * 2500
}

//继续为department绑定一个同名的salary方法
func (dept department) salary() int64 {
	return 2500 * int64(len(dept.code))
}

func main() {
    ··· 

	fmt.Printf("jack's salary is %v now\n", jack.salary()) //这里就会报错了：ambiguous selector
}
```
为了解决这种情况，要么避免使用到同名方法，要么就在单独为“父类型”`employee`显式的声明一个`salary()`方法。
```go
    ··· 

//来解决“命名冲突”
func (e employee) salary() int64 {
	return e.account.salary()
}

func main() {
    ··· 

	fmt.Printf("jack's salary is %v now\n", jack.salary()) //报错消失
}
```

## 使用组合还是继承
引经据典，大佬们这么说的：   
· 优先使用对象组合而不是类的继承
> Favor object composition over class inheritance   ——Gang of Four

· 对传统的继承不是必须的；所有使用继承解决的问题都可以使用其他方法解决
> Use of classical inheritance is always optional;every problem that it solves can be solved another way.   ——Sandi Metz
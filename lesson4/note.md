# 基本类型

## 实数

### 浮点类型
go语言中有两种浮点类型：
1. float64（占8字节内存）,默认的浮点类型，就是说不显式地指定float32，都是使用float64来定义一个带小数点的变量
2. float32（占4字节内存）

days := 365.2425    //短声明，days会被Go编译器推断为浮点类型默认的float64类型
如果需要使用float32，我们必须指定类型
var pi64 = math.Pi      
var pi32 float32 = math.Pi  //这样声明的变量才会是float32

### 零值
也就是浮点类型的默认值，如果只声明变量不赋值，那么便是零值
var price float
等同于
price := 0.0

### 格式化输出浮点值
格式化输出需要使用到fmt.Printf()函数：
fmt.Printf("%f\n", third)     //0.333333
fmt.Printf("%.2f\n", third)   //0.33，.2f就是表示小数点后保留2位
fmt.Printf("%4.2f\n", third)  //0.33，4.2f表示总宽（长）度为4，小数点后保留2位
fmt.Printf("%5.2f\n", third)  // 0.33，5.2f表示总宽（长）度为5，小数点后保留2位，长度不够使用空格来补
fmt.Printf("%05.2f\n", third) //00.33，05.2f表示总宽（长）度为5，小数点后保留2位，长度不够使用“0”来补

### 浮点类型的精确性
由于计算机只能通过0和1来表示浮点数，所以浮点数会经常受到舍入错误的影响
比如：
piggyBank := 0.1
piggyBank += 0.2
fmt.Println(piggyBank) //0.30000000000000004

由上面提到的浮点类型的精确度问题，就会导致浮点数的比较
fmt.Println(piggyBank == 0.3) //false
一个折中的解决方案就是使用一定精确度来判断是否相等
fmt.Println(math.Abs(piggyBank-0.3) < 0.0001) //true

那么说到底，避免浮点数精确度问题的最佳方案就是：不使用浮点数
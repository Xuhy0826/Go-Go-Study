# Go操作PostgreSql：CURD

## 连接到数据库
Go连接数据要用到的包
1. database/sql
2. 相应的数据库的驱动（PostgreSql：github.com/bmizerany/pq）

#### 连接到数据库
* 要连接到数据库，首先要加载数据库的驱动，驱动里包含与数据库交互的逻辑。
* 方法：sql.Open("%数据库驱动名%", "%数据源名称%")，返回值是指向`sql.DB`的指针。
* Open函数不会连接到数据库，并且不进行参数校验，只是将后续进行数据库操作的结构体给创建
* 连接是懒加载的
* sql.DB是线程安全的，多个goroutine可以同时操作
* sql.DB不需要手动关闭
* sql.DB是个抽象包含了数据库连接的池，并会自己进行维护

#### 如何获得驱动
* `sql.Register("%数据库驱动名%", %实现了driver.Driver的结构体%)`
* 驱动一般是在包的Init函数中进行注册的，也就是包能够自我注册。

#### 安装数据库驱动
* `go get github.com/bmizerany/pq`

#### Ping函数
* 通过Ping函数来测试下连接是否有效

#### PingContext函数
* 用来验证数据库的连接是否有效，如有必要则建立一个连接
* 函数需要传入一个`Context`类型
* 使用`context.Background()`方法获取的是一个非nil的Context。改Context没有截止时间，没有值，不会被取消。它通常在测试，或者main函数，或初始化中用的顶级Context

## 简单查询
使用sql.DB这个结构体来进行操作
* Query ：查询多行数据
* QueryRow ：查询单笔数据
* QueryContext ：Query的上下文版本
* QueryRowContext ：QueryRow的上下文版本

#### Query
* 返回类型：`type Rows struct{}`
* `Rows`的方法
* * func(re *Rows) Close() error
* * func(re *Rows) ColumnTypes() ([]*ColumnType, error)
* * func(re *Rows) Column() ([]string, error)  //返回列名
* * func(re *Rows) Err() error
* * func(re *Rows) Next() bool
* * func(re *Rows) Next() bool
* * func(re *Rows) NextResultSet() bool
* * func(re *Rows) Scan(dest ...interface{}) error  //将查到的数据一一赋值到Scan的入参中，有点类似RowMapper功能

#### QueryRow
* 返回类型：`type Row struct{}`
* `Row`的方法
* * func(re *Row) Err() error
* * func(re *Row) Scan(dest ...interface{}) error
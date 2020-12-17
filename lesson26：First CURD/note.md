# Go操作PostgreSql：CURD
Go如何与关系型数据库进行简单的交互操作。

## 连接到数据库
Go连接数据要用到的包
1. 内置包：database/sql
2. 相应的数据库的驱动，需要选择并下载（PostgreSql：github.com/lib/pq）

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
* `go get github.com/lib/pq`

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
查询，返回多笔数据。直接看例子简单明了
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
```
func queryByDate(minDate time.Time) (entityCollection []testEntity, err error) {
	sqlStr := "SELECT id, msg, create_time FROM public.test where create_time > $1"
	rows, err := db.Query(sqlStr, minDate)
	//使用Next方法读数据，类似ADO.NET的Read()方法
	for rows.Next() {
		entity := testEntity{}
		err = rows.Scan(&entity.id, &entity.msg, &entity.createTime)
		if err != nil {
			return
		}
		entityCollection = append(entityCollection, entity)
	}
	return
}
```

#### QueryRow
查询，返回单笔数据。直接看例子简单明了
* 返回类型：`type Row struct{}`
* `Row`的方法
* * func(re *Row) Err() error
* * func(re *Row) Scan(dest ...interface{}) error
```
func queryByID(id int) (entity testEntity, err error) {
	entity = testEntity{}
	sqlStr := "select t.id, t.msg, t.create_time from public.test t where t.id = $1"
	err = db.QueryRow(sqlStr, id).Scan(&entity.id, &entity.msg, &entity.createTime)
	return
}
```

#### Exec
执行命令，直接看例子简单明了
* 返回类型：`type Result struct{}`
```
func updateEntity(entity testEntity) (newEntity testEntity, err error) {
	sqlStr := "UPDATE public.test SET msg=$1, create_time=$2 WHERE id=$3"
	_, err = db.Exec(sqlStr, entity.msg, time.Now(), entity.id)
	if err != nil {
		return
	}
	newEntity, err = queryByID(entity.id)
	return
}
```

#### Prepare
这个比较特殊，可以成为预处理，下面[摘抄来自](https://blog.csdn.net/qq_34857250/article/details/100569676)
> * 什么是预处理
> 普通SQL语句执行过程：
> 1. 客户端对SQL语句进行占位符替换得到完整的SQL语句。
> 2. 客户端发送完整SQL语句到数据库服务端
> 3. 数据库服务端执行完整的SQL语句并将结果返回给客户端。
> 
> * 预处理执行过程：
> 1. 把SQL语句分成两部分，命令部分与数据部分。
> 2. 先把命令部分发送给数据库服务端，数据库服务端进行SQL预处理。
> 3. 然后把数据部分发送给MySQL服务端，数据库服务端对SQL语句进行占位符替换。
> 4. 数据库服务端执行完整的SQL语句并将结果返回给客户端。
> 
> * 为什么需要 Prepare
> 1. 优化数据库服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
> 2. 避免SQL注入问题

```
func insertEntities(entityCollection []testEntity) (err error) {
	sqlStr :=
		`
	INSERT INTO public.test(
	id, msg, create_time
	)
	VALUES (
	$1, $2, $3
	)
	`
	stmt, err := db.Prepare(sqlStr) //进行预处理
	defer stmt.Close()              //由于statement需要关闭
	if err != nil {
		fmt.Println("prepare failed ,", err.Error())
		return
	}

	for _, e := range entityCollection {
		_, err = stmt.Exec(e.id, e.msg, time.Now())
		if err != nil {
			fmt.Println("insert failed , id = ", e.id, err.Error())
			return
		}
	}
	return
}
```

#### 事务
简单的几个要素：
1. 开启事务：`db.Begin()`，返回`sql.Tx`，后续使用其进行sql操作
2. 执行操作：`tx.Exec()` 或 `tx.Query()`等
3. 提交事务：`tx.Commit()`
4. 回滚：`tx.Rollback()`
```
func execByTransaction(entity testEntity) (err error) {
	tx, err := db.Begin() //开启事务
	if err != nil {
		fmt.Println("open Transaction failed,", err.Error())
	}
	sqlStr1 := "INSERT INTO public.test(id, msg, create_time) VALUES ($1, $2, $3)"
	sqlStr2 := "Update public.test SET msg=$1, create_time=$2 WHERE id=$3"

	_, err = tx.Exec(sqlStr1)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Println("exec sqlStr1 failed,", err.Error())
		return
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Println("exec sqlStr2 failed,", err.Error())
		return
	}
	err = tx.Commit() //提交事务
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Println("commit failed", err.Error())
		return
	}
	return
}
```
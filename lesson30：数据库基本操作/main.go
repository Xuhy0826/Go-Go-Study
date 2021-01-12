package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

//定义全局变量
var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "biz"
)

type testEntity struct {
	id         int
	msg        string
	createTime time.Time
}

func (entity testEntity) String() string {
	return fmt.Sprintf("entity : {id: %v, msg: %v, createTime: %v }", entity.id, entity.msg, entity.createTime.Format("2006-01-02 15:04:05"))
}

func main() {
	//（1）连接数据库
	err := initDb()
	if err == nil {
		fmt.Println("connect successfully")

		//（2）查询数据一
		entity, err := queryByID(3)
		if err != nil {
			fmt.Println("query failed,", err.Error())
		} else {
			fmt.Println(entity)
		}
		fmt.Println("============================================")
		//（3）查询数据二
		date := time.Date(2020, 12, 16, 0, 0, 0, 0, time.Local)
		entityCollection, err := queryByDate(date)
		if err != nil {
			fmt.Println("query failed,", err.Error())
		} else {
			for _, e := range entityCollection {
				fmt.Println(e)
			}
		}
		fmt.Println("============================================")
		//（4）更新操作
		entity.msg = entity.msg + " !!!"
		newEntity, err := updateEntity(entity)
		if err != nil {
			fmt.Println("execute failed,", err.Error())
		} else {
			fmt.Println(newEntity)
		}
		//（5）批量操作
		startIndex := 10
		entities := make([]testEntity, 0, 3)
		for i := startIndex; i < startIndex+4; i++ {
			entities = append(entities, testEntity{
				id:  i,
				msg: fmt.Sprintf("%d shot", i),
			})
		}
		err = insertEntities(entities)
		if err != nil {
			fmt.Println("insert failed,", err.Error())
		} else {
			for _, e := range entities {
				entity, _ := queryByID(e.id)
				fmt.Println(entity)
			}
		}
	}
}

//initDb 初始化与连接
func initDb() error {
	connStr :=
		fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	fmt.Println(connStr)

	var err error
	db, err = sql.Open("postgres", connStr) //驱动名称为：postgres

	if err != nil {
		fmt.Println("Connected failed.", err.Error())
		return err
	}
	//ctx := context.Background()
	//err = db.PingContext(ctx)

	err = db.Ping()
	if err != nil {
		fmt.Println("ping failed.", err.Error())
	}
	return nil
}

//queryById 通过Id查询单笔数据
func queryByID(id int) (entity testEntity, err error) {
	entity = testEntity{}
	sqlStr := "select t.id, t.msg, t.create_time from public.test t where t.id = $1"
	err = db.QueryRow(sqlStr, id).Scan(&entity.id, &entity.msg, &entity.createTime)
	return
}

//queryByDate 通过创建时间查询多笔数据
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

//updateEntity 更新操作
func updateEntity(entity testEntity) (newEntity testEntity, err error) {
	sqlStr := "UPDATE public.test SET msg=$1, create_time=$2 WHERE id=$3"
	_, err = db.Exec(sqlStr, entity.msg, time.Now(), entity.id)
	if err != nil {
		return
	}
	newEntity, err = queryByID(entity.id)
	return
}

//insertEntities 插入多条数据，使用Prepare
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

//execByTransaction 执行事务
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

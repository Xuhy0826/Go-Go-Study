package main

import (
	"database/sql"
	"fmt"

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

func main() {
	//（1）连接数据库
	err := initDb()
	if err == nil {
		fmt.Println("connect successfully")

		//（2）查询数据
		entity, err := queryByID(3)
		if err != nil {
			fmt.Println("query failer,", err.Error())
		} else {
			fmt.Printf("%+v", entity)
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
	sql := fmt.Sprintf("select t.id, t.msg, t.create_time from public.test t where t.id = $1")
	err = db.QueryRow(sql, id).Scan(&entity.id, &entity.msg, &entity.createTime)
	return
}

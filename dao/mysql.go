package dao

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

//定义基础连接库
var (
	userName  string = "dlan"
	password  string = "root123"
	ipAddrees string = "172.16.0.38"
	port      int    = 3318
	dbName    string = "lepus"
	charset   string = "utf8"
)

func InitDB() (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddrees, port, dbName, charset)
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err)

	}
	Db = db
	return
}

//给归档定义的连接信息

func ConDB(username, password, host, dbname, charset string, port int) (conn *sqlx.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", username, password, host, port, dbname, charset)
	conn, err = sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = conn.Ping()
	if err != nil {
		fmt.Println("connect failed:", err)
	}
	return
}

func GoInception()(conn *sql.DB,err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", "gouser", "go134", "172.16.0.11", 4000)
	conn, err = sql.Open("mysql", dsn)
	if err !=nil{
		panic(err)
	}
	err =conn.Ping()
	if err !=nil{
		return
	}
	return

}

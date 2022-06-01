package dao

import (
	"fmt"
    "log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/syohex/go-texttable"
	"database/sql"
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


func GoInception(sqltext string) {
	//use  test;
	//create table t1(id int primary key);
	//alter table t1 add index idx_id (id);
	//create table t2(jid int primary key);
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", "gouser", "go134", "172.16.0.11", 2121)
	gdb, err := sql.Open("mysql", dsn)
	defer gdb.Close()
	sql := fmt.Sprintf(`/*--user=%s;--password=%s;--host=%s;--port=%d;--check=1;*/
    inception_magic_start;
    %v
    inception_magic_commit;`, userName, password, ipAddrees, port, sqltext)
	rows, err := gdb.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	tbl := &texttable.TextTable{}
	for rows.Next() {
		var order_id, affected_rows, stage, error_level, stage_status, error_message, sql, sequence, backup_dbname, execute_time, sqlsha1, backup_time []uint8
		err = rows.Scan(&order_id, &stage, &error_level, &stage_status, &error_message, &sql, &affected_rows, &sequence, &backup_dbname, &execute_time, &sqlsha1, &backup_time)
		fmt.Println(order_id, affected_rows, stage, error_level, stage_status, error_message, sql, sequence, backup_dbname, execute_time, sqlsha1, backup_time)
		//tbl.AddRow(&order_id, &stage, &error_level, &stage_status, &error_message, &sql, &affected_rows, &sequence, &backup_dbname, &execute_time, &sqlsha1, &backup_time)
		tbl.AddRow(string(order_id), string(affected_rows), string(stage), string(error_level), string(stage_status), string(error_message), string(sql), string(sequence), string(backup_dbname), string(execute_time))

	}
	fmt.Println(tbl.Draw())

}
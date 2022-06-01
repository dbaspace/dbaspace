package service

//"github.com/jmoiron/sqlx"
import (
	"awesomeProject/db-monitorProject/dao"
	"awesomeProject/db-monitorProject/model"
	"database/sql"
	"fmt"
	"strings"
)

func Putddlschema(ddl *model.Tbl_add_ddl_task) (res int64, err error) {
	res, err = dao.Insertddlschema(ddl)
	if err != nil {
		fmt.Println("ddl", err)
		return
	}

	return
}

func CharCheck(sc string, num int) int {
	num = strings.Count(sc, "varchar")
	//"tinyint", "smallint", "int", "integer", "bigint", "float", "double", "decimal", "enum", "bit", "date", "time", "year", "datetime", "timestamp", "char", "varchar", "tinyblob", "tinytext", "blob", "text", "mediumblob", "mediumtext", "longblob", "longtext")
	return num
}

func Alterddl(sql string) {
	//SQL格式处理:换行符
	//"alter tABLE aa add column  name varchar(255) not null default '0' comment 'name,adfdf,',
	//add column cname varchar(255) not null default 0 comment '1';"
	if !strings.HasSuffix(sql, ";") {
		fmt.Println("sql has not suffix [:]")
		return
	}
	if !strings.HasPrefix(strings.ToLower(sql), "alter") {
		fmt.Println("not full [alter table]")
		return
	}
	//获取alter table 后面的数据
	data := strings.Split(sql, " ")
	tmpdata := ""
	for _, kk := range data {
		if kk != "" {
			tmpdata = tmpdata + " " + kk
		}

	}
	tmp := strings.Split(strings.TrimSpace(tmpdata), " ")[3:]
	tmpdata2 := ""
	for _, vv := range tmp {
		tmpdata2 = tmpdata2 + " " + vv
	}
	fmt.Println(tmpdata2)
	tmpdata3 := strings.Split(tmpdata2, "add")
	fmt.Println(tmpdata3)
}

var (
	userName  string = "dlan"
	password  string = "root123"
	ipAddrees string = "172.16.0.38"
	port      int    = 3318
	dbName    string = "lepus"
	charset   string = "utf8"
)
func InceptionCheckSQL(sqltext string)(rows *sql.Rows,err error){
    conn,err:=dao.GoInception()
	if err !=nil{
		return
	}
	sqlexe := fmt.Sprintf(`/*--user=%s;--password=%s;--host=%s;--port=%d;--enable-check=1;*/
    inception_magic_start;
    %v
    inception_magic_commit;`, userName, password, ipAddrees, port, sqltext)
	rows,err=conn.Query(sqlexe)
	if err !=nil{
		return
	}
	return



}
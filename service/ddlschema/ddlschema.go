package service

//"github.com/jmoiron/sqlx"
import (
	"awesomeProject/db-monitorProject/dao"
	"awesomeProject/db-monitorProject/model"
	"database/sql"
	"fmt"
	"strings"
)

var (
	userName  string = "dlan"
	password  string = "root123"
	ipAddrees string = "172.16.0.38"
	port      int    = 3318
	dbName    string = "lepus"
	charset   string = "utf8"
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
	//工单审核通过提交后,处理流程：1获取未执行工单 2、根据输入的条件进行判断 执行对应的操作
	getlist, err := dao.GetDdlList()
	if err != nil {
		return
	}

	for _, li := range getlist {
		if li.Cmd_exe == 1 {
			sqltext := strings.Replace(li.Sqltext, "`", "", -1)
			sqllist := strings.Split(sqltext, " ")
			gg := []string{"truncate", "drop"}
			if res := in(strings.ToLower(sqllist[0]), gg); res {
				fmt.Println("危险命令，禁止操作.....")
				return
			}
			tmp := strings.Join(sqllist[3:], " ")
			fmt.Println(tmp)
		}
		if li.Exe_type == 1 && li.Db_type != 6 {
			var info []model.Tbl_dbinfo_ddllist
			getli := "select c_host,c_portfrom tbl_dbinfo_ddllist where db_type=?"
			err := dao.Db.Select(&info, getli, li.Db_type)
			if err != nil {
				fmt.Println("get dblist failed", err)
			}
			fmt.Println(info)
			for _, key := range info {
				dst := key.C_host
				dot := key.C_port
				if li.Cmd_exe != 6 {
					fmt.Println("exc add column|add index", dst, dot)
				} else {
					fmt.Println("add tablename", dst, dot)
				}
			}
		} else {
			dhost := li.Shost
			dport := li.Sport
			if li.Cmd_exe != 6 {
				fmt.Println("exc add column|add index", dhost, dport)
			} else {
				fmt.Println("add tablename", dhost, dport)
			}
		}

	}
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

func in(s string, gg []string) {
	panic("unimplemented")
}

func InceptionCheckSQL(sqltext string) (rows *sql.Rows, err error) {
	conn, err := dao.GoInception()
	if err != nil {
		return
	}
	sqlexe := fmt.Sprintf(`/*--user=%s;--password=%s;--host=%s;--port=%d;--enable-check=1;*/
    inception_magic_start;
    %v
    inception_magic_commit;`, userName, password, ipAddrees, port, sqltext)
	//	fmt.Println(sqlexe)
	rows, err = conn.Query(sqlexe)
	if err != nil {
		fmt.Println("exe auth failed", err)
		return
	}
	return

}

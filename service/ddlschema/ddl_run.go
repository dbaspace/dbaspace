package service

import (
	"awesomeProject/db-monitorProject/dao"
	"awesomeProject/db-monitorProject/model"
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
)

func AddTaskRun(host string, port int, table, sql, dbname, db_typename string, cmd_exe, command_exe, cmd_idc int) {
	//username, password, host, dbname, charset string, port int
	//获取目标端
	dsconn, err := dao.ConDB("dlan", "root123", host, "mysql", "utf8", port)
	if err != nil {
		panic(err)
	}
	sqlstr := `SELECT table_schema,table_name FROM information_schema.tables  WHERE table_schema<>'mysql' AND table_schema<>'information_schema' AND table_schema<>'performance_schema' AND 
			 table_schema<>'sys'  and table_schema like '?%' and table_name like '?' GROUP BY table_schema,table_name`
	var dname []*model.Dblist
	var dbt string
	//匹配模式
	if db_typename != "shopcrm" {
		dbt = fmt.Sprintf(`?%`, table[0])
	} else {
		dbt = fmt.Sprintf(`?`, table[0])
	}
	err = dsconn.Select(&dname, sqlstr, dbname[0], dbt)
	if err != nil {
		fmt.Println("get dbname failed:::", err)
		return
	}
	for _, i := range dname {
		ts := i.Dbname
		tb := i.Tablename
		var pcmd string = ""
		start_time := time.Now().Format("2006-01-02 15:04:05")
		if command_exe == 1 {
			if cmd_exe == 1 && cmd_idc == 1 {
				pcmd = fmt.Sprintf(`/usr/bin/gh-ost  --alter="?" --database='?' --table='?' 
			--host='?' --user='?' --password='?' --port='?' --allow-on-master --assume-master-host='?:?' --max-load=Threads_running=350 --critical-load=Threads_running=300 --chunk-size=5000 --max-lag-millis=150000 --assume-rbr --timestamp-old-table --verbose --cut-over=default --concurrent-rowcount --default-retries=120 
			--initially-drop-ghost-table --ok-to-drop-table --panic-flag-file=/tmp/ghost.panic_?_?.flag --execute`, sql[0], ts, tb, host, "username", "password", port, host, port, ts, tb)
			} else if cmd_exe == 1 && cmd_idc == 2 {
				pcmd = fmt.Sprintf(`/usr/bin/gh-ost --aliyun-rds  --alter="?" --database='?' --table='?' 
			--host='?' --user='?' --password='?' --port='?' --allow-on-master --assume-master-host='?:?' --max-load=Threads_running=350 --critical-load=Threads_running=300 --chunk-size=5000 --max-lag-millis=150000 --assume-rbr --timestamp-old-table --verbose --cut-over=default --concurrent-rowcount --default-retries=120 
			--initially-drop-ghost-table --ok-to-drop-table --panic-flag-file=/tmp/ghost.panic_?_?.flag --execute`, sql[0], ts, tb, host, "username", "password", port, host, port, ts, tb)
			}
			fmt.Println("EXE GH-OST")
			c := exec.Command("bash", "-c", pcmd)
			stdout, err := c.StdoutPipe()
			if err != nil {
				fmt.Println(err)
				return
			}
			reader := bufio.NewReader(stdout)
			for {
				readStr, err := reader.ReadString('\n')
				if err != nil || err == io.EOF {
					break
				}
				if strings.Contains(readStr, "Done") {
					fmt.Println("EXE SUCCESS....")
				} else {
					fmt.Println("EXE FAILED....")
				}

			}
			end_time := time.Now().Format("2006-01-02 15:04:05")
			sql_insert := "insert into db_osc_all(c_host,c_port,dbname,tablename,info,start_time,end_time)values(?,?,?,?,?,?,?)"
			_, err = dao.Db.Exec(sql_insert, host, port, ts, tb, sql[0], start_time, end_time)
			if err != nil {
				fmt.Println("exec log failed", err)
			}

		} else {
			pcmd = fmt.Sprintf(`alter table ?.? ?;`, ts, tb, sql[0])
			_, err := dsconn.Exec(pcmd)
			if err != nil {
				fmt.Println("EXE ALTER FAILED::::", err)
			}
			//记录日志

		}
	}
}

func CreateTable(host string, port int, table, sql, dbname, db_typename string, cmd_exe, command_exe, cmd_idc int) {
	status := "00"
	start_time := time.Now().Format("2006-01-02 15:04:05")
	dsconn, err := dao.ConDB("dlan", "root123", host, "mysql", "utf8", port)
	if err != nil {
		status = "01"
		fmt.Println("conn desc db:::", err)
		panic(err)
	}
	sqlstr := `SELECT table_schema,table_name FROM information_schema.tables  WHERE table_schema<>'mysql' AND table_schema<>'information_schema' AND table_schema<>'performance_schema' AND 
			 table_schema<>'sys'  and table_schema like '?%' GROUP BY table_schema`
	var dname []*model.Dblist

	err = dsconn.Select(&dname, sqlstr, dbname[0])
	if err != nil {
		status = "02"
		fmt.Println("get dbname failed:::", err)
		return
	}
	for _, li := range dname {
		tsc := fmt.Sprintf("use  ?;?", li.Dbname, sql[0])
		ret, err := dsconn.Exec(tsc)
		if err != nil {
			status = "02"
			fmt.Println("exec  create table failed::", err)
		}
		_, err = ret.RowsAffected()
		if err != nil {
			status = "03"
			fmt.Println("create table failed:::", err)
		}
		status = "05"
		//Tbl_osc_all_record
		sqlr := "insert into tbl_osc_all_record(taskid,c_host,c_port,dbname,tablename,info,status,start_time,end_time)values(?,?,?,?,?,?,?,?)"
		//dao.InsertServiceLog(sqlr, host, port, li.Dbname, li.Tablename, sql, status)
		end_time := time.Now().Format("2006-01-02 15:04:05")
		_, err := dao.Db.Exec(sqlr, host, port, li.Dbname, li.Tablename, sql, status, start_time, end_time)
		if err != nil {
			fmt.Println("record log failed:::::", err)
		}
	}
}

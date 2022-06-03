package service

import (
	"awesomeProject/db-monitorProject/dao"
	"awesomeProject/db-monitorProject/model"
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
)

func GetALLPtarchiverList() (ptlist []*model.Tbl_dbarchiver_task, err error) {
	//获取工单提交的工单列表
	ptlist, err = dao.GetPtarchiveList()
	if err != nil {
		return
	}
	if len(ptlist) <= 0 {
		return
	}
	return
}

//根据工单号查询
func GetPtarchiverTaskId(taskid string) (ptlist *model.Tbl_dbarchiver_task, err error) {
	ptlist, err = dao.GetPtarchiverTaskId(taskid)
	if err != nil {
		return
	}
	if ptlist == nil {
		return
	}

	return
}

//提交归档工单
func PutPtarchiver(pt *model.Tbl_dbarchiver_task) (ptarchiverIds int64, err error) {
	if pt == nil {
		return
	}
	if flage := CheckArgs(pt); flage {
		ptarchiverIds, err = dao.InsertPtarchiver(pt)
		if err != nil {
			return
		}
	}

	return
}

//任务执行：1判断任务类型（delete|pt）2、判断任务是立即 还是定时
//任务触发 可以手动 自动 定时任务
func TaskRun() {
	taskList, err := dao.GetPtarchiverExecList()
	if err != nil {
		return
	}
	if taskList == nil {
		return
	}
	for _, key := range taskList {
		Taskid := key.Taskid
		Shost := key.Shost
		Sport := key.Sport
		Sschname := key.Sschname
		Stablename := key.Stablename
		Dschename := key.Dschename
		Dtabname := key.Dtabname
		Is_idx := key.Is_idx
		Is_condition := key.Is_condition
		Dhost := key.Dhost
		Dport := key.Dport
		exdtype := key.Exdtype
		//获取归档实例实例
		conn, err := dao.ConDB("dlan", "root123", Dhost, Dschename, "utf8", Dport)
		if err != nil {
			return
		}
		err = conn.Ping()
		if err != nil {
			return
		}
		defer conn.Close()
		conn.SetMaxOpenConns(1)
		conn.SetMaxIdleConns(1)

		//是否存在指定索引校验
		sqlstr := fmt.Sprintf("select INDEX_NAME from information_schema.statistics where table_schema=? and table_name=? and index_name=? ")
		var index string
		err = conn.Get(&index, sqlstr, Dschename, Dtabname, Is_idx)
		if err != nil {
			fmt.Println("查询索引异常", err)
			return
		}

		switch exdtype {
		case 1:

			fmt.Println("exec delete")
			//执行单次提交的的delete
			//taskid,shost,sport,sschname,stablename,dhost,dport,dschename,dtabname,is_condition,is_state,is_store,delete_count,insert_count,start_time,end_time
			sqlstr = fmt.Sprintf("delete from %s.%s  where %s limit 10000", Dschename, Dtabname, Is_condition)
			delTaskSingle(sqlstr, Taskid, Shost, Sport, Sschname, Stablename, Dhost, exdtype, Dport, Dschename, Dtabname, Is_idx, Is_condition, conn)
			if err != nil {
				fmt.Println("delete failed....")
			}
			//fmt.Println(del_count)
		case 2:
			del_count, err := ptTaskSingle(Taskid, Shost, Sport, Sschname, Stablename, Dhost, exdtype, Dport, Dschename, Dtabname, Is_idx, Is_condition)
			if err != nil {
				fmt.Println("pt delete failed")
			}
			fmt.Println(del_count)
			fmt.Println("exec pt")
		case 3:
			//批量DELETE处理
			MultiTaskDelete(sqlstr, Taskid, Shost, Sport, Sschname, Stablename, Dhost, exdtype, Dport, Dschename, Dtabname, Is_idx, Is_condition, conn)
			fmt.Println("mutil delete")
		default:
			fmt.Println("err")
		}

	}
}

//处理delete事件体，支持单条提交和批量提交处理
func Schesql(sqlstr string, Taskid, Shost string, Sport int, Sschname, Stablename, Dhost string, exdtype, Dport int, Dschename, Dtabname, Is_idx, Is_condition string, conn *sqlx.DB) {
	start_time := time.Now().Format("2006-01-02 15:04:05")
	is_state := 0
	del_count := 0
	for {
		stmt, err := conn.Preparex(sqlstr)
		if err != nil {
			break
		}
		result, err := stmt.Exec()
		if err != nil {
			break
		}
		i, err := result.RowsAffected()
		if err != nil {
			break
		}
		del_count = del_count + int(i)
		if i == 0 {
			is_state = 1
			break
		}
		fmt.Println(sqlstr)
		time.Sleep(1 * time.Second)
	}
	//记录完成日志
	end_time := time.Now().Format("2006-01-02 15:04:05")
	insql := "insert into tbl_dbarchiver_task_log(taskid,shost,sport,sschname,stablename,dhost,dport,dschename,dtabname,is_condition,is_state,is_store,exdtype,delete_count,insert_count,start_time,end_time)values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	//res, err := dao.Db.Exec(insql, Taskid, Shost, Sport, Sschname, Stablename, Dhost, Dport, Dschename, Dtabname, Is_condition, is_state, 1, exdtype, del_count, del_count, start_time, end_time)
	res, err := dao.InsertServiceLog(insql, Taskid, Shost, Sport, Sschname, Stablename, Dhost, Dport, Dschename, Dtabname, Is_condition, is_state, 1, exdtype, del_count, del_count, start_time, end_time)
	if err != nil {
		fmt.Println("exec record log failed")
		return
	}
	_, err = res.LastInsertId()
	if err != nil {
		fmt.Printf("记录日志失败[%v   %v    %v     %v]", Shost, Sport, Sschname, Stablename)
		return
	}
	//return
}
func delTaskSingle(sqlstr, Taskid, Shost string, Sport int, Sschname, Stablename, Dhost string, exdtype, Dport int, Dschename, Dtabname, Is_idx, Is_condition string, conn *sqlx.DB) {
	//sqlstr = fmt.Sprintf("delete from %s.%s  where %s limit 10000", Dschename, Dtabname, Is_condition)
	Schesql(sqlstr, Taskid, Shost, Sport, Sschname, Stablename, Dhost, exdtype, Dport, Dschename, Dtabname, Is_idx, Is_condition, conn)
	//return
}

//Dhost, Dschename, Dtabname, Is_idx, Is_condition string, Dport int
//pt归档 启动协程处理
func ptTaskSingle(Taskid, Shost string, Sport int, Sschname, Stablename, Dhost string, exdtype, Dport int, Dschename, Dtabname, Is_idx, Is_condition string) (del_count int64, err error) {
	//start_time := time.Now().Format("2006-01-02 15:04:05")
	pcmd := fmt.Sprintf("/usr/bin/pt-archiver --source h=%s,P=%d,u=%s,p=%s,D=%s,t=%s,i=%s --dest h=%s,P=%d,u=%s,p=%s,D=%s,t=%s --where %s --progress 10000 --limit=10000 --txn-size 10000 --bulk-insert --bulk-delete --no-check-charset --noversion-check --statistics --purge", Shost, Sport, "dlan", "root123", Sschname, Stablename, Is_idx, Dhost, Dport, "dlan", "root123", Dschename, Dtabname, Is_condition)
	c := exec.Command("bash", "-c", pcmd)
	stdout, err := c.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	start_time := time.Now().Format("2006-01-02 15:04:05")
	insert_count := 0
	delete_count := 0
	select_count := 0
	is_state := 0
	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup, insert_count, delete_count, select_count, is_state *int) {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readStr, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				break
			}
			i := strings.Split(readStr, " ")
			Num := i[1]
			p, _ := strconv.Atoi(strings.Replace(Num, "\n", "", -1))
			switch {
			case strings.Contains(readStr, "INSERT"):
				*insert_count = p
			case strings.Contains(readStr, "DELETE"):
				*delete_count = p
			case strings.Contains(readStr, "SELECT"):
				*select_count = p
			}
			if *insert_count == *delete_count && *insert_count > 1 {
				*is_state = 1
			} else if *insert_count == *delete_count {
				*is_state = 2
			} else {
				*is_state = 3
			}
		}
		end_time := time.Now().Format("2006-01-02 15:04:05")
		insql := "insert into tbl_dbarchiver_task_log(taskid,shost,sport,sschname,stablename,dhost,dport,dschename,dtabname,is_condition,is_state,is_store,exdtype,delete_count,insert_count,start_time,end_time)values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
		//res, err := dao.Db.Exec(insql, Taskid, Shost, Sport, Sschname, Stablename, Dhost, Dport, Dschename, Dtabname, Is_condition, is_state, 1, exdtype, del_count, del_count, start_time, end_time)
		res, err := dao.InsertServiceLog(insql, Taskid, Shost, Sport, Sschname, Stablename, Dhost, Dport, Dschename, Dtabname, Is_condition, is_state, 1, exdtype, del_count, del_count, start_time, end_time)
		if err != nil {
			fmt.Println("exec record log failed")
			return
		}
		_, err = res.LastInsertId()
		if err != nil {
			return
		}
	}(&wg, &insert_count, &delete_count, &select_count, &is_state)
	err = c.Start()
	wg.Wait()

	return
}

//批量处理的delete sql 需要根据分号隔开“;”
func MultiTaskDelete(sqlstr, Taskid, Shost string, Sport int, Sschname, Stablename, Dhost string, exdtype, Dport int, Dschename, Dtabname, Is_idx, Is_condition string, conn *sqlx.DB) {
	fmt.Println("处理批量的delet SQL文件.....")
	sqllist := strings.Split(sqlstr, ";")
	for _, v := range sqllist {
		newsql := fmt.Sprintf(v+"%v", " limit 10000;")
		fmt.Println(newsql)
		Schesql(sqlstr, Taskid, Shost, Sport, Sschname, Stablename, Dhost, exdtype, Dport, Dschename, Dtabname, Is_idx, Is_condition, conn)
	}

}
